package itemservice

import (
	"fmt"

	"github.com/oldjon/gutil"
	gcsv "github.com/oldjon/gutil/csv"
	gdm "github.com/oldjon/gutil/dirmonitor"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"gitlab.com/morbackend/mor_services/util"
	"go.uber.org/zap"
)

const (
	//csvSuffix   = ".csv"
	baseCSVPath = "./resources/item/"

	itemCSV            = "Items.csv"
	baseEquipLevelsCSV = "BaseEquipLevels.csv"
	baseEquipStarsCSV  = "BaseEquipStars.csv"
	botsCSV            = "Bots.csv"
)

type itemResourceMgr struct {
	logger *zap.Logger
	dm     *gdm.DirMonitor
	mtr    *util.ServiceMetrics

	items             map[uint32]*mpb.ItemRsc
	baseEquipStars    map[mpb.EItem_BaseEquipType]map[uint32]*mpb.BaseEquipStarRsc
	baseEquipMaxStar  map[mpb.EItem_BaseEquipType]uint32
	baseEquipLevels   map[mpb.EItem_BaseEquipType]map[uint32]*mpb.BaseEquipLevelRsc
	baseEquipMaxLevel map[mpb.EItem_BaseEquipType]uint32
	bots              map[uint64]*mpb.BotRsc
}

func newItemResourceMgr(logger *zap.Logger, mtr *util.ServiceMetrics) (*itemResourceMgr, error) {
	rMgr := &itemResourceMgr{
		logger: logger,
		mtr:    mtr,
	}

	var err error
	rMgr.dm, err = gdm.NewDirMonitor(baseCSVPath)
	if err != nil {
		return nil, err
	}

	err = rMgr.load()
	if err != nil {
		return nil, err
	}

	err = rMgr.watch()
	if err != nil {
		return nil, err
	}

	return rMgr, nil
}

func (rm *itemResourceMgr) watch() error {
	return rm.dm.StartWatch()
}

func (rm *itemResourceMgr) load() error {
	err := rm.dm.BindAndExec(itemCSV, rm.loadItems)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(baseEquipStarsCSV, rm.loadBaseEquipStars)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(baseEquipLevelsCSV, rm.loadBaseEquipLevels)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(botsCSV, rm.loadBots)
	if err != nil {
		fmt.Println(0)
		return err
	}
	return nil
}

func (rm *itemResourceMgr) loadItems(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32]*mpb.ItemRsc)
	for _, data := range datas {
		node := &mpb.ItemRsc{
			ItemId:        gcsv.StrToUint32(data["itemid"]),
			ItemType:      mpb.EItem_Type(gcsv.StrToUint32(data["itemtype"])),
			NameStringKey: data["namestringkey"],
			IsUnique:      gcsv.StrToBool(data["isunique"]),
			OriginId:      gcsv.StrToUint32(data["originid"]),
			ExpireTime:    gcsv.StrToInt64(data["expiretime"]),
		}

		m[node.ItemId] = node
		rm.logger.Debug("loadItems read:", zap.Any("row", node))
	}

	rm.logger.Debug("loadItems read finish:", zap.Any("rows", m))
	rm.items = m

	return nil
}

func (rm *itemResourceMgr) getItemRsc(itemId uint32) *mpb.ItemRsc {
	return rm.items[itemId]
}

func (rm *itemResourceMgr) loadBaseEquipStars(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[mpb.EItem_BaseEquipType]map[uint32]*mpb.BaseEquipStarRsc)
	maxM := make(map[mpb.EItem_BaseEquipType]uint32)
	for _, data := range datas {
		node := &mpb.BaseEquipStarRsc{
			EquipType:          mpb.EItem_BaseEquipType(gcsv.StrToUint32(data["equiptype"])),
			Star:               gcsv.StrToUint32(data["star"]),
			Attrs:              com.ReadAttrs(data),
			UpgradeSuccessRate: gcsv.StrToUint32(data["upgradesuccessrate"]),
			ProtectSuccessNum:  gcsv.StrToUint32(data["protectsuccessnum"]),
		}
		node.UpgradeConsumeItems, err = com.ReadItems(data["upgradeconsumeitems"])
		if err != nil {
			rm.logger.Error("loadBaseEquipStars parse upgradeconsumeitems failed",
				zap.String("upgradeconsumeitems", data["upgradeconsumeitems"]))
			return err
		}

		subm := m[node.EquipType]
		if subm == nil {
			subm = map[uint32]*mpb.BaseEquipStarRsc{}
			m[node.EquipType] = subm
		}
		subm[node.Star] = node
		maxM[node.EquipType] = gutil.Max(maxM[node.EquipType], node.Star)
		rm.logger.Debug("loadBaseEquipStars read:", zap.Any("row", node))
	}

	rm.logger.Debug("loadBaseEquipStars read finish:", zap.Any("rows", m))
	rm.baseEquipStars = m
	rm.baseEquipMaxStar = maxM
	return nil
}

func (rm *itemResourceMgr) getBaseEquipStarRsc(equipType mpb.EItem_BaseEquipType, star uint32) *mpb.BaseEquipStarRsc {
	return rm.baseEquipStars[equipType][star]
}

func (rm *itemResourceMgr) getBaseEquipMaxStar(equipType mpb.EItem_BaseEquipType) uint32 {
	return rm.baseEquipMaxStar[equipType]
}

func (rm *itemResourceMgr) loadBaseEquipLevels(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[mpb.EItem_BaseEquipType]map[uint32]*mpb.BaseEquipLevelRsc)
	maxM := make(map[mpb.EItem_BaseEquipType]uint32)
	for _, data := range datas {
		node := &mpb.BaseEquipLevelRsc{
			EquipType:          mpb.EItem_BaseEquipType(gcsv.StrToUint32(data["equiptype"])),
			Level:              gcsv.StrToUint32(data["level"]),
			Attrs:              com.ReadAttrs(data),
			UpgradeSuccessRate: gcsv.StrToUint32(data["upgradesuccessrate"]),
			ProtectSuccessNum:  gcsv.StrToUint32(data["protectsuccessnum"]),
		}

		node.UpgradeConsumeItems, err = com.ReadItems(data["upgradeconsumeitems"])
		if err != nil {
			rm.logger.Error("loadBaseEquipStars parse upgradeconsumeitems failed",
				zap.String("upgradeconsumeitems", data["upgradeconsumeitems"]))
			return err
		}

		subm := m[node.EquipType]
		if subm == nil {
			subm = map[uint32]*mpb.BaseEquipLevelRsc{}
			m[node.EquipType] = subm
		}
		subm[node.Level] = node
		maxM[node.EquipType] = gutil.Max(maxM[node.EquipType], node.Level)
		rm.logger.Debug("loadBaseEquipLevels read:", zap.Any("row", node))
	}

	rm.logger.Debug("loadBaseEquipLevels read finish:", zap.Any("rows", m))
	rm.baseEquipLevels = m
	rm.baseEquipMaxLevel = maxM
	return nil
}

func (rm *itemResourceMgr) getBaseEquipLevelRsc(equipType mpb.EItem_BaseEquipType, level uint32) *mpb.BaseEquipLevelRsc {
	return rm.baseEquipLevels[equipType][level]
}

func (rm *itemResourceMgr) getBaseEquipMaxLevel(equipType mpb.EItem_BaseEquipType) uint32 {
	return rm.baseEquipMaxLevel[equipType]
}

func (rm *itemResourceMgr) loadBots(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint64]*mpb.BotRsc)
	for _, data := range datas {
		node := &mpb.BotRsc{
			Id:       gcsv.StrToUint64(data["id"]),
			Nickname: data["nickname"],
			Icon:     data["icon"],
		}
		baseEquipsLevels := gcsv.StrToUint32Slice(data["baseequipslevels"], ";")
		if len(baseEquipsLevels) != baseEquipCount {
			rm.logger.Error("loadBots failed", zap.String("baseequipslevels", data["baseequipslevels"]))
			return mpberr.ErrConfig
		}
		baseEquipsStars := gcsv.StrToUint32Slice(data["baseequipsstars"], ";")
		if len(baseEquipsStars) != baseEquipCount {
			rm.logger.Error("loadBots failed", zap.String("baseequipsstars", data["baseequipsstars"]))
			return mpberr.ErrConfig
		}
		for i := range baseEquipsLevels {
			node.BaseEquips = append(node.BaseEquips, &mpb.BaseEquip{
				EquipType: mpb.EItem_BaseEquipType(i + 1),
				Level:     baseEquipsLevels[i],
				Star:      baseEquipsStars[i],
			})
		}
		m[node.Id] = node
	}
	rm.bots = m
	rm.logger.Debug("loadBots read finish:", zap.Any("row", m))
	return nil
}

func (rm *itemResourceMgr) getBotRsc(userId uint64) *mpb.BotRsc {
	return rm.bots[userId]
}
