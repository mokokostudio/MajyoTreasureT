package timerservice

import (
	"context"
	"strings"
	"time"

	"github.com/oldjon/gutil"
	"github.com/oldjon/gutil/gdb"
	grmux "github.com/oldjon/gutil/redismutex"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"gitlab.com/morbackend/mor_services/util"
	"go.uber.org/zap"
)

type timerDAO struct {
	logger *zap.Logger
	rm     *timerResourceMgr
	rMux   *grmux.RedisMutex
	gameDB *gdb.DB
	pvpDB  *gdb.DB
	tmpDB  *gdb.DB
}

func newTimerDAO(svc *TimerService, rMux *grmux.RedisMutex, gameRedis, pvpRedis, tmpRedis gdb.RedisClient) *timerDAO {
	return &timerDAO{
		logger: svc.logger,
		rm:     svc.rm,
		rMux:   rMux,
		gameDB: gdb.NewDB(gameRedis, nil),
		pvpDB:  gdb.NewDB(pvpRedis, nil),
		tmpDB:  gdb.NewDB(tmpRedis, nil),
	}
}

func (dao *timerDAO) checkPVPDailyNotSettled(ctx context.Context, date string) bool {
	ok, err := dao.pvpDB.SetNX(ctx, com.PVPSettleMarkKey(date), 1)
	if err != nil {
		dao.logger.Error("checkPVPDailyNotSettled failed", zap.String("date", date), zap.Error(err))
		return false
	}
	return ok
}

func (dao *timerDAO) handlePVPDailySettle(ctx context.Context, season uint32, date string,
	rscs []*mpb.PVPRankRewardsRsc) error {
	// init rewards by ranks
	var shardIndexes = make(map[uint32]bool)
	date = strings.ReplaceAll(date, "-", "")
	for _, rsc := range rscs {
		for i := rsc.RankRange[0] - 1; i <= rsc.RankRange[1]-1; i++ {
			shardIndexes[i/com.PVPRankShardUserCnt] = true
		}
	}

	var rank = make(map[uint32]map[uint32]uint64)
	for shardIndex := range shardIndexes {
		key := com.PVPRanksKey(shardIndex)
		uids, err := gdb.ToUint64Slice(dao.pvpDB.ZRange(ctx, com.PVPRanksKey(shardIndex), 0, -1))
		if err != nil {
			dao.logger.Error("handlePVPDailySettle get rank failed", zap.String("key", key), zap.Error(err))
			return mpberr.ErrDB
		}
		m := make(map[uint32]uint64)
		for i, uid := range uids {
			m[uint32(i)+shardIndex*com.PVPRankShardUserCnt] = uid
		}
		rank[shardIndex] = m
	}

	var rank0Uid uint64
	for _, rsc := range rscs {
		var keys []string
		var values []uint32
		for i := rsc.RankRange[0] - 1; i <= rsc.RankRange[1]-1; i++ {
			m := rank[i/com.PVPRankShardUserCnt]
			if util.IsBotUId(m[i]) {
				continue
			}
			if i == 0 {
				rank0Uid = m[i]
			}
			keys = append(keys, com.UserPVPInfoKey(m[i]))
			values = append(values, i)
		}
		err := dao.batchSetPVPRankRewards(ctx, season, date, keys, values)
		if err != nil {
			return err
		}
	}

	// save rank 0 to pvp history
	dao.savePVPHistory(ctx, date, rank0Uid)

	return nil
}

func (dao *timerDAO) batchSetPVPRankRewards(ctx context.Context, season uint32, date string, keys []string,
	values []uint32) error {
	if len(keys) != len(values) {
		return mpberr.ErrParam
	}

	var err error
	mk := gutil.Uint32ToString(season) + "_" + date
	for i, key := range keys {
		err = dao.rMux.Safely(ctx, key, func() error {
			var userPVPInfo = &mpb.DBUserPVPInfo{}
			err := dao.pvpDB.GetObject(ctx, key, userPVPInfo)
			if err != nil {
				dao.logger.Error("batchSetPVPRankRewards GetObject failed",
					zap.String("key", key), zap.Error(err))
				return mpberr.ErrDB
			}

			if userPVPInfo == nil {
				userPVPInfo = &mpb.DBUserPVPInfo{}
			}
			if userPVPInfo.RankRewards == nil {
				userPVPInfo.RankRewards = make(map[string]*mpb.DBUserPVPDailyRewardsNode)
			}

			_, ok := userPVPInfo.RankRewards[mk]
			if ok {
				return nil
			}

			userPVPInfo.RankRewards[mk] = &mpb.DBUserPVPDailyRewardsNode{
				Rank:   values[i],
				Status: 1,
			}
			err = dao.pvpDB.SetObject(ctx, key, userPVPInfo)
			if err != nil {
				dao.logger.Error("batchSetPVPRankRewards SetObject failed",
					zap.String("key", key), zap.Error(err))
				return mpberr.ErrDB
			}
			return nil
		})
		if err != nil {
			dao.logger.Error("batchSetPVPRankRewards Safely failed",
				zap.String("key", key), zap.Error(err))
			continue
		}
	}

	return nil
}

func (dao *timerDAO) savePVPHistory(ctx context.Context, dateStr string, rank0Uid uint64) {
	if rank0Uid == 0 {
		return
	}

	zeroTime, _ := time.ParseInLocation("20060102 15:04:05", dateStr+" 00:00:00", util.TimeZone_GMT)
	zeroTime = zeroTime.AddDate(0, 0, -20)
	dateStart := gutil.StrToUint32(zeroTime.Format("20060102"))

	err := dao.rMux.Safely(ctx, com.PVPHistoryKey(), func() error {
		dbHis := &mpb.DBPVPHistory{}
		err := dao.pvpDB.GetObject(ctx, com.PVPHistoryKey(), dbHis)
		if err != nil && !dao.pvpDB.IsErrNil(err) {
			dao.logger.Error("savePVPHistory GetObject failed", zap.Error(err))
			return err
		}
		if dbHis.History == nil {
			dbHis.History = make(map[uint32]uint64)
		}
		for k := range dbHis.History {
			if k < dateStart {
				delete(dbHis.History, k)
			}
		}
		dbHis.History[gutil.StrToUint32(dateStr)] = rank0Uid
		err = dao.pvpDB.SetObject(ctx, com.PVPHistoryKey(), dbHis)
		if err != nil {
			dao.logger.Error("savePVPHistory SetObject failed", zap.Error(err))
			return err
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("savePVPHistory failed", zap.Error(err))
	}
}
