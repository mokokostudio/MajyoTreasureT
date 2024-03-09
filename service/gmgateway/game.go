package gmgateway

import (
	"net/http"

	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
)

func (gg *GMGateway) adminRecoverEnergy(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	req := &mpb.CReqAdminRecoverEnergy{}
	err := gg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, gg)
	if err != nil {
		return err
	}

	rpcRes, err := client.AdminRecoverEnergy(ctx, &mpb.ReqUserId{
		UserId: req.UserId,
	})
	if err != nil {
		return err
	}
	return gg.writeHTTPRes(w, rpcRes)
}
