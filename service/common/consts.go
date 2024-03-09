package common

import "time"

const (
	JWTGatewayTokenKey   = "gateway_token_key"
	TCPMarshallerType    = "tcp_msg_marshaller_type"
	GMJWTGatewayTokenKey = "gm_gateway_token_key"
)

const (
	UserIdStart           = 10000000
	DefaultPassword       = "123456"
	PasswordSalt          = "lalalademala@#!asd"
	MinAccountLength      = 4
	MaxAccountLength      = 30
	MinAdminAccountLength = 4
	MaxAdminAccountLength = 30
	RateBase              = 10000
	ManaFactor            = 10000
)

const (
	Secs1Min                  = 60
	Secs10Mins                = 10 * Secs1Min
	Dur10Mins                 = 10 * time.Minute
	Secs1Hour                 = 6 * Secs1Min
	Secs1Day                  = 24 * Secs1Hour
	Dur1Day                   = 24 * time.Hour
	Secs7Days                 = Secs1Day * 7
	Secs1Week                 = Secs7Days
	ResetHour                 = 4
	SecsRestHour              = Secs1Hour * ResetHour
	CtxTimeout                = 10 * time.Second
	TokenExpireDuration       = 7 * 86400 * time.Second
	LoginTokenExpiredDuration = 60 * time.Minute
)

const (
	DBBatchNum100   = 100
	DBBatchNum10000 = 10000
	//DBBatchNum100000 = 100000
	KeepTTL = time.Duration(-1)
)

const (
	// snowflake name
	SnowflakeItemUUID        = "item_uuid"
	SnowflakeBossUUID        = "boss_uuid"
	SnowflakeTransactionUUID = "transaction_uuid"
	SnowflakeMailUUID        = "mail_uuid"
	SnowflakeMarketOrderUUID = "market_order_uuid"
)

const (
	NonceLen             = 6
	VCodeLen             = 6
	EmailSendDailyLimit  = 50
	PasswordLen          = 32
	DefaultWebNFTPageNum = 10
	MaxWebNFTPageNum     = 200
)

const (
	AptosNFTTokenIdLen = 64
)

const (
	TGGameBot         = "fishilltegame"
	TGCmd_Echo        = "/echo"
	TGCmd_Menu        = "/menu"
	TGCmd_HiddenBoss  = "/hiddenboss"
	TGCmd_Game        = "/game"
	TGCBQ_Echo        = "/echo"
	TGCBQ_SendGame    = "/cbqsendgame"
	TGCBQ_LaunchGame  = "/cbqlaunchgame"
	TGCBQ_FightHBoss  = "/cbqfighthboss"
	TGCBQ_SendSetting = "/cbqsendsetting"
	TGCBQ_SetLan      = "/cbqsetlan"
)

const (
	TGReplyType_Ok          = "ok"
	TGReplyType_Err         = "err"
	TGReplySubType_Answer   = "answer"
	TGReplySubType_NotFound = "notfound"
)

const (
	PVPRankShardUserCnt = 1000
)
