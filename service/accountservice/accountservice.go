package accountservice

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/oldjon/gutil"
	"github.com/oldjon/gutil/env"
	"github.com/oldjon/gutil/gdb"
	gprotocol "github.com/oldjon/gutil/protocol"
	grmux "github.com/oldjon/gutil/redismutex"
	gxgrpc "github.com/oldjon/gx/modules/grpc"
	"github.com/oldjon/gx/service"
	"github.com/pkg/errors"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"gitlab.com/morbackend/mor_services/util"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AccountService struct {
	mpb.UnimplementedAccountServiceServer
	name            string
	logger          *zap.Logger
	config          env.ModuleConfig
	etcdClient      *etcd.Client
	host            service.Host
	connMgr         *gxgrpc.ConnManager
	signingMethod   jwt.SigningMethod
	signingDuration time.Duration
	signingKey      []byte
	rm              *accountResourceMgr
	kvm             *service.KVMgr
	serverEnv       uint32
	sm              *util.ServiceMetrics
	dao             *accountDAO
	tcpMsgCoder     gprotocol.FrameCoder
}

// NewAccountService create an accountservice entity
func NewAccountService(driver service.ModuleDriver) (gxgrpc.GRPCServer, error) {
	svc := &AccountService{
		name:            driver.ModuleName(),
		logger:          driver.Logger(),
		config:          driver.ModuleConfig(),
		etcdClient:      driver.Host().EtcdSession().Client(),
		host:            driver.Host(),
		kvm:             driver.Host().KVManager(),
		sm:              util.NewServiceMetrics(driver),
		signingMethod:   jwt.SigningMethodHS256,
		signingDuration: 24 * 30 * time.Hour,
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
	svc.rm, err = newAccountResourceMgr(svc.logger, svc.sm)
	if err != nil {
		return nil, err
	}

	redisMux, err := grmux.NewRedisMux(svc.config.SubConfig("redis_mutex"), nil, svc.logger, driver.Tracer())
	if err != nil {
		return nil, err
	}

	accRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("acc_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	tmpRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("tmp_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	svc.dao = newAccountDAO(svc.logger, redisMux, accRedis, tmpRedis)

	svc.serverEnv = uint32(svc.config.GetInt64("server_env"))
	svc.tcpMsgCoder = gprotocol.NewFrameCoder(svc.config.GetString("protocol_code"))

	return svc, err
}

func (svc *AccountService) Register(grpcServer *grpc.Server) {
	mpb.RegisterAccountServiceServer(grpcServer, svc)
}

func (svc *AccountService) Serve(ctx context.Context) error {

	signingKey, err := svc.kvm.GetOrGenerate(ctx, com.JWTGatewayTokenKey, 32)
	if err != nil {
		return errors.WithStack(err)
	}
	svc.signingKey = signingKey

	<-ctx.Done()
	return ctx.Err()
}

func (svc *AccountService) Logger() *zap.Logger {
	return svc.logger
}

func (svc *AccountService) ConnMgr() *gxgrpc.ConnManager {
	return svc.connMgr
}

func (svc *AccountService) Name() string {
	return svc.name
}

func (svc *AccountService) generateLoginToken(userId uint64, account string,
	clientVersion, walletAddr string) (string, error) {
	var sToken string
	now := time.Now()
	claim := &mpb.JWTClaims{}
	claim.UserId = userId
	claim.Account = account
	claim.ClientVersion = clientVersion
	claim.RegisteredClaims = jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(now.Add(svc.signingDuration)),
	}
	claim.WalletAddr = walletAddr
	token := jwt.NewWithClaims(svc.signingMethod, claim)
	sToken, err := token.SignedString(svc.signingKey)
	if err != nil {
		return "", err
	}
	return sToken, nil
}

func (svc *AccountService) LoginTest(ctx context.Context, req *mpb.ReqUserId) (*mpb.ResLoginTest, error) {
	dbAcc, err := svc.dao.getAccountInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	token, err := svc.generateLoginToken(dbAcc.UserId, dbAcc.Account, "", dbAcc.WalletAddr)
	if err != nil {
		return nil, err
	}

	gameInfo, err := svc.getGameInfo(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	res := &mpb.ResLoginTest{
		Account:           svc.DBAccountInfo2AccountInfo(dbAcc),
		Token:             token,
		Energy:            gameInfo.Energy,
		EnergyUpdateAt:    gameInfo.UpdateAt,
		BossDefeatHistory: gameInfo.BossDefeatHistory,
		BuffCards:         gameInfo.BuffCards,
		BuffCardStatus:    gameInfo.BuffCardStatus,
	}
	return res, nil
}

func (svc *AccountService) getEnergy(ctx context.Context, userId uint64) (*mpb.ResGetEnergy, error) {
	client, err := com.GetGameServiceClient(ctx, svc)
	if err != nil {
		return nil, err
	}
	res, err := client.GetEnergy(ctx, &mpb.ReqUserId{
		UserId: userId,
	})
	return res, err
}

func (svc *AccountService) getGameInfo(ctx context.Context, userId uint64) (*mpb.ResGetGameInfo, error) {
	client, err := com.GetGameServiceClient(ctx, svc)
	if err != nil {
		return nil, err
	}
	res, err := client.GetGameInfo(ctx, &mpb.ReqUserId{
		UserId: userId,
	})
	return res, err
}

func (svc *AccountService) generateToken(tgId uint64) string {
	id := tgId + uint64(time.Now().UnixMicro()) + uint64(rand.Int63n(1000000))
	str := strconv.Itoa(int(id))
	return gutil.MD5(str)
}

func (svc *AccountService) GenerateLoginToken(ctx context.Context, req *mpb.ReqGenerateLoginToken) (*mpb.ResGenerateLoginToken, error) {
	if req.TgId == 0 {
		return nil, mpberr.ErrParam
	}
	token := svc.generateToken(req.TgId)
	err := svc.dao.saveLoginToken(ctx, token, req.TgId, req.FirstName, req.LastName, req.LanguageCode)
	if err != nil {
		return nil, err
	}
	return &mpb.ResGenerateLoginToken{Token: token}, err
}

func (svc *AccountService) LoginByToken(ctx context.Context, req *mpb.ReqLoginByToken) (*mpb.ResLoginByToken, error) {
	if req.Token == "" {
		return nil, mpberr.ErrParam
	}
	dbToken, err := svc.dao.getLoginToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	account := "TG" + strconv.Itoa(int(dbToken.TgId))
	dbAcc, err := svc.dao.getAccountInfoByAccount(ctx, account)
	if err != nil && !errors.Is(err, mpberr.ErrAccountNotExist) {
		return nil, err
	}
	if err != nil {
		dbAcc, err = svc.dao.initTGAccount(ctx, account, dbToken.TgId, dbToken.FirstName+" "+dbToken.LastName,
			dbToken.LanguageCode)
		if err != nil {
			return nil, err
		}
	}

	err = svc.dao.delLoginToken(ctx, req.Token)
	if err != nil {
		return nil, err
	}

	token, err := svc.generateLoginToken(dbAcc.UserId, dbAcc.Account, "", dbAcc.WalletAddr)
	if err != nil {
		return nil, err
	}

	gameInfo, err := svc.getGameInfo(ctx, dbAcc.UserId)
	if err != nil {
		return nil, err
	}
	res := &mpb.ResLoginByToken{
		Account:           svc.DBAccountInfo2AccountInfo(dbAcc),
		Token:             token,
		Energy:            gameInfo.Energy,
		EnergyUpdateAt:    gameInfo.UpdateAt,
		BossDefeatHistory: gameInfo.BossDefeatHistory,
		BuffCards:         gameInfo.BuffCards,
		BuffCardStatus:    gameInfo.BuffCardStatus,
	}
	return res, nil
}

func (svc *AccountService) GetAccountByTGUser(ctx context.Context, req *mpb.ReqGetAccountByTGUser) (*mpb.ResGetAccountByTGUser, error) {
	account := "TG" + strconv.Itoa(int(req.TgId))
	dbAcc, err := svc.dao.getAccountInfoByAccount(ctx, account)
	if err != nil && !errors.Is(err, mpberr.ErrAccountNotExist) {
		return nil, err
	}
	if err != nil {
		dbAcc, err = svc.dao.initTGAccount(ctx, account, req.TgId, req.FirstName+" "+req.LastName,
			req.LanguageCode)
		if err != nil {
			return nil, err
		}
	}

	if dbAcc.Nickname != req.FirstName+" "+req.LastName {
		dbAcc.Nickname = req.FirstName + " " + req.LastName
		err = svc.dao.updateNickname(ctx, dbAcc.UserId, dbAcc.Nickname)
	}

	energy, err := svc.getEnergy(ctx, dbAcc.UserId)
	if err != nil {
		return nil, err
	}
	res := &mpb.ResGetAccountByTGUser{
		Account:        svc.DBAccountInfo2AccountInfo(dbAcc),
		Energy:         energy.Energy,
		EnergyUpdateAt: energy.UpdateAt,
	}
	return res, nil
}

func (svc *AccountService) GetAccountByUserId(ctx context.Context, req *mpb.ReqUserId) (*mpb.ResGetAccountByUserId, error) {
	if util.IsBotUId(req.UserId) {
		return svc.getBotAccountByUserId(req.UserId)
	}
	dbAcc, err := svc.dao.getAccountInfo(ctx, req.UserId)
	if err != nil && !errors.Is(err, mpberr.ErrAccountNotExist) {
		return nil, err
	}

	res := &mpb.ResGetAccountByUserId{
		Account: svc.DBAccountInfo2AccountInfo(dbAcc),
	}
	return res, nil
}

func (svc *AccountService) getBotAccountByUserId(userId uint64) (*mpb.ResGetAccountByUserId, error) {
	botRsc := svc.rm.getBotRsc(userId)
	if botRsc == nil {
		return nil, mpberr.ErrConfig
	}
	res := &mpb.ResGetAccountByUserId{
		Account: &mpb.AccountInfo{
			UserId:   userId,
			Nickname: botRsc.Nickname,
			Icon:     botRsc.Icon,
		},
	}
	return res, nil
}

func (svc *AccountService) SetAccountTGLan(ctx context.Context, req *mpb.ReqSetAccountTGLan) (*mpb.Empty, error) {
	account := "TG" + strconv.Itoa(int(req.TgId))
	dbAcc, err := svc.dao.getAccountInfoByAccount(ctx, account)
	if err != nil {
		return nil, err
	}

	if dbAcc.LanguageCode == req.LanguageCode {
		return &mpb.Empty{}, nil
	}

	err = svc.dao.updateLanguage(ctx, dbAcc.UserId, req.LanguageCode)
	if err != nil {
		return nil, err
	}
	return &mpb.Empty{}, nil
}

func (svc *AccountService) BatchGetAccountsByUserIds(ctx context.Context, req *mpb.ReqUserIds) (*mpb.ResBatchGetAccountsByUserIds, error) {
	if len(req.UserIds) == 0 {
		return &mpb.ResBatchGetAccountsByUserIds{}, nil
	}

	res := &mpb.ResBatchGetAccountsByUserIds{
		Accounts: make(map[uint64]*mpb.AccountInfo),
	}
	realPlayerUserIds := make([]uint64, 0, len(req.UserIds))
	for _, v := range req.UserIds {
		if !util.IsBotUId(v) {
			realPlayerUserIds = append(realPlayerUserIds, v)
			continue
		}
		rr, err := svc.getBotAccountByUserId(v)
		if err != nil {
			svc.logger.Error("BatchGetAccountsByUserIds getBotAccountByUserId failed",
				zap.Uint64("bot_user_id", v), zap.Error(err))
			return nil, err
		}
		res.Accounts[v] = rr.Account
	}

	if len(realPlayerUserIds) == 0 {
		return res, nil
	}

	dbAccs, err := svc.dao.batchGetAccounts(ctx, realPlayerUserIds)
	if err != nil {
		return nil, err
	}

	for _, dbAcc := range dbAccs {
		res.Accounts[dbAcc.UserId] = svc.DBAccountInfo2AccountInfo(dbAcc)
	}

	return res, nil
}
