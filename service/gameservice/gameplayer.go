package gameservice

import (
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
)

type gamePlayer struct {
	roleTyp         mpb.ERole_Type
	userId          uint64
	baseEquips      []*mpb.BaseEquip
	nftEquips       []*mpb.NFTEquip
	attrs           *mpb.Attrs
	totalHP         uint64
	hp              uint64
	atk             uint64
	baseAtk         uint64
	atkGap          int64
	criRate         uint64
	criDmgRate      uint64
	hitRate         uint64
	dodgeRate       uint64
	dmgAddRate      uint64
	dmgReduceRate   uint64
	atkBuffRate     uint64
	defenseBuffRate uint64
	roleBuffs
}

func newPlayer(svc *GameService, userId uint64, role mpb.ERole_Type, buffCards []buffCardI) *gamePlayer {
	player := &gamePlayer{
		roleTyp: role,
		userId:  userId,
		attrs:   svc.rm.getPlayerInitAttrs(),
	}
	player.buffCards = buffCards
	return player
}

func (p *gamePlayer) updateEquips(baseEquips []*mpb.BaseEquip, nftEquips []*mpb.NFTEquip) {
	p.baseEquips = baseEquips
	p.nftEquips = nftEquips
	p.calcPlayerAttrs()
}

func (p *gamePlayer) calcPlayerAttrs() {
	for _, v := range p.baseEquips {
		p.attrs = com.AddAttrs(p.attrs, v.Attrs)
	}
	for _, v := range p.nftEquips {
		p.attrs = com.AddAttrs(p.attrs, v.Attrs)
	}
	p.hp = p.attrs.Hp
	p.hp = p.hp * (com.RateBase + p.attrs.HpAddRate) / com.RateBase
	p.totalHP = p.hp
	p.baseAtk = p.attrs.Atk
	p.atk = p.attrs.Atk
	p.atk = p.atk * (com.RateBase + p.attrs.AtkAddRate) / com.RateBase
	p.atkGap = p.attrs.AtkGap
	p.atkGap = p.atkGap * com.RateBase / int64(com.RateBase+p.attrs.AtkSpeedAddRate)
	p.criRate = p.attrs.CriRate
	p.criDmgRate = p.attrs.CriDmgRate
	p.hitRate = p.attrs.HitRate
	p.dodgeRate = p.attrs.DodgeRate
	p.dmgAddRate = p.attrs.DmgAddRate
	p.dmgReduceRate = p.attrs.DmgReduceRate
	p.atkBuffRate = p.attrs.AtkBuffRate
	p.defenseBuffRate = p.attrs.DefenseBuffRate
}

func (p *gamePlayer) RoleType() mpb.ERole_Type                  { return p.roleTyp }
func (p *gamePlayer) HP() uint64                                { return p.hp }
func (p *gamePlayer) SetHP(hp uint64)                           { p.hp = hp }
func (p *gamePlayer) ATK() uint64                               { return p.atk }
func (p *gamePlayer) BaseATK() uint64                           { return p.baseAtk }
func (p *gamePlayer) ATKGap() int64                             { return p.atkGap }
func (p *gamePlayer) CriRate() uint64                           { return p.criRate }
func (p *gamePlayer) SetCriRate(criRate uint64)                 { p.criRate = criRate }
func (p *gamePlayer) CriDmgRate() uint64                        { return p.criDmgRate }
func (p *gamePlayer) SetCriDmgRate(criDmgRate uint64)           { p.criDmgRate = criDmgRate }
func (p *gamePlayer) HitRate() uint64                           { return p.hitRate }
func (p *gamePlayer) DodgeRate() uint64                         { return p.dodgeRate }
func (p *gamePlayer) SetDodgeRate(dodgeRate uint64)             { p.dodgeRate = dodgeRate }
func (p *gamePlayer) DmgAddRate() uint64                        { return p.dmgAddRate }
func (p *gamePlayer) DmgReduceRate() uint64                     { return p.dmgReduceRate }
func (p *gamePlayer) ATKBuffRate() uint64                       { return p.atkBuffRate }
func (p *gamePlayer) SetATKBuffRate(atkBuffRate uint64)         { p.atkBuffRate = atkBuffRate }
func (p *gamePlayer) DefenseBuffRate() uint64                   { return p.defenseBuffRate }
func (p *gamePlayer) SetDefenseBuffRate(defenseBuffRate uint64) { p.defenseBuffRate = defenseBuffRate }
