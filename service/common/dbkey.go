package common

import (
	"fmt"
	"gitlab.com/morbackend/mor_services/mpb"
	"reflect"
	"strings"
)

const (
	// tcp
	tcpGatewayNodeFmt = "tcpnode:%d"

	// account
	userIdIndexKeyFmt         = "uidindex"
	accountKeyFmt             = "acc:%d"
	accountUIDKeyFmt          = "accuid:%s"
	nonceKeyFmt               = "nonce:%s"
	walletAccKeyFmt           = "walletacc:%s"
	emailSendDailyLimitKeyFmt = "esdl:%s:%s"
	emailBindCodeKeyFmt       = "ebc:%s"
	emailAcc                  = "emailacc:%s"
	emailResetPWCodeKeyFmt    = "erpwc:%s"
	emailResetPWNonceKeyFmt   = "erpwn:%s"
	tokenKeyFmt               = "token:%s"
	uidTokenKeyFmt            = "uidtoken:%d"
	deviceAccountsKeyFmt      = "devaccs:%s"
	loginInfoKeyFmt           = "login:%d"
	loginTokenKeyFmt          = "logintoken:%s"

	// game
	energyKeyFmt                = "energy:%d"
	fightHistoryKeyFmt          = "fighthis:%d"
	userHiddenBossCDKeyFmt      = "uhbosscd:%d"
	hiddenBossKeyFmt            = "hboss:%d"
	hiddenBossFindHistoryKeyFmt = "hbossfindhis:%d"
	hiddenBossPoolKeyFmt        = "hbosspool"
	bossDefeatHistoryKeyFmt     = "bdhis:%d"
	userPVPInfoKeyFmt           = "pvpinfo:%d"
	userPVPRankKeyFmt           = "upvprank:%d"
	pvpRanksKeyFmt              = "pvpranks:%d"
	pvpRankIndexKeyFmt          = "pvprankindex"
	pvpRankRewardMarkKeyFmt     = "pvprrmark:%d:%s"
	pvpManaPoolKeyFmt           = "pvpmanapool:%d:%d"
	pvpHistoryKey               = "pvphis"
	buffCardsValidKey           = "buffcardsvalid:%d"
	buffCardOptionsKey          = "buffcardopts:%d"
	buffCardChoseStatusKey      = "buffcardchosestatus:%d"

	// item
	itemsKeyFmt       = "items:%d"
	itemsShardKeyFmt  = "items:%d"
	itemsUShardKeyFmt = "uitems:%d"
	walletKeyFmt      = "wallet:%d"
	baseEquipsKeyFmt  = "baseequips:%d"

	// market
	allOrdersOnSellKeyFmt   = "allordersonsell"
	goodsOrdersOnSellKeyFmt = "itemordersonsell:%d"
	myOrdersOnSellKeyFmt    = "myordersonsell:%d"
	myOrdersSoldKeyFmt      = "myorderssold:%d"
	myOrderPurchasedKeyFmt  = "myorderpurchased:%d"
	goodsOrderInfoKeyFmt    = "goodsorderinfo:%d"

	// timer
	pvpSettleMarkKeyFmt = "pvpsettlemark:%s"
)

type KeyObjTypeMapping struct {
	mapping map[string]string
}

func (kom *KeyObjTypeMapping) Mapping() map[string]string {
	return kom.mapping
}

var DBKeyObjTypeMapping = &KeyObjTypeMapping{
	mapping: map[string]string{
		// account
		accountKeyFmt:           reflect.TypeOf(mpb.DBAccountInfo{}).String(),
		accountUIDKeyFmt:        reflect.TypeOf(uint64(0)).String(),
		nonceKeyFmt:             reflect.TypeOf(0).String(),
		walletAccKeyFmt:         reflect.TypeOf(mpb.DBWalletAcc{}).String(),
		emailBindCodeKeyFmt:     reflect.TypeOf("").String(),
		emailAcc:                reflect.TypeOf(uint64(0)).String(),
		emailResetPWCodeKeyFmt:  reflect.TypeOf("").String(),
		emailResetPWNonceKeyFmt: reflect.TypeOf("").String(),
		// login
		tokenKeyFmt:    reflect.TypeOf(mpb.DBTokenInfo{}).String(),
		uidTokenKeyFmt: reflect.TypeOf("").String(),

		// game
		energyKeyFmt:                reflect.TypeOf(mpb.DBEnergy{}).String(),
		fightHistoryKeyFmt:          reflect.TypeOf(mpb.DBFightHistory{}).String(),
		userHiddenBossCDKeyFmt:      reflect.TypeOf(int64(0)).String(),
		hiddenBossKeyFmt:            reflect.TypeOf(mpb.DBHiddenBoss{}).String(),
		hiddenBossFindHistoryKeyFmt: reflect.TypeOf(mpb.DBHiddenBossFindHistory{}).String(),

		// item
		walletKeyFmt:     reflect.TypeOf(mpb.DBWallet{}).String(),
		baseEquipsKeyFmt: reflect.TypeOf(mpb.DBBaseEquips{}).String(),
	},
}

// Get key
// tcp
func TCPGatewayNodeKey(userId uint64) string {
	return fmt.Sprintf(tcpGatewayNodeFmt, userId)
}

// account
func UserIdIndexKey() string {
	return userIdIndexKeyFmt
}

func AccountKey(userId uint64) string {
	return fmt.Sprintf(accountKeyFmt, userId)
}

func AccountUIDKey(acc string) string {
	return fmt.Sprintf(accountUIDKeyFmt, acc)
}

func NonceKey(nonce string) string {
	return fmt.Sprintf(nonceKeyFmt, nonce)
}

func WalletAccKey(walletAddr string) string {
	return fmt.Sprintf(walletAccKeyFmt, walletAddr)
}

func EmailSendDailyLimitKey(emailAddr string, date string) string {
	return fmt.Sprintf(emailSendDailyLimitKeyFmt, emailAddr, date)
}

func EmailBindCodeKey(emailAddr string) string {
	return fmt.Sprintf(emailBindCodeKeyFmt, emailAddr)
}

func EmailAccKey(emailAddr string) string {
	return fmt.Sprintf(emailAcc, emailAddr)
}

func EmailResetPasswordValidationCodeKey(emailAddr string) string {
	return fmt.Sprintf(emailResetPWCodeKeyFmt, emailAddr)
}

func EmailResetPasswordNonceKey(emailAddr string) string {
	return fmt.Sprintf(emailResetPWNonceKeyFmt, emailAddr)
}

func TokenKey(token string) string {
	return fmt.Sprintf(tokenKeyFmt, token)
}

func UIDTokenKey(userId uint64) string {
	return fmt.Sprintf(uidTokenKeyFmt, userId)
}

func DeviceAccountsKey(deviceId string) string {
	return fmt.Sprintf(deviceAccountsKeyFmt, deviceId)
}

func LoginInfoKey(userId uint64) string {
	return fmt.Sprintf(loginInfoKeyFmt, userId)
}

func LoginTokenKey(token string) string {
	return fmt.Sprintf(loginTokenKeyFmt, token)
}

// user

// item
func ItemsKey(userId uint64) string {
	return fmt.Sprintf(itemsKeyFmt, userId)
}

func ItemsShardKeys(ids []uint32) []string {
	if len(ids) == 0 {
		return nil
	}
	ret := make([]string, 0, len(ids))
	for _, id := range ids {
		ret = append(ret, fmt.Sprintf(itemsShardKeyFmt, id))
	}
	return ret
}

func ItemsShardKey(id uint32) string {
	return fmt.Sprintf(itemsShardKeyFmt, id)
}

func ItemsShardKeysByShardCnt(cnt uint32) []string {
	if cnt == 0 {
		return nil
	}
	ret := make([]string, cnt)
	for i := uint32(0); i < cnt; i++ {
		ret[i] = fmt.Sprintf(itemsShardKeyFmt, i)
	}
	return ret
}

func UItemsShardKeys(ids []uint32) []string {
	if len(ids) == 0 {
		return nil
	}
	ret := make([]string, 0, len(ids))
	for _, id := range ids {
		ret = append(ret, fmt.Sprintf(itemsUShardKeyFmt, id))
	}
	return ret
}

func UItemsShardKey(id uint32) string {
	return fmt.Sprintf(itemsUShardKeyFmt, id)
}

func UItemsShardKeysByShardCnt(cnt uint32) []string {
	if cnt == 0 {
		return nil
	}
	ret := make([]string, cnt)
	for i := uint32(0); i < cnt; i++ {
		ret[i] = fmt.Sprintf(itemsUShardKeyFmt, i)
	}
	return ret
}

func WalletKey(userId uint64) string {
	return fmt.Sprintf(walletKeyFmt, userId)
}

func BaseEquipsKey(userId uint64) string {
	return fmt.Sprintf(baseEquipsKeyFmt, userId)
}

// mail

// social

// nft

// game
func EnergyKey(userId uint64) string {
	return fmt.Sprintf(energyKeyFmt, userId)
}

func FightHistoryKey(userId uint64) string {
	return fmt.Sprintf(fightHistoryKeyFmt, userId)
}

func UserHiddenBossCDKey(userId uint64) string {
	return fmt.Sprintf(userHiddenBossCDKeyFmt, userId)
}

func HiddenBossKey(bossUUID uint64) string {
	return fmt.Sprintf(hiddenBossKeyFmt, bossUUID)
}

func HiddenBossPoolKey() string {
	return hiddenBossPoolKeyFmt
}

func HiddenBossFindHistoryKey(userId uint64) string {
	return fmt.Sprintf(hiddenBossFindHistoryKeyFmt, userId)
}

func BossDefeatHistory(userId uint64) string {
	return fmt.Sprintf(bossDefeatHistoryKeyFmt, userId)
}

func UserPVPInfoKey(userId uint64) string {
	return fmt.Sprintf(userPVPInfoKeyFmt, userId)
}

func UserPVPRankKey(userId uint64) string {
	return fmt.Sprintf(userPVPRankKeyFmt, userId)
}

func PVPRanksKey(shardId uint32) string {
	return fmt.Sprintf(pvpRanksKeyFmt, shardId)
}

func PVPRankIndexKey() string {
	return pvpRankIndexKeyFmt
}

func PVPSettleMarkKey(date string) string {
	return fmt.Sprintf(pvpSettleMarkKeyFmt, date)
}

func PVPRankRewardMarkKey(userId uint64, date string) string {
	date = strings.ReplaceAll(date, "-", "")
	return fmt.Sprintf(pvpRankRewardMarkKeyFmt, userId, date)
}

func PVPManaPoolKey(season, date uint32) string {
	return fmt.Sprintf(pvpManaPoolKeyFmt, season, date)
}

func PVPHistoryKey() string {
	return pvpHistoryKey
}

func BuffCardsValidKey(userId uint64) string {
	return fmt.Sprintf(buffCardsValidKey, userId)
}

func BuffCardOptionsKey(userId uint64) string {
	return fmt.Sprintf(buffCardOptionsKey, userId)
}

func BuffCardChoseStatusKey(userId uint64) string {
	return fmt.Sprintf(buffCardChoseStatusKey, userId)
}

func AllOrderOnSellKey() string {
	return allOrdersOnSellKeyFmt
}

func GoodsOrdersOnSellKey(itemId uint32) string {
	return fmt.Sprintf(goodsOrdersOnSellKeyFmt, itemId)
}

func MyOrdersOnSellKey(userId uint64) string {
	return fmt.Sprintf(myOrdersOnSellKeyFmt, userId)
}

func MyOrdersSoldKey(userId uint64) string {
	return fmt.Sprintf(myOrdersSoldKeyFmt, userId)
}

func MyOrderPurchasedKey(userId uint64) string {
	return fmt.Sprintf(myOrderPurchasedKeyFmt, userId)
}

func GoodsOrderInfoKey(orderUuid uint64) string {
	return fmt.Sprintf(goodsOrderInfoKeyFmt, orderUuid)
}
