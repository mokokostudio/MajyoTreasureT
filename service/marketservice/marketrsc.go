package marketservice

import (
	"fmt"

	gcsv "github.com/oldjon/gutil/csv"
	gdm "github.com/oldjon/gutil/dirmonitor"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"gitlab.com/morbackend/mor_services/util"
	"go.uber.org/zap"
)

const (
	//csvSuffix   = ".csv"
	baseCSVPath = "./resources/market/"

	marketConfigCSV = "MarketConfig.csv"
	marketGoodsCSV  = "MarketGoods.csv"
)

type marketResourceMgr struct {
	logger       *zap.Logger
	dm           *gdm.DirMonitor
	mtr          *util.ServiceMetrics
	marketConfig *mpb.MarketConfigRsc
	goodsMap     map[uint32]*mpb.MarketGoodsRsc
}

func newMarketResourceMgr(logger *zap.Logger, mtr *util.ServiceMetrics) (*marketResourceMgr, error) {
	rMgr := &marketResourceMgr{
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

func (rm *marketResourceMgr) watch() error {
	return rm.dm.StartWatch()
}

func (rm *marketResourceMgr) load() error {
	err := rm.dm.BindAndExec(marketConfigCSV, rm.loadMarketConfig)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(marketGoodsCSV, rm.loadMarketGoods)
	if err != nil {
		return err
	}
	return nil
}

func (rm *marketResourceMgr) loadMarketConfig(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	if len(datas) != 1 {
		rm.logger.Error(fmt.Sprintf("load %s failed: config row num %d", csvPath, len(datas)))
		return mpberr.ErrConfig
	}
	cfg := &mpb.MarketConfigRsc{
		OnSellOrderCnt: gcsv.StrToUint32(datas[0]["onsellordercnt"]),
		OpenSellAfter:  gcsv.StrToInt64(datas[0]["opensellafter"]),
		MinPrice:       gcsv.StrToUint64(datas[0]["minprice"]),
		MaxPrice:       gcsv.StrToUint64(datas[0]["maxprice"]),
		MaxSellCnt:     gcsv.StrToUint64(datas[0]["maxsellcnt"]),
	}

	rm.marketConfig = cfg
	rm.logger.Debug("loadMarketConfig read finish:", zap.Any("row", cfg))
	return nil
}

func (rm *marketResourceMgr) getOnSellOrderCnt() uint32 {
	return rm.marketConfig.GetOnSellOrderCnt()
}

func (rm *marketResourceMgr) getOpenSellAfter() int64 {
	return rm.marketConfig.GetOpenSellAfter()
}

func (rm *marketResourceMgr) getMinPrice() uint64 {
	return rm.marketConfig.GetMinPrice()
}

func (rm *marketResourceMgr) getMaxPrice() uint64 {
	return rm.marketConfig.GetMaxPrice()
}

func (rm *marketResourceMgr) getMaxSellCnt() uint64 {
	return rm.marketConfig.GetMaxSellCnt()
}

func (rm *marketResourceMgr) loadMarketGoods(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32]*mpb.MarketGoodsRsc)
	for _, data := range datas {
		node := &mpb.MarketGoodsRsc{
			GoodsId: gcsv.StrToUint32(data["goodsid"]),
			GasRate: gcsv.StrToUint64(data["gasrate"]),
		}

		m[node.GoodsId] = node
		rm.logger.Debug("loadMarketGoods read:", zap.Any("row", node))
	}

	rm.goodsMap = m
	rm.logger.Debug("loadMarketGoods read finish:", zap.Any("rows", m))
	return nil
}

func (rm *marketResourceMgr) getGoodsRsc(goodsId uint32) *mpb.MarketGoodsRsc {
	return rm.goodsMap[goodsId]
}
