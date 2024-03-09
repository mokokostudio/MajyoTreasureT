package gameservice

import (
	"github.com/oldjon/gutil"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"go.uber.org/zap"
	"math/rand"
)

const (
	FightMaxRound = 100000
)

type gameLevelMgr struct {
	svc    *GameService
	logger *zap.Logger
}

func newGameLevelMgr(svc *GameService) *gameLevelMgr {
	return &gameLevelMgr{
		svc:    svc,
		logger: svc.logger,
	}
}

func (gm *gameLevelMgr) newGameLevel(roleA, roleB roleI) *gameLevel {
	return &gameLevel{
		mgr:   gm,
		roleA: roleA,
		roleB: roleB,
	}
}

type gameLevel struct {
	mgr   *gameLevelMgr
	roleA roleI
	roleB roleI
}

type roleI interface {
	RoleType() mpb.ERole_Type
	BuffCards() []buffCardI
	HP() uint64
	SetHP(uint64)
	ATK() uint64
	BaseATK() uint64
	ATKGap() int64
	CriRate() uint64
	SetCriRate(uint64)
	CriDmgRate() uint64
	SetCriDmgRate(uint64)
	//CriDmgAddBuffs() []uint32
	//SetCriDmgAddBuff(uint32)
	HitRate() uint64
	DodgeRate() uint64
	SetDodgeRate(uint64)
	//DodgeAtkBuffs() []buffCardI
	//SetDodgeAtkBuff(buffCardI)
	DmgAddRate() uint64
	DmgReduceRate() uint64
	ATKBuffRate() uint64
	SetATKBuffRate(uint64)
	DefenseBuffRate() uint64
	SetDefenseBuffRate(uint64)
	//CriAwardsBuffs() []buffCardI
	//SetCriAwardsBuff(buff buffCardI)
	//DodgeStealBuffs() []buffCardI
	//SetDodgeStealBuff(buff buffCardI)
}

func (gl *gameLevel) atk(ms int64, from, to roleI) *mpb.FightDetail {
	hp := to.HP()
	if hp == 0 {
		return nil
	}
	detail := &mpb.FightDetail{
		AttackerType: from.RoleType(),
		BeAttacker:   to.RoleType(),
		HpBefore:     hp,
		HpAfter:      hp,
		AttackTime:   ms,
	}
	hitRate := gutil.Bound(
		gutil.If(from.HitRate() > to.DodgeRate(), from.HitRate()-to.DodgeRate(), 0),
		30, com.RateBase)
	if hitRate < uint64(rand.Int31n(com.RateBase)+1) { // miss
		detail.IsMiss = true
		return detail
	}
	var dmg uint64
	dmg, detail.IsCri = gl.calcDmg(from, to, true)

	dmg = gutil.Min(dmg, hp)
	hp = hp - dmg
	to.SetHP(hp)
	detail.HpAfter = hp
	detail.Dmg = dmg
	return detail
}

func (gl *gameLevel) atkDmgRate(ms int64, from, to roleI, dmgRate uint64) *mpb.FightDetail {
	hp := to.HP()
	if hp == 0 {
		return nil
	}
	detail := &mpb.FightDetail{
		AttackerType: from.RoleType(),
		BeAttacker:   to.RoleType(),
		HpBefore:     hp,
		HpAfter:      hp,
		AttackTime:   ms,
	}
	var dmg uint64
	dmg, detail.IsCri = gl.calcDmg(from, to, false)
	dmg = dmg * dmgRate / com.RateBase

	dmg = gutil.Min(dmg, hp)
	hp = hp - dmg
	to.SetHP(hp)
	detail.HpAfter = hp
	detail.Dmg = dmg
	return detail
}

func (gl *gameLevel) calcDmg(from, to roleI, calcCri bool) (uint64, bool) {
	dmg := from.ATK() *
		gutil.If(com.RateBase+from.DmgAddRate() > to.DmgReduceRate(),
			com.RateBase+from.DmgAddRate()-to.DmgReduceRate(), 0) *
		gutil.If(com.RateBase+from.ATKBuffRate() > to.DefenseBuffRate(),
			com.RateBase+from.ATKBuffRate()-to.DefenseBuffRate(), 0) /
		(com.RateBase * com.RateBase)
	dmg = dmg * (90 + uint64(rand.Int63n(21))) / 100
	if !calcCri {
		return dmg, false
	}
	isCri := from.CriRate() >= uint64(rand.Int31n(com.RateBase)+1)
	if isCri {
		dmg = dmg * (com.RateBase + from.CriDmgRate()) / com.RateBase
	}
	return dmg, isCri
}

func (gl *gameLevel) fight(player *gamePlayer, boss *gameBoss) (win bool, details []*mpb.FightDetail,
	dmg uint64, dmgRate uint64, bossDie bool, awardsAdd uint64) {
	var ms int64
	var round int64
	details = make([]*mpb.FightDetail, 0, 20)
	hpBefore := boss.hp
	for player.hp > 0 && boss.hp > 0 && round < FightMaxRound {
		ms++
		if ms%player.atkGap == 0 {
			// do player atk
			detail := gl.atk(ms, player, boss)
			if detail != nil {
				details = append(details, detail)
				for _, v := range player.BuffCards() {
					awardsAdd += v.effectWhenAtk(detail)
				}
			}
			round++
		}
		if boss.hp == 0 {
			boss.killedBy = player.userId
			break
		}
		if ms%boss.atkGap == 0 {
			// do boss atk
			detail := gl.atk(ms, boss, player)
			if detail != nil {
				details = append(details, detail)
				for _, v := range player.BuffCards() {
					tmpDetail, tmpAwardAdd := v.effectWhenBeenAtk(gl, true, detail)
					if tmpDetail != nil {
						details = append(details, tmpDetail)
					}
					awardsAdd += tmpAwardAdd
				}
			}
			round++
		}

		if boss.hp == 0 {
			boss.killedBy = player.userId
			break
		}
	}

	dmg = hpBefore - boss.hp
	dmgRate = dmg * com.RateBase / boss.totalHP
	win = boss.hp <= (boss.totalHP - boss.winLoseHp)
	return win, details, dmg, dmgRate, boss.hp == 0, awardsAdd
}

func (gl *gameLevel) fightPVP(player *gamePlayer, defender *gamePlayer) (win bool, details []*mpb.FightDetail) {
	var ms int64
	var round int64
	details = make([]*mpb.FightDetail, 0, 20)
	for player.hp > 0 && defender.hp > 0 && round < FightMaxRound {
		ms++
		if ms%player.atkGap == 0 {
			// do player atk
			detail := gl.atk(ms, player, defender)
			if detail != nil {
				details = append(details, detail)
			}
			round++
		}
		if defender.hp == 0 {
			break
		}
		if ms%defender.atkGap == 0 {
			// do boss atk
			detail := gl.atk(ms, defender, player)
			if detail != nil {
				details = append(details, detail)
			}
			round++
		}
	}

	win = defender.hp == 0
	return win, details
}
