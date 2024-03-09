package gameservice

import "gitlab.com/morbackend/mor_services/mpb"

func (svc *GameService) dbHiddenBoss2HiddenBoss(dbBoss *mpb.DBHiddenBoss, withFightHis bool) (*mpb.HiddenBoss,
	[]*mpb.HiddenBossFightHistory) {
	if dbBoss == nil {
		return nil, nil
	}
	boss := &mpb.HiddenBoss{
		BossId:   dbBoss.BossId,
		BossUuid: dbBoss.BossUuid,
		Hp:       dbBoss.Hp,
		Finder:   dbBoss.Finder,
		ExpireAt: dbBoss.ExpiredAt,
	}
	if len(dbBoss.FightHistories) == 0 || !withFightHis {
		return boss, nil
	}
	histories := make([]*mpb.HiddenBossFightHistory, 0, len(dbBoss.FightHistories))
	for _, v := range dbBoss.FightHistories {
		histories = append(histories, &mpb.HiddenBossFightHistory{
			Nickname: v.Nickname,
			DmgRate:  v.DmgRate,
		})
	}
	return boss, histories
}

func (svc *GameService) dbBossDefeatHistory2BossDefeatHistory(dbHis *mpb.DBBossDefeatHistory) *mpb.BossDefeatHistory {
	res := &mpb.BossDefeatHistory{}
	for k, v := range dbHis.BossDefeatHistory {
		res.List = append(res.List, &mpb.BossDefeatHistoryNode{
			BossClass: k,
			BossId:    v.BossId,
		})
	}
	return res
}

func (svc *GameService) dbBuffCards2BuffCards(dbBC *mpb.DBBuffCardsValid) []*mpb.BuffCard {
	if dbBC == nil {
		return nil
	}
	ret := make([]*mpb.BuffCard, 0, len(dbBC.BuffCards))
	for _, v := range dbBC.BuffCards {
		ret = append(ret, &mpb.BuffCard{
			BuffCardId: v.BuffCardId,
			LeftRound:  v.LeftRound,
		})
	}
	return ret
}
