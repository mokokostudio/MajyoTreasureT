package marketservice

import (
	"context"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"

	"github.com/oldjon/gutil/gdb"
	grmux "github.com/oldjon/gutil/redismutex"
	"go.uber.org/zap"
)

const (
	OrderPageCnt = 20
)

type marketDAO struct {
	logger   *zap.Logger
	rMux     *grmux.RedisMutex
	marketDB *gdb.DB
	tmpDB    *gdb.DB
}

func newMarketDAO(logger *zap.Logger, rMux *grmux.RedisMutex, marketRedis gdb.RedisClient, tmpRedis gdb.RedisClient,
) *marketDAO {
	return &marketDAO{
		logger:   logger,
		rMux:     rMux,
		marketDB: gdb.NewDB(marketRedis, nil),
		tmpDB:    gdb.NewDB(tmpRedis, nil),
	}
}

func (dao *marketDAO) getOnSellOrderUuids(ctx context.Context, pageNum uint32) (orderUuids []uint64, totalNum uint32,
	err error) {
	key := com.AllOrderOnSellKey()
	orderUuids, err = gdb.ToUint64Slice(dao.marketDB.ZRange(ctx, key, int64(pageNum)*OrderPageCnt,
		int64(pageNum+1)*OrderPageCnt-1))
	if err != nil {
		dao.logger.Error("getOnSellOrderUuids ZRange failed", zap.String("key", key),
			zap.Uint32("page_num", pageNum), zap.Error(err))
		return nil, 0, mpberr.ErrDB
	}
	totalNum, err = gdb.ToUint32(dao.marketDB.ZCard(ctx, key))
	if err != nil {
		dao.logger.Error("getOnSellOrderUuids ZCard failed", zap.String("key", key), zap.Error(err))
		return nil, 0, mpberr.ErrDB
	}

	return orderUuids, totalNum, nil
}

func (dao *marketDAO) getOnSellOrderUuidsByGoodsId(ctx context.Context, goodsId uint32, pageNum uint32) (
	orderUuids []uint64, totalNum uint32, err error) {
	key := com.GoodsOrdersOnSellKey(goodsId)
	orderUuids, err = gdb.ToUint64Slice(dao.marketDB.ZRange(ctx, key, int64(pageNum)*OrderPageCnt,
		int64(pageNum+1)*OrderPageCnt-1))
	if err != nil {
		dao.logger.Error("getOnSellOrderUuidsByGoodsId ZRange failed", zap.String("key", key),
			zap.Uint32("page_num", pageNum), zap.Error(err))
		return nil, 0, mpberr.ErrDB
	}
	totalNum, err = gdb.ToUint32(dao.marketDB.ZCard(ctx, key))
	if err != nil {
		dao.logger.Error("getOnSellOrderUuidsByGoodsId ZCard failed", zap.String("key", key), zap.Error(err))
		return nil, 0, mpberr.ErrDB
	}
	return orderUuids, totalNum, nil
}

func (dao *marketDAO) getOrderInfo(ctx context.Context, uuid uint64) (*mpb.DBGoodsOrder, error) {
	dbOrder := &mpb.DBGoodsOrder{}
	key := com.GoodsOrderInfoKey(uuid)
	err := dao.marketDB.GetObject(ctx, key, dbOrder)
	if err != nil {
		dao.logger.Error("getOrderInfo GetObject failed", zap.Uint64("uuid", uuid), zap.Error(err))
		return nil, mpberr.ErrDB
	}
	return dbOrder, nil
}

func (dao *marketDAO) getOrderInfos(ctx context.Context, uuids []uint64) (orders []*mpb.DBGoodsOrder, err error) {
	orders = make([]*mpb.DBGoodsOrder, 0, len(uuids))
	if len(uuids) == 0 {
		return orders, nil
	}
	keys := make([]string, 0, len(uuids))
	for _, v := range uuids {
		keys = append(keys, com.GoodsOrderInfoKey(v))
		orders = append(orders, &mpb.DBGoodsOrder{})
	}
	err = dao.marketDB.GetObjects(ctx, keys, orders)
	if err != nil {
		dao.logger.Error("getOrderInfos GetObjects failed", zap.Any("uuids", uuids), zap.Error(err))
		return nil, mpberr.ErrDB
	}
	return orders, nil
}

func (dao *marketDAO) getMyOnSellGoodsOrderUuids(ctx context.Context, userId uint64) (orderUuids []uint64, err error) {
	key := com.MyOrdersOnSellKey(userId)
	orderUuids, err = gdb.ToUint64Slice(dao.marketDB.ZRange(ctx, key, 0, -1))
	if err != nil {
		dao.logger.Error("getMyOnSellGoodsOrderUuids ZRange failed", zap.String("key", key),
			zap.Uint64("user_id", userId), zap.Error(err))
		return nil, mpberr.ErrDB
	}
	return orderUuids, nil
}

func (dao *marketDAO) getMyOnSellGoodsOrderCnt(ctx context.Context, userId uint64) (cnt uint32, err error) {
	key := com.MyOrdersOnSellKey(userId)
	cnt, err = gdb.ToUint32(dao.marketDB.ZCard(ctx, key))
	if err != nil {
		dao.logger.Error("getMyOnSellGoodsOrderCnt ZCard failed", zap.String("key", key),
			zap.Uint64("user_id", userId), zap.Error(err))
		return 0, mpberr.ErrDB
	}
	return cnt, nil
}

func (dao *marketDAO) publishOrder(ctx context.Context, userId uint64, dbOrder *mpb.DBGoodsOrder) error {
	key := com.GoodsOrderInfoKey(dbOrder.OrderUuid)
	err := dao.marketDB.SetObject(ctx, key, dbOrder)
	if err != nil {
		dao.logger.Error("publishOrder SetObject failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return mpberr.ErrDB
	}

	key = com.MyOrdersOnSellKey(userId)
	_, err = dao.marketDB.ZAdd(ctx, key, dbOrder.PublishAt, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("publishOrder ZAdd failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return err
	}

	score := float64(dbOrder.Price) / float64(dbOrder.GoodsNum)
	key = com.GoodsOrdersOnSellKey(dbOrder.GoodsId)
	_, err = dao.marketDB.ZAdd(ctx, key, score, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("publishOrder ZAdd failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return err
	}

	key = com.AllOrderOnSellKey()
	_, err = dao.marketDB.ZAdd(ctx, key, score, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("publishOrder ZAdd failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return err
	}
	return nil
}

func (dao *marketDAO) purchaseOrder(ctx context.Context, userId uint64, orderUuid uint64, nowUnix int64) (
	*mpb.DBGoodsOrder, error) {

	key := com.GoodsOrderInfoKey(orderUuid)
	dbOrder := &mpb.DBGoodsOrder{}
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.marketDB.GetObject(ctx, key, dbOrder)
		if dao.marketDB.IsErrNil(err) {
			return mpberr.ErrOrderNotExist
		} else if err != nil {
			dao.logger.Error("purchaseOrder GetObject failed", zap.String("key", key),
				zap.Uint64("order_uuid", orderUuid), zap.Error(err))
			return mpberr.ErrDB
		}
		switch dbOrder.Status {
		case uint32(mpb.EMarket_OrderStatus_NEW):
			if dbOrder.OpenAt > nowUnix {
				return mpberr.ErrOrderCantPurchaseNow
			}
		case uint32(mpb.EMarket_OrderStatus_OPEN_SELL):
		case uint32(mpb.EMarket_OrderStatus_SOLD_OUT):
			return mpberr.ErrOrderSoldOut
		case uint32(mpb.EMarket_OrderStatus_TAKE_OFF):
			return mpberr.ErrOrderTakeOff
		default:
			return mpberr.ErrDB
		}
		dbOrder.Status = uint32(mpb.EMarket_OrderStatus_SOLD_OUT)
		dbOrder.SoldAt = nowUnix
		dbOrder.Buyer = userId
		err = dao.marketDB.SetObject(ctx, key, dbOrder)
		if err != nil {
			dao.logger.Error("purchaseOrder SetObject failed", zap.String("key", key),
				zap.Any("db_order", dbOrder), zap.Error(err))
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	key = com.MyOrdersOnSellKey(userId)
	_, err = dao.marketDB.ZRem(ctx, key, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("purchaseOrder ZRem failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	key = com.MyOrdersSoldKey(dbOrder.Seller)
	_, err = dao.marketDB.ZAdd(ctx, key, dbOrder.SoldAt, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("purchaseOrder ZAdd failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	_, err = dao.marketDB.ZRemRangeByRank(ctx, key, -101, -101)
	if err != nil {
		dao.logger.Error("purchaseOrder ZRemRangeByRank failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	key = com.MyOrderPurchasedKey(userId)
	_, err = dao.marketDB.ZAdd(ctx, key, dbOrder.SoldAt, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("purchaseOrder ZAdd failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	_, err = dao.marketDB.ZRemRangeByRank(ctx, key, -101, -101)
	if err != nil {
		dao.logger.Error("purchaseOrder ZRemRangeByRank failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	key = com.GoodsOrdersOnSellKey(dbOrder.GoodsId)
	_, err = dao.marketDB.ZRem(ctx, key, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("purchaseOrder ZRem failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	key = com.AllOrderOnSellKey()
	_, err = dao.marketDB.ZRem(ctx, key, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("purchaseOrder ZRem failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	return dbOrder, nil
}

func (dao *marketDAO) takeOffGoodsOrder(ctx context.Context, userId uint64, orderUuid uint64) (*mpb.DBGoodsOrder, error) {
	dbOrder := &mpb.DBGoodsOrder{}
	key := com.GoodsOrderInfoKey(orderUuid)
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.marketDB.GetObject(ctx, key, dbOrder)
		if dao.marketDB.IsErrNil(err) {
			return mpberr.ErrOrderNotExist
		}
		if err != nil {
			dao.logger.Error("takeOffGoodsOrder GetObject failed", zap.Uint64("user_id", userId),
				zap.Uint64("order_uuid", orderUuid), zap.String("key", key), zap.Error(err))
			return mpberr.ErrDB
		}
		if dbOrder.Seller != userId {
			return mpberr.ErrOrderNotExist
		}
		switch dbOrder.Status {
		case uint32(mpb.EMarket_OrderStatus_NEW), uint32(mpb.EMarket_OrderStatus_OPEN_SELL):
		case uint32(mpb.EMarket_OrderStatus_SOLD_OUT):
			return mpberr.ErrOrderSoldOut
		case uint32(mpb.EMarket_OrderStatus_TAKE_OFF):
			return mpberr.ErrOrderTakeOff
		default:
			return mpberr.ErrDB
		}

		dbOrder.Status = uint32(mpb.EMarket_OrderStatus_TAKE_OFF)
		err = dao.marketDB.SetObject(ctx, key, dbOrder)
		if err != nil {
			dao.logger.Error("takeOffGoodsOrder SetObject failed", zap.Uint64("user_id", userId),
				zap.Uint64("order_uuid", orderUuid), zap.String("key", key), zap.Error(err))
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	// remove from my orderInfos on sell
	key = com.MyOrdersOnSellKey(userId)
	_, err = dao.marketDB.ZRem(ctx, key, orderUuid)
	if err != nil {
		dao.logger.Error("takeOffGoodsOrder ZRem failed", zap.Uint64("user_id", userId),
			zap.String("key", key), zap.Uint64("order_uuid", orderUuid), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	key = com.GoodsOrdersOnSellKey(dbOrder.GoodsId)
	_, err = dao.marketDB.ZRem(ctx, key, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("takeOffGoodsOrder ZRem failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	key = com.AllOrderOnSellKey()
	_, err = dao.marketDB.ZRem(ctx, key, dbOrder.OrderUuid)
	if err != nil {
		dao.logger.Error("takeOffGoodsOrder ZRem failed", zap.String("key", key),
			zap.Any("db_order", dbOrder), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	return dbOrder, nil
}
