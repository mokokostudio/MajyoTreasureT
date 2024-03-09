package common

import (
	"gitlab.com/morbackend/mor_services/mpberr"
	"strings"

	"github.com/oldjon/gutil"
	"gitlab.com/morbackend/mor_services/mpb"
)

func ReadAttrs(data map[string]string) *mpb.Attrs {
	if data == nil {
		return &mpb.Attrs{}
	}
	return &mpb.Attrs{
		Hp:              gutil.StrToUint64(data["basehp"]),
		Atk:             gutil.StrToUint64(data["baseatk"]),
		AtkGap:          gutil.StrToInt64(data["baseatkgap"]) * 1000,
		HpAddRate:       gutil.StrToUint64(data["hpaddrate"]),
		AtkAddRate:      gutil.StrToUint64(data["atkaddrate"]),
		AtkSpeedAddRate: gutil.StrToUint64(data["atkspeedaddrate"]),
		CriRate:         gutil.Bound(gutil.StrToUint64(data["crirate"]), 0, 10000),
		CriDmgRate:      gutil.StrToUint64(data["cridmgrate"]),
		HitRate:         gutil.Bound(gutil.StrToUint64(data["hitrate"]), 0, 10000),
		DodgeRate:       gutil.Bound(gutil.StrToUint64(data["dodgerate"]), 0, 10000),
		DmgAddRate:      gutil.StrToUint64(data["dmgaddrate"]),
		DmgReduceRate:   gutil.StrToUint64(data["dmgreducerate"]),
		AtkBuffRate:     gutil.StrToUint64(data["atkbuffrate"]),
		DefenseBuffRate: gutil.StrToUint64(data["defensebuffrate"]),
	}
}

func AddAttrs(a, b *mpb.Attrs) *mpb.Attrs {
	if a == nil {
		a = &mpb.Attrs{}
	}
	if b == nil {
		b = &mpb.Attrs{}
	}
	return &mpb.Attrs{
		Hp:              a.Hp + b.Hp,
		Atk:             a.Atk + b.Atk,
		AtkGap:          a.AtkGap + b.AtkGap,
		HpAddRate:       a.HpAddRate + b.HpAddRate,
		AtkAddRate:      a.AtkAddRate + b.AtkAddRate,
		AtkSpeedAddRate: a.AtkSpeedAddRate + b.AtkSpeedAddRate,
		CriRate:         a.CriRate + b.CriRate,
		CriDmgRate:      a.CriDmgRate + b.CriDmgRate,
		HitRate:         a.HitRate + b.HitRate,
		DodgeRate:       a.DodgeRate + b.DodgeRate,
		DmgAddRate:      a.DmgAddRate + b.DmgAddRate,
		DmgReduceRate:   a.DmgReduceRate + b.DmgReduceRate,
		AtkBuffRate:     a.AtkBuffRate + b.AtkBuffRate,
		DefenseBuffRate: a.DefenseBuffRate + b.DefenseBuffRate,
	}
}

func CloneAttrs(a *mpb.Attrs) *mpb.Attrs {
	if a == nil {
		return nil
	}
	return &mpb.Attrs{
		Hp:              a.Hp,
		Atk:             a.Atk,
		AtkGap:          a.AtkGap,
		HpAddRate:       a.HpAddRate,
		AtkAddRate:      a.AtkAddRate,
		AtkSpeedAddRate: a.AtkSpeedAddRate,
		CriRate:         a.CriRate,
		CriDmgRate:      a.CriDmgRate,
		HitRate:         a.HitRate,
		DodgeRate:       a.DodgeRate,
		DmgAddRate:      a.DmgAddRate,
		DmgReduceRate:   a.DmgReduceRate,
		AtkBuffRate:     a.AtkBuffRate,
		DefenseBuffRate: a.DefenseBuffRate,
	}
}

func ReadAwardsRsc(awardsStr string) ([]*mpb.AwardRsc, error) {
	if len(awardsStr) == 0 {
		return nil, nil
	}
	awards := make([]*mpb.AwardRsc, 0, 1)
	for _, awardStr := range strings.Split(awardsStr, ";") {
		strs := strings.Split(awardStr, ":")
		if len(strs) != 2 {
			return nil, mpberr.ErrConfig
		}
		award := &mpb.AwardRsc{
			ItemId: gutil.StrToUint32(strs[0]),
		}
		strs = strings.Split(strs[1], "|")
		if len(strs) == 1 {
			award.Num = gutil.StrToUint64(strs[0])
			if award.ItemId == uint32(mpb.EItem_ItemId_Mana) {
				award.Num *= ManaFactor
			}
		} else if len(strs) == 2 {
			award.NumRange = []uint64{gutil.StrToUint64(strs[0]), gutil.StrToUint64(strs[1])}
			if award.NumRange[0] > award.NumRange[1] {
				award.NumRange[0], award.NumRange[1] = award.NumRange[1], award.NumRange[0]
			}
			if award.ItemId == uint32(mpb.EItem_ItemId_Mana) {
				award.NumRange[0] *= ManaFactor
				award.NumRange[1] *= ManaFactor
			}
		} else {
			return nil, mpberr.ErrConfig
		}

		awards = append(awards, award)
	}
	return awards, nil
}

func ReadItems(itemsStr string) ([]*mpb.Item, error) {
	if len(itemsStr) == 0 {
		return nil, nil
	}
	items := make([]*mpb.Item, 0, 1)
	for _, itemStr := range strings.Split(itemsStr, ";") {
		strs := strings.Split(itemStr, ":")
		if len(strs) != 2 {
			return nil, mpberr.ErrConfig
		}
		item := &mpb.Item{
			ItemId: gutil.StrToUint32(strs[0]),
			Num:    gutil.StrToUint64(strs[1]),
		}
		if item.ItemId == uint32(mpb.EItem_ItemId_Mana) {
			item.Num *= ManaFactor
		}
		items = append(items, item)
	}
	return items, nil
}
