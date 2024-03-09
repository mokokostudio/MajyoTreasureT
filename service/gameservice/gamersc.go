package gameservice

import (
	"fmt"
	gtime "github.com/oldjon/gutil/timeutil"
	"strings"
	"time"

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
	baseCSVPath = "./resources/game/"

	gameConfigCSV        = "GameConfig.csv"
	playerInitAttrsCSV   = "PlayerInitAttrs.csv"
	bossCSV              = "Boss.csv"
	hiddenBossTriggerCSV = "HiddenBossTrigger.csv"
	pvpScheduleCSV       = "PVPSchedule.csv"
	pvpRankRewardsCSV    = "PVPRankRewards.csv"
	pvpConfigCSV         = "PVPConfig.csv"
	buffCardCSV          = "BuffCard.csv"
	buffCardRandPoolCSV  = "BuffCardRandPool.csv"
)

type gameResourceMgr struct {
	logger *zap.Logger
	dm     *gdm.DirMonitor
	mtr    *util.ServiceMetrics

	gameConfig        *mpb.GameConfigRsc
	playerInitAttrs   *mpb.PlayerInitAttrs
	bosses            map[uint32]*mpb.BossRsc
	bossClassesMap    map[uint32][]*mpb.BossRsc
	hiddenBossTrigger map[uint32]*mpb.HiddenBossTriggerRsc
	pvpSchedules      map[uint32]*mpb.PVPScheduleRsc
	pvpRankRewards    map[uint32][]*mpb.PVPRankRewardsRsc
	pvpConfig         map[uint32]*mpb.PVPConfigRsc
	buffCards         map[uint32]*mpb.BuffCardRsc
	buffCardTypeMap   map[mpb.EGame_BuffCardType][]*mpb.BuffCardRsc
	buffCardRandPools map[uint32]*mpb.BuffCardRandPoolRsc
}

func newGameResourceMgr(logger *zap.Logger, mtr *util.ServiceMetrics) (*gameResourceMgr, error) {
	rMgr := &gameResourceMgr{
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

func (rm *gameResourceMgr) watch() error {
	return rm.dm.StartWatch()
}

func (rm *gameResourceMgr) load() error {
	err := rm.dm.BindAndExec(gameConfigCSV, rm.loadGameConfig)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(playerInitAttrsCSV, rm.loadPlayerInitAttrs)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(bossCSV, rm.loadBoss)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(hiddenBossTriggerCSV, rm.loadHiddenBossTrigger)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(pvpScheduleCSV, rm.loadPVPSchedule)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(pvpRankRewardsCSV, rm.loadPVPRankRewards)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(pvpConfigCSV, rm.loadPVPConfig)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(buffCardCSV, rm.loadBuffCard)
	if err != nil {
		return err
	}
	err = rm.dm.BindAndExec(buffCardRandPoolCSV, rm.loadBuffCardRandPool)
	if err != nil {
		return err
	}
	return nil
}

func (rm *gameResourceMgr) loadGameConfig(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	if len(datas) != 1 {
		rm.logger.Error(fmt.Sprintf("load %s failed: config row num %d", csvPath, len(datas)))
		return mpberr.ErrConfig
	}
	node := &mpb.GameConfigRsc{
		EnergyLimit:       gutil.StrToUint32(datas[0]["energylimit"]),
		EnergyRecoverTime: gutil.StrToInt64(datas[0]["energyrecovertime"]),
		FightHiddenBossCd: gutil.StrToInt64(datas[0]["fighthiddenbosscd"]),
	}

	rm.gameConfig = node
	rm.logger.Debug("loadGameConfig read finish:", zap.Any("row", node))
	return nil
}

func (rm *gameResourceMgr) getEnergyLimit() uint32 {
	return rm.gameConfig.GetEnergyLimit()
}

func (rm *gameResourceMgr) getEnergyRecoverTime() int64 {
	return rm.gameConfig.GetEnergyRecoverTime()
}

func (rm *gameResourceMgr) getFightHiddenBossCd() int64 {
	return rm.gameConfig.GetFightHiddenBossCd()
}

func (rm *gameResourceMgr) loadPlayerInitAttrs(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	if len(datas) != 1 {
		rm.logger.Error(fmt.Sprintf("load %s failed: config row num %d", csvPath, len(datas)))
		return mpberr.ErrConfig
	}
	node := &mpb.PlayerInitAttrs{
		Attrs: com.ReadAttrs(datas[0]),
	}

	rm.logger.Debug("loadPlayerInitAttrs read finish:", zap.Any("row", node))
	rm.playerInitAttrs = node
	return nil
}

func (rm *gameResourceMgr) getPlayerInitAttrs() *mpb.Attrs {
	return com.CloneAttrs(rm.playerInitAttrs.Attrs)
}

func (rm *gameResourceMgr) loadBoss(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32]*mpb.BossRsc)
	cm := make(map[uint32][]*mpb.BossRsc)
	for _, data := range datas {
		node := &mpb.BossRsc{
			BossId:           gutil.StrToUint32(data["bossid"]),
			BossType:         mpb.ERole_BossType(gutil.StrToUint32(data["bosstype"])),
			Name:             data["name"],
			Photo:            data["photo"],
			Class:            gutil.StrToUint32(data["class"]),
			Level:            gutil.StrToUint32(data["level"]),
			LevelShow:        data["levelshow"],
			LiveTime:         gutil.StrToInt64(data["livetime"]),
			PreBoss:          gutil.StrToUint32(data["preboss"]),
			NftEquips:        util.ReadUint32Slice(data["nftequips"], ";"),
			NftEquipsLevel:   gutil.StrToUint32(data["nftequipslevel"]),
			EnergyCost:       gutil.StrToUint32(data["energycost"]),
			Attrs:            com.ReadAttrs(data),
			WinDmgRate:       gutil.StrToUint64(data["windmgrate"]),
			DmgAwardsCoe1:    gutil.StrToUint64(data["dmgawardscoe1"]),
			DmgAwardsCoe2:    gutil.StrToUint64(data["dmgawardscoe2"]),
			BuffCardRandPool: gutil.StrToUint32(data["buffcardrandpool"]),
		}
		node.FirstWinAwards, err = com.ReadAwardsRsc(data["firstwinawards"])
		if err != nil {
			rm.logger.Error("loadBoss parse firstwinawards failed",
				zap.String("firstwinawards", data["firstwinawards"]))
			return err
		}
		node.Awards, err = com.ReadAwardsRsc(data["awards"])
		if err != nil {
			rm.logger.Error("loadBoss parse awards failed",
				zap.String("awards", data["awards"]))
			return err
		}
		node.FinderAwards, err = com.ReadAwardsRsc(data["finderawards"])
		if err != nil {
			rm.logger.Error("loadBoss parse finderawards failed",
				zap.String("finderawards", data["finderawards"]))
			return err
		}
		node.KillerAwards, err = com.ReadAwardsRsc(data["killerawards"])
		if err != nil {
			rm.logger.Error("loadBoss parse killerawards failed",
				zap.String("killerawards", data["killerawards"]))
			return err
		}
		node.DmgAwards, err = com.ReadAwardsRsc(data["dmgawards"])
		if err != nil {
			rm.logger.Error("loadBoss parse dmgawards failed",
				zap.String("dmgawards", data["dmgawards"]))
			return err
		}
		if len(data["manaawards"]) > 0 {
			manas := gcsv.StrToUint32Slice(data["manaawards"], ":")
			if len(manas) != 2 || manas[0] < manas[1] {
				rm.logger.Error("loadBoss parse manaawards failed",
					zap.String("manaawards", data["manaawards"]))
			}
			node.ManaAwardsBossFight = uint64(manas[0]) * com.ManaFactor
			node.ManaAwardsPvpPool = uint64(manas[1]) * com.ManaFactor
		}
		m[node.BossId] = node
		rm.logger.Debug("loadBoss read:", zap.Any("row", node))
	}

	rm.bosses = m
	rm.bossClassesMap = cm
	rm.logger.Debug("loadBoss read finish:", zap.Any("rows", m))

	return nil
}

func (rm *gameResourceMgr) getBossClassNextLevelBossRsc(bossId uint32) *mpb.BossRsc {
	cb := rm.bosses[bossId]
	if cb == nil {
		return nil
	}
	var b *mpb.BossRsc
	for _, v := range rm.bossClassesMap[cb.Class] {
		if v.Level <= cb.Level {
			continue
		}
		if b == nil {
			b = v
		} else if b.Level > v.Level {
			b = v
		}
	}
	return b
}

func (rm *gameResourceMgr) getBossRsc(bossId uint32) *mpb.BossRsc {
	return rm.bosses[bossId]
}

func (rm *gameResourceMgr) loadHiddenBossTrigger(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32]*mpb.HiddenBossTriggerRsc)
	for _, data := range datas {
		node := &mpb.HiddenBossTriggerRsc{
			BossId:       gutil.StrToUint32(data["bossid"]),
			TriggerRate:  gutil.StrToUint32(data["triggerrate"]),
			HiddenBossId: gutil.StrToUint32(data["hiddenbossid"]),
		}

		m[node.BossId] = node
		rm.logger.Debug("loadHiddenBossTrigger read:", zap.Any("row", node))
	}

	rm.hiddenBossTrigger = m
	rm.logger.Debug("loadHiddenBossTrigger read finish:", zap.Any("rows", m))
	return nil
}

func (rm *gameResourceMgr) getHiddenBossTriggerRsc(bossId uint32) *mpb.HiddenBossTriggerRsc {
	return rm.hiddenBossTrigger[bossId]
}

func (rm *gameResourceMgr) loadPVPSchedule(csvPath string) error {
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

func (rm *gameResourceMgr) isPVPSettle(now time.Time) (uint32, string, bool) {
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

func (rm *gameResourceMgr) getPVPSeason(nowUnix int64) uint32 {
	for _, v := range rm.pvpSchedules {
		if v.StartTime <= nowUnix && nowUnix <= v.EndTime {
			return v.PvpSeasonId
		}
	}
	return 0
}

func (rm *gameResourceMgr) getPVPSeasonDate(now time.Time) (uint32, uint32) {
	now = now.In(util.TimeZone_GMT)
	nowUnix := now.Unix()
	var rsc *mpb.PVPScheduleRsc
	for _, v := range rm.pvpSchedules {
		if v.StartTime <= nowUnix && nowUnix <= v.EndTime {
			rsc = v
			break
		}
	}
	if rsc == nil {
		return 0, 0
	}

	settleStartAt := rsc.SettleStartAt
	settleEndAt := rsc.SettleEndAt

	zeroUnix := gtime.TodayXHourTimeUnix(0, util.TimeZone_GMT)
	nowUnix = nowUnix - zeroUnix

	if settleStartAt <= settleEndAt {
		if nowUnix >= settleEndAt {
			dateStr := gtime.TimeToDate(now.Add(24 * time.Hour)) // use next day date
			date := gutil.StrToUint32(strings.ReplaceAll(dateStr, "-", ""))
			return rsc.PvpSeasonId, date
		}
		dateStr := gtime.TimeToDate(now) // use next day date
		date := gutil.StrToUint32(strings.ReplaceAll(dateStr, "-", ""))
		return rsc.PvpSeasonId, date
	}

	if nowUnix <= settleEndAt {
		dateStr := gtime.TimeToDate(now.Add(-24 * time.Hour)) // use next day date
		date := gutil.StrToUint32(strings.ReplaceAll(dateStr, "-", ""))
		return rsc.PvpSeasonId, date
	}

	dateStr := gtime.TimeToDate(now) // use next day date
	date := gutil.StrToUint32(strings.ReplaceAll(dateStr, "-", ""))

	return rsc.PvpSeasonId, date
}

func (rm *gameResourceMgr) isPVPBattle(now time.Time) (uint32, uint32, bool) {
	now = now.In(util.TimeZone_GMT)
	nowUnix := now.Unix()
	var rsc *mpb.PVPScheduleRsc
	for _, v := range rm.pvpSchedules {
		if v.StartTime <= nowUnix && nowUnix <= v.EndTime {
			rsc = v
			break
		}
	}
	if rsc == nil {
		return 0, 0, false
	}
	settleStartAt := rsc.SettleStartAt
	settleEndAt := rsc.SettleEndAt

	zeroUnix := gtime.TodayXHourTimeUnix(0, util.TimeZone_GMT)
	nowUnix = nowUnix - zeroUnix

	if settleStartAt <= settleEndAt {
		if nowUnix < settleStartAt {
			dateStr := gtime.TimeToDate(now)
			date := gutil.StrToUint32(strings.ReplaceAll(dateStr, "-", ""))
			return rsc.PvpSeasonId, date, true
		}
		if nowUnix <= settleEndAt {
			return 0, 0, false
		}
		dateStr := gtime.TimeToDate(now.Add(24 * time.Hour)) // use next day date
		date := gutil.StrToUint32(strings.ReplaceAll(dateStr, "-", ""))
		return rsc.PvpSeasonId, date, true
	}
	if nowUnix <= settleEndAt {
		return 0, 0, false
	}
	if nowUnix < settleStartAt {
		dateStr := gtime.TimeToDate(now)
		date := gutil.StrToUint32(strings.ReplaceAll(dateStr, "-", ""))
		return rsc.PvpSeasonId, date, true
	}
	return 0, 0, false
}

func (rm *gameResourceMgr) loadPVPRankRewards(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32][]*mpb.PVPRankRewardsRsc)
	for _, data := range datas {
		node := &mpb.PVPRankRewardsRsc{
			PvpSeasonId:         gcsv.StrToUint32(data["pvpseasonid"]),
			RankRange:           gcsv.StrToUint32Slice(data["rankrange"], ":"),
			ManaRewardsPoolRate: gcsv.StrToUint32(data["manarewardspoolrate"]),
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

func (rm *gameResourceMgr) getPVPSeasonRankRewardsRscs(seasonId uint32, rank uint32) ([]*mpb.AwardRsc, uint32) {
	seasonRewards := rm.pvpRankRewards[seasonId]
	for _, v := range seasonRewards {
		if v.RankRange[0] <= rank && rank <= v.RankRange[1] {
			return v.Rewards, v.ManaRewardsPoolRate
		}
	}
	return nil, 0
}

func (rm *gameResourceMgr) loadPVPConfig(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32]*mpb.PVPConfigRsc)
	for _, data := range datas {
		node := &mpb.PVPConfigRsc{
			PvpSeasonId:              gcsv.StrToUint32(data["pvpseasonid"]),
			DailyPvpFreeChallengeCnt: gutil.StrToUint32(data["dailypvpfreechallengecnt"]),
			PvpManaConsume:           gutil.StrToUint32(data["pvpmanaconsume"]) * com.ManaFactor,
			ManaAwardsPool:           gutil.StrToUint64(data["manaawardspool"]) * com.ManaFactor,
		}
		if len(data["rankranges"]) > 0 {
			for _, rangeStr := range strings.Split(data["rankranges"], ";") {
				r := util.ReadUint32Range(rangeStr, ":")
				if r == nil {
					rm.logger.Error("loadPVPConfig parse rank range failed",
						zap.String("rankranges", data["rankranges"]))
					return mpberr.ErrConfig
				}
				node.RankRanges = append(node.RankRanges, r)
			}
		}
		node.ChallengeTarget1Coe = util.ReadFloat64Range(data["challengetarget1coe"], ":")
		if node.ChallengeTarget1Coe == nil {
			rm.logger.Error("loadPVPConfig parse challengetarget1coe failed",
				zap.String("challengetarget1coe", data["challengetarget1coe"]))
			return mpberr.ErrConfig
		}

		node.ChallengeTarget2Coe = util.ReadFloat64Range(data["challengetarget2coe"], ":")
		if node.ChallengeTarget2Coe == nil {
			rm.logger.Error("loadPVPConfig parse challengetarget2coe failed",
				zap.String("challengetarget2coe", data["challengetarget2coe"]))
			return mpberr.ErrConfig
		}

		node.ChallengeTarget3Coe = util.ReadFloat64Range(data["challengetarget3coe"], ":")
		if node.ChallengeTarget3Coe == nil {
			rm.logger.Error("loadPVPConfig parse challengetarget3coe failed",
				zap.String("challengetarget3coe", data["challengetarget3coe"]))
			return mpberr.ErrConfig
		}

		m[node.PvpSeasonId] = node
		rm.logger.Debug("loadPVPConfig read:", zap.Any("row", node))
	}

	rm.logger.Debug("loadPVPConfig read finish:", zap.Any("rows", m))

	rm.pvpConfig = m
	return nil
}

func (rm *gameResourceMgr) getDailyPVPFreeChallengeCnt(seasonId uint32) uint32 {
	return rm.pvpConfig[seasonId].GetDailyPvpFreeChallengeCnt()
}

func (rm *gameResourceMgr) getPVPManaConsume(seasonId uint32) uint32 {
	return rm.pvpConfig[seasonId].GetPvpManaConsume()
}

func (rm *gameResourceMgr) getPVPConfigRsc(seasonId uint32) *mpb.PVPConfigRsc {
	return rm.pvpConfig[seasonId]
}

func (rm *gameResourceMgr) getChallengeTargetRank(rank, targetNum uint32, pvpConfigRsc *mpb.PVPConfigRsc) (uint32, bool) {
	rank++
	switch targetNum {
	case 1:
		newRank := uint32(float64(rank) * util.RandomInFloat64Range(pvpConfigRsc.ChallengeTarget1Coe))
		return newRank - 1, newRank > 0
	case 2:
		newRank := uint32(float64(rank) * util.RandomInFloat64Range(pvpConfigRsc.ChallengeTarget2Coe))
		return newRank - 1, newRank > 0
	case 3:
		newRank := uint32(float64(rank) * util.RandomInFloat64Range(pvpConfigRsc.ChallengeTarget3Coe))
		return newRank - 1, newRank > 0
	case 4:
		if rank == 1 {
			return 0, false
		}
		for _, r := range pvpConfigRsc.RankRanges {
			if rank > r.Max || rank < r.Min {
				continue
			}
			return (r.Min - 1) - 1, true
		}
	}
	return 0, false
}

func (rm *gameResourceMgr) loadBuffCard(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32]*mpb.BuffCardRsc)
	tm := make(map[mpb.EGame_BuffCardType][]*mpb.BuffCardRsc)
	for _, data := range datas {
		node := &mpb.BuffCardRsc{
			CardId:    gcsv.StrToUint32(data["cardid"]),
			CardType:  mpb.EGame_BuffCardType(gcsv.StrToUint32(data["cardtype"])),
			CardLevel: gutil.StrToUint32(data["cardlevel"]),
			Round:     gcsv.StrToUint32(data["round"]),
		}

		switch node.CardType {
		case mpb.EGame_BuffCardType_1:
			node.AtcAdd = gcsv.StrToUint64(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_2:
			node.DefenseAdd = gcsv.StrToUint64(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_3:
			node.CriRateAdd = gcsv.StrToUint64(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_4:
			node.DodgeRateAdd = gcsv.StrToUint64(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_5:
			node.CriDmgAdd = gcsv.StrToUint64(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_6:
			node.DodgeAtk = gcsv.StrToUint64(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_7:
			node.BuffRoundAdd = gcsv.StrToUint64(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_8:
			us := gcsv.StrToUint32Slice(data["buffaffectvalues"], ";")
			if len(us) != 2 {
				rm.logger.Error("loadBuffCard parse buffaffectvalues failed",
					zap.Uint32("card_id", node.CardId), zap.Any("card_type", node.CardType),
					zap.String("buffaffectvalues", data["buffaffectvalues"]))
				return mpberr.ErrConfig
			}
			node.CriAwardsAdd = uint64(us[0])
			node.TriggerCntPerRound = us[1]
		case mpb.EGame_BuffCardType_9:
			us := gcsv.StrToUint32Slice(data["buffaffectvalues"], ";")
			if len(us) != 3 {
				rm.logger.Error("loadBuffCard parse buffaffectvalues failed",
					zap.Uint32("card_id", node.CardId), zap.Any("card_type", node.CardType),
					zap.String("buffaffectvalues", data["buffaffectvalues"]))
				return mpberr.ErrConfig
			}
			node.DodgeStealAtk = uint64(us[0])
			node.TriggerCntPerRound = us[1]
			node.DodgeAwardsAdd = uint64(us[2])
		case mpb.EGame_BuffCardType_10:
			node.BuffCardRandPool = gcsv.StrToUint32(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_11:
			node.BuffCardRandPool = gcsv.StrToUint32(data["buffaffectvalues"])
		case mpb.EGame_BuffCardType_12:
			us := gcsv.StrToUint32Slice(data["buffaffectvalues"], ";")
			if len(us) != 2 {
				rm.logger.Error("loadBuffCard parse buffaffectvalues failed",
					zap.Uint32("card_id", node.CardId), zap.Any("card_type", node.CardType),
					zap.String("buffaffectvalues", data["buffaffectvalues"]))
				return mpberr.ErrConfig
			}
			node.BuffCardRandPool = us[0]
			node.BuffLevelAdd = us[1]
		}

		m[node.CardId] = node
		tm[node.CardType] = append(tm[node.CardType], node)
		rm.logger.Debug("loadBuffCard read:", zap.Any("row", node))
	}

	rm.logger.Debug("loadBuffCard read finish:", zap.Any("rows", m))

	rm.buffCards = m
	rm.buffCardTypeMap = tm
	return nil
}

func (rm *gameResourceMgr) getBuffCardRsc(buffCardId uint32) *mpb.BuffCardRsc {
	return rm.buffCards[buffCardId]
}

func (rm *gameResourceMgr) loadBuffCardRandPool(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[uint32]*mpb.BuffCardRandPoolRsc)
	for _, data := range datas {
		node := &mpb.BuffCardRandPoolRsc{
			PoolId: gcsv.StrToUint32(data["poolid"]),
		}

		node.BuffCards, err = util.ReadWeightRandomPoolFromString(data["rates"])
		if err != nil {
			rm.logger.Error("loadBuffCardRandPool parse pool rate failed",
				zap.Uint32("pool_id", node.PoolId),
				zap.String("rates", data["rates"]))
			return mpberr.ErrConfig
		}

		m[node.PoolId] = node
		rm.logger.Debug("loadBuffCardRandPool read:", zap.Any("row", node))
	}

	rm.logger.Debug("loadBuffCardRandPool read finish:", zap.Any("rows", m))

	rm.buffCardRandPools = m
	return nil
}

func (rm *gameResourceMgr) getBuffCardRandPool(poolId uint32) *mpb.BuffCardRandPoolRsc {
	return rm.buffCardRandPools[poolId]
}

func (rm *gameResourceMgr) upgradeBuffCard(buffCard uint32) uint32 {
	buffCardRsc := rm.getBuffCardRsc(buffCard)
	if buffCardRsc == nil {
		return buffCard
	}
	for _, v := range rm.buffCardTypeMap[buffCardRsc.CardType] {
		if v.CardLevel == buffCardRsc.CardLevel+1 {
			return v.CardId
		}
	}

	return buffCard
}
