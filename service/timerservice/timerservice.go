package timerservice

import (
	"context"
	"fmt"
	"time"

	"github.com/oldjon/gutil/env"
	"github.com/oldjon/gutil/gdb"
	gprotocol "github.com/oldjon/gutil/protocol"
	grmux "github.com/oldjon/gutil/redismutex"
	gxgrpc "github.com/oldjon/gx/modules/grpc"
	"github.com/oldjon/gx/service"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/util"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type TimerService struct {
	mpb.UnimplementedTimerServiceServer
	name        string
	logger      *zap.Logger
	config      env.ModuleConfig
	etcdClient  *etcd.Client
	host        service.Host
	connMgr     *gxgrpc.ConnManager
	tcpMsgCoder gprotocol.FrameCoder
	rm          *timerResourceMgr
	kvm         *service.KVMgr
	serverEnv   uint32
	sm          *util.ServiceMetrics
	dao         *timerDAO
}

// NewTimerService create a TimerService entity
func NewTimerService(driver service.ModuleDriver) (gxgrpc.GRPCServer, error) {
	svc := &TimerService{
		name:       driver.ModuleName(),
		logger:     driver.Logger(),
		config:     driver.ModuleConfig(),
		etcdClient: driver.Host().EtcdSession().Client(),
		host:       driver.Host(),
		kvm:        driver.Host().KVManager(),
		sm:         util.NewServiceMetrics(driver),
	}

	dialer := gxgrpc.Dialer{
		HostName:   driver.Host().Name(),
		Tracer:     driver.Tracer(),
		EtcdClient: svc.etcdClient,
		Logger:     svc.logger,
		EnableTLS:  svc.config.GetBool("enable_tls"),
		CAFile:     svc.config.GetString("ca_file"),
		CertFile:   svc.config.GetString("cert_file"),
		KeyFile:    svc.config.GetString("key_file"),
	}
	svc.connMgr = gxgrpc.NewConnManager(&dialer)

	var err error
	svc.rm, err = newTimerResourceMgr(svc.logger, svc.sm)
	if err != nil {
		return nil, err
	}

	redisMux, err := grmux.NewRedisMux(svc.config.SubConfig("redis_mutex"), nil, svc.logger, driver.Tracer())
	if err != nil {
		return nil, err
	}

	gameRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("game_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	pvpRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("pvp_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	tmpRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("tmp_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	svc.dao = newTimerDAO(svc, redisMux, gameRedis, pvpRedis, tmpRedis)

	svc.serverEnv = uint32(svc.config.GetInt64("server_env"))
	svc.tcpMsgCoder = gprotocol.NewFrameCoder(svc.config.GetString("protocol_code"))

	return svc, err
}

func (svc *TimerService) Register(grpcServer *grpc.Server) {
	mpb.RegisterTimerServiceServer(grpcServer, svc)
}

func (svc *TimerService) Serve(ctx context.Context) error {
	go svc.TimerLoop(ctx)
	<-ctx.Done()
	return ctx.Err()
}

func (svc *TimerService) Logger() *zap.Logger {
	return svc.logger
}

func (svc *TimerService) ConnMgr() *gxgrpc.ConnManager {
	return svc.connMgr
}

func (svc *TimerService) Name() string {
	return svc.name
}

func (svc *TimerService) TimerLoop(ctx context.Context) {
	defer func() {
		err := recover()
		if err != nil {
			svc.logger.Error("TimerLoop panic", zap.Any("err", err))
			go svc.TimerLoop(ctx)
		}
	}()

	secondTicker := time.NewTicker(time.Second)
	minuteTicker := time.NewTicker(time.Minute)
	defer func() {
		secondTicker.Stop()
		minuteTicker.Stop()
	}()
	var (
		PVPDailySettleChan        = make(chan int, 1)
		secondLoop         uint64 = 0
		minuteLoop         uint64 = 0
	)
	for {
		select {
		case <-ctx.Done():
			return
		case <-secondTicker.C:
			secondLoop++
		case <-minuteTicker.C:
			select {
			case PVPDailySettleChan <- 1:
				go svc.handlePVPDailySettle(PVPDailySettleChan)
			default:
			}
			minuteLoop++
		default:
		}
	}
}

func (svc *TimerService) handlePVPDailySettle(ch chan int) {
	defer func() {
		<-ch
	}()
	fmt.Println("try handlePVPDailySettle")
	season, date, ok := svc.rm.canPVPSettle(time.Now())
	if !ok {
		return
	}
	fmt.Println("handlePVPDailySettle settle time", season, date)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	time.Sleep(10 * time.Second) // make sure pvp is end and all pvp request is handled

	if !svc.dao.checkPVPDailyNotSettled(ctx, date) {
		fmt.Println("handlePVPDailySettle settled", season, date)
		return
	}

	fmt.Println("handlePVPDailySettle start settle", season, date)

	rankRewardsRscs := svc.rm.getPVPSeasonRankRewardsRscs(season)
	if len(rankRewardsRscs) == 0 {
		return
	}

	err := svc.dao.handlePVPDailySettle(ctx, season, date, rankRewardsRscs)
	if err != nil {
		svc.logger.Error("handlePVPDailySettle failed", zap.Error(err))
	}
}
