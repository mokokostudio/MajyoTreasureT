package marketservice

import (
	"context"
	"time"

	"github.com/oldjon/gutil"
	"github.com/oldjon/gutil/env"
	"github.com/oldjon/gutil/gdb"
	gprotocol "github.com/oldjon/gutil/protocol"
	grmux "github.com/oldjon/gutil/redismutex"
	gxgrpc "github.com/oldjon/gx/modules/grpc"
	"github.com/oldjon/gx/service"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"gitlab.com/morbackend/mor_services/util"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type MarketService struct {
	mpb.UnimplementedMarketServiceServer
	name        string
	logger      *zap.Logger
	config      env.ModuleConfig
	etcdClient  *etcd.Client
	host        service.Host
	connMgr     *gxgrpc.ConnManager
	rm          *marketResourceMgr
	kvm         *service.KVMgr
	serverEnv   uint32
	sm          *util.ServiceMetrics
	dao         *marketDAO
	tcpMsgCoder gprotocol.FrameCoder
	orderUUIDSF service.Snowflake
	cm          *marketCacheMgr
}

// NewMarketService create a marketservice entity
func NewMarketService(driver service.ModuleDriver) (gxgrpc.GRPCServer, error) {
	svc := &MarketService{
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
	svc.rm, err = newMarketResourceMgr(svc.logger, svc.sm)
	if err != nil {
		return nil, err
	}

	redisMux, err := grmux.NewRedisMux(svc.config.SubConfig("redis_mutex"), nil, svc.logger, driver.Tracer())
	if err != nil {
		return nil, err
	}

	marketRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("market_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	tmpRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("tmp_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	svc.dao = newMarketDAO(svc.logger, redisMux, marketRedis, tmpRedis)
	svc.cm = newMarketCacheMgr(svc)

	svc.serverEnv = uint32(svc.config.GetInt64("server_env"))
	svc.tcpMsgCoder = gprotocol.NewFrameCoder(svc.config.GetString("protocol_code"))

	return svc, err
}

func (svc *MarketService) Register(grpcServer *grpc.Server) {
	mpb.RegisterMarketServiceServer(grpcServer, svc)
}

func (svc *MarketService) Serve(ctx context.Context) error {
	var err error
	svc.orderUUIDSF, err = svc.host.Snowflake(ctx, com.SnowflakeMarketOrderUUID, service.SnowflakeType_53)
	if err != nil {
		svc.logger.Error("failed to create snowflake", zap.Error(err))
		return err
	}
	<-ctx.Done()
	return ctx.Err()
}

func (svc *MarketService) Logger() *zap.Logger {
	return svc.logger
}

func (svc *MarketService) ConnMgr() *gxgrpc.ConnManager {
	return svc.connMgr
}

func (svc *MarketService) Name() string {
	return svc.name
}

func (svc *MarketService) GetGoodsOrdersOnSell(ctx context.Context, req *mpb.ReqGetGoodsOrdersOnSell) (*mpb.ResGetGoodsOrdersOnSell, error) {
	res := &mpb.ResGetGoodsOrdersOnSell{
		PageNum: req.PageNum,
		ItemId:  req.ItemId,
	}
	var err error
	if req.ItemId == 0 {
		res.Orders, res.OrderCnt, err = svc.cm.getPageOrdersCache(ctx, req.PageNum)
		if err != nil {
			return nil, err
		}
	} else {
		res.Orders, res.OrderCnt, err = svc.cm.getPageItemOrdersCache(ctx, req.ItemId, req.PageNum)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (svc *MarketService) GetMyGoodsOrdersOnSell(ctx context.Context, req *mpb.ReqUserId) (*mpb.ResGetMyGoodsOrdersOnSell, error) {
	orderUuids, err := svc.dao.getMyOnSellGoodsOrderUuids(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	res := &mpb.ResGetMyGoodsOrdersOnSell{}
	for _, v := range orderUuids {
		order, err := svc.cm.getOrderCache(ctx, v)
		if err != nil {
			return nil, err
		}
		res.Orders = append(res.Orders, order)
	}

	return res, nil
}

func (svc *MarketService) PublishGoodsOrder(ctx context.Context, req *mpb.ReqPublishGoodsOrder) (*mpb.Empty, error) {
	// check whether item can sell
	goodsRsc := svc.rm.getGoodsRsc(req.ItemId)
	if goodsRsc == nil {
		return nil, mpberr.ErrParam
	}

	if req.Num == 0 || req.Num > svc.rm.getMaxSellCnt() ||
		req.Price < svc.rm.getMinPrice() || req.Price > svc.rm.getMaxPrice() {
		return nil, mpberr.ErrParam
	}

	// check whether user can publish new order
	curOrderCnt, err := svc.dao.getMyOnSellGoodsOrderCnt(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if curOrderCnt >= svc.rm.getOnSellOrderCnt() {
		return nil, mpberr.ErrMaxOnSellOrders
	}

	nowUnx := time.Now().Unix()
	dbOrder := &mpb.DBGoodsOrder{
		GoodsId:   req.ItemId,
		GoodsNum:  req.Num,
		Price:     req.Price,
		Gas:       gutil.Max(1, req.Price*goodsRsc.GasRate/com.ManaFactor),
		PublishAt: nowUnx,
		OpenAt:    nowUnx + svc.rm.getOpenSellAfter(),
		Seller:    req.UserId,
	}

	dbOrder.OrderUuid, err = svc.orderUUIDSF.Next()
	if err != nil {
		svc.logger.Error("PublishGoodsOrder new order uuid failed", zap.Uint64("user_id", req.UserId),
			zap.Error(err))
		return nil, err
	}

	// sub item from backpack
	_, _, _, err = com.ExchangeItems(ctx, svc, req.UserId, nil,
		[]*mpb.Item{{ItemId: req.ItemId, Num: req.Num}}, 0,
		mpb.EItem_TransReason_PublishGoodsOrder, 0)
	if err != nil {
		return nil, err
	}

	err = svc.dao.publishOrder(ctx, req.UserId, dbOrder)
	if err != nil {
		return nil, err
	}

	return &mpb.Empty{}, nil
}

func (svc *MarketService) PurchaseGoodsOrder(ctx context.Context, req *mpb.ReqPurchaseGoodsOrder) (*mpb.ResPurchaseGoodsOrder, error) {
	order, err := svc.cm.getOrderCache(ctx, req.OrderUuid)
	if err != nil {
		return nil, err
	}
	if order.OrderUuid == 0 {
		return nil, mpberr.ErrOrderNotExist
	}
	nowUnix := time.Now().Unix()
	switch order.Status {
	case mpb.EMarket_OrderStatus_NEW:
		if order.OpenAt > nowUnix {
			return nil, mpberr.ErrOrderCantPurchaseNow
		}
	case mpb.EMarket_OrderStatus_OPEN_SELL:
	case mpb.EMarket_OrderStatus_SOLD_OUT:
		return nil, mpberr.ErrOrderSoldOut
	case mpb.EMarket_OrderStatus_TAKE_OFF:
		return nil, mpberr.ErrOrderTakeOff
	default:
		return nil, mpberr.ErrDB
	}

	// buyer check and sub mana
	mana, err := com.ExchangeMana(ctx, svc, req.UserId, -(int64(order.Price)),
		mpb.EItem_TransReason_PurchaseGoodsOrder, uint64(mpb.EItem_TransSubReason_PurchaseGoodsOrderCostMana))
	if err != nil {
		return nil, err
	}

	dbOrder, err := svc.dao.purchaseOrder(ctx, req.UserId, req.OrderUuid, nowUnix)
	if err != nil {
		// roll back mana
		_, _ = com.ExchangeMana(ctx, svc, req.UserId, int64(order.Price),
			mpb.EItem_TransReason_PurchaseGoodsOrder, uint64(mpb.EItem_TransSubReason_PurchaseGoodsOrderRollbackMana))
		return nil, err
	}

	// buyer add items
	_, _, _, err = com.ExchangeItems(ctx, svc, req.UserId,
		[]*mpb.Item{{ItemId: dbOrder.GoodsId, Num: dbOrder.GoodsNum}}, nil, 0,
		mpb.EItem_TransReason_PurchaseGoodsOrder, uint64(mpb.EItem_TransSubReason_PurchaseGoodsOrderAddGoods))
	if err != nil {
		svc.logger.Error("PurchaseGoodsOrder add goods to buyer failed", zap.Uint64("user_id", req.UserId),
			zap.Any("order", dbOrder), zap.Error(err))
		return nil, err
	}
	// seller add mana
	_, err = com.ExchangeMana(ctx, svc, dbOrder.Seller, int64(dbOrder.Price-dbOrder.Gas),
		mpb.EItem_TransReason_PurchaseGoodsOrder, uint64(mpb.EItem_TransSubReason_PurchaseGoodsOrderAddMana))
	if err != nil {
		svc.logger.Error("PurchaseGoodsOrder add mana to seller failed",
			zap.Uint64("seller_id", dbOrder.Seller), zap.Any("order", dbOrder), zap.Error(err))
		return nil, err
	}

	return &mpb.ResPurchaseGoodsOrder{
		Goods: &mpb.Item{
			ItemId: dbOrder.GoodsId,
			Num:    dbOrder.GoodsNum,
		},
		ManaCost: dbOrder.Price,
		ManaLeft: mana,
	}, nil
}

func (svc *MarketService) TakeOffGoodsOrder(ctx context.Context, req *mpb.ReqTakeOffGoodsOrder) (*mpb.ResTakeOffGoodsOrder, error) {
	dbOrder, err := svc.dao.takeOffGoodsOrder(ctx, req.UserId, req.OrderUuid)
	if err != nil {
		return nil, err
	}
	// rollback items
	addItems, _, updateItems, err := com.ExchangeItems(ctx, svc, req.UserId,
		[]*mpb.Item{{ItemId: dbOrder.GoodsId, Num: dbOrder.GoodsNum}}, nil, 0,
		mpb.EItem_TransReason_TakeOffGoodsOrder, 0)
	if err != nil {
		svc.logger.Error("TakeOffGoodsOrder rollback item in order failed", zap.Uint64("user_id", req.UserId),
			zap.Any("order", dbOrder), zap.Error(err))
		return nil, err
	}
	if len(addItems) != 1 || len(updateItems) != 1 {
		return nil, mpberr.ErrDB
	}
	return &mpb.ResTakeOffGoodsOrder{
		AddItem:    addItems[0],
		UpdateItem: updateItems[0],
	}, nil
}
