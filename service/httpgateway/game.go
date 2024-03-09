package httpgateway

import (
	"net/http"

	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/util"
)

func (hg *HTTPGateway) fight(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqFight{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.Fight(ctx, &mpb.ReqFight{
		UserId:   claim.UserId,
		BossId:   req.BossId,
		BossUuid: req.BossUuid,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResFight{
		Win:             rpcRes.Win,
		BossDie:         rpcRes.BossDie,
		Details:         rpcRes.Details,
		Awards:          rpcRes.Awards,
		EnergyCost:      rpcRes.EnergyCost,
		Energy:          rpcRes.Energy,
		EnergyRecoverAt: rpcRes.EnergyRecoverAt,
		Dmg:             rpcRes.Dmg,
		DmgRate:         rpcRes.DmgRate,
		HiddenBoss:      rpcRes.HiddenBoss,
		BossHp:          rpcRes.BossHp,
		PlayerHp:        rpcRes.PlayerHp,
		BuffCards:       rpcRes.BuffCards,
		BuffCardStatus:  rpcRes.BuffCardStatus,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) getHiddenBoss(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqGetHiddenBoss{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.GetHiddenBoss(ctx, &mpb.ReqGetHiddenBoss{
		UserId:   claim.UserId,
		BossUuid: req.BossUuid,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResGetHiddenBoss{
		HiddenBoss: rpcRes.HiddenBoss,
		Fought:     rpcRes.Fought,
		FightCd:    rpcRes.FightCd,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) newHiddenBoss(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	req := &mpb.CReqNewHiddenBoss{}
	err := hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.NewHiddenBoss(ctx, &mpb.ReqNewHiddenBoss{UserId: req.Tguser, BossId: req.BossId})
	if err != nil {
		return err
	}
	res := &mpb.CResNewHiddenBoss{
		HiddenBoss: rpcRes.HiddenBoss,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) getEnergy(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.GetEnergy(ctx, &mpb.ReqUserId{
		UserId: claim.UserId,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResGetEnergy{
		Energy:   rpcRes.Energy,
		UpdateAt: rpcRes.UpdateAt,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) fightPVP(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqFightPVP{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.FightPVP(ctx, &mpb.ReqFightPVP{
		UserId:   claim.UserId,
		TargetId: req.TargetId,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResFightPVP{
		Win:                     rpcRes.Win,
		OldRank:                 rpcRes.OldRank,
		NewRank:                 rpcRes.NewRank,
		ChallengerHp:            rpcRes.ChallengerHp,
		DefenderHp:              rpcRes.DefenderHp,
		Details:                 rpcRes.Details,
		PvpChallengeCnt:         rpcRes.PvpChallengeCnt,
		PvpChallengeCntUpdateAt: rpcRes.PvpChallengeCntUpdateAt,
		Mana:                    rpcRes.Mana,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) getPVPInfo(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.GetPVPInfo(ctx, &mpb.ReqUserId{
		UserId: claim.UserId,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResGetPVPInfo{
		Rank:                  rpcRes.Rank,
		ChallengerCnt:         rpcRes.ChallengerCnt,
		ChallengerCntUpdateAt: rpcRes.ChallengerCntUpdateAt,
		PvpSettleRewards:      rpcRes.PvpSettleRewards,
		PvpSeasonDate:         rpcRes.PvpSeasonDate,
		PvpManaAwardsPool:     rpcRes.PvpManaAwardsPool,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) getPVPRanks(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	req := &mpb.CReqGetPVPRanks{}
	err := hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}

	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.GetPVPRanks(ctx, &mpb.ReqGetPVPRanks{
		PageNum: req.PageNum,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResGetPVPRanks{
		PageNum:  rpcRes.PageNum,
		RankList: rpcRes.RankList,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) getPVPChallengeTargets(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.GetPVPChallengeTargets(ctx, &mpb.ReqUserId{
		UserId: claim.UserId,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResGetPVPChallengeTargets{
		TargetList: rpcRes.TargetList,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) getPVPHistory(w http.ResponseWriter, r *http.Request) error {
	ctx := r.Context()
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.GetPVPHistory(ctx, &mpb.Empty{})
	if err != nil {
		return err
	}

	res := &mpb.CResGetPVPHistory{
		List: rpcRes.List,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) randomBuffCards(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqRandomBuffCards{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.RandomBuffCards(ctx, &mpb.ReqRandomBuffCards{
		UserId: claim.UserId,
		BossId: req.BossId,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResRandomBuffCards{
		BuffCards: rpcRes.BuffCards,
	}
	return hg.writeHTTPRes(w, res)
}

func (hg *HTTPGateway) choseBuffCard(w http.ResponseWriter, r *http.Request) error {
	claim, ctx, err := util.ClaimFromContext(r.Context())
	if err != nil {
		return err
	}
	req := &mpb.CReqChoseBuffCard{}
	err = hg.readHTTPReq(w, r, req)
	if err != nil {
		return err
	}
	client, err := com.GetGameServiceClient(ctx, hg)
	if err != nil {
		return err
	}

	rpcRes, err := client.ChoseBuffCard(ctx, &mpb.ReqChoseBuffCard{
		UserId:   claim.UserId,
		BuffCard: req.BuffCard,
	})
	if err != nil {
		return err
	}

	res := &mpb.CResChoseBuffCard{
		BuffCards: rpcRes.BuffCards,
	}
	return hg.writeHTTPRes(w, res)
}
