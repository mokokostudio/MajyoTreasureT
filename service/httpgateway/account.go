package httpgateway

import (
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"net/http"
)

func (hg *HTTPGateway) helloWorld(w http.ResponseWriter, r *http.Request) error {
	_, err := w.Write([]byte("Hello world"))
	return err
}

func (hg *HTTPGateway) loginTest(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	req := &mpb.CReqLoginTest{}
	err := hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetAccountServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.LoginTest(ctx, &mpb.ReqUserId{
		UserId: req.UserId,
	})
	if err != nil {
		return err
	}
	res := &mpb.CResLoginTest{
		Account:           rpcRes.Account,
		Token:             rpcRes.Token,
		Energy:            rpcRes.Energy,
		EnergyUpdateAt:    rpcRes.EnergyUpdateAt,
		BossDefeatHistory: rpcRes.BossDefeatHistory,
		BuffCards:         rpcRes.BuffCards,
		BuffCardStatus:    rpcRes.BuffCardStatus,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) loginByToken(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	req := &mpb.CReqLoginByToken{}
	err := hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetAccountServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.LoginByToken(ctx, &mpb.ReqLoginByToken{
		Token: req.Token,
	})
	if err != nil {
		return err
	}
	res := &mpb.CResLoginByToken{
		Account:           rpcRes.Account,
		Token:             rpcRes.Token,
		Energy:            rpcRes.Energy,
		EnergyUpdateAt:    rpcRes.EnergyUpdateAt,
		BossDefeatHistory: rpcRes.BossDefeatHistory,
		BuffCards:         rpcRes.BuffCards,
		BuffCardStatus:    rpcRes.BuffCardStatus,
	}
	return hg.writeHTTPRes(w, res)
}
