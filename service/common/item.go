package common

import (
	"context"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
)

func GetItemType(itemId uint32) mpb.EItem_Type {
	switch mpb.EItem_ItemId(itemId) {
	case mpb.EItem_ItemId_BaseEquipLevelUpgradeMaterial10:
		return mpb.EItem_Type_BaseEquipLevelUpgradeMaterial
	case mpb.EItem_ItemId_BaseEquipStarUpgradeMaterial11,
		mpb.EItem_ItemId_BaseEquipStarUpgradeMaterial12,
		mpb.EItem_ItemId_BaseEquipStarUpgradeMaterial13,
		mpb.EItem_ItemId_BaseEquipStarUpgradeMaterial14,
		mpb.EItem_ItemId_BaseEquipStarUpgradeMaterial15:
		return mpb.EItem_Type_BaseEquipStarUpgradeMaterial
	}
	return mpb.EItem_Type(itemId / 1000000)
}

func AddItemsFromAwardRsc(ctx context.Context, svc GRPCClientGetter, userId uint64, awardRscs [][]*mpb.AwardRsc,
	reason mpb.EItem_TransReason, subReason uint64) (awards *mpb.CAwards, err error) {
	if len(awardRscs) == 0 {
		return nil, nil
	}
	addItems := make([]*mpb.Item, 0, len(awardRscs))
	for i, as := range awardRscs {
		for _, a := range as {
			addItems = append(addItems, &mpb.Item{
				ItemId:  a.ItemId,
				Num:     uint64(a.Num),
				BatchId: uint32(i),
			})
		}
	}
	addItems, _, updateItems, err := ExchangeItems(ctx, svc, userId, addItems, nil, 0, reason,
		subReason)
	if err != nil {
		return nil, err
	}
	return &mpb.CAwards{
		AddItems:    addItems,
		UpdateItems: updateItems,
	}, err
}

func ExchangeMana(ctx context.Context, svc GRPCClientGetter, userId uint64, deltaMana int64,
	reason mpb.EItem_TransReason, subReason uint64) (uint64, error) {
	_, _, updates, err := ExchangeItems(ctx, svc, userId, nil, nil, deltaMana, reason, subReason)
	if err != nil {
		return 0, err
	}
	if len(updates) != 1 {
		return 0, mpberr.ErrDB
	}
	return updates[0].GetNum(), nil
}

func ExchangeItems(ctx context.Context, svc GRPCClientGetter, userId uint64, addItems []*mpb.Item, delItems []*mpb.Item,
	deltaMana int64, reason mpb.EItem_TransReason, subReason uint64) (adds []*mpb.Item, dels []*mpb.Item,
	updates []*mpb.Item, err error) {
	client, err := GetItemServiceClient(ctx, svc)
	if err != nil {
		return nil, nil, nil, err
	}
	res, err := client.ExchangeItems(ctx, &mpb.ReqExchangeItems{
		UserId:         userId,
		AddItems:       addItems,
		DelItems:       delItems,
		DeltaMana:      deltaMana,
		TransReason:    reason,
		TransSubReason: subReason,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	return res.AddItems, res.DelItems, res.UpdateItems, nil
}

func BatchAddItemsFromAwardRsc(ctx context.Context, svc GRPCClientGetter, awardRscs map[uint64][][]*mpb.AwardRsc,
	reason mpb.EItem_TransReason, subReason uint64) (awards map[uint64]*mpb.CAwards, err error) {
	addItems := make(map[uint64]*mpb.Items)
	for uid, v := range awardRscs {
		for i, vv := range v {
			for _, vvv := range vv {
				items := addItems[uid]
				if items == nil {
					items = &mpb.Items{}
					addItems[uid] = items
				}
				items.Items = append(items.Items, &mpb.Item{
					ItemId:  vvv.ItemId,
					Num:     uint64(vvv.Num),
					BatchId: uint32(i),
				})
			}
		}
	}
	addItems, updateItems, err := BatchAddItems(ctx, svc, addItems, nil, reason, subReason)
	if err != nil {
		return nil, err
	}
	awards = make(map[uint64]*mpb.CAwards)
	for uid, v := range addItems {
		awards[uid] = &mpb.CAwards{
			AddItems:    v.GetItems(),                // avoid nil
			UpdateItems: updateItems[uid].GetItems(), // avoid nil
		}
	}
	return awards, err
}

func BatchAddItems(ctx context.Context, svc GRPCClientGetter, addItems map[uint64]*mpb.Items,
	addManas map[uint64]uint64, reason mpb.EItem_TransReason, subReason uint64) (map[uint64]*mpb.Items,
	map[uint64]*mpb.Items, error) {
	client, err := GetItemServiceClient(ctx, svc)
	if err != nil {
		return nil, nil, err
	}
	res, err := client.BatchAddItems(ctx, &mpb.ReqBatchAddItems{
		AddItems:       addItems,
		AddManas:       addManas,
		TransReason:    reason,
		TransSubReason: subReason,
	})
	if err != nil {
		return nil, nil, err
	}

	return res.AddItems, res.UpdateItems, nil
}
