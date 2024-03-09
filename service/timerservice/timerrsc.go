package timerservice

import (
	"fmt"
	gcsv "github.com/oldjon/gutil/csv"
	gdm "github.com/oldjon/gutil/dirmonitor"
	gtime "github.com/oldjon/gutil/timeutil"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"gitlab.com/morbackend/mor_services/util"
	"go.uber.org/zap"
	"strings"
	"time"
)

const (
	baseCSVPath = "./resources/timer/"

	pvpScheduleCSV    = "PVPSchedule.csv"
	pvpRankRewardsCSV = "PVPRankRewards.csv"
)

type timerResourceMgr struct {
	logger *zap.Logger
	dm     *gdm.DirMonitor
	mtr    *util.ServiceMetrics

	pvpSchedules   map[uint32]*mpb.PVPScheduleRsc
	pvpRankRewards map[uint32][]*mpb.PVPRankRewardsRsc
}

func newTimerResourceMgr(logger *zap.Logger, mtr *util.ServiceMetrics) (*timerResourceMgr, error) {
	rMgr := &timerResourceMgr{
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

func (rm *timerResourceMgr) watch() error {
	return rm.dm.StartWatch()
}

func (rm *timerResourceMgr) load() error {
	err := rm.dm.BindAndExec(pvpScheduleCSV, rm.loadPVPSchedule)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(pvpRankRewardsCSV, rm.loadPVPRankRewards)
	if err != nil {
		return err
	}
	return nil
}

func (rm *timerResourceMgr) loadPVPSchedule(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32]*mpb.PVPScheduleRsc)
	for _, data := range datas {
		node := &mpb.PVPScheduleRsc{
			PvpSeasonId: gcsv.StrToUint32(data["pvpseasonid"]),
		}
		node.StartTime, err = gtime.DateTimeStrToTimeUnixWithErr(data["starttime"], util.TimeZone_GMT)
		if err != nil {
			rm.logger.Error("loadPVPSchedule parse starttime failed",
				zap.String("starttime", data["starttime"]))
			return mpberr.ErrConfig
		}
		node.EndTime, err = gtime.DateTimeStrToTimeUnixWithErr(data["endtime"], util.TimeZone_GMT)
		if err != nil {
			rm.logger.Error("loadPVPSchedule parse endtime failed",
				zap.String("endtime", data["endtime"]))
			return mpberr.ErrConfig
		}
		strs := strings.Split(data["settlestartat"], ":")
		if len(strs) != 3 {
			rm.logger.Error("loadPVPSchedule parse settle_start_at failed",
				zap.String("settle_start_at", data["settlestartat"]))
			return mpberr.ErrConfig
		}
		node.SettleStartAt = gcsv.StrToInt64(strs[0])*3600 + gcsv.StrToInt64(strs[1])*60 + gcsv.StrToInt64(strs[2])
		strs = strings.Split(data["settleendat"], ":")
		if len(strs) != 3 {
			rm.logger.Error("loadPVPSchedule parse settle_end_at failed",
				zap.String("settle_end_at", data["settleendat"]))
			return mpberr.ErrConfig
		}
		node.SettleEndAt = gcsv.StrToInt64(strs[0])*3600 + gcsv.StrToInt64(strs[1])*60 + gcsv.StrToInt64(strs[2])
		m[node.PvpSeasonId] = node
		rm.logger.Debug("loadPVPSchedule read:", zap.Any("row", node))
	}

	rm.pvpSchedules = m
	rm.logger.Debug("loadPVPSchedule read finish:", zap.Any("rows", m))
	return nil
}

func (rm *timerResourceMgr) canPVPSettle(now time.Time) (uint32, string, bool) {
	var rsc *mpb.PVPScheduleRsc
	nowUnix := now.Unix()
	for _, v := range rm.pvpSchedules {
		if v.StartTime <= nowUnix && nowUnix <= v.EndTime {
			rsc = v
			break
		}
	}
	if rsc == nil {
		return 0, "", false
	}

	settleStartAt := rsc.SettleStartAt
	settleEndAt := rsc.SettleEndAt

	zeroUnix := gtime.TodayXHourTimeUnix(0, util.TimeZone_GMT)
	nowUnix = nowUnix - zeroUnix

	if settleStartAt <= settleEndAt { // start_at and end_at in same day
		if settleStartAt <= nowUnix && nowUnix <= settleEndAt {
			return rsc.PvpSeasonId, gtime.TimeToDate(now.In(util.TimeZone_GMT)), true
		}
		return 0, "", false
	}
	if nowUnix <= settleEndAt {
		return rsc.PvpSeasonId, gtime.TimeToDate(now.In(util.TimeZone_GMT).Add(-24 * time.Hour)), true
	} else if nowUnix >= settleStartAt {
		return rsc.PvpSeasonId, gtime.TimeToDate(now.In(util.TimeZone_GMT)), true
	}
	return 0, "", false
}

func (rm *timerResourceMgr) loadPVPRankRewards(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32][]*mpb.PVPRankRewardsRsc)
	for _, data := range datas {
		node := &mpb.PVPRankRewardsRsc{
			PvpSeasonId: gcsv.StrToUint32(data["pvpseasonid"]),
			RankRange:   gcsv.StrToUint32Slice(data["rankrange"], ":"),
		}
		if len(node.RankRange) != 2 || node.RankRange[0] > node.RankRange[1] {
			rm.logger.Error("loadPVPRankRewards parse rank range failed",
				zap.String("rankrange", data["rankrange"]))
			return mpberr.ErrConfig
		}
		node.Rewards, err = com.ReadAwardsRsc(data["rewards"])
		if err != nil {
			rm.logger.Error("loadPVPRankRewards parse rewards failed",
				zap.String("rewards", data["rewards"]))
			return mpberr.ErrConfig
		}

		m[node.PvpSeasonId] = append(m[node.PvpSeasonId], node)
		rm.logger.Debug("loadPVPRankRewards read:", zap.Any("row", node))
	}

	rm.pvpRankRewards = m
	rm.logger.Debug("loadPVPRankRewards read finish:", zap.Any("rows", m))
	return nil
}

func (rm *timerResourceMgr) getPVPSeasonRankRewardsRscs(seasonId uint32) []*mpb.PVPRankRewardsRsc {
	return rm.pvpRankRewards[seasonId]
}
