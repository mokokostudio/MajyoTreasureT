package httpgateway

import (
	"fmt"
	"net/http"

	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/util"
)

func (hg *HTTPGateway) getGoodsOrdersOnSell(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	req := &mpb.CReqGetGoodsOrdersOnSell{}
	err := hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetMarketServiceClient(ctx, hg)
	if err != nil {
		return err
	}
	rpcRes, err := client.GetGoodsOrdersOnSell(ctx, &mpb.ReqGetGoodsOrdersOnSell{
		ItemId:  req.ItemId,
		PageNum: req.PageNum,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResGetGoodsOrdersOnSell{
		ItemId:   rpcRes.ItemId,
		PageNum:  rpcRes.PageNum,
		OrderCnt: rpcRes.OrderCnt,
		Orders:   rpcRes.Orders,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) getMyGoodsOrdersOnSell(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	client, err := com.GetMarketServiceClient(ctx, hg)
	if err != nil {
		return err
	}
	rpcRes, err := client.GetMyGoodsOrdersOnSell(ctx, &mpb.ReqUserId{
		UserId: claim.UserId,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResGetMyGoodsOrdersOnSell{
		Orders: rpcRes.Orders,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) publishGoodsOrder(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	fmt.Println(11)
	req := &mpb.CReqPublishGoodsOrder{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	fmt.Println(12)
	client, err := com.GetMarketServiceClient(ctx, hg)
	if err != nil {
		return err
	}
	fmt.Println(13)
	rpcRes, err := client.PublishGoodsOrder(ctx, &mpb.ReqPublishGoodsOrder{
		UserId: claim.UserId,
		ItemId: req.ItemId,
		Num:    req.Num,
		Price:  req.Price,
	})
	if err != nil {
		return err
	}
	fmt.Println(14)
	return hg.writeHTTPRes(w, rpcRes)
}

func (hg *HTTPGateway) purchaseGoodsOrder(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqPurchaseGoodsOrder{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetMarketServiceClient(ctx, hg)
	if err != nil {
		return err
	}
	rpcRes, err := client.PurchaseGoodsOrder(ctx, &mpb.ReqPurchaseGoodsOrder{
		UserId:    claim.UserId,
		OrderUuid: req.OrderUuid,
	})
	if err != nil {
		return err
	}
	res := &mpb.CResPurchaseGoodsOrder{
		Goods:    rpcRes.Goods,
		ManaCost: rpcRes.ManaCost,
		ManaLeft: rpcRes.ManaLeft,
	}

	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) takeOffGoodsOrder(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqTakeOffGoodsOrder{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetMarketServiceClient(ctx, hg)
	if err != nil {
		return err
	}
	rpcRes, err := client.TakeOffGoodsOrder(ctx, &mpb.ReqTakeOffGoodsOrder{
		UserId:    claim.UserId,
		OrderUuid: req.OrderUuid,
	})
	if err != nil {
		return err
	}
	res := &mpb.CResTakeOffGoodsOrder{
		AddItem:    rpcRes.AddItem,
		UpdateItem: rpcRes.UpdateItem,
	}
	return hg.writeHTTPRes(w, res)
}
