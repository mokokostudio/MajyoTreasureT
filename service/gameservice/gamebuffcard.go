package gameservice

import (
	"context"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"go.uber.org/zap"
)

type buffCardMgr struct {
	svc    *GameService
	logger *zap.Logger
	dao    *gameDAO
	rm     *gameResourceMgr
}

func newBuffCardMgr(svc *GameService) *buffCardMgr {
	return &buffCardMgr{
		svc:    svc,
		logger: svc.logger,
		dao:    svc.dao,
		rm:     svc.rm,
	}
}

type buffCardI interface {
	getBuffCardId() uint32
	getBuffCardType() mpb.EGame_BuffCardType
	effectBeforeGame(ctx context.Context) error
	effectWhenAtk(detail *mpb.FightDetail) uint64
	effectWhenBeenAtk(gl *gameLevel, roleABeenAtk bool, detail *mpb.FightDetail) (*mpb.FightDetail, uint64)
	effectBeforeFight(role roleI)
	results() interface{}
}

func (bcm *buffCardMgr) effectBeforeGame(ctx context.Context, userId uint64, buffCardRsc *mpb.BuffCardRsc) (
	newCards []uint32, err error) {
	bc := bcm.newBuffCard(userId, buffCardRsc)
	if bc == nil {
		return nil, mpberr.ErrConfig
	}
	err = bc.effectBeforeGame(ctx)
	if err != nil {
		return nil, err
	}
	r, ok := bc.results().([]uint32)
	if !ok {
		return nil, nil
	}
	return r, nil
}

func (bcm *buffCardMgr) newBuffCard(userId uint64, rsc *mpb.BuffCardRsc) buffCardI {
	switch rsc.CardType {
	case mpb.EGame_BuffCardType_1:
		bc := &buffCard1{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_2:
		bc := &buffCard2{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_3:
		bc := &buffCard3{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_4:
		bc := &buffCard4{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_5:
		bc := &buffCard5{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_6:
		bc := &buffCard6{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_7:
		bc := &buffCard7{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_8:
		bc := &buffCard8{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_9:
		bc := &buffCard9{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_10:
		bc := &buffCard10{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_11:
		bc := &buffCard11{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	case mpb.EGame_BuffCardType_12:
		bc := &buffCard12{}
		bc.mgr = bcm
		bc.userId = userId
		bc.rsc = rsc
		return bc
	}
	return nil
}

func (bcm *buffCardMgr) newBuffCardsByIds(userId uint64, buffCardIds []uint32) []buffCardI {
	buffCards := make([]buffCardI, 0, len(buffCardIds))
	for _, id := range buffCardIds {
		rsc := bcm.rm.getBuffCardRsc(id)
		if rsc == nil {
			continue
		}
		buffCards = append(buffCards, bcm.newBuffCard(userId, rsc))
	}
	return buffCards
}

func (bcm *buffCardMgr) effectBeforeFight(role roleI) {
	for _, v := range role.BuffCards() {
		v.effectBeforeFight(role)
	}
}

type buffCardBase struct {
	mgr    *buffCardMgr
	userId uint64
	rsc    *mpb.BuffCardRsc
}

func (bc *buffCardBase) getBuffCardId() uint32                    { return bc.rsc.GetCardId() }
func (bc *buffCardBase) getBuffCardType() mpb.EGame_BuffCardType  { return bc.rsc.GetCardType() }
func (bc *buffCardBase) effectBeforeGame(_ context.Context) error { return nil }
func (bc *buffCardBase) effectBeforeFight(_ roleI)                {}
func (bc *buffCardBase) effectWhenAtk(_ *mpb.FightDetail) uint64  { return 0 }
func (bc *buffCardBase) effectWhenBeenAtk(_ *gameLevel, _ bool, _ *mpb.FightDetail) (*mpb.FightDetail, uint64) {
	return nil, 0
}
func (bc *buffCardBase) results() interface{} { return nil }

// buff card type 1
type buffCard1 struct {
	buffCardBase
}

func (bc *buffCard1) effectBeforeFight(role roleI) {
	role.SetATKBuffRate(role.ATKBuffRate() + bc.rsc.GetAtcAdd())
}

func (bc *buffCard1) effectWhenAtk(detail *mpb.FightDetail) uint64 {
	if detail == nil {
		return 0
	}
	detail.AtkBuffAddBuffCards = append(detail.AtkBuffAddBuffCards, bc.getBuffCardId())
	return 0
}

// buff card type 2
type buffCard2 struct {
	buffCardBase
}

func (bc *buffCard2) effectBeforeFight(role roleI) {
	role.SetDefenseBuffRate(role.DefenseBuffRate() + bc.rsc.GetDefenseAdd())
}

func (bc *buffCard2) effectWhenBeenAtk(_ *gameLevel, _ bool, detail *mpb.FightDetail) (*mpb.FightDetail, uint64) {
	if detail == nil {
		return nil, 0
	}
	if detail.IsMiss {
		return nil, 0
	}
	detail.DefenceBuffAddBuffCards = append(detail.DefenceBuffAddBuffCards, bc.getBuffCardId())
	return nil, 0
}

// buff card type 3
type buffCard3 struct {
	buffCardBase
}

func (bc *buffCard3) effectBeforeFight(role roleI) {
	role.SetCriRate(role.CriRate() + bc.rsc.GetCriRateAdd())
}

func (bc *buffCard3) effectWhenAtk(detail *mpb.FightDetail) uint64 {
	if detail == nil {
		return 0
	}
	if !detail.IsCri {
		return 0
	}
	detail.CriRateAddBuffCards = append(detail.CriRateAddBuffCards, bc.getBuffCardId())
	return 0
}

// buff card type 4
type buffCard4 struct {
	buffCardBase
}

func (bc *buffCard4) effectBeforeFight(role roleI) {
	role.SetDodgeRate(role.DodgeRate() + bc.rsc.GetDodgeRateAdd())
}

func (bc *buffCard4) effectWhenBeenAtk(_ *gameLevel, _ bool, detail *mpb.FightDetail) (*mpb.FightDetail, uint64) {
	if detail == nil {
		return nil, 0
	}
	if !detail.IsMiss {
		return nil, 0
	}
	detail.DodgeRateAddBuffCards = append(detail.DodgeRateAddBuffCards, bc.getBuffCardId())
	return nil, 0
}

// buff card type 5
type buffCard5 struct {
	buffCardBase
}

func (bc *buffCard5) effectBeforeFight(role roleI) {
	role.SetCriDmgRate(role.CriDmgRate() + bc.rsc.GetCriDmgAdd())
}

func (bc *buffCard5) effectWhenAtk(detail *mpb.FightDetail) uint64 {
	if detail == nil {
		return 0
	}
	if !detail.IsCri {
		return 0
	}
	detail.CriDmgAddBuffCards = append(detail.CriDmgAddBuffCards, bc.getBuffCardId())
	return 0
}

// buff card type 6
type buffCard6 struct {
	buffCardBase
}

func (bc *buffCard6) effectWhenBeenAtk(gl *gameLevel, roleABeenAtk bool, detail *mpb.FightDetail) (*mpb.FightDetail,
	uint64) {
	if detail == nil {
		return nil, 0
	}
	if !detail.IsMiss {
		return nil, 0
	}
	detail.TriggerDodgeAtk = true
	roleA, roleB := gl.roleA, gl.roleB
	if !roleABeenAtk {
		roleA, roleB = roleB, roleA
	}
	tmpDetail := gl.atkDmgRate(detail.AttackTime, roleA, roleB, bc.rsc.GetDodgeAtk())
	if tmpDetail == nil {
		return nil, 0
	}
	tmpDetail.DodgeAtkBuffCard = bc.getBuffCardId()
	return tmpDetail, 0
}

// buff card type 7
type buffCard7 struct {
	buffCardBase
}

func (bc *buffCard7) effectBeforeGame(ctx context.Context) error {
	return bc.mgr.svc.dao.addValidBuffCardsRound(ctx, bc.userId)
}

// buff card type 8
type buffCard8 struct {
	buffCardBase
	value1 uint32
}

func (bc *buffCard8) effectWhenAtk(detail *mpb.FightDetail) uint64 {
	if detail == nil {
		return 0
	}
	if !detail.IsCri {
		return 0
	}
	if bc.value1 >= bc.rsc.GetTriggerCntPerRound() {
		return 0
	}
	bc.value1++
	detail.CriAwardsBuffCards = append(detail.CriAwardsBuffCards, bc.getBuffCardId())
	return bc.rsc.GetCriAwardsAdd()
}

// buff card type 9
type buffCard9 struct {
	buffCardBase
	value1 uint32
}

func (bc *buffCard9) effectWhenBeenAtk(gl *gameLevel, roleABeenAtk bool, detail *mpb.FightDetail) (*mpb.FightDetail,
	uint64) {
	if detail == nil {
		return nil, 0
	}
	if !detail.IsMiss {
		return nil, 0
	}
	roleA, roleB := gl.roleA, gl.roleB
	if !roleABeenAtk {
		roleA, roleB = roleB, roleA
	}
	if bc.value1 >= bc.rsc.GetTriggerCntPerRound() {
		return nil, 0
	}
	detail.TriggerDodgeSteal = true
	bc.value1++
	tmpDetail := gl.atkDmgRate(detail.AttackTime, roleA, roleB, bc.rsc.GetDodgeStealAtk())
	if tmpDetail == nil {
		return nil, 0
	}
	tmpDetail.DodgeStealBuffCard = bc.getBuffCardId()
	return tmpDetail, bc.rsc.GetDodgeAwardsAdd()
}

// buff card type 10
type buffCard10 struct {
	buffCardBase
	newBuffCardOpts []uint32
}

func (bc *buffCard10) effectBeforeGame(ctx context.Context) error {
	var err error
	bc.newBuffCardOpts, err = bc.mgr.svc.rerandomBuffCard(ctx, bc.userId, bc.rsc.BuffCardRandPool)
	return err
}
func (bc *buffCard10) results() interface{} { return bc.newBuffCardOpts }

// buff card type 11
type buffCard11 struct {
	buffCardBase
	newBuffCardOpts []uint32
}

func (bc *buffCard11) effectBeforeGame(ctx context.Context) error {
	var err error
	bc.newBuffCardOpts, err = bc.mgr.svc.rerandomBuffCard(ctx, bc.userId, bc.rsc.BuffCardRandPool)
	return err
}
func (bc *buffCard11) results() interface{} { return bc.newBuffCardOpts }

// buff card type 12
type buffCard12 struct {
	buffCardBase
	newBuffCardOpts []uint32
}

func (bc *buffCard12) effectBeforeGame(ctx context.Context) error {
	var err error
	bc.newBuffCardOpts, err = bc.mgr.svc.rerandomBuffCard(ctx, bc.userId, bc.rsc.BuffCardRandPool)
	if err != nil {
		return err
	}
	// upgrade owned buff cards level
	err = bc.mgr.svc.dao.upgradeValidBuffCardsLevel(ctx, bc.userId, bc.rsc.BuffLevelAdd)
	return err
}
func (bc *buffCard12) results() interface{} { return bc.newBuffCardOpts }

type roleBuffs struct {
	buffCards []buffCardI
}

func (r *roleBuffs) GetBuffCard(buffCardId uint32) buffCardI {
	for _, v := range r.buffCards {
		if v.getBuffCardId() == buffCardId {
			return v
		}
	}
	return nil
}

func (r *roleBuffs) BuffCards() []buffCardI { return r.buffCards }
