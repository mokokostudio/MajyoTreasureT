package accountservice

import (
	"fmt"

	"github.com/oldjon/gutil"
	gcsv "github.com/oldjon/gutil/csv"
	gdm "github.com/oldjon/gutil/dirmonitor"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/util"
	"go.uber.org/zap"
)

const (
	//csvSuffix   = ".csv"
	baseCSVPath = "./resources/account/"

	botsCSV = "Bots.csv"
)

type accountResourceMgr struct {
	logger *zap.Logger
	dm     *gdm.DirMonitor
	mtr    *util.ServiceMetrics
	bots   map[uint64]*mpb.BotRsc
}

func newAccountResourceMgr(logger *zap.Logger, mtr *util.ServiceMetrics) (*accountResourceMgr, error) {
	rMgr := &accountResourceMgr{
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

func (rm *accountResourceMgr) watch() error {
	return rm.dm.StartWatch()
}

func (rm *accountResourceMgr) load() error {
	err := rm.dm.BindAndExec(botsCSV, rm.loadBots)
	if err != nil {
		return err
	}
	return nil
}

func (rm *accountResourceMgr) loadBots(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint64]*mpb.BotRsc)
	for _, data := range datas {
		node := &mpb.BotRsc{
			Id:       gutil.StrToUint64(data["id"]),
			Nickname: data["nickname"],
			Icon:     data["icon"],
		}
		m[node.Id] = node
	}

	rm.bots = m
	rm.logger.Debug("loadBots read finish:", zap.Any("row", m))
	return nil
}

func (rm *accountResourceMgr) getBotRsc(userId uint64) *mpb.BotRsc {
	return rm.bots[userId]
}
