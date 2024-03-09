package gameservice

import (
	"context"
	"errors"
	gcsv "github.com/oldjon/gutil/csv"
	"sort"
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

const (
	PVPRankPageUserCnt = 50
	PVPRankPageCnt     = 1
)

type gameDAO struct {
	logger *zap.Logger
	rm     *gameResourceMgr
	rMux   *grmux.RedisMutex
	gameDB *gdb.DB
	pvpDB  *gdb.DB
	tmpDB  *gdb.DB
}

func newGameDAO(svc *GameService, rMux *grmux.RedisMutex, gameRedis, pvpRedis, tmpRedis gdb.RedisClient) *gameDAO {
	return &gameDAO{
		logger: svc.logger,
		rm:     svc.rm,
		rMux:   rMux,
		gameDB: gdb.NewDB(gameRedis, nil),
		pvpDB:  gdb.NewDB(pvpRedis, nil),
		tmpDB:  gdb.NewDB(tmpRedis, nil),
	}
}

func (dao *gameDAO) getEnergy(ctx context.Context, userId uint64, nowUnix int64) (*mpb.DBEnergy, error) {
	dbEnergy := &mpb.DBEnergy{}
	err := dao.gameDB.GetObject(ctx, com.EnergyKey(userId), dbEnergy)
	if dao.gameDB.IsErrNil(err) {
		dbEnergy.Energy = dao.rm.getEnergyLimit()
		dbEnergy.RecoverAt = nowUnix
		return dbEnergy, nil
	}
	if err != nil {
		dao.logger.Error("get energy failed", zap.Uint64("user_id", userId), zap.Error(err))
		return nil, mpberr.ErrDB
	}
	if dbEnergy.Energy >= dao.rm.getEnergyLimit() {
		dbEnergy.RecoverAt = time.Now().Unix()
		return dbEnergy, nil
	}
	er := uint32((nowUnix - dbEnergy.RecoverAt) / dao.rm.getEnergyRecoverTime())
	dbEnergy.Energy = gutil.Min(dbEnergy.Energy+er, dao.rm.getEnergyLimit())
	dbEnergy.RecoverAt = gutil.If(dbEnergy.Energy == dao.rm.getEnergyLimit(),
		nowUnix,
		dbEnergy.RecoverAt+int64(er)*dao.rm.getEnergyRecoverTime())
	return dbEnergy, err
}

func (dao *gameDAO) recoverEnergy(ctx context.Context, userId uint64) error {
	_, err := dao.gameDB.Del(ctx, com.EnergyKey(userId))
	if err != nil {
		dao.logger.Error("recoverEnergy del key failed", zap.Uint64("user_id", userId), zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) consumeEnergy(ctx context.Context, userId uint64, consumeEnergy uint32, nowUnix int64) (
	*mpb.DBEnergy, error) {
	var dbEnergy *mpb.DBEnergy
	key := com.EnergyKey(userId)
	err := dao.rMux.Safely(ctx, key, func() error {
		var err error
		dbEnergy, err = dao.getEnergy(ctx, userId, nowUnix)
		if err != nil {
			return err
		}
		if dbEnergy.Energy < consumeEnergy {
			return mpberr.ErrEnergyNotEnough
		}
		dbEnergy.Energy -= consumeEnergy
		err = dao.gameDB.SetObject(ctx, key, dbEnergy)
		if err != nil {
			dao.logger.Error("consumeEnergy update db failed", zap.Uint64("user_id", userId),
				zap.Any("energy", dbEnergy), zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("consumeEnergy failed", zap.Uint64("user_id", userId), zap.Error(err))
		return nil, err
	}
	return dbEnergy, nil
}

func (dao *gameDAO) addEnergy(ctx context.Context, userId uint64, addEnergy uint32, nowUnix int64) (*mpb.DBEnergy,
	error) {
	var dbEnergy *mpb.DBEnergy
	key := com.EnergyKey(userId)
	err := dao.rMux.Safely(ctx, key, func() error {
		var err error
		dbEnergy, err = dao.getEnergy(ctx, userId, nowUnix)
		if err != nil {
			return err
		}
		dbEnergy.Energy += addEnergy
		err = dao.gameDB.SetObject(ctx, key, dbEnergy)
		if err != nil {
			dao.logger.Error("addEnergy update db failed", zap.Uint64("user_id", userId),
				zap.Any("energy", dbEnergy), zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("addEnergy failed", zap.Uint64("user_id", userId), zap.Error(err))
		return nil, err
	}
	return dbEnergy, nil
}

func (dao *gameDAO) getFightHistory(ctx context.Context, userId uint64) (*mpb.DBFightHistory, error) {
	dbHis := &mpb.DBFightHistory{}
	err := dao.gameDB.GetObject(ctx, com.FightHistoryKey(userId), dbHis)
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("getFightHistory get failed", zap.Uint64("user_id", userId), zap.Error(err))
		return nil, mpberr.ErrDB
	}
	if dbHis.WinTimes == nil {
		dbHis.WinTimes = make(map[uint32]uint32)
	}
	return dbHis, nil
}

func (dao *gameDAO) updateFightHistory(ctx context.Context, userId uint64, bossId uint32) (*mpb.DBFightHistory,
	error) {
	dbHis := &mpb.DBFightHistory{}
	key := com.FightHistoryKey(userId)
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.gameDB.GetObject(ctx, key, dbHis)
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("updateFightHistory get obj from db failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		if dbHis.WinTimes == nil {
			dbHis.WinTimes = make(map[uint32]uint32)
		}
		dbHis.WinTimes[bossId] += 1
		err = dao.gameDB.SetObject(ctx, key, dbHis)
		if err != nil {
			dao.logger.Error("updateFightHistory set obj failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("updateFightHistory failed", zap.Uint64("user_id", userId), zap.Error(err))
		return nil, err
	}
	return dbHis, nil
}

func (dao *gameDAO) getHiddenBoss(ctx context.Context, bossUUID uint64) (*mpb.DBHiddenBoss, error) {
	dbHiddenBoss := &mpb.DBHiddenBoss{}
	err := dao.gameDB.GetObject(ctx, com.HiddenBossKey(bossUUID), dbHiddenBoss)
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("getHiddenBoss get obj failed", zap.Uint64("boss_uuid", bossUUID), zap.Error(err))
		return nil, mpberr.ErrDB
	}
	if dao.gameDB.IsErrNil(err) {
		return nil, mpberr.ErrHiddenBossNotFound
	}
	if dbHiddenBoss.LastFightTime == nil {
		dbHiddenBoss.LastFightTime = make(map[uint64]int64)
	}
	if dbHiddenBoss.Dmgs == nil {
		dbHiddenBoss.Dmgs = make(map[uint64]uint64)
	}
	return dbHiddenBoss, nil
}

func (dao *gameDAO) updateHiddenBoss(ctx context.Context, dbHiddenBoss *mpb.DBHiddenBoss, expiration time.Duration,
) error {
	err := dao.gameDB.SetObjectEX(ctx, com.HiddenBossKey(dbHiddenBoss.BossUuid), dbHiddenBoss, expiration)
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("updateHiddenBoss set obj failed",
			zap.Uint64("boss_uuid", dbHiddenBoss.BossUuid),
			zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) fightHiddenBoss(ctx context.Context, bossUUId uint64, fightFunc func() error) error {
	err := dao.rMux.Safely(ctx, com.HiddenBossKey(bossUUId), fightFunc)
	if err != nil {
		dao.logger.Error("fightHiddenBoss failed", zap.Uint64("boss_uuid", bossUUId),
			zap.Error(err))
		return err
	}
	return nil
}

func (dao *gameDAO) getHiddenBossFightCD(ctx context.Context, userId uint64) (int64, error) {
	t, err := gdb.ToUint64(dao.gameDB.Get(ctx, com.UserHiddenBossCDKey(userId)))
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("getHiddenBossFightCD get cd failed",
			zap.Uint64("user_id", userId), zap.Error(err))
		return 0, mpberr.ErrDB
	}
	return int64(t), nil
}

func (dao *gameDAO) updateHiddenBossFightCD(ctx context.Context, userId uint64, cd int64) error {
	err := dao.gameDB.Set(ctx, com.UserHiddenBossCDKey(userId), cd)
	if err != nil {
		dao.logger.Error("updateHiddenBossFightCD failed", zap.Uint64("user_id", userId),
			zap.Int64("cd", cd), zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) getHiddenBossFindHistory(ctx context.Context, userId uint64) (*mpb.DBHiddenBossFindHistory, error) {
	dbHis := &mpb.DBHiddenBossFindHistory{}
	err := dao.gameDB.GetObject(ctx, com.HiddenBossFindHistoryKey(userId), dbHis)
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("getHiddenBossFindHistory get cd failed",
			zap.Uint64("user_id", userId), zap.Error(err))
		return nil, mpberr.ErrDB
	}
	if dbHis.BossExpireAt == nil {
		dbHis.BossExpireAt = make(map[uint64]int64)
	}
	return dbHis, nil
}

func (dao *gameDAO) updateHiddenBossFindHistory(ctx context.Context, userId uint64, dbHis *mpb.DBHiddenBossFindHistory,
) error {
	err := dao.gameDB.SetObject(ctx, com.HiddenBossFindHistoryKey(userId), dbHis)
	if err != nil {
		dao.logger.Error("updateHiddenBossFindHistory failed",
			zap.Uint64("user_id", userId), zap.Any("db_his", dbHis),
			zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) pushIntoHiddenBossPool(ctx context.Context, bossUUId uint64, expireAt int64) error {
	_, err := dao.gameDB.ZAdd(ctx, com.HiddenBossPoolKey(), expireAt, bossUUId)
	if err != nil {
		dao.logger.Error("pushHiddenBossToPool failed",
			zap.Uint64("boss_uuid", bossUUId), zap.Int64("expire_at", expireAt), zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) delFromHiddenBossPool(ctx context.Context, bossUUIds []uint64) error {
	list := make([]any, len(bossUUIds))
	for i, v := range bossUUIds {
		list[i] = v
	}
	_, err := dao.gameDB.ZRem(ctx, com.HiddenBossPoolKey(), list...)
	if err != nil {
		dao.logger.Error("delFromHiddenBossPool failed",
			zap.Any("boss_uuid", bossUUIds), zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) randomFromHiddenBossPool(ctx context.Context, nowUnix int64) (uint64, error) {
	values, err := gdb.ToUint64Slice(dao.gameDB.ZRangeWithScores(ctx, com.HiddenBossPoolKey(), 0, -1))
	if err != nil {
		dao.logger.Error("randomFromHiddenBossPool failed", zap.Error(err))
		return 0, mpberr.ErrDB
	}

	bossUUids := make([]uint64, 0, len(values)/2)
	expiredBoss := make([]uint64, 0, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		v := values[i]
		s := values[i+1]
		if s > uint64(nowUnix) {
			bossUUids = append(bossUUids, v)
		} else {
			expiredBoss = append(expiredBoss, v)
		}
	}

	if len(expiredBoss) > 0 {
		_ = dao.delFromHiddenBossPool(ctx, expiredBoss)
	}

	if len(bossUUids) == 0 {
		return 0, nil
	}

	return util.RandFromPool(bossUUids)
}

func (dao *gameDAO) getBossDefeatHistory(ctx context.Context, userId uint64) (*mpb.DBBossDefeatHistory, error) {
	dbHis := &mpb.DBBossDefeatHistory{}
	err := dao.gameDB.GetObject(ctx, com.BossDefeatHistory(userId), dbHis)
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("getBossDefeatHistory get cd failed",
			zap.Uint64("user_id", userId), zap.Error(err))
		return nil, mpberr.ErrDB
	}
	if dbHis.BossDefeatHistory == nil {
		dbHis.BossDefeatHistory = make(map[uint32]*mpb.DBBossDefeatHistoryNode)
	}
	return dbHis, nil
}

func (dao *gameDAO) updateBossDefeatHistory(ctx context.Context, userId uint64, bossRsc *mpb.BossRsc) error {
	key := com.BossDefeatHistory(userId)
	err := dao.rMux.Safely(ctx, key, func() error {
		dbHis := &mpb.DBBossDefeatHistory{}
		err := dao.gameDB.GetObject(ctx, com.BossDefeatHistory(userId), dbHis)
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("updateBossDefeatHistory get cd failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		if dbHis.BossDefeatHistory == nil {
			dbHis.BossDefeatHistory = make(map[uint32]*mpb.DBBossDefeatHistoryNode)
		}
		node := dbHis.BossDefeatHistory[bossRsc.Class]
		if node != nil && node.Level >= bossRsc.Level {
			return nil
		}
		if node == nil {
			node = &mpb.DBBossDefeatHistoryNode{
				BossId: bossRsc.BossId,
				Level:  bossRsc.Level,
			}
		} else {
			node.BossId = bossRsc.BossId
			node.Level = bossRsc.Level
		}
		dbHis.BossDefeatHistory[bossRsc.Class] = node
		err = dao.gameDB.SetObject(ctx, key, dbHis)
		if err != nil {
			dao.logger.Error("updateBossDefeatHistory SetObject failed", zap.Uint64("user_id", userId),
				zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("updateBossDefeatHistory safely failed", zap.Uint64("user_id", userId),
			zap.Error(err))
		return err
	}
	return nil
}

func (dao *gameDAO) getUserPVPInfoAndClaimRankRewards(ctx context.Context, season uint32, userId uint64, nowUnix int64) (
	*mpb.DBUserPVPInfo, [][3]uint32, error) {
	key := com.UserPVPInfoKey(userId)
	dbInfo := &mpb.DBUserPVPInfo{}
	rewards := make([][3]uint32, 0)
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.pvpDB.GetObject(ctx, key, dbInfo)
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("getUserPVPInfo GetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		var needUpdate bool
		if dbInfo == nil {
			dbInfo = &mpb.DBUserPVPInfo{
				ChallengeCnt: dao.rm.getDailyPVPFreeChallengeCnt(season),
				UpdateAt:     nowUnix,
			}
			needUpdate = true
		}
		if dbInfo.UpdateAt == 0 {
			dbInfo.ChallengeCnt = dao.rm.getDailyPVPFreeChallengeCnt(season)
			dbInfo.UpdateAt = nowUnix
			needUpdate = true
		}

		if dbInfo.RankRewards == nil {
			dbInfo.RankRewards = make(map[string]*mpb.DBUserPVPDailyRewardsNode)
			needUpdate = true
		}

		// TODO(fishillte): TODEL for test
		//dbInfo.RankRewards["1_20140126"] = &mpb.DBUserPVPDailyRewardsNode{
		//	Rank:   1,
		//	Status: 1,
		//}
		//dbInfo.RankRewards["1_20140125"] = &mpb.DBUserPVPDailyRewardsNode{
		//	Rank:   1,
		//	Status: 1,
		//}
		// TODO(fishillte): TODEL for test

		for k, v := range dbInfo.RankRewards {
			switch v.Status {
			case 1:
				tmps := gcsv.StrToUint32Slice(k, "_")
				if len(tmps) == 2 {
					rewards = append(rewards, [3]uint32{tmps[0], tmps[1], v.Rank})
				}
				delete(dbInfo.RankRewards, k)
				needUpdate = true
			default:
				delete(dbInfo.RankRewards, k)
				needUpdate = true
			}
		}

		if !needUpdate {
			return nil
		}

		err = dao.pvpDB.SetObject(ctx, key, dbInfo)
		if err != nil {
			dao.logger.Error("getUserPVPInfoAndClaimRankRewards SetObject failed",
				zap.Uint64("user_id", userId), zap.Any("db_info", dbInfo), zap.Error(err))
			return err
		}

		return nil
	})
	if err != nil {
		dao.logger.Error("getUserPVPInfoAndClaimRankRewards Safely failed", zap.Uint64("user_id", userId),
			zap.Error(err))
		return nil, nil, err
	}

	if len(rewards) == 0 {
		return dbInfo, nil, nil
	}

	sort.Slice(rewards, func(i, j int) bool {
		if rewards[i][0] != rewards[j][0] {
			return rewards[i][0] < rewards[j][0]
		}
		return rewards[i][1] < rewards[j][1]
	})

	return dbInfo, rewards, nil
}

func (dao *gameDAO) checkAndCostPVPChallengeCnt(ctx context.Context, season uint32, userId uint64, nowUnix int64) (*mpb.DBUserPVPInfo,
	error) {
	key := com.UserPVPInfoKey(userId)
	dbInfo := &mpb.DBUserPVPInfo{}
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.pvpDB.GetObject(ctx, key, dbInfo)
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("checkAndCostPVPChallengeCnt GetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		if dbInfo == nil {
			dbInfo = &mpb.DBUserPVPInfo{
				ChallengeCnt: dao.rm.getDailyPVPFreeChallengeCnt(season),
				UpdateAt:     nowUnix,
			}
		}
		if dbInfo.UpdateAt == 0 {
			dbInfo.ChallengeCnt = dao.rm.getDailyPVPFreeChallengeCnt(season)
			dbInfo.UpdateAt = nowUnix
		}
		if !util.IsTimestampInSameDayInLocation(dbInfo.UpdateAt, nowUnix, util.TimeZone_GMT) {
			dbInfo = &mpb.DBUserPVPInfo{
				ChallengeCnt: dao.rm.getDailyPVPFreeChallengeCnt(season),
				UpdateAt:     nowUnix,
			}
		}
		if dbInfo.ChallengeCnt == 0 {
			return mpberr.ErrNotEnoughPVPChallengeCnt
		}
		dbInfo.ChallengeCnt--
		dbInfo.UpdateAt = nowUnix
		err = dao.pvpDB.SetObject(ctx, key, dbInfo)
		if err != nil {
			dao.logger.Error("checkAndCostPVPChallengeCnt SetObject failed", zap.Uint64("user_id", userId),
				zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		if !errors.Is(err, mpberr.ErrNotEnoughPVPChallengeCnt) {
			dao.logger.Error("checkAndCostPVPChallengeCnt failed", zap.Uint64("user_id", userId),
				zap.Error(err))
		}
		return dbInfo, err
	}

	return dbInfo, nil
}

// initPVPRank rank start from 0
func (dao *gameDAO) initPVPRank(ctx context.Context, userId uint64) (uint32, error) {
	rank, err := gdb.ToUint32(dao.pvpDB.IncrBy(ctx, com.PVPRankIndexKey(), 1))
	if err != nil {
		dao.logger.Error("initPVPRank failed", zap.Error(err))
		return 0, mpberr.ErrDB
	}
	shardIndex := rank / com.PVPRankShardUserCnt
	_, err = dao.pvpDB.ZAdd(ctx, com.PVPRanksKey(shardIndex), rank, userId)
	if err != nil {
		dao.logger.Error("initPVPRank ZAdd failed", zap.Uint64("user_id", userId), zap.Error(err))
		return 0, mpberr.ErrDB
	}
	err = dao.pvpDB.Set(ctx, com.UserPVPRankKey(userId), rank)
	if err != nil {
		dao.logger.Error("initPVPRank Set rank failed", zap.Uint64("user_id", userId), zap.Error(err))
		return 0, mpberr.ErrDB
	}
	return rank, nil
}

func (dao *gameDAO) getUserPVPRank(ctx context.Context, userId uint64) (uint32, error) {
	userRank, err := gdb.ToUint32(dao.pvpDB.Get(ctx, com.UserPVPRankKey(userId)))
	if dao.pvpDB.IsErrNil(err) {
		userRank, err = dao.initPVPRank(ctx, userId)
		if err != nil {
			return 0, err
		}
	} else if err != nil {
		dao.logger.Error("getUserPVPRank failed", zap.Uint64("user_id", userId), zap.Error(err))
		return 0, mpberr.ErrDB
	}
	return userRank, nil
}

func (dao *gameDAO) fightPVP(ctx context.Context, userId uint64, defenderId uint64, f func() bool) (newRank uint32,
	oldRank uint32, err error) {
	uidA := gutil.Min(userId, defenderId)
	uidB := gutil.Max(userId, defenderId)

	keyUserAPVPRank := com.UserPVPRankKey(uidA)
	keyUserBPVPRank := com.UserPVPRankKey(uidB)
	keyUserPVPRank := com.UserPVPRankKey(userId)
	keyDefenderPVPRank := com.UserPVPRankKey(defenderId)

	if !dao.rMux.Lock(ctx, keyUserAPVPRank) { // lock smaller uid
		dao.logger.Error("fightPVP lock key failed", zap.String("key", keyUserAPVPRank))
		return 0, 0, mpberr.ErrDB
	}
	defer func() { dao.rMux.Unlock(ctx, keyUserAPVPRank) }()

	if !dao.rMux.Lock(ctx, keyUserBPVPRank) { // lock bigger uid
		dao.logger.Error("fightPVP lock key failed", zap.String("key", keyUserBPVPRank))
		return 0, 0, mpberr.ErrDB
	}
	defer func() { dao.rMux.Unlock(ctx, keyUserBPVPRank) }()

	userRank, err := gdb.ToUint32(dao.pvpDB.Get(ctx, keyUserPVPRank))
	if dao.pvpDB.IsErrNil(err) {
		userRank, err = dao.initPVPRank(ctx, userId)
		if err != nil {
			return 0, 0, err
		}
	} else if err != nil {
		return 0, 0, mpberr.ErrDB
	}
	defenderRank, err := gdb.ToUint32(dao.pvpDB.Get(ctx, keyDefenderPVPRank))
	if dao.pvpDB.IsErrNil(err) {
		defenderRank, err = dao.initPVPRank(ctx, defenderId)
		if err != nil {
			return 0, 0, err
		}
	} else if err != nil {
		return 0, 0, mpberr.ErrDB
	}
	if defenderRank > userRank {
		return 0, 0, mpberr.ErrPVPTargetRankChanged
	}

	userRankShardKey := com.PVPRanksKey(userRank / com.PVPRankShardUserCnt)
	defenderRankShardKey := com.PVPRanksKey(defenderRank / com.PVPRankShardUserCnt)

	if !dao.rMux.Lock(ctx, defenderRankShardKey) { // lock smaller uid
		dao.logger.Error("fightPVP lock key failed", zap.String("key", defenderRankShardKey))
		return 0, 0, mpberr.ErrDB
	}
	defer func() { dao.rMux.Unlock(ctx, defenderRankShardKey) }()

	if defenderRankShardKey != userRankShardKey {
		if !dao.rMux.Lock(ctx, userRankShardKey) { // lock bigger uid
			dao.logger.Error("fightPVP lock key failed", zap.String("key", userRankShardKey))
			return 0, 0, mpberr.ErrDB
		}
		defer func() { dao.rMux.Unlock(ctx, userRankShardKey) }()
	}

	// check rank is right
	userShardRank, err := gdb.ToUint32(dao.pvpDB.ZRank(ctx, userRankShardKey, userId))
	if err != nil {
		return 0, 0, mpberr.ErrDB
	}
	if userShardRank != userRank%com.PVPRankShardUserCnt {
		dao.logger.Error("fightPVP user rank is not match with rank in rank shard key",
			zap.Uint64("user_id", userId))
		return 0, 0, mpberr.ErrDB
	}

	defenderShardRank, err := gdb.ToUint32(dao.pvpDB.ZRank(ctx, defenderRankShardKey, defenderId))
	if err != nil {
		return 0, 0, mpberr.ErrDB
	}
	if defenderShardRank != defenderRank%com.PVPRankShardUserCnt {
		dao.logger.Error("fightPVP user rank is not match with rank in rank shard key",
			zap.Uint64("user_id", defenderId))
		return 0, 0, mpberr.ErrDB
	}

	win := f()

	if !win { // lose
		return userRank + 1, userRank + 1, nil
	}

	_, err = dao.pvpDB.ZRem(ctx, defenderRankShardKey, defenderId)
	if err != nil {
		dao.logger.Error("fightPVP remove defender from rank shard key failed",
			zap.Uint64("defender_id", defenderId), zap.String("key", defenderRankShardKey), zap.Error(err))
		return 0, 0, mpberr.ErrDB
	}

	_, err = dao.pvpDB.ZRem(ctx, userRankShardKey, userId)
	if err != nil {
		dao.logger.Error("fightPVP remove user from rank shard key failed",
			zap.Uint64("user_id", userId), zap.String("key", userRankShardKey), zap.Error(err))
		return 0, 0, mpberr.ErrDB
	}

	_, err = dao.pvpDB.ZAdd(ctx, defenderRankShardKey, defenderRank, userId)
	if err != nil {
		dao.logger.Error("fightPVP add user to rank shard key failed",
			zap.Uint64("user_id", userId), zap.String("key", defenderRankShardKey),
			zap.Uint32("rank", defenderShardRank), zap.Error(err))
		return 0, 0, mpberr.ErrDB
	}

	_, err = dao.pvpDB.ZAdd(ctx, userRankShardKey, userRank, defenderId)
	if err != nil {
		dao.logger.Error("fightPVP add defender to rank shard key failed",
			zap.Uint64("defender_id", defenderId), zap.String("key", userRankShardKey),
			zap.Uint32("rank", userRank), zap.Error(err))
		return 0, 0, mpberr.ErrDB
	}

	err = dao.pvpDB.Set(ctx, keyUserPVPRank, defenderRank)
	if err != nil {
		dao.logger.Error("fightPVP set user rank failed",
			zap.Uint64("user_id", userId), zap.String("key", keyUserPVPRank),
			zap.Uint32("rank", defenderRank), zap.Error(err))
		return 0, 0, mpberr.ErrDB
	}

	err = dao.pvpDB.Set(ctx, keyDefenderPVPRank, userRank)
	if err != nil {
		dao.logger.Error("fightPVP set defender rank failed",
			zap.Uint64("defender_id", userId), zap.String("key", keyDefenderPVPRank),
			zap.Uint32("rank", userRank), zap.Error(err))
		return 0, 0, mpberr.ErrDB
	}

	return defenderRank + 1, userRank + 1, nil
}

func (dao *gameDAO) getPVPRanks(ctx context.Context, pageNum uint32) ([]uint64, error) {
	startRank := pageNum * PVPRankPageUserCnt
	shardIndex := startRank / com.PVPRankShardUserCnt

	userIds, err := gdb.ToUint64Slice(dao.pvpDB.ZRange(ctx, com.PVPRanksKey(shardIndex),
		int64(startRank%com.PVPRankShardUserCnt), int64(startRank+PVPRankPageUserCnt-1)%com.PVPRankShardUserCnt))
	if err != nil {
		dao.logger.Error("getPVPRanks failed", zap.Uint32("page_num", pageNum), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	return userIds, nil
}

func (dao *gameDAO) getPVPRankById(ctx context.Context, rank uint32) (uint64, error) {
	shardIndex := rank / com.PVPRankShardUserCnt
	startRank := rank % com.PVPRankShardUserCnt
	userIds, err := gdb.ToUint64Slice(dao.pvpDB.ZRange(ctx, com.PVPRanksKey(shardIndex),
		int64(startRank), int64(startRank)))
	if err != nil {
		dao.logger.Error("getPVPRankById failed", zap.Uint32("rank", rank), zap.Error(err))
		return 0, mpberr.ErrDB
	}

	if len(userIds) == 0 {
		return 0, nil
	}

	return userIds[0], nil
}

func (dao *gameDAO) putManaIntoPVPManaPool(ctx context.Context, season, date uint32, mana uint64) error {
	key := com.PVPManaPoolKey(season, date)
	_, err := dao.pvpDB.IncrBy(ctx, key, int64(mana))
	if err != nil {
		dao.logger.Error("putManaIntoPVPManaPool failed", zap.String("key", key), zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) getPVPManaPool(ctx context.Context, season, date uint32) (uint64, error) {
	key := com.PVPManaPoolKey(season, date)
	mana, err := gdb.ToUint64(dao.pvpDB.Get(ctx, key))
	if err != nil && !dao.pvpDB.IsErrNil(err) {
		dao.logger.Error("getPVPManaPool failed", zap.String("key", key), zap.Error(err))
		return 0, mpberr.ErrDB
	}
	return mana, nil
}

func (dao *gameDAO) getPVPHistory(ctx context.Context) (map[uint32]uint64, error) {
	dbHis := &mpb.DBPVPHistory{}
	err := dao.pvpDB.GetObject(ctx, com.PVPHistoryKey(), dbHis)
	if err != nil && !dao.pvpDB.IsErrNil(err) {
		dao.logger.Error("getPVPHistory GetObject failed", zap.Error(err))
		return nil, mpberr.ErrDB
	}
	if dbHis.History == nil {
		dbHis.History = make(map[uint32]uint64)
	}

	return dbHis.History, nil
}

func (dao *gameDAO) setBuffCardOptions(ctx context.Context, userId uint64, bossId uint32, buffCards []uint32) error {
	dbBuffCardOpts := &mpb.DBBuffCardOptions{BossId: bossId, BuffCards: buffCards}
	key := com.BuffCardOptionsKey(userId)
	err := dao.gameDB.SetObject(ctx, key, dbBuffCardOpts)
	if err != nil {
		dao.logger.Error("setBuffCardOptions failed", zap.Uint64("user_id", userId),
			zap.Uint32("boss_id", bossId), zap.Any("buff_cards", buffCards), zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) updateBuffCardOptions(ctx context.Context, userId uint64, buffCards []uint32) error {
	dbBuffCardOpts := &mpb.DBBuffCardOptions{}
	key := com.BuffCardOptionsKey(userId)
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.gameDB.GetObject(ctx, key, dbBuffCardOpts)
		if err != nil {
			dao.logger.Error("updateBuffCardOptions GetObject failed", zap.Uint64("user_id", userId),
				zap.Any("buff_cards", buffCards), zap.Error(err))
			return mpberr.ErrDB
		}
		dbBuffCardOpts.BuffCards = buffCards
		err = dao.gameDB.SetObject(ctx, key, dbBuffCardOpts)
		if err != nil {
			dao.logger.Error("updateBuffCardOptions SetObject failed", zap.Uint64("user_id", userId),
				zap.Any("buff_cards", buffCards), zap.Error(err))
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("updateBuffCardOptions Safely failed", zap.Uint64("user_id", userId),
			zap.Any("buff_cards", buffCards), zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) getBuffCardOptions(ctx context.Context, userId uint64) (*mpb.DBBuffCardOptions, error) {
	dbBuffCardOpts := &mpb.DBBuffCardOptions{}
	key := com.BuffCardOptionsKey(userId)
	err := dao.gameDB.GetObject(ctx, key, dbBuffCardOpts)
	if err != nil {
		dao.logger.Error("getBuffCardOptions failed", zap.Uint64("user_id", userId),
			zap.Error(err))
		return nil, mpberr.ErrDB
	}
	return dbBuffCardOpts, nil
}

func (dao *gameDAO) checkAndUpdateBuffCardChoseStatus(ctx context.Context, userId uint64, oldStatus []uint32,
	newStatus uint32) (bool, error) {
	key := com.BuffCardChoseStatusKey(userId)
	var ok bool
	err := dao.rMux.Safely(ctx, key, func() error {
		status, err := gdb.ToUint32(dao.gameDB.Get(ctx, key))
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("checkAndUpdateBuffCardChoseStatus Get failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		for _, v := range oldStatus {
			if status == v {
				ok = true
			}
		}
		if !ok {
			return nil
		}
		err = dao.gameDB.Set(ctx, key, newStatus)
		if err != nil {
			dao.logger.Error("checkAndUpdateBuffCardChoseStatus Set failed",
				zap.Uint64("user_id", userId), zap.Error(err))
		}
		return nil
	})
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("checkAndUpdateBuffCardChoseStatus Get failed", zap.Uint64("user_id", userId), zap.Error(err))
		return false, mpberr.ErrDB
	}
	return ok, nil
}

func (dao *gameDAO) checkBuffCardChoseStatus(ctx context.Context, userId uint64, oldStatus []uint32) (bool, error) {
	key := com.BuffCardChoseStatusKey(userId)
	status, err := gdb.ToUint32(dao.gameDB.Get(ctx, key))
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("checkBuffCardChoseStatus Get failed",
			zap.Uint64("user_id", userId), zap.Error(err))
		return false, mpberr.ErrDB
	}
	for _, v := range oldStatus {
		if status == v {
			return true, nil
		}
	}
	return false, nil
}

func (dao *gameDAO) updateBuffCardChoseStatus(ctx context.Context, userId uint64, newStatus uint32) error {
	key := com.BuffCardChoseStatusKey(userId)
	err := dao.gameDB.Set(ctx, key, newStatus)
	if err != nil {
		dao.logger.Error("updateBuffCardChoseStatus Get failed",
			zap.Uint64("user_id", userId), zap.Error(err))
		return mpberr.ErrDB
	}
	return nil
}

func (dao *gameDAO) getBuffCardChoseStatus(ctx context.Context, userId uint64) (uint32, error) {
	key := com.BuffCardChoseStatusKey(userId)
	status, err := gdb.ToUint32(dao.gameDB.Get(ctx, key))
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("checkBuffCardChoseStatus Get failed",
			zap.Uint64("user_id", userId), zap.Error(err))
		return 0, mpberr.ErrDB
	}
	return status, nil
}

func (dao *gameDAO) useBuffCardOneRound(ctx context.Context, userId uint64) ([]uint32, []*mpb.BuffCard, error) {
	key := com.BuffCardsValidKey(userId)
	dbBC := &mpb.DBBuffCardsValid{}
	curBuffCards := make([]uint32, 0, 1)
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.gameDB.GetObject(ctx, key, dbBC)
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("useBuffCardOneRound GetObject failed", zap.Uint64("user_id", userId),
				zap.Error(err))
			return mpberr.ErrDB
		}

		tmp := make([]*mpb.DBBuffCardsValid_Node, 0, len(dbBC.BuffCards))
		for _, v := range dbBC.BuffCards {
			if v.LeftRound > 0 {
				curBuffCards = append(curBuffCards, v.BuffCardId)
			}
			v.LeftRound--
			if v.LeftRound > 0 {
				tmp = append(tmp, v)
			}
		}
		dbBC.BuffCards = tmp

		err = dao.gameDB.SetObject(ctx, key, dbBC)
		if err != nil {
			dao.logger.Error("useBuffCardOneRound SetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("useBuffCardOneRound Safely failed", zap.Uint64("user_id", userId), zap.Error(err))
		return nil, nil, err
	}
	ret := make([]*mpb.BuffCard, 0, len(dbBC.BuffCards))
	for _, v := range dbBC.BuffCards {
		ret = append(ret, &mpb.BuffCard{
			BuffCardId: v.BuffCardId,
			LeftRound:  v.LeftRound,
		})
	}
	return curBuffCards, ret, nil
}

func (dao *gameDAO) upgradeValidBuffCardsLevel(ctx context.Context, userId uint64, levelAdd uint32) error { // TODO(fishillte):
	key := com.BuffCardsValidKey(userId)
	dbBC := &mpb.DBBuffCardsValid{}
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.gameDB.GetObject(ctx, key, dbBC)
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("addValidBuffCardsRound GetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}

		if len(dbBC.BuffCards) == 0 {
			return nil
		}

		for _, v := range dbBC.BuffCards {
			for i := uint32(0); i < levelAdd; i++ {
				v.BuffCardId = dao.rm.upgradeBuffCard(v.BuffCardId)
			}
		}

		err = dao.gameDB.SetObject(ctx, key, dbBC)
		if err != nil {
			dao.logger.Error("addValidBuffCardsRound SetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("useBuffCardOneRound Safely failed", zap.Uint64("user_id", userId), zap.Error(err))
		return err
	}
	return nil
}

func (dao *gameDAO) addValidBuffCardsRound(ctx context.Context, userId uint64) error {
	key := com.BuffCardsValidKey(userId)
	dbBC := &mpb.DBBuffCardsValid{}
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.gameDB.GetObject(ctx, key, dbBC)
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("addValidBuffCardsRound GetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}

		if len(dbBC.BuffCards) == 0 {
			return nil
		}

		for _, v := range dbBC.BuffCards {
			v.LeftRound++
		}

		err = dao.gameDB.SetObject(ctx, key, dbBC)
		if err != nil {
			dao.logger.Error("addValidBuffCardsRound SetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("useBuffCardOneRound Safely failed", zap.Uint64("user_id", userId), zap.Error(err))
		return err
	}
	return nil
}

func (dao *gameDAO) addBuffCards(ctx context.Context, userId uint64, buffCardRsc *mpb.BuffCardRsc) error {
	key := com.BuffCardsValidKey(userId)
	dbBC := &mpb.DBBuffCardsValid{}
	err := dao.rMux.Safely(ctx, key, func() error {
		err := dao.gameDB.GetObject(ctx, key, dbBC)
		if err != nil && !dao.gameDB.IsErrNil(err) {
			dao.logger.Error("addValidBuffCardsRound GetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}

		dbBC.BuffCards = append(dbBC.BuffCards, &mpb.DBBuffCardsValid_Node{
			BuffCardId: buffCardRsc.CardId,
			LeftRound:  buffCardRsc.Round,
		})

		err = dao.gameDB.SetObject(ctx, key, dbBC)
		if err != nil {
			dao.logger.Error("addValidBuffCardsRound SetObject failed",
				zap.Uint64("user_id", userId), zap.Error(err))
			return mpberr.ErrDB
		}
		return nil
	})
	if err != nil {
		dao.logger.Error("useBuffCardOneRound Safely failed", zap.Uint64("user_id", userId), zap.Error(err))
		return err
	}
	return nil
}

func (dao *gameDAO) getBuffCards(ctx context.Context, userId uint64) (*mpb.DBBuffCardsValid, error) {
	key := com.BuffCardsValidKey(userId)
	dbBC := &mpb.DBBuffCardsValid{}
	err := dao.gameDB.GetObject(ctx, key, dbBC)
	if err != nil && !dao.gameDB.IsErrNil(err) {
		dao.logger.Error("addValidBuffCardsRound GetObject failed",
			zap.Uint64("user_id", userId), zap.Error(err))
		return nil, mpberr.ErrDB
	}

	return dbBC, nil
}
