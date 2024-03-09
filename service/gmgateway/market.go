package gmgateway

import (
	"net/http"

	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/util"
)

func (gg *GMGateway) adminFreezeAccountTrade(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqAdminFreezeAccountTrade{}
	err = gg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetMarketServiceClient(ctx, gg)
	if err != nil {
		return err
	}
	rpcRes, err := client.AdminFreezeAccountTrade(ctx, &mpb.ReqUserId{
		UserId: claim.UserId,
	})
	if err != nil {
		return err
	}

	return gg.writeHTTPRes(w, rpcRes)
}
