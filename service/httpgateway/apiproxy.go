package httpgateway

import (
	"net/http"

	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/util"
)

func (hg *HTTPGateway) generateShareHBossLink(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqGenerateShareHBossLink{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetAPIProxyGRPCClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.GenerateShareHBossLink(ctx, &mpb.ReqGenerateShareHBossLink{
		UserId:   claim.UserId,
		BossUuid: req.BossUuid,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResGenerateShareHBossLink{
		ShareLink: rpcRes.ShareLink,
	}
	return hg.writeHTTPRes(w, res)
}
