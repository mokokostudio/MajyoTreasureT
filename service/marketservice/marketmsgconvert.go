package marketservice

import "gitlab.com/morbackend/mor_services/mpb"

func (svc *MarketService) DBGoodsOrderToGuidOrder(dbOrder *mpb.DBGoodsOrder) *mpb.GoodsOrder {
	if dbOrder == nil {
		return nil
	}
	return &mpb.GoodsOrder{
		OrderUuid: dbOrder.OrderUuid,
		Goods: &mpb.Item{
			ItemId: dbOrder.GoodsId,
			Num:    dbOrder.GoodsNum,
		},
		Price:     dbOrder.Price,
		Gas:       dbOrder.Gas,
		PublishAt: dbOrder.PublishAt,
		OpenAt:    dbOrder.OpenAt,
		SoldAt:    dbOrder.SoldAt,
		Status:    mpb.EMarket_OrderStatus(dbOrder.Status),
	}
}
