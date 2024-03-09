package apiproxy

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/oldjon/gutil"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"gitlab.com/morbackend/mor_services/mpberr"
	"gitlab.com/morbackend/mor_services/util"
	"go.uber.org/zap"
)

type telegramAPIHandler func(ctx context.Context, w http.ResponseWriter, msg *mpb.TGMsgRecv) error

type TelegramManager struct {
	svc         *APIProxyGRPCService
	logger      *zap.Logger
	sendUrl     string
	botTokens   map[string]string
	cmdHandlers map[string]telegramAPIHandler
	cbqHandlers map[string]telegramAPIHandler
	rm          *apiProxyResourceMgr
}

func TelegramManagerGetMe() *TelegramManager {
	tm := APIProxyGRPCGetMe().tgMgr
	if tm.cmdHandlers == nil {
		tm.initCmdHandlers()
	}
	if tm.cbqHandlers == nil {
		tm.initCBQHandlers()
	}
	return tm
}

func newTelegramManager(as *APIProxyGRPCService, sendUrl string, botTokens map[string]string,
) *TelegramManager {
	return &TelegramManager{
		svc:       as,
		logger:    as.logger,
		sendUrl:   sendUrl,
		botTokens: botTokens,
		rm:        as.rm,
	}
}

func (tm *TelegramManager) initCmdHandlers() {
	tm.cmdHandlers = make(map[string]telegramAPIHandler)
	tm.cmdHandlers[com.TGCmd_Echo] = tm.tgAPICmdEcho
	tm.cmdHandlers[com.TGCmd_Menu] = tm.tgAPICmdMenu
	tm.cmdHandlers[com.TGCmd_Game] = tm.tgAPICmdGame
	tm.cmdHandlers[com.TGCmd_HiddenBoss] = tm.tgAPICmdHiddenBoss
	return
}

func (tm *TelegramManager) initCBQHandlers() {
	tm.cbqHandlers = make(map[string]telegramAPIHandler)
	tm.cbqHandlers[com.TGCBQ_SendGame] = tm.tgAPICBQSendGame
	tm.cbqHandlers[com.TGCBQ_LaunchGame] = tm.tgAPICBQLaunchGame
	tm.cbqHandlers[com.TGCBQ_FightHBoss] = tm.tgAPICBQFightHiddenBoss
	tm.cbqHandlers[com.TGCBQ_SendSetting] = tm.tgAPICBQSendSetting
	tm.cbqHandlers[com.TGCBQ_SetLan] = tm.tgAPICBQSetLan

	return
}

func (tm *TelegramManager) getCmdHandler(cmd string) (telegramAPIHandler, bool) {
	fn, ok := tm.cmdHandlers[cmd]
	return fn, ok
}

func (tm *TelegramManager) getCBQHandler(cmd string) (telegramAPIHandler, bool) {
	fn, ok := tm.cbqHandlers[cmd]
	return fn, ok
}

func (tm *TelegramManager) sendMsgToTelegram(ctx context.Context, bot string, msg []byte) error {
	botToken, ok := tm.botTokens[bot]
	if !ok {
		return mpberr.ErrNoTelegramBot
	}
	urlStr := tm.sendUrl + botToken + "/"
	headers := map[string]string{"Content-Type": "application/json"}
	tm.logger.Debug("sendMsgToTelegram url", zap.String("", urlStr), zap.String("data", string(msg)))
	resp, err := util.HttpsPost(ctx, urlStr, headers, msg)
	if err != nil {
		tm.logger.Info("sendMsgToTelegram failed", zap.Error(err))
		return err
	}
	tm.logger.Info("sendMsgToTelegram result", zap.String("", string(resp)))
	return nil
}

func (tm *TelegramManager) sendCmdReply(ctx context.Context, msg interface{}) error {
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return tm.sendMsgToTelegram(ctx, com.TGGameBot, msgBytes)
}

// -----------------------------------------Telegram command handlers---------------------------------------------------

func (tm *TelegramManager) tgAPICmdEcho(ctx context.Context, _ http.ResponseWriter, msg *mpb.TGMsgRecv) error {
	if msg == nil || msg.Message == nil || msg.Message.Chat == nil {
		return mpberr.ErrParam
	}

	replyRsc := tm.rm.getTGReplyRsc(com.TGCmd_Echo, com.TGReplyType_Ok, "")
	if replyRsc == nil {
		tm.logger.Error("tgAPICmdMenu cmd reply rsc not found")
		return mpberr.ErrConfig
	}
	replyMsg := &mpb.TGReply{
		Method: replyRsc.Method,
		ChatId: msg.Message.Chat.Id,
		Text:   msg.Message.Text,
	}
	err := tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICmdEcho failed", zap.Error(err))
		return err
	}
	return nil
}

func (tm *TelegramManager) tgAPICmdMenu(ctx context.Context, w http.ResponseWriter, msg *mpb.TGMsgRecv) error {
	if msg == nil || msg.Message == nil || msg.Message.Chat == nil {
		return mpberr.ErrParam
	}

	acc, err := tm.getAccountByTGUser(ctx, msg.Message.From)
	if err != nil {
		return err
	}

	replyRsc := tm.rm.getTGReplyRsc(com.TGCmd_Menu, com.TGReplyType_Ok, "")
	if replyRsc == nil {
		tm.logger.Error("tgAPICmdMenu cmd reply rsc not found")
		return mpberr.ErrConfig
	}
	replyMsg := &mpb.TGReply{
		Method:      replyRsc.Method,
		ChatId:      msg.Message.Chat.Id,
		Text:        replyRsc.Text,
		Photo:       replyRsc.Photo,
		ReplyMarkup: tm.rm.getTGInlineKeyBoard(com.TGCmd_Menu, acc.Lan),
	}
	err = tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICmdMenu failed", zap.Error(err))
		return err
	}
	return nil
}

func (tm *TelegramManager) tgAPICmdGame(ctx context.Context, _ http.ResponseWriter, msg *mpb.TGMsgRecv) error {
	if msg == nil || msg.Message == nil || msg.Message.Chat == nil {
		return mpberr.ErrParam
	}
	replyRsc := tm.rm.getTGReplyRsc(com.TGCmd_Game, com.TGReplyType_Ok, "")
	if replyRsc == nil {
		tm.logger.Error("tgAPICmdGame cmd reply rsc not found")
		return mpberr.ErrConfig
	}
	replyMsg := &mpb.TGReply{
		Method:        replyRsc.Method,
		ChatId:        msg.Message.Chat.Id,
		GameShortName: replyRsc.GameShortName,
	}
	err := tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICmdGame failed", zap.Error(err))
		return err
	}
	return nil
}

func (tm *TelegramManager) tgAPICmdHiddenBoss(ctx context.Context, _ http.ResponseWriter, msg *mpb.TGMsgRecv) error {
	if msg == nil || msg.Message == nil || msg.Message.From == nil || msg.Message.Chat == nil {
		return mpberr.ErrParam
	}
	_, paramsStr := util.ParseTGMsg(msg.Message.Text)

	if paramsStr == "" { // if there is no params, send a random hidden boss
		return tm.tgAPICmdHiddenBossRandomBoss(ctx, msg)
	}

	params, ok := tm.parseSpecialLinkParams(paramsStr)
	if !ok {
		return tm.tgAPICmdHiddenBossRandomBoss(ctx, msg)
	}

	account, err := tm.getAccountByTGUser(ctx, msg.Message.From)
	if err != nil {
		return err
	}

	sendBossNotFound := func() error {
		replyRsc := tm.rm.getTGReplyRsc(com.TGCmd_HiddenBoss, com.TGReplyType_Err, com.TGReplySubType_NotFound)
		if replyRsc == nil {
			tm.logger.Error("tgAPICmdHiddenBoss cmd reply rsc not found")
			return mpberr.ErrConfig
		}

		replyMsg := &mpb.TGReply{
			Method: replyRsc.Method,
			ChatId: msg.Message.Chat.Id,
			Text:   replyRsc.Text,
		}
		err = tm.sendCmdReply(ctx, replyMsg)
		if err != nil {
			tm.logger.Error("tgAPICmdEcho failed", zap.Error(err))
			return err
		}
		return nil
	}

	bossUUID := gutil.StrToUint64(params["bossuuid"])
	if bossUUID == 0 {
		return sendBossNotFound()
	}

	nowUnix := time.Now().Unix()
	boss, err := tm.getHiddenBoss(ctx, account.UserId, bossUUID)
	if util.GRPCErrIs(err, mpberr.ErrHiddenBossNotFound) || boss.HiddenBoss.GetHp() == 0 ||
		boss.HiddenBoss.ExpireAt < nowUnix {
		return tm.tgAPICmdHiddenBossRandomBoss(ctx, msg)
	}
	if err != nil {
		return err
	}
	replyRsc := tm.rm.getTGReplyRsc(com.TGCmd_HiddenBoss, com.TGReplyType_Ok, "")
	if replyRsc == nil {
		tm.logger.Error("tgAPICmdHiddenBoss cmd ok reply rsc not found")
		return mpberr.ErrConfig
	}

	replyMsg := &mpb.TGReply{
		Method: replyRsc.Method,
		ChatId: msg.Message.Chat.Id,
		Photo:  boss.BossRsc.GetPhoto(),
	}

	sharerUid := gutil.StrToUint64(params["sharer"])
	var sharer string
	if sharerUid > 0 {
		acc, err := tm.getAccount(ctx, sharerUid)
		if err != nil {
			return err
		}
		sharer = acc.Nickname
	}

	// Caption
	replyMsg.Caption, replyMsg.ParseMode, err = tm.generateHiddenBossCaption(ctx, boss, sharer, account.Lan)
	if err != nil {
		return err
	}
	// ReplyMarkup
	replyMsg.ReplyMarkup = &mpb.TGInlineKeyboardMarkup{
		InlineKeyboard: [][]*mpb.TGInlineKeyboardButton{{{
			Text: tm.rm.getKeyString(com.STRING_KEY_BUTTON_FIGHT, account.Lan),
			CallbackData: com.TGCBQ_FightHBoss + " " + strconv.Itoa(int(boss.HiddenBoss.BossId)) + " " +
				strconv.Itoa(int(boss.HiddenBoss.BossUuid)) + " " + params["sharer"]}}},
	}

	err = tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICmdMenu failed", zap.Error(err))
		return err
	}

	return nil
}

func (tm *TelegramManager) generateHiddenBossCaption(ctx context.Context, boss *mpb.ResGetHiddenBoss, sharer, lan string) (string, string, error) {
	var str string
	//if len(sharer) > 0 {
	//	str = fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_INVITE, lan), sharer,
	//		tm.rm.getKeyString(boss.BossRsc.GetName(), lan)) + "\n"
	//} else {
	str = fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_APPEARED, lan),
		tm.rm.getKeyString(boss.BossRsc.GetName(), lan)) + "\n"
	//}

	str += tm.rm.getKeyString(com.STRING_KEY_BOSS_COME_TO_FIGHT, lan) + "\n"

	// boss level
	str += fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_LEVEL, lan), boss.BossRsc.GetLevelShow()) + "\n"

	// boss hp
	if boss.HiddenBoss.GetTotalHp() > 0 {
		str += fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_HP, lan), boss.HiddenBoss.GetHp(),
			boss.HiddenBoss.GetTotalHp(),
			strconv.Itoa(int(boss.HiddenBoss.GetHp()*100/boss.HiddenBoss.GetTotalHp()))+"%") + "\n"
	}
	// expire time
	str += fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_EXPIRE_AT, lan),
		tm.calcTimeDurationStr(time.Now().Unix(), boss.HiddenBoss.GetExpireAt()))

	// rewards can get
	itemIds := make([]uint32, 0, len(boss.BossRsc.GetDmgAwards()))
	for _, v := range boss.BossRsc.GetDmgAwards() {
		itemIds = append(itemIds, v.ItemId)
	}

	itemsRsc, err := tm.getItemsRsc(ctx, itemIds)
	if err != nil {
		return "", "", err
	}

	if len(boss.BossRsc.GetDmgAwards()) > 0 {
		str += "\n\n" + tm.rm.getKeyString(com.STRING_KEY_BOSS_REWARDS_TITLE, lan) + "\n"
		for i, item := range boss.BossRsc.GetDmgAwards() {
			itemRsc, ok := itemsRsc[item.ItemId]
			if !ok {
				continue
			}
			if i != len(boss.BossRsc.GetDmgAwards())-1 {
				str += fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_REWARDS, lan),
					tm.rm.getKeyString(itemRsc.NameStringKey, lan)+":"+strconv.Itoa(int(item.Num))) + "\n"
			} else {
				str += fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_REWARDS, lan),
					tm.rm.getKeyString(itemRsc.NameStringKey, lan)+":"+strconv.Itoa(int(item.Num)))
			}
		}
	}

	if len(boss.Histories) != 0 {
		// fight detail
		str += "\n \n" + tm.rm.getKeyString(com.STRING_KEY_BOSS_FIGHT_TITLE, lan)
		for _, v := range boss.Histories {
			str += "\n" + fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_DMG_CASED_BY_PLAYER, lan), v.Nickname,
				strconv.Itoa(int(float64(v.DmgRate)/float64(com.RateBase)*100))+"%")
		}
	}

	return str, "HTML", nil
}

func (tm *TelegramManager) generateHiddenBossNotFoundCaption(lan string) (string, string) {
	return tm.rm.getKeyString(com.STRING_KEY_BOSS_DIED_OR_ESCAPED, lan), "HTML"
}

func (tm *TelegramManager) generateHiddenBossDefeatedCaption(ctx context.Context, boss *mpb.ResGetHiddenBoss,
	result *mpb.ResFight, sharer string, acc *mpb.AccountInfo) (string, string, error) {
	str, _, err := tm.generateHiddenBossCaption(ctx, boss, sharer, acc.Lan)
	if err != nil {
		return "", "", err
	}
	str += "\n" + fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_HAS_BEEN_DEFEATED, acc.Lan), acc.Nickname,
		tm.rm.getKeyString(boss.BossRsc.GetName(), acc.Lan))
	if result.Awards == nil || len(result.Awards.AddItems) == 0 {
		str += "\n" + fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_WIN_NO_REWARDS, acc.Lan), acc.Nickname)
		return str, "HTML", nil
	}

	addItemsList := make(map[uint32]*mpb.Item) //, 0, len(result.Awards.AddItems))
	for _, v := range result.Awards.AddItems {
		item := addItemsList[v.ItemId]
		if item == nil {
			addItemsList[v.ItemId] = v
		} else {
			item.Num += v.Num
		}
	}

	itemIds := make([]uint32, 0, len(addItemsList))
	for _, v := range addItemsList {
		itemIds = append(itemIds, v.ItemId)
	}

	itemsRsc, err := tm.getItemsRsc(ctx, itemIds)
	if err != nil {
		return "", "", err
	}

	var rewardsStr string
	for _, v := range addItemsList {
		itemRsc := itemsRsc[v.ItemId]
		if itemRsc == nil {
			continue
		}
		rewardsStr += tm.rm.getKeyString(itemRsc.NameStringKey, acc.Lan) + "*" + strconv.Itoa(int(v.Num)) + " "
	}

	str += "\n" + fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_BOSS_WIN_REWARDS, acc.Lan), acc.Nickname, rewardsStr)

	return str, "HTML", nil
}

func (tm *TelegramManager) tgAPICmdHiddenBossRandomBoss(ctx context.Context, msg *mpb.TGMsgRecv) error {
	tm.logger.Debug("tgAPICmdHiddenBossRandomBoss")
	if msg == nil || msg.Message == nil || msg.Message.Chat == nil || msg.Message.From == nil {
		return mpberr.ErrParam
	}

	account, err := tm.getAccountByTGUser(ctx, msg.Message.From)
	if err != nil {
		return err
	}

	// random an exist hidden boss
	boss, err := tm.randomHiddenBoss(ctx, account.UserId)
	if err != nil {
		return err
	}

	if boss.HiddenBoss == nil {
		return nil
	}

	replyRsc := tm.rm.getTGReplyRsc(com.TGCmd_HiddenBoss, com.TGReplyType_Ok, "")
	if replyRsc == nil {
		tm.logger.Error("tgAPICmdHiddenBoss cmd ok reply rsc not found")
		return mpberr.ErrConfig
	}
	replyMsg := &mpb.TGReply{
		Method: replyRsc.Method,
		ChatId: msg.Message.Chat.Id,
		Photo:  boss.BossRsc.GetPhoto(),
	}

	// Caption
	replyMsg.Caption, replyMsg.ParseMode, err = tm.generateHiddenBossCaption(ctx, boss, "", account.Lan)
	if err != nil {
		return err
	}
	// ReplyMarkup
	replyMsg.ReplyMarkup = &mpb.TGInlineKeyboardMarkup{
		InlineKeyboard: [][]*mpb.TGInlineKeyboardButton{{{
			Text: tm.rm.getKeyString(com.STRING_KEY_BUTTON_FIGHT, account.Lan),
			CallbackData: com.TGCBQ_FightHBoss + " " + strconv.Itoa(int(boss.HiddenBoss.BossId)) + " " +
				strconv.Itoa(int(boss.HiddenBoss.BossUuid))},
		}},
	}

	err = tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICmdMenu failed", zap.Error(err))
		return err
	}
	return nil
}

//--------------------------------------------Telegram callback query handlers------------------------------------------

func (tm *TelegramManager) tgAPICBQSendGame(ctx context.Context, _ http.ResponseWriter, msg *mpb.TGMsgRecv) error {
	if msg == nil || msg.CallbackQuery == nil ||
		msg.CallbackQuery.Message == nil || msg.CallbackQuery.Message.Chat == nil {
		return mpberr.ErrParam
	}

	answerRsc := tm.rm.getTGReplyRsc(com.TGCBQ_SendGame, com.TGReplyType_Ok, com.TGReplySubType_Answer)
	if answerRsc != nil {
		answerMsg := &mpb.TGAnswerCallbackQuery{
			Method:          "answerCallbackQuery",
			CallbackQueryId: msg.CallbackQuery.Id,
		}
		err := tm.sendCmdReply(ctx, answerMsg)
		if err != nil {
			tm.logger.Error("tgAPICBQSendGame failed", zap.Error(err))
			return err
		}
	}

	acc, err := tm.getAccountByTGUser(ctx, msg.CallbackQuery.From)
	if err != nil {
		return err
	}

	replyRsc := tm.rm.getTGReplyRsc(com.TGCBQ_SendGame, com.TGReplyType_Ok, "")
	if replyRsc == nil {
		tm.logger.Error("tgAPICBQSendGame cmd reply rsc not found")
		return mpberr.ErrConfig
	}

	replyMsg := &mpb.TGReply{
		Method:        replyRsc.Method,
		ChatId:        msg.CallbackQuery.Message.Chat.Id,
		GameShortName: replyRsc.GameShortName,
		ReplyMarkup:   tm.rm.getTGInlineKeyBoard(com.TGCBQ_SendGame, acc.Lan),
	}
	err = tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICBQSendGame failed", zap.Error(err))
		return err
	}
	return nil
}

func (tm *TelegramManager) tgAPICBQLaunchGame(ctx context.Context, _ http.ResponseWriter, msg *mpb.TGMsgRecv) error {
	if msg == nil || msg.CallbackQuery == nil || msg.CallbackQuery.From == nil {
		return nil
	}
	answerRsc := tm.rm.getTGReplyRsc(com.TGCBQ_LaunchGame, com.TGReplyType_Ok, com.TGReplySubType_Answer)
	if answerRsc == nil {
		tm.logger.Error("tgAPICBQLaunchGame cmd reply rsc not found")
		return mpberr.ErrConfig
	}

	_, err := tm.getAccountByTGUser(ctx, msg.CallbackQuery.From)
	if err != nil {
		return err
	}

	token, err := tm.generateLoginToken(ctx, msg.CallbackQuery.From)
	if err != nil {
		return err
	}
	answerMsg := &mpb.TGAnswerCallbackQuery{
		Method:          "answerCallbackQuery",
		CallbackQueryId: msg.CallbackQuery.Id,
		Url:             answerRsc.Url + "?token=" + token, //正式流传递token
		//Url: answerRsc.Url + "?tguser=" + strconv.Itoa(int(acc.UserId)),
	}
	if msg.CallbackQuery.GameShortName == "majyonfttest" {
		answerMsg.Url = "https://club.loveat.cn/game/index.html#/index?v=13"
	}
	err = tm.sendCmdReply(ctx, answerMsg)
	if err != nil {
		tm.logger.Error("tgAPICBQLaunchGame failed", zap.Error(err))
		return err
	}
	return nil
}

func (tm *TelegramManager) tgAPICBQFightHiddenBoss(ctx context.Context, _ http.ResponseWriter, msg *mpb.TGMsgRecv,
) error {
	if msg == nil || msg.CallbackQuery == nil || msg.CallbackQuery.From == nil || msg.CallbackQuery.Message == nil {
		return nil
	}
	answerRsc := tm.rm.getTGReplyRsc(com.TGCBQ_FightHBoss, com.TGReplyType_Ok, com.TGReplySubType_Answer)
	if answerRsc == nil {
		tm.logger.Error("tgAPICBQFightHiddenBoss cmd reply rsc not found")
		return mpberr.ErrConfig
	}
	answerMsg := &mpb.TGAnswerCallbackQuery{
		Method:          "answerCallbackQuery",
		CallbackQueryId: msg.CallbackQuery.Id,
	}
	err := tm.sendCmdReply(ctx, answerMsg)
	if err != nil {
		tm.logger.Error("tgAPICBQFightHiddenBoss failed", zap.Error(err))
		return err
	}
	tm.logger.Debug("tgAPICBQFightHiddenBoss", zap.String("query data", msg.CallbackQuery.Data))

	acc, err := tm.getAccountByTGUser(ctx, msg.CallbackQuery.From)
	if err != nil {
		return err
	}

	strs := strings.Split(msg.CallbackQuery.Data, " ")
	if len(strs) < 3 {
		return mpberr.ErrParam
	}
	bossId := gutil.StrToUint32(strs[1])
	bossUuid := gutil.StrToUint64(strs[2])
	var sharerName string
	if len(strs) == 4 {
		sharerUid := gutil.StrToUint64(strs[3])
		sharer, err := tm.getAccount(ctx, sharerUid)
		if err == nil {
			sharerName = sharer.Nickname
		}
	}
	if bossId == 0 || bossUuid == 0 {
		return mpberr.ErrParam
	}

	fightResult, err := tm.fightHBoss(ctx, acc.UserId, bossId, bossUuid)
	if err != nil {
		return tm.handleFightBossErr(ctx, acc, msg, bossUuid, sharerName, err)
	}

	// if boss still alive, just update the caption.
	boss, err := tm.getHiddenBoss(ctx, acc.UserId, bossUuid)
	if err != nil {
		return err
	}

	// if boss die may remove the inlinekeyboard, change the caption.
	if fightResult.BossDie {
		replyMsg := &mpb.TGReply{
			Method:    "editMessageCaption",
			ChatId:    msg.CallbackQuery.Message.Chat.Id,
			MessageId: msg.CallbackQuery.Message.MessageId,
		}
		replyMsg.Caption, replyMsg.ParseMode, err = tm.generateHiddenBossDefeatedCaption(ctx, boss, fightResult,
			sharerName, acc)
		if err != nil {
			return err
		}

		err = tm.sendCmdReply(ctx, replyMsg)
		if err != nil {
			tm.logger.Error("tgAPICBQFightHiddenBoss send editMessageCaption failed", zap.Error(err))
			return err
		}
		return nil
	}
	// if boss still alive, just update the caption.
	replyMsg := &mpb.TGReply{
		Method:    "editMessageCaption",
		ChatId:    msg.CallbackQuery.Message.Chat.Id,
		MessageId: msg.CallbackQuery.Message.MessageId,
	}

	// Caption
	replyMsg.Caption, replyMsg.ParseMode, err = tm.generateHiddenBossCaption(ctx, boss, sharerName, acc.Lan)
	if err != nil {
		return err
	}
	// ReplyMarkup
	replyMsg.ReplyMarkup = &mpb.TGInlineKeyboardMarkup{
		InlineKeyboard: [][]*mpb.TGInlineKeyboardButton{{{
			Text:         tm.rm.getKeyString(com.STRING_KEY_BUTTON_FIGHT, acc.Lan),
			CallbackData: msg.CallbackQuery.Data,
		}}},
	}

	err = tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICBQFightHiddenBoss failed", zap.Error(err))
		return err
	}

	return nil
}

func (tm *TelegramManager) tgAPICBQSendSetting(ctx context.Context, _ http.ResponseWriter, msg *mpb.TGMsgRecv,
) error {
	if msg == nil || msg.CallbackQuery == nil || msg.CallbackQuery.From == nil || msg.CallbackQuery.Message == nil {
		return nil
	}

	acc, err := tm.getAccountByTGUser(ctx, msg.CallbackQuery.From)
	if err != nil {
		return err
	}

	answerRsc := tm.rm.getTGReplyRsc(com.TGCBQ_SendSetting, com.TGReplyType_Ok, com.TGReplySubType_Answer)
	if answerRsc != nil {
		answerMsg := &mpb.TGAnswerCallbackQuery{
			Method:          answerRsc.Method,
			CallbackQueryId: msg.CallbackQuery.Id,
		}
		err := tm.sendCmdReply(ctx, answerMsg)
		if err != nil {
			tm.logger.Error("tgAPICBQSendSetting failed", zap.Error(err))
			return err
		}
	}

	replyMsg := &mpb.TGReply{
		Method:      "editMessageReplyMarkup",
		ChatId:      msg.CallbackQuery.Message.Chat.Id,
		MessageId:   msg.CallbackQuery.Message.MessageId,
		ReplyMarkup: tm.rm.getTGInlineKeyBoard(com.TGCBQ_SendSetting, acc.Lan),
	}

	err = tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICBQSendSetting failed", zap.Error(err))
		return err
	}

	return nil
}

func (tm *TelegramManager) tgAPICBQSetLan(ctx context.Context, _ http.ResponseWriter, msg *mpb.TGMsgRecv,
) error {
	if msg == nil || msg.CallbackQuery == nil || msg.CallbackQuery.From == nil || msg.CallbackQuery.Message == nil {
		return nil
	}
	datas := strings.Split(msg.CallbackQuery.Data, "?")
	if len(datas) != 2 {
		return mpberr.ErrParam
	}
	values, err := url.ParseQuery(datas[1])
	if err != nil {
		tm.logger.Error("tgAPICBQSetLan parse callback_query_data_failed", zap.Error(err))
		return err
	}
	tm.logger.Debug("values", zap.Any("values", values))
	if values == nil || len(values["lan"]) == 0 {
		return err
	}
	lan := values["lan"][0]

	answerRsc := tm.rm.getTGReplyRsc(com.TGCBQ_SetLan, com.TGReplyType_Ok, com.TGReplySubType_Answer)
	if answerRsc != nil {
		answerMsg := &mpb.TGAnswerCallbackQuery{
			Method:          answerRsc.Method,
			CallbackQueryId: msg.CallbackQuery.Id,
		}
		err := tm.sendCmdReply(ctx, answerMsg)
		if err != nil {
			tm.logger.Error("tgAPICBQSetLan failed", zap.Error(err))
			return err
		}
	}

	//if account not exit, new an account for tg user
	_, err = tm.getAccountByTGUser(ctx, msg.CallbackQuery.From)
	if err != nil {
		return err
	}

	err = tm.setAccountTGLan(ctx, msg.CallbackQuery.From.Id, lan)
	if err != nil {
		return err
	}
	replyMsg := &mpb.TGReply{
		Method:      "editMessageReplyMarkup",
		ChatId:      msg.CallbackQuery.Message.Chat.Id,
		MessageId:   msg.CallbackQuery.Message.MessageId,
		ReplyMarkup: tm.rm.getTGInlineKeyBoard(com.TGCmd_Menu, lan),
	}
	err = tm.sendCmdReply(ctx, replyMsg)
	if err != nil {
		tm.logger.Error("tgAPICBQSetLan failed", zap.Error(err))
		return err
	}

	return nil
}

func (tm *TelegramManager) generateLoginToken(ctx context.Context, tgUser *mpb.TGUser) (string, error) {
	client, err := com.GetAccountServiceClient(ctx, tm.svc)
	if err != nil {
		return "", err
	}
	res, err := client.GenerateLoginToken(ctx, &mpb.ReqGenerateLoginToken{
		TgId:         tgUser.Id,
		FirstName:    tgUser.FirstName,
		LastName:     tgUser.LastName,
		LanguageCode: tgUser.LanguageCode,
	})
	if err != nil {
		return "", err
	}

	return res.Token, nil
}

func (tm *TelegramManager) generateSpecialLinkParams(params map[string]string) string {
	paramsStr := ""
	for k, v := range params {
		paramsStr += k + ":" + v + ";"
	}
	if len(paramsStr) == 0 {
		return ""
	}

	paramsStr = paramsStr[:len(paramsStr)-1]
	md5Str := gutil.MD5(paramsStr)[:6]
	paramsStr += md5Str
	paramsStr = base64.StdEncoding.EncodeToString([]byte(paramsStr))

	return paramsStr
}

func (tm *TelegramManager) parseSpecialLinkParams(text string) (params map[string]string, ok bool) {
	bys, err := base64.StdEncoding.DecodeString(text)
	if err != nil {
		return
	}
	text = string(bys)
	md5Str := text[len(text)-6:]
	fmt.Println(md5Str)
	text = text[:len(text)-6]
	if md5Str != gutil.MD5(text)[:6] {
		return
	}

	ok = true
	params = make(map[string]string)
	for _, v := range strings.Split(text, ";") {
		vs := strings.Split(v, ":")
		if len(vs) != 2 {
			continue
		}
		params[vs[0]] = vs[1]
	}
	return
}

func (tm *TelegramManager) getHiddenBoss(ctx context.Context, userId uint64, bossUUID uint64) (*mpb.ResGetHiddenBoss,
	error) {
	client, err := com.GetGameServiceClient(ctx, tm.svc)
	if err != nil {
		return nil, err
	}
	res, err := client.GetHiddenBoss(ctx, &mpb.ReqGetHiddenBoss{
		UserId:         userId,
		BossUuid:       bossUUID,
		WithBossDetail: true,
	})
	if err != nil {
		tm.logger.Error("getHiddenBoss failed", zap.Uint64("user_id", userId),
			zap.Uint64("boss_uuid", bossUUID), zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (tm *TelegramManager) randomHiddenBoss(ctx context.Context, userId uint64) (*mpb.ResGetHiddenBoss, error) {
	client, err := com.GetGameServiceClient(ctx, tm.svc)
	if err != nil {
		return nil, err
	}
	res, err := client.GetRandomHiddenBoss(ctx, &mpb.ReqGetHiddenBoss{
		UserId:         userId,
		WithBossDetail: true,
	})
	if err != nil {
		tm.logger.Error("randomHiddenBoss failed", zap.Uint64("user_id", userId), zap.Error(err))
		return nil, err
	}
	return res, nil
}

func (tm *TelegramManager) getAccountByTGUser(ctx context.Context, tgUser *mpb.TGUser) (*mpb.AccountInfo, error) {
	client, err := com.GetAccountServiceClient(ctx, tm.svc)
	if err != nil {
		return nil, err
	}
	res, err := client.GetAccountByTGUser(ctx, &mpb.ReqGetAccountByTGUser{
		TgId:         tgUser.Id,
		FirstName:    tgUser.FirstName,
		LastName:     tgUser.LastName,
		LanguageCode: tgUser.LanguageCode,
	})
	if err != nil {
		tm.logger.Error("getAccountByTGUser failed", zap.Any("tg_user", tgUser), zap.Error(err))
		return nil, err
	}
	return res.Account, nil
}

func (tm *TelegramManager) setAccountTGLan(ctx context.Context, tgUserId uint64, lan string) error {
	client, err := com.GetAccountServiceClient(ctx, tm.svc)
	if err != nil {
		return err
	}
	_, err = client.SetAccountTGLan(ctx, &mpb.ReqSetAccountTGLan{
		TgId:         tgUserId,
		LanguageCode: lan,
	})
	if err != nil {
		tm.logger.Error("setAccountTGLan failed", zap.Uint64("tg_user_id", tgUserId), zap.Error(err))
		return err
	}
	return nil
}

func (tm *TelegramManager) getAccount(ctx context.Context, userId uint64) (*mpb.AccountInfo, error) {
	client, err := com.GetAccountServiceClient(ctx, tm.svc)
	if err != nil {
		return nil, err
	}
	res, err := client.GetAccountByUserId(ctx, &mpb.ReqUserId{
		UserId: userId,
	})
	if err != nil {
		tm.logger.Error("getAccount failed", zap.Error(err))
		return nil, err
	}
	return res.Account, nil
}

func (tm *TelegramManager) calcTimeDurationStr(unixA, unixB int64) string {
	if unixA > unixB {
		unixA, unixB = unixB, unixA
	}
	duration := time.Duration(unixB-unixA) * time.Second

	return fmt.Sprintf("%02d:%02d:%02d", int(duration.Hours()), int(duration.Minutes())%60,
		int(duration.Seconds())%60)
}

func (tm *TelegramManager) handleFightBossErr(ctx context.Context, acc *mpb.AccountInfo, msg *mpb.TGMsgRecv,
	bossUuid uint64, sharerName string, inErr error) error {
	if !util.GRPCErrIs(inErr, mpberr.ErrHiddenBossFightCD) &&
		!util.GRPCErrIs(inErr, mpberr.ErrNFTWeaponLevel) &&
		!util.GRPCErrIs(inErr, mpberr.ErrNFTWeaponNotEquipped) &&
		!util.GRPCErrIs(inErr, mpberr.ErrHiddenBossExpired) &&
		!util.GRPCErrIs(inErr, mpberr.ErrHiddenBossDied) &&
		!util.GRPCErrIs(inErr, mpberr.ErrHiddenBossFought) &&
		!util.GRPCErrIs(inErr, mpberr.ErrHiddenBossNotFound) {
		tm.logger.Debug("handleFightBossErr can not handler inErr", zap.Error(inErr),
			zap.Bool("", util.GRPCErrIs(inErr, mpberr.ErrHiddenBossFought)))
		return inErr
	}
	replyMsg := &mpb.TGReply{
		Method:    "editMessageCaption",
		ChatId:    msg.CallbackQuery.Message.Chat.Id,
		MessageId: msg.CallbackQuery.Message.MessageId,
	}
HandleFightBossErrNotFound:
	if util.GRPCErrIs(inErr, mpberr.ErrHiddenBossNotFound) ||
		util.GRPCErrIs(inErr, mpberr.ErrHiddenBossExpired) ||
		util.GRPCErrIs(inErr, mpberr.ErrHiddenBossDied) {
		replyMsg.Caption, replyMsg.ParseMode = tm.generateHiddenBossNotFoundCaption(acc.Lan)
		inErr = tm.sendCmdReply(ctx, replyMsg)
		if inErr != nil {
			tm.logger.Error("handleFightBossErr failed", zap.Error(inErr))
			return inErr
		}
		return nil
	}

	boss, newErr := tm.getHiddenBoss(ctx, acc.UserId, bossUuid)
	if newErr != nil {
		if util.GRPCErrIs(newErr, mpberr.ErrHiddenBossNotFound) {
			inErr = newErr
			goto HandleFightBossErrNotFound
		}
		return newErr
	}
	replyMsg.Caption, replyMsg.ParseMode, newErr = tm.generateHiddenBossCaption(ctx, boss, sharerName, acc.Lan)
	if newErr != nil {
		return newErr
	}
	tm.logger.Info("inErr", zap.Error(inErr))
	if util.GRPCErrIs(inErr, mpberr.ErrHiddenBossFightCD) {
		tm.logger.Info("1")
		replyMsg.Caption += "\n" + fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_ERR_IN_FIGHT_CD, acc.Lan),
			acc.Nickname)
	} else if util.GRPCErrIs(inErr, mpberr.ErrNFTWeaponNotEquipped) {
		tm.logger.Info("2")
		replyMsg.Caption += "\n" + fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_ERR_NO_REQUESTED_NFT_WEAPON, acc.Lan),
			acc.Nickname)
	} else if util.GRPCErrIs(inErr, mpberr.ErrNFTWeaponLevel) {
		tm.logger.Info("3")
		replyMsg.Caption += "\n" + fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_ERR_NOT_MEET_NFT_WEAPON_LEVEL, acc.Lan),
			acc.Nickname)
	} else if util.GRPCErrIs(inErr, mpberr.ErrHiddenBossFought) {
		tm.logger.Info("4")
		replyMsg.Caption += "\n" + fmt.Sprintf(tm.rm.getKeyString(com.STRING_KEY_ERR_ALREADY_FOUGHT, acc.Lan),
			acc.Nickname)
	}
	// ReplyMarkup
	replyMsg.ReplyMarkup = &mpb.TGInlineKeyboardMarkup{
		InlineKeyboard: [][]*mpb.TGInlineKeyboardButton{{{
			Text:         tm.rm.getKeyString(com.STRING_KEY_BUTTON_FIGHT, acc.Lan),
			CallbackData: msg.CallbackQuery.Data},
		}},
	}

	inErr = tm.sendCmdReply(ctx, replyMsg)
	if inErr != nil {
		tm.logger.Error("handleFightBossErr failed", zap.Error(inErr))
		return inErr
	}

	return inErr
}

func (tm *TelegramManager) fightHBoss(ctx context.Context, userId uint64, bossId uint32, bossUuid uint64) (
	*mpb.ResFight, error) {
	client, err := com.GetGameServiceClient(ctx, tm.svc)
	if err != nil {
		return nil, err
	}
	rpcRes, err := client.Fight(ctx, &mpb.ReqFight{UserId: userId, BossId: bossId, BossUuid: bossUuid,
		WithBossDetail: false})
	if err != nil {
		return nil, err
	}
	return rpcRes, err
}

func (tm *TelegramManager) getItemsRsc(ctx context.Context, itemIds []uint32) (map[uint32]*mpb.ItemRsc, error) {
	if len(itemIds) == 0 {
		return make(map[uint32]*mpb.ItemRsc), nil
	}
	client, err := com.GetItemServiceClient(ctx, tm.svc)
	if err != nil {
		return nil, err
	}

	res, err := client.GetItemsRsc(ctx, &mpb.ReqGetItemsRsc{ItemIds: itemIds})
	if err != nil {
		return nil, err
	}

	if res.ItemsRsc == nil {
		res.ItemsRsc = make(map[uint32]*mpb.ItemRsc)
	}

	return res.ItemsRsc, nil
}
