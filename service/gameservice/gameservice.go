package gameservice

import (
	"context"
	"errors"
	"fmt"
	gtime "github.com/oldjon/gutil/timeutil"
	"time"

	"github.com/oldjon/gutil"
	"github.com/oldjon/gutil/env"
	"github.com/oldjon/gutil/gdb"
	gprotocol "github.com/oldjon/gutil/protocol"
	grmux "github.com/oldjon/gutil/redismutex"
	gxgrpc "github.com/oldjon/gx/modules/grpc"
	"github.com/oldjon/gx/service"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"gitlab.com/morbackend/mor_services/util"
	etcd "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	buffCardCnt = 3
)

type GameService struct {
	mpb.UnimplementedGameServiceServer
	name        string
	logger      *zap.Logger
	config      env.ModuleConfig
	etcdClient  *etcd.Client
	host        service.Host
	connMgr     *gxgrpc.ConnManager
	tcpMsgCoder gprotocol.FrameCoder
	rm          *gameResourceMgr
	kvm         *service.KVMgr
	serverEnv   uint32
	sm          *util.ServiceMetrics
	dao         *gameDAO
	glm         *gameLevelMgr
	bossUUIDSF  service.Snowflake
	gcm         *gameCacheMgr
	bcm         *buffCardMgr
}

// NewGameService create a GameService entity
func NewGameService(driver service.ModuleDriver) (gxgrpc.GRPCServer, error) {
	svc := &GameService{
		name:       driver.ModuleName(),
		logger:     driver.Logger(),
		config:     driver.ModuleConfig(),
		etcdClient: driver.Host().EtcdSession().Client(),
		host:       driver.Host(),
		kvm:        driver.Host().KVManager(),
		sm:         util.NewServiceMetrics(driver),
		gcm:        newGameCacheMgr(),
	}

	dialer := gxgrpc.Dialer{
		HostName:   driver.Host().Name(),
		Tracer:     driver.Tracer(),
		EtcdClient: svc.etcdClient,
		Logger:     svc.logger,
		EnableTLS:  svc.config.GetBool("enable_tls"),
		CAFile:     svc.config.GetString("ca_file"),
		CertFile:   svc.config.GetString("cert_file"),
		KeyFile:    svc.config.GetString("key_file"),
	}
	svc.connMgr = gxgrpc.NewConnManager(&dialer)
	svc.glm = newGameLevelMgr(svc)

	var err error
	svc.rm, err = newGameResourceMgr(svc.logger, svc.sm)
	if err != nil {
		return nil, err
	}

	redisMux, err := grmux.NewRedisMux(svc.config.SubConfig("redis_mutex"), nil, svc.logger, driver.Tracer())
	if err != nil {
		return nil, err
	}

	gameRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("game_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	pvpRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("pvp_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	tmpRedis, err := gdb.NewRedisClientByConfig(svc.config.SubConfig("tmp_redis"),
		svc.config.GetString("db_marshaller"), driver.Tracer())
	if err != nil {
		return nil, err
	}

	svc.dao = newGameDAO(svc, redisMux, gameRedis, pvpRedis, tmpRedis)

	svc.serverEnv = uint32(svc.config.GetInt64("server_env"))
	svc.tcpMsgCoder = gprotocol.NewFrameCoder(svc.config.GetString("protocol_code"))

	svc.bcm = newBuffCardMgr(svc)

	return svc, err
}

func (svc *GameService) Register(grpcServer *grpc.Server) {
	mpb.RegisterGameServiceServer(grpcServer, svc)
}

func (svc *GameService) Serve(ctx context.Context) error {
	var err error
	svc.bossUUIDSF, err = svc.host.Snowflake(ctx, com.SnowflakeBossUUID, service.SnowflakeType_53)
	if err != nil {
		svc.logger.Error("failed to create snowflake", zap.Error(err))
		return err
	}

	<-ctx.Done()
	return ctx.Err()
}

func (svc *GameService) Logger() *zap.Logger {
	return svc.logger
}

func (svc *GameService) ConnMgr() *gxgrpc.ConnManager {
	return svc.connMgr
}

func (svc *GameService) Name() string {
	return svc.name
}

func (svc *GameService) Fight(ctx context.Context, req *mpb.ReqFight) (*mpb.ResFight, error) {
	// check energy
	now := time.Now()

	bossRsc := svc.rm.getBossRsc(req.BossId)
	if bossRsc == nil {
		return nil, mpberr.ErrParam
	}

	if bossRsc.PreBoss > 0 { // check pre boss
		fightHis, err := svc.dao.getFightHistory(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		if fightHis.WinTimes[bossRsc.PreBoss] == 0 {
			return nil, mpberr.ErrParam
		}
	}

	var res *mpb.ResFight
	var err error

	if bossRsc.BossType != mpb.ERole_BossType_Hidden {
		res, err = svc.fightBoss(ctx, req, bossRsc, now)
		if err != nil {
			return nil, err
		}
	} else { // normal boss or nft boss
		if req.BossUuid == 0 {
			return nil, mpberr.ErrParam
		}
		res, err = svc.fightHiddenBoss(ctx, req, bossRsc, now)
		if err != nil {
			return nil, err
		}
	}

	err = svc.dao.updateBuffCardChoseStatus(ctx, req.UserId, uint32(mpb.EGame_BuffCardStatus_None))
	if err != nil {
		return nil, err
	}

	res.BuffCardStatus = mpb.EGame_BuffCardStatus_None

	if !res.Win {
		return res, nil
	}

	// trigger hidden boss
	res.HiddenBoss, err = svc.triggerHiddenBoss(ctx, req.UserId, req.BossId, now)
	if err != nil {
		return nil, err
	}

	if req.WithBossDetail {
		res.BossRsc = bossRsc
	}

	return res, nil
}

func (svc *GameService) fightBoss(ctx context.Context, req *mpb.ReqFight, bossRsc *mpb.BossRsc, now time.Time) (
	*mpb.ResFight, error) {
	dbEnergy, err := svc.dao.getEnergy(ctx, req.UserId, now.Unix())
	if err != nil {
		return nil, err
	}
	if dbEnergy.Energy < bossRsc.EnergyCost {
		return nil, mpberr.ErrEnergyNotEnough
	}

	curBuffCards, leftBuffCards, err := svc.dao.useBuffCardOneRound(ctx, req.UserId)

	equips, err := svc.getUserEquips(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if len(bossRsc.NftEquips) > 0 {
		var totalLevel uint32
		for _, equipId := range bossRsc.NftEquips {
			var equip *mpb.NFTEquip
			for _, v := range equips.NftEquips {
				if v.EquipId == equipId {
					equip = v
					break
				}
			}
			if equip == nil {
				return nil, mpberr.ErrNFTWeaponNotEquipped
			}
			totalLevel += equip.Level
		}
		if totalLevel < bossRsc.NftEquipsLevel {
			return nil, mpberr.ErrBaseEquipMaxLevel
		}
	}

	boss := newGameBoss(bossRsc)
	player := newPlayer(svc, req.UserId, mpb.ERole_RoleType_Player, svc.bcm.newBuffCardsByIds(req.UserId, curBuffCards))
	player.updateEquips(equips.BaseEquips, equips.NftEquips)
	svc.bcm.effectBeforeFight(player)
	res := &mpb.ResFight{
		BossHp:    boss.totalHP,
		PlayerHp:  player.totalHP,
		BuffCards: leftBuffCards,
	}
	var awardAddRate uint64
	gl := svc.glm.newGameLevel(player, boss)
	res.Win, res.Details, res.Dmg, res.DmgRate, res.BossDie, awardAddRate = gl.fight(player, boss)
	if !res.Win {
		return res, nil
	}
	// win
	// cost energy
	dbEnergy, err = svc.dao.consumeEnergy(ctx, req.UserId, bossRsc.EnergyCost, now.Unix())
	if err != nil {
		return nil, err
	}
	res.EnergyCost = bossRsc.EnergyCost
	res.Energy = dbEnergy.Energy
	res.EnergyRecoverAt = dbEnergy.RecoverAt
	var awards []*mpb.AwardRsc
	// update win times
	if res.BossDie {
		dbFightHis, err := svc.dao.updateFightHistory(ctx, req.UserId, req.BossId)
		if err != nil {
			return nil, err
		}
		if dbFightHis.WinTimes[req.BossId] == 1 { // first win and boss died
			awards = bossRsc.FirstWinAwards
		}
	}
	// give awards
	for _, v := range bossRsc.Awards {
		award := &mpb.AwardRsc{
			ItemId: v.ItemId,
		}
		award.Num = uint64(v.NumRange[1]-v.NumRange[0]) * (boss.totalHP - boss.hp - boss.winLoseHp) /
			(boss.totalHP - boss.winLoseHp)
		award.Num += v.NumRange[0]
		awards = append(awards, award)
	}

	if bossRsc.ManaAwardsBossFight > 0 {
		awards = append(awards, &mpb.AwardRsc{
			ItemId: uint32(mpb.EItem_ItemId_Mana),
			Num:    bossRsc.ManaAwardsBossFight,
		})
	}

	if bossRsc.ManaAwardsPvpPool > 0 {
		if season, date, ok := svc.rm.isPVPBattle(now); ok {
			err = svc.dao.putManaIntoPVPManaPool(ctx, season, date, bossRsc.ManaAwardsPvpPool)
		}
	}
	if res.BossDie {
		err = svc.dao.updateBossDefeatHistory(ctx, req.UserId, bossRsc)
		if err != nil {
			return nil, err
		}
	}

	for _, v := range awards {
		if com.GetItemType(v.ItemId) == mpb.EItem_Type_BaseEquipStarUpgradeMaterial ||
			com.GetItemType(v.ItemId) == mpb.EItem_Type_BaseEquipLevelUpgradeMaterial {
			v.Num = (v.Num) * (com.RateBase + awardAddRate) / com.RateBase
		}
	}

	res.Awards, err = com.AddItemsFromAwardRsc(ctx, svc, req.UserId, [][]*mpb.AwardRsc{awards},
		mpb.EItem_TransReason_GameFight, uint64(boss.bossId))
	if err != nil {
		svc.logger.Error("fight boss give awards failed", zap.Uint64("user_id", req.UserId),
			zap.Any("awards", awards), zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (svc *GameService) fightHiddenBoss(ctx context.Context, req *mpb.ReqFight, bossRsc *mpb.BossRsc, now time.Time) (
	*mpb.ResFight, error) {

	nowUnix := now.Unix()
	// check fight cd
	if svc.inHiddenBossFightCD(ctx, req.UserId, nowUnix) {
		return nil, mpberr.ErrHiddenBossFightCD
	}

	equips, err := svc.getUserEquips(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if len(bossRsc.NftEquips) > 0 {
		var totalLevel uint32
		for _, equipId := range bossRsc.NftEquips {
			var equip *mpb.NFTEquip
			for _, v := range equips.NftEquips {
				if v.EquipId == equipId {
					equip = v
					break
				}
			}
			if equip == nil {
				return nil, mpberr.ErrNFTWeaponNotEquipped
			}
			totalLevel += equip.Level
		}
		if totalLevel < bossRsc.NftEquipsLevel {
			return nil, mpberr.ErrNFTWeaponLevel
		}
	}

	curBuffCards, leftBuffCards, err := svc.dao.useBuffCardOneRound(ctx, req.UserId)

	boss := newGameBoss(bossRsc)
	player := newPlayer(svc, req.UserId, mpb.ERole_RoleType_Player, svc.bcm.newBuffCardsByIds(req.UserId, curBuffCards))
	player.updateEquips(equips.BaseEquips, equips.NftEquips)
	svc.bcm.effectBeforeFight(player)
	res := &mpb.ResFight{
		BossHp:    boss.totalHP,
		PlayerHp:  player.totalHP,
		BuffCards: leftBuffCards,
	}
	var dbBoss *mpb.DBHiddenBoss
	err = svc.dao.fightHiddenBoss(ctx, req.BossUuid, func() error {
		dbBoss, err = svc.dao.getHiddenBoss(ctx, req.BossUuid)
		if err != nil {
			return err
		}
		if dbBoss.ExpiredAt < nowUnix {
			return mpberr.ErrHiddenBossExpired
		}
		if dbBoss.Hp == 0 {
			return mpberr.ErrHiddenBossDied
		}
		if dbBoss.LastFightTime[req.UserId] > 0 {
			return mpberr.ErrHiddenBossFought
		}
		boss.hp = dbBoss.Hp
		// fight
		var awardAddRate uint64
		gl := svc.glm.newGameLevel(player, boss)
		res.Win, res.Details, res.Dmg, res.DmgRate, res.BossDie, awardAddRate = gl.fight(player, boss)
		// update dbBoss
		dbBoss.Hp = boss.hp
		if dbBoss.Hp == 0 {
			dbBoss.Killer = boss.killedBy
		}
		dbBoss.LastFightTime[req.UserId] = nowUnix
		if res.Dmg >= boss.totalHP/100 {
			dbBoss.Dmgs[req.UserId] = res.Dmg
			dbBoss.AwardAddRates[req.UserId] = awardAddRate
			acc, err := svc.getAccountByUserId(ctx, req.UserId)
			if err != nil {
				return err
			}
			dbBoss.FightHistories = append(dbBoss.FightHistories, &mpb.DBHiddenBossFightHistory{
				Nickname: acc.Nickname,
				DmgRate:  res.DmgRate,
			})
		}

		err = svc.dao.updateHiddenBoss(ctx, dbBoss, com.KeepTTL)
		if err != nil {
			svc.logger.Error("fightHiddenBoss update boss failed", zap.Uint64("user_id", req.UserId),
				zap.Any("boss", dbBoss), zap.Error(err))
		}
		return nil
	})
	if err != nil {
		svc.logger.Error("fightHiddenBoss failed", zap.Uint64("user_id", req.UserId),
			zap.Uint32("boss_id", bossRsc.BossId), zap.Uint64("boss_uuid", req.BossUuid),
			zap.Error(err))
		return nil, err
	}
	// cost energy
	dbEnergy, err := svc.dao.consumeEnergy(ctx, req.UserId, bossRsc.EnergyCost, nowUnix)
	if err != nil {
		return nil, err
	}
	res.EnergyCost = bossRsc.EnergyCost
	res.Energy = dbEnergy.Energy
	res.EnergyRecoverAt = dbEnergy.RecoverAt

	err = svc.dao.updateHiddenBossFightCD(ctx, req.UserId, svc.rm.getFightHiddenBossCd()+nowUnix)
	if err != nil {
		return nil, err
	}

	if boss.hp > 0 { // boss still alive
		return res, nil
	}

	_ = svc.dao.delFromHiddenBossPool(ctx, []uint64{dbBoss.BossUuid})

	// boss die
	awards := make(map[uint64][][]*mpb.AwardRsc)
	// 1.finder awards
	awards[dbBoss.Finder] = [][]*mpb.AwardRsc{bossRsc.FinderAwards}
	// 2.killer awards
	awards[dbBoss.Killer] = [][]*mpb.AwardRsc{bossRsc.KillerAwards}
	// 3.dmg awards
	for uid, dmg := range dbBoss.Dmgs {
		dmgAwards := make([]*mpb.AwardRsc, 0, 1)
		for _, v := range bossRsc.DmgAwards {
			a := &mpb.AwardRsc{
				ItemId: v.ItemId,
				Num: uint64((bossRsc.DmgAwardsCoe1*uint64(v.Num)*dmg +
					bossRsc.DmgAwardsCoe2*uint64(v.Num)*boss.totalHP +
					boss.totalHP*com.RateBase/2) / // for round
					(boss.totalHP * com.RateBase)),
			}
			a.Num += uint64(a.Num) * (com.RateBase + dbBoss.AwardAddRates[uid]) / com.RateBase
			dmgAwards = append(dmgAwards, a)
		}
		uAwards := awards[uid]
		if len(uAwards) == 0 {
			uAwards = append(uAwards, [][]*mpb.AwardRsc{dmgAwards}...)
		} else {
			uAwards[0] = append(uAwards[0], dmgAwards...)
		}
		awards[uid] = uAwards
	}
	cAwards, err := com.BatchAddItemsFromAwardRsc(ctx, svc, awards, mpb.EItem_TransReason_GameFight, req.BossUuid)
	if err != nil {
		svc.logger.Error("fightHiddenBoss give awards failed", zap.Uint64("user_id", req.UserId),
			zap.Any("awards", awards), zap.Error(err))
		return nil, err
	}
	res.Awards = cAwards[req.UserId]
	return res, nil
}

func (svc *GameService) getUserEquips(ctx context.Context, userId uint64) (*mpb.ResGetEquips, error) {
	itemClient, err := com.GetItemServiceClient(ctx, svc)
	if err != nil {
		return nil, err
	}
	res, err := itemClient.GetEquips(ctx, &mpb.ReqUserId{
		UserId: userId,
	})
	if err != nil {
		svc.logger.Error("getUserEquips get user equips failed", zap.Uint64("user_id", userId),
			zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (svc *GameService) batchGetUsersEquips(ctx context.Context, userIds []uint64) (map[uint64]*mpb.UserEquips, error) {
	itemClient, err := com.GetItemServiceClient(ctx, svc)
	if err != nil {
		return nil, err
	}
	res, err := itemClient.BatchGetEquips(ctx, &mpb.ReqUserIds{
		UserIds: userIds,
	})
	if err != nil {
		svc.logger.Error("getUserEquips get user equips failed", zap.Any("user_ids", userIds),
			zap.Error(err))
		return nil, err
	}
	if res.Equips == nil {
		res.Equips = make(map[uint64]*mpb.UserEquips)
	}
	return res.Equips, nil
}

func (svc *GameService) triggerHiddenBoss(ctx context.Context, userId uint64, bossId uint32, now time.Time) (
	hiddenBoss *mpb.HiddenBoss, err error) {
	nowUnix := now.Unix()
	triggerRsc := svc.rm.getHiddenBossTriggerRsc(bossId)
	if triggerRsc == nil {
		return nil, nil
	}

	if !util.IsPick(triggerRsc.TriggerRate, com.RateBase) {
		return nil, nil
	}

	bossRsc := svc.rm.getBossRsc(triggerRsc.HiddenBossId)
	if bossRsc == nil {
		return nil, nil
	}

	// check whether user can trigger new boss or not
	dbFindHis, err := svc.dao.getHiddenBossFindHistory(ctx, userId)
	if err != nil {
		return nil, err
	}

	for bossUUID, expireAt := range dbFindHis.BossExpireAt {
		if expireAt < nowUnix {
			delete(dbFindHis.BossExpireAt, bossUUID)
			continue
		}
		dbBoss, err := svc.dao.getHiddenBoss(ctx, bossUUID)
		if err != nil {
			return nil, err
		}
		curBossRsc := svc.rm.getBossRsc(dbBoss.BossId)
		if curBossRsc == nil {
			return nil, mpberr.ErrConfig
		}
		boss := newGameBoss(curBossRsc)
		if dbBoss.Hp > (boss.totalHP - boss.winLoseHp) {
			return nil, nil
		}
		delete(dbFindHis.BossExpireAt, bossUUID)
	}

	// new a hidden boss
	bossUUID, err := svc.bossUUIDSF.Next()
	if err != nil {
		return nil, err
	}

	boss := newGameBoss(bossRsc)
	dbBoss := &mpb.DBHiddenBoss{
		BossUuid:  bossUUID,
		BossId:    bossRsc.BossId,
		Finder:    userId,
		Hp:        boss.hp,
		ExpiredAt: bossRsc.LiveTime + nowUnix,
	}
	err = svc.dao.updateHiddenBoss(ctx, dbBoss, time.Duration(bossRsc.LiveTime+com.Secs1Week)*time.Second)
	if err != nil {
		return nil, err
	}

	// update find history
	dbFindHis.BossExpireAt[bossUUID] = dbBoss.ExpiredAt
	err = svc.dao.updateHiddenBossFindHistory(ctx, userId, dbFindHis)
	if err != nil {
		return nil, err
	}

	// push new boss to pool
	err = svc.dao.pushIntoHiddenBossPool(ctx, dbBoss.BossUuid, dbBoss.ExpiredAt)
	if err != nil {
		return nil, err
	}
	resBoss, _ := svc.dbHiddenBoss2HiddenBoss(dbBoss, false)
	return resBoss, nil
}

func (svc *GameService) inHiddenBossFightCD(ctx context.Context, userId uint64, nowUnix int64) bool {
	cd, err := svc.dao.getHiddenBossFightCD(ctx, userId)
	if err != nil {
		return true
	}
	return cd > nowUnix
}

func (svc *GameService) AddEnergy(ctx context.Context, req *mpb.ReqAddEnergy) (*mpb.ResAddEnergy, error) {
	dbEnergy, err := svc.dao.addEnergy(ctx, req.UserId, req.Energy, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	return &mpb.ResAddEnergy{
		Energy:   dbEnergy.Energy,
		UpdateAt: dbEnergy.RecoverAt,
	}, nil
}

func (svc *GameService) GetHiddenBoss(ctx context.Context, req *mpb.ReqGetHiddenBoss) (*mpb.ResGetHiddenBoss, error) {
	res := &mpb.ResGetHiddenBoss{}
	dbHiddenBoss, err := svc.dao.getHiddenBoss(ctx, req.BossUuid)
	if err != nil {
		return nil, err
	}
	res.HiddenBoss, res.Histories = svc.dbHiddenBoss2HiddenBoss(dbHiddenBoss, req.WithBossDetail)
	bossRsc := svc.rm.getBossRsc(dbHiddenBoss.BossId)
	if bossRsc == nil {
		return nil, mpberr.ErrConfig
	}
	if req.WithBossDetail {
		res.BossRsc = bossRsc
	}
	boss := newGameBoss(bossRsc)
	res.HiddenBoss.TotalHp = boss.totalHP
	res.Fought = dbHiddenBoss.LastFightTime[req.UserId] > 0
	dbFightCD, err := svc.dao.getHiddenBossFightCD(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	res.FightCd = dbFightCD
	return res, nil
}

func (svc *GameService) GetRandomHiddenBoss(ctx context.Context, req *mpb.ReqGetHiddenBoss) (*mpb.ResGetHiddenBoss, error) {
	res := &mpb.ResGetHiddenBoss{}
	dbFightCD, err := svc.dao.getHiddenBossFightCD(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	res.FightCd = dbFightCD

	bossUUid, err := svc.dao.randomFromHiddenBossPool(ctx, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	if bossUUid == 0 { // there is no valid boss
		return res, nil
	}

	dbHiddenBoss, err := svc.dao.getHiddenBoss(ctx, bossUUid)
	if err != nil {
		return nil, err
	}
	res.HiddenBoss, res.Histories = svc.dbHiddenBoss2HiddenBoss(dbHiddenBoss, req.WithBossDetail)
	bossRsc := svc.rm.getBossRsc(dbHiddenBoss.BossId)
	if bossRsc == nil {
		return nil, mpberr.ErrConfig
	}
	if req.WithBossDetail {
		res.BossRsc = bossRsc
	}
	boss := newGameBoss(bossRsc)
	res.HiddenBoss.TotalHp = boss.totalHP

	res.Fought = dbHiddenBoss.LastFightTime[req.UserId] > 0
	return res, nil
}

func (svc *GameService) GetEnergy(ctx context.Context, req *mpb.ReqUserId) (*mpb.ResGetEnergy, error) {
	dbEnergy, err := svc.dao.getEnergy(ctx, req.UserId, time.Now().Unix())
	if err != nil {
		return nil, err
	}

	return &mpb.ResGetEnergy{
		Energy:   dbEnergy.Energy,
		UpdateAt: dbEnergy.RecoverAt,
	}, nil
}

func (svc *GameService) getAccountByUserId(ctx context.Context, userId uint64) (*mpb.AccountInfo, error) {
	client, err := com.GetAccountServiceClient(ctx, svc)
	if err != nil {
		return nil, err
	}
	res, err := client.GetAccountByUserId(ctx, &mpb.ReqUserId{UserId: userId})
	if err != nil {
		svc.logger.Error("getAccountByUserId failed", zap.Uint64("user_id", userId), zap.Error(err))
		return nil, err
	}
	return res.Account, nil
}

func (svc *GameService) NewHiddenBoss(ctx context.Context, req *mpb.ReqNewHiddenBoss) (*mpb.ResNewHiddenBoss, error) {
	hiddenBoss, err := svc.triggerHiddenBoss(ctx, req.UserId, gutil.If(req.BossId > 0, req.BossId, 10001),
		time.Now())
	if err != nil {
		return nil, err
	}
	res := &mpb.ResNewHiddenBoss{
		HiddenBoss: hiddenBoss,
	}
	return res, nil
}

func (svc *GameService) GetGameInfo(ctx context.Context, req *mpb.ReqUserId) (*mpb.ResGetGameInfo, error) {
	engRes, err := svc.GetEnergy(ctx, req)
	if err != nil {
		return nil, err
	}

	res := &mpb.ResGetGameInfo{
		Energy:   engRes.Energy,
		UpdateAt: engRes.UpdateAt,
	}

	dbHist, err := svc.dao.getBossDefeatHistory(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	dbBC, err := svc.dao.getBuffCards(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	status, err := svc.dao.getBuffCardChoseStatus(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	res.BossDefeatHistory = svc.dbBossDefeatHistory2BossDefeatHistory(dbHist)
	res.BuffCards = svc.dbBuffCards2BuffCards(dbBC)
	res.BuffCardStatus = mpb.EGame_BuffCardStatus(status)
	return res, nil
}

func (svc *GameService) FightPVP(ctx context.Context, req *mpb.ReqFightPVP) (*mpb.ResFightPVP, error) {
	if req.UserId == req.TargetId {
		return nil, mpberr.ErrParam
	}
	now := time.Now()
	if _, _, ok := svc.rm.isPVPSettle(now); ok {
		return nil, mpberr.ErrPVPSettle
	}
	nowUnix := time.Now().Unix()
	season := svc.rm.getPVPSeason(nowUnix)
	var manaConsume int64
	var leftMana uint64
	// check challenge cnt
	dbPVPInfo, err := svc.dao.checkAndCostPVPChallengeCnt(ctx, season, req.UserId, nowUnix)
	if err != nil {
		if !errors.Is(err, mpberr.ErrNotEnoughPVPChallengeCnt) {
			return nil, err
		}
		// try cost mana
		pvpSeason := svc.rm.getPVPSeason(nowUnix)
		manaConsume = int64(svc.rm.getPVPManaConsume(pvpSeason))
		_, _, updateItem, err := com.ExchangeItems(ctx, svc, req.UserId, nil, nil,
			-1*manaConsume,
			mpb.EItem_TransReason_GameFightPVP, 0)
		if err != nil {
			svc.logger.Error("FightPVP cost mana failed", zap.Uint64("user_id", req.UserId),
				zap.Int64("mana_cost", -1*manaConsume), zap.Error(err))
			return nil, mpberr.ErrNotEnoughPVPChallengeCnt
		}
		if len(updateItem) == 1 && updateItem[0].ItemId == uint32(mpb.EItem_ItemId_Mana) {
			leftMana = updateItem[0].Num
		}
	} else {
		w, err := svc.getWallet(ctx, req.UserId)
		if err != nil {
			return nil, err
		}
		leftMana = w.Mana
	}

	res := &mpb.ResFightPVP{
		PvpChallengeCnt:         dbPVPInfo.ChallengeCnt,
		PvpChallengeCntUpdateAt: dbPVPInfo.UpdateAt,
		Mana:                    leftMana,
	}

	// get equips
	equips, err := svc.batchGetUsersEquips(ctx, []uint64{req.UserId, req.TargetId})
	if err != nil {
		return nil, err
	}

	player := newPlayer(svc, req.UserId, mpb.ERole_RoleType_Player, nil)
	if e, ok := equips[req.UserId]; ok {
		player.updateEquips(e.BaseEquips, e.NftEquips)
	}
	defender := newPlayer(svc, req.UserId, mpb.ERole_RoleType_Defender, nil)
	if e, ok := equips[req.TargetId]; ok {
		defender.updateEquips(e.BaseEquips, e.NftEquips)
	}

	res.ChallengerHp = player.hp
	res.DefenderHp = defender.hp
	gl := svc.glm.newGameLevel(player, defender)
	res.NewRank, res.OldRank, err = svc.dao.fightPVP(ctx, req.UserId, req.TargetId, func() bool {
		res.Win, res.Details = gl.fightPVP(player, defender)
		return res.Win
	})
	if err != nil {
		if errors.Is(err, mpberr.ErrPVPTargetRankChanged) && manaConsume > 0 {
			// roll back mana consume
			_, _, _, err := com.ExchangeItems(ctx, svc, req.UserId, nil, nil, manaConsume,
				mpb.EItem_TransReason_GameFightPVPRollBack, 0)
			if err != nil {
				svc.logger.Error("FightPVP roll back mana consume failed", zap.Uint64("user_id", req.UserId),
					zap.Int64("mana_consume", -1*manaConsume), zap.Error(err))
			}
		}

		return nil, err
	}

	return res, nil
}

func (svc *GameService) batchGetAccountsByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*mpb.AccountInfo,
	error) {
	client, err := com.GetAccountServiceClient(ctx, svc)
	if err != nil {
		return nil, err
	}
	res, err := client.BatchGetAccountsByUserIds(ctx, &mpb.ReqUserIds{UserIds: userIds})
	if err != nil {
		svc.logger.Error("batchGetAccountsByUserIds failed", zap.Any("user_ids", userIds), zap.Error(err))
		return nil, err
	}
	if res.Accounts == nil {
		res.Accounts = make(map[uint64]*mpb.AccountInfo)
	}
	return res.Accounts, nil
}

func (svc *GameService) GetPVPRanks(ctx context.Context, req *mpb.ReqGetPVPRanks) (*mpb.ResGetPVPRanks, error) {
	if req.PageNum >= PVPRankPageCnt { // page num start from 0
		return nil, mpberr.ErrParam
	}
	res := &mpb.ResGetPVPRanks{PageNum: req.PageNum}
	users, err := svc.dao.getPVPRanks(ctx, req.PageNum)
	if err != nil {
		return nil, err
	}
	accs, err := svc.batchGetAccountsByUserIds(ctx, users)
	if err != nil {
		return nil, err
	}

	equips, err := svc.batchGetUsersEquips(ctx, users)
	if err != nil {
		return nil, err
	}

	for i, uid := range users {
		player := newPlayer(svc, uid, mpb.ERole_RoleType_Player, nil)
		if e, ok := equips[uid]; ok {
			player.updateEquips(e.BaseEquips, e.NftEquips)
		}
		res.RankList = append(res.RankList, &mpb.PVPRankNode{
			Rank:        uint32(i) + 1,
			AccountInfo: accs[uid],
			BaseEquips:  equips[uid].GetBaseEquips(),
			NftEquips:   equips[uid].GetNftEquips(),
			Attr:        player.attrs,
		})
	}
	return res, nil
}

func (svc *GameService) GetPVPInfo(ctx context.Context, req *mpb.ReqUserId) (*mpb.ResGetPVPInfo, error) {
	now := time.Now()
	nowUnix := now.Unix()
	season, date := svc.rm.getPVPSeasonDate(now)
	if season == 0 {
		return nil, mpberr.ErrPVPSettle
	}
	dbPVPInfo, rankRewards, err := svc.dao.getUserPVPInfoAndClaimRankRewards(ctx, season, req.UserId, nowUnix)
	if err != nil {
		return nil, err
	}

	rank, err := svc.dao.getUserPVPRank(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	if dbPVPInfo.UpdateAt < gtime.TodayXHourTimeUnix(0, util.TimeZone_GMT) {
		dbPVPInfo.ChallengeCnt = svc.rm.getPVPConfigRsc(season).GetDailyPvpFreeChallengeCnt()
		dbPVPInfo.UpdateAt = nowUnix
	}

	res := &mpb.ResGetPVPInfo{
		Rank:                  rank + 1,
		ChallengerCnt:         dbPVPInfo.ChallengeCnt,
		ChallengerCntUpdateAt: dbPVPInfo.UpdateAt,
		PvpSeasonDate:         date,
	}

	res.PvpManaAwardsPool, err = svc.dao.getPVPManaPool(ctx, season, date)
	if err != nil {
		return nil, err
	}
	res.PvpManaAwardsPool += svc.rm.getPVPConfigRsc(season).GetManaAwardsPool()

	var totalRewards [][]*mpb.AwardRsc
	for _, v := range rankRewards {
		resRankRewards := &mpb.PVPSettleRewards{
			SeasonId: v[0],
			Date:     v[1],
			Rank:     v[2],
			Rewards:  &mpb.Items{},
		}
		rewards, manaRate := svc.rm.getPVPSeasonRankRewardsRscs(v[0], v[2])
		totalRewards = append(totalRewards, rewards)
		for _, vv := range rewards {
			resRankRewards.Rewards.Items = append(resRankRewards.Rewards.Items, &mpb.Item{
				ItemId: vv.ItemId,
				Num:    vv.Num,
			})
		}

		// calc mana pool
		pvpConfigRsc := svc.rm.getPVPConfigRsc(resRankRewards.SeasonId)
		manaPool, err := svc.dao.getPVPManaPool(ctx, resRankRewards.SeasonId, resRankRewards.Date)
		if err != nil {
			return nil, err
		}
		manaPool += pvpConfigRsc.GetManaAwardsPool()
		mana := manaPool * uint64(manaRate) / com.RateBase
		if mana > 0 {
			resRankRewards.Rewards.Items = append(resRankRewards.Rewards.Items, &mpb.Item{
				ItemId: uint32(mpb.EItem_ItemId_Mana),
				Num:    mana,
			})
			totalRewards = append(totalRewards,
				[]*mpb.AwardRsc{&mpb.AwardRsc{ItemId: uint32(mpb.EItem_ItemId_Mana), Num: mana}})
		}

		res.PvpSettleRewards = append(res.PvpSettleRewards, resRankRewards)
	}

	_, err = com.AddItemsFromAwardRsc(ctx, svc, req.UserId, totalRewards, mpb.EItem_TransReason_PVPRankSettle,
		0)
	if err != nil {
		svc.logger.Error("GetPVPInfo claim pvp rank settle rewards ", zap.Uint64("user_id", req.UserId),
			zap.Any("rewards", totalRewards), zap.Error(err))
		return nil, mpberr.ErrNotEnoughPVPChallengeCnt
	}

	return res, nil
}

func (svc *GameService) GetPVPChallengeTargets(ctx context.Context, req *mpb.ReqUserId) (*mpb.ResGetPVPChallengeTargets, error) {
	res := &mpb.ResGetPVPChallengeTargets{}
	rank, err := svc.dao.getUserPVPRank(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	pvpConfig := svc.rm.getPVPConfigRsc(svc.rm.getPVPSeason(time.Now().Unix()))
	if pvpConfig == nil {
		return res, nil
	}

	targets := make([]uint64, 0, 4)
	targetRanks := make([]uint32, 0, 4)
	m := make(map[uint32]bool)
	targetRank, ok := svc.rm.getChallengeTargetRank(rank, 1, pvpConfig)
	if ok {
		targetUid, err := svc.dao.getPVPRankById(ctx, targetRank)
		if err != nil {
			return nil, err
		}
		if targetUid > 0 {
			targets = append(targets, targetUid)
			targetRanks = append(targetRanks, targetRank)
		}
	}
	m[targetRank] = true
	targetRank, ok = svc.rm.getChallengeTargetRank(rank, 2, pvpConfig)
	if ok && !m[targetRank] {
		targetUid, err := svc.dao.getPVPRankById(ctx, targetRank)
		if err != nil {
			return nil, err
		}
		if targetUid > 0 {
			targets = append(targets, targetUid)
			targetRanks = append(targetRanks, targetRank)
		}
	}
	m[targetRank] = true
	targetRank, ok = svc.rm.getChallengeTargetRank(rank, 3, pvpConfig)
	if ok && !m[targetRank] {
		targetUid, err := svc.dao.getPVPRankById(ctx, targetRank)
		if err != nil {
			return nil, err
		}
		if targetUid > 0 {
			targets = append(targets, targetUid)
			targetRanks = append(targetRanks, targetRank)
		}
	}
	m[targetRank] = true
	targetRank, ok = svc.rm.getChallengeTargetRank(rank, 4, pvpConfig)
	if ok && !m[targetRank] {
		targetUid, err := svc.dao.getPVPRankById(ctx, targetRank)
		if err != nil {
			return nil, err
		}
		if targetUid > 0 {
			targets = append(targets, targetUid)
			targetRanks = append(targetRanks, targetRank)
		}
	}

	if len(targets) == 0 {
		return res, nil
	}

	accs, err := svc.batchGetAccountsByUserIds(ctx, targets)
	if err != nil {
		return nil, err
	}

	equips, err := svc.batchGetUsersEquips(ctx, targets)
	if err != nil {
		return nil, err
	}

	for i, uid := range targets {
		player := newPlayer(svc, uid, mpb.ERole_RoleType_Player, nil)
		if e, ok := equips[uid]; ok {
			player.updateEquips(e.BaseEquips, e.NftEquips)
		}
		res.TargetList = append(res.TargetList, &mpb.PVPRankNode{
			Rank:        targetRanks[i] + 1,
			AccountInfo: accs[uid],
			BaseEquips:  equips[uid].GetBaseEquips(),
			NftEquips:   equips[uid].GetNftEquips(),
			Attr:        player.attrs,
		})
	}

	return res, nil
}

func (svc *GameService) GetPVPHistory(ctx context.Context, _ *mpb.Empty) (*mpb.ResGetPVPHistory, error) {
	list, ok := svc.gcm.getPVPHistoryCache()
	if ok {
		fmt.Println("get from cache")
		return &mpb.ResGetPVPHistory{
			List: list,
		}, nil
	}
	fmt.Println("get from db")
	// get from db
	his, err := svc.dao.getPVPHistory(ctx)
	if err != nil {
		return nil, err
	}

	var uids = make([]uint64, 0, len(his))
	for _, uid := range his {
		uids = append(uids, uid)
	}

	accs, err := svc.batchGetAccountsByUserIds(ctx, uids)
	if err != nil {
		return nil, err
	}

	for date, uid := range his {
		list = append(list, &mpb.PVPHistory{
			Date:    date,
			Rank:    1,
			Account: accs[uid],
		})
	}

	svc.gcm.setPVPHistoryCache(list)

	return &mpb.ResGetPVPHistory{
		List: list,
	}, nil
}

func (svc *GameService) getWallet(ctx context.Context, userId uint64) (*mpb.ResGetWallet, error) {
	itemClient, err := com.GetItemServiceClient(ctx, svc)
	if err != nil {
		return nil, err
	}
	res, err := itemClient.GetWallet(ctx, &mpb.ReqUserId{
		UserId: userId,
	})
	if err != nil {
		svc.logger.Error("getWallet get user equips failed", zap.Any("user_id", userId),
			zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (svc *GameService) RandomBuffCards(ctx context.Context, req *mpb.ReqRandomBuffCards) (*mpb.ResRandomBuffCards, error) {
	bossRsc := svc.rm.getBossRsc(req.BossId)
	if bossRsc == nil || bossRsc.BuffCardRandPool == 0 {
		return nil, mpberr.ErrParam
	}

	poolRsc := svc.rm.getBuffCardRandPool(bossRsc.BuffCardRandPool)
	if poolRsc == nil {
		return nil, mpberr.ErrParam
	}

	dbHist, err := svc.dao.getBossDefeatHistory(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	canFightBoss := svc.rm.getBossClassNextLevelBossRsc(dbHist.BossDefeatHistory[bossRsc.Class].GetBossId())
	if canFightBoss != nil && canFightBoss.Level < bossRsc.Level {
		return nil, mpberr.ErrParam
	}

	buffCards, err := util.RandSliceByWeight(poolRsc.BuffCards, buffCardCnt, func(i int) uint32 {
		return poolRsc.BuffCards[i].Weight
	})
	if err != nil {
		svc.logger.Error("RandomBuffCards rand failed", zap.Error(err))
		return nil, mpberr.ErrParam
	}

	ok, err := svc.dao.checkAndUpdateBuffCardChoseStatus(ctx, req.UserId,
		[]uint32{uint32(mpb.EGame_BuffCardStatus_None), uint32(mpb.EGame_BuffCardStatus_Randomed)},
		uint32(mpb.EGame_BuffCardStatus_Randomed))
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, mpberr.ErrBuffCardStatusWrong
	}

	res := &mpb.ResRandomBuffCards{
		BossId: req.BossId,
	}
	for _, v := range buffCards {
		res.BuffCards = append(res.BuffCards, v.Id)
	}

	// write db
	err = svc.dao.setBuffCardOptions(ctx, req.UserId, req.BossId, res.BuffCards)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (svc *GameService) ChoseBuffCard(ctx context.Context, req *mpb.ReqChoseBuffCard) (*mpb.ResChoseBuffCard, error) {
	buffCardRsc := svc.rm.getBuffCardRsc(req.BuffCard)
	if buffCardRsc == nil {
		return nil, mpberr.ErrConfig
	}

	dbBuffCardOpts, err := svc.dao.getBuffCardOptions(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	var ok bool
	for _, v := range dbBuffCardOpts.BuffCards {
		if v == req.BuffCard {
			ok = true
			break
		}
	}
	if !ok {
		return nil, mpberr.ErrParam
	}

	ok, err = svc.dao.checkBuffCardChoseStatus(ctx, req.UserId, []uint32{uint32(mpb.EGame_BuffCardStatus_Randomed)})
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, mpberr.ErrBuffCardStatusWrong
	}

	newBuffCards, err := svc.bcm.effectBeforeGame(ctx, req.UserId, buffCardRsc)
	if err != nil {
		return nil, err
	}

	if len(newBuffCards) == 0 {
		err = svc.dao.updateBuffCardChoseStatus(ctx, req.UserId, uint32(mpb.EGame_BuffCardStatus_Chosed))
		if err != nil {
			return nil, err
		}
		err = svc.dao.addBuffCards(ctx, req.UserId, buffCardRsc)
		if err != nil {
			return nil, err
		}
	}
	return &mpb.ResChoseBuffCard{
		BuffCards: newBuffCards,
	}, nil
}

func (svc *GameService) rerandomBuffCard(ctx context.Context, userId uint64, poolId uint32) ([]uint32, error) {
	poolRsc := svc.rm.getBuffCardRandPool(poolId)
	if poolRsc == nil {
		return nil, mpberr.ErrParam
	}

	buffCards, err := util.RandSliceByWeight(poolRsc.BuffCards, buffCardCnt, func(i int) uint32 {
		return poolRsc.BuffCards[i].Weight
	})
	if err != nil {
		svc.logger.Error("RandomBuffCards rand failed", zap.Error(err))
		return nil, mpberr.ErrParam
	}

	newBuffCards := make([]uint32, 0, buffCardCnt)
	for _, v := range buffCards {
		newBuffCards = append(newBuffCards, v.Id)
	}

	// write db
	err = svc.dao.updateBuffCardOptions(ctx, userId, newBuffCards)
	if err != nil {
		return nil, err
	}

	return newBuffCards, nil
}
