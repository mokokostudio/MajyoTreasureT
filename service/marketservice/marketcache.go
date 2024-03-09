package marketservice

import (
	"context"
	"sync"
	"time"

	"gitlab.com/morbackend/mor_services/mpb"
)

const (
	cacheTime3Secs = 3
)

type marketCacheMgr struct {
	svc                   *MarketService
	dao                   *marketDAO
	pageOrderMux          sync.RWMutex
	pageOrderUpdateAt     map[uint32]int64
	pageOrders            map[uint32][]*mpb.GoodsOrder
	orderTotalNum         uint32
	pageItemOrderMux      sync.RWMutex
	pageItemOrderUpdateAt map[uint32]map[uint32]int64             // item_id,page_num
	pageItemOrders        map[uint32]map[uint32][]*mpb.GoodsOrder // item_id,page_num
	itemOrderTotalNum     map[uint32]uint32                       // item_id
	orderInfoMux          sync.RWMutex
	orderInfoUpdateAt     map[uint64]int64
	orderInfos            map[uint64]*mpb.GoodsOrder
}

func newMarketCacheMgr(svc *MarketService) *marketCacheMgr {
	return &marketCacheMgr{
		svc:                   svc,
		dao:                   svc.dao,
		pageOrderUpdateAt:     make(map[uint32]int64),
		pageOrders:            make(map[uint32][]*mpb.GoodsOrder),
		pageItemOrderUpdateAt: make(map[uint32]map[uint32]int64),
		pageItemOrders:        make(map[uint32]map[uint32][]*mpb.GoodsOrder),
		itemOrderTotalNum:     make(map[uint32]uint32),
		orderInfoUpdateAt:     make(map[uint64]int64),
		orderInfos:            make(map[uint64]*mpb.GoodsOrder),
	}
}

func (cm *marketCacheMgr) getPageOrdersCache(ctx context.Context, pageNum uint32) ([]*mpb.GoodsOrder, uint32, error) {
	nowUnix := time.Now().Unix()
	orders := make([]*mpb.GoodsOrder, 0, OrderPageCnt)
	defer func() {
		for _, v := range orders {
			if v.Status == mpb.EMarket_OrderStatus_NEW && v.OpenAt <= nowUnix {
				v.Status = mpb.EMarket_OrderStatus_OPEN_SELL
			}
		}
	}()

	var ok bool
	cm.pageOrderMux.RLock()
	if cm.pageOrderUpdateAt[pageNum] > nowUnix-cacheTime3Secs {
		orders, ok = cm.pageOrders[pageNum]
		if ok {
			cm.pageOrderMux.RUnlock()
			return orders, cm.orderTotalNum, nil
		}
	}
	cm.pageOrderMux.RUnlock()

	cm.pageOrderMux.Lock()

	if cm.pageOrderUpdateAt[pageNum] > nowUnix-cacheTime3Secs {
		orders, ok = cm.pageOrders[pageNum]
		if ok { // double check, if cache set by other goroutines
			cm.pageOrderMux.Unlock()
			return orders, cm.orderTotalNum, nil
		}
	}

	orderUuids, totalNum, err := cm.dao.getOnSellOrderUuids(ctx, pageNum)
	if err != nil {
		cm.pageOrderMux.Unlock()
		return nil, 0, err
	}
	cm.orderTotalNum = totalNum

	dbOrders, err := cm.dao.getOrderInfos(ctx, orderUuids)
	if err != nil {
		cm.pageOrderMux.Unlock()
		return nil, 0, err
	}

	for _, v := range dbOrders {
		orders = append(orders, cm.svc.DBGoodsOrderToGuidOrder(v))
	}
	cm.pageOrders[pageNum] = orders
	cm.pageOrderUpdateAt[pageNum] = nowUnix

	// remove some expired cache
	maxRemoveCnt := 10
	needRemovePages := make([]uint32, 0, 10)
	for p, at := range cm.pageOrderUpdateAt {
		if at >= nowUnix-cacheTime3Secs {
			continue
		}
		needRemovePages = append(needRemovePages, p)
		maxRemoveCnt--
		if maxRemoveCnt == 0 {
			break
		}
	}
	for _, v := range needRemovePages {
		delete(cm.pageOrderUpdateAt, v)
		delete(cm.pageOrders, v)
	}

	cm.pageOrderMux.Unlock()

	// update orders info cache by the way
	cm.updateOrdersCache(orders)

	return orders, totalNum, nil
}

func (cm *marketCacheMgr) getPageItemOrdersCache(ctx context.Context, itemId, pageNum uint32) ([]*mpb.GoodsOrder,
	uint32, error) {
	nowUnix := time.Now().Unix()
	orders := make([]*mpb.GoodsOrder, 0, OrderPageCnt)
	defer func() {
		for _, v := range orders {
			if v.Status == mpb.EMarket_OrderStatus_NEW && v.OpenAt <= nowUnix {
				v.Status = mpb.EMarket_OrderStatus_OPEN_SELL
			}
		}
	}()

	var ok bool
	cm.pageItemOrderMux.RLock()
	if cm.pageItemOrderUpdateAt[itemId][pageNum] > nowUnix-cacheTime3Secs {
		orders, ok = cm.pageItemOrders[itemId][pageNum]
		if ok {
			cm.pageItemOrderMux.RUnlock()
			return orders, cm.itemOrderTotalNum[itemId], nil
		}
	}
	cm.pageItemOrderMux.RUnlock()

	cm.pageItemOrderMux.Lock()
	if cm.pageItemOrderUpdateAt[itemId][pageNum] > nowUnix-cacheTime3Secs {
		orders, ok = cm.pageItemOrders[itemId][pageNum]
		if ok { // double check, if cache set by other goroutines
			cm.pageItemOrderMux.Unlock()
			return orders, cm.itemOrderTotalNum[itemId], nil
		}
	}

	orderUuids, totalNum, err := cm.dao.getOnSellOrderUuidsByGoodsId(ctx, itemId, pageNum)
	if err != nil {
		cm.pageItemOrderMux.Unlock()
		return nil, 0, err
	}
	cm.itemOrderTotalNum[itemId] = totalNum

	dbOrders, err := cm.dao.getOrderInfos(ctx, orderUuids)
	if err != nil {
		cm.pageItemOrderMux.Unlock()
		return nil, 0, err
	}

	for _, v := range dbOrders {
		orders = append(orders, cm.svc.DBGoodsOrderToGuidOrder(v))
	}
	sm := cm.pageItemOrders[itemId]
	if sm == nil {
		sm = make(map[uint32][]*mpb.GoodsOrder, 0)
		cm.pageItemOrders[itemId] = sm
	}
	sm[pageNum] = orders

	usm := cm.pageItemOrderUpdateAt[itemId]
	if usm == nil {
		usm = make(map[uint32]int64)
		cm.pageItemOrderUpdateAt[itemId] = usm
	}
	usm[pageNum] = nowUnix

	// remove some expired cache
	maxRemoveCnt := 10
	needRemovePages := make([]uint32, 0, 10)
	for p, at := range usm {
		if at >= nowUnix-cacheTime3Secs {
			continue
		}
		needRemovePages = append(needRemovePages, p)
		maxRemoveCnt--
		if maxRemoveCnt == 0 {
			break
		}
	}
	for _, v := range needRemovePages {
		delete(sm, v)
		delete(usm, v)
	}

	cm.pageItemOrderMux.Unlock()

	// update orders info cache by the way
	cm.updateOrdersCache(orders)

	return orders, totalNum, nil
}

func (cm *marketCacheMgr) getOrderCache(ctx context.Context, orderUuid uint64) (*mpb.GoodsOrder, error) {
	var order *mpb.GoodsOrder
	nowUnix := time.Now().Unix()
	defer func() {
		if order.Status == mpb.EMarket_OrderStatus_NEW && order.OpenAt <= nowUnix {
			order.Status = mpb.EMarket_OrderStatus_OPEN_SELL
		}
	}()
	cm.orderInfoMux.RLock()
	var ok bool
	if cm.orderInfoUpdateAt[orderUuid] >= nowUnix-cacheTime3Secs {
		order, ok = cm.orderInfos[orderUuid]
		if ok {
			cm.orderInfoMux.RUnlock()
			return order, nil
		}
	}
	cm.orderInfoMux.RUnlock()

	cm.orderInfoMux.Lock()
	defer cm.orderInfoMux.Unlock()
	order, ok = cm.orderInfos[orderUuid]
	if ok {
		return order, nil
	}

	dbOrder, err := cm.dao.getOrderInfo(ctx, orderUuid)
	if err != nil {
		return nil, err
	}

	order = cm.svc.DBGoodsOrderToGuidOrder(dbOrder)
	cm.orderInfos[orderUuid] = order
	cm.orderInfoUpdateAt[orderUuid] = nowUnix

	// remove some expired cache
	maxRemoveCnt := 100
	needRemoveUuids := make([]uint64, 0, 100)
	for uuid, at := range cm.orderInfoUpdateAt {
		if at >= nowUnix-cacheTime3Secs {
			continue
		}
		needRemoveUuids = append(needRemoveUuids, uuid)
		maxRemoveCnt--
		if maxRemoveCnt == 0 {
			break
		}
	}
	for _, v := range needRemoveUuids {
		delete(cm.orderInfoUpdateAt, v)
		delete(cm.orderInfos, v)
	}

	return order, nil
}

func (cm *marketCacheMgr) updateOrdersCache(orders []*mpb.GoodsOrder) {
	nowUnix := time.Now().Unix()
	cm.orderInfoMux.Lock()
	defer cm.orderInfoMux.Unlock()

	// remove some expired cache
	maxRemoveCnt := 100
	needRemoveUuids := make([]uint64, 0, 100)
	for uuid, at := range cm.orderInfoUpdateAt {
		if at >= nowUnix-cacheTime3Secs {
			continue
		}
		needRemoveUuids = append(needRemoveUuids, uuid)
		maxRemoveCnt--
		if maxRemoveCnt == 0 {
			break
		}
	}
	for _, v := range needRemoveUuids {
		delete(cm.orderInfoUpdateAt, v)
		delete(cm.orderInfos, v)
	}

	for _, v := range orders {
		cm.orderInfoUpdateAt[v.OrderUuid] = nowUnix
		cm.orderInfos[v.OrderUuid] = v
	}
}

func (cm *marketCacheMgr) delOrderCache(orderUuid uint64) {
	cm.orderInfoMux.Lock()
	defer cm.orderInfoMux.Unlock()
	delete(cm.orderInfoUpdateAt, orderUuid)
	delete(cm.orderInfos, orderUuid)
}
