package apiproxy

import (
	"fmt"
	"github.com/oldjon/gutil"
	gcsv "github.com/oldjon/gutil/csv"
	gdm "github.com/oldjon/gutil/dirmonitor"
	com "gitlab.com/morbackend/mor_services/common"
	"gitlab.com/morbackend/mor_services/mpb"
	"go.uber.org/zap"
	"strings"
)

const (
	//csvSuffix   = ".csv"
	baseCSVPath = "./resources/apiproxy/"

	tgReplyCSV          = "TelegramReply.csv"
	tgInlineKeyBoardCSV = "TelegramInlineKeyboard.csv"
	tgGameCSV           = "TelegramGame.csv"
	tgSpecialLinkCSV    = "TelegramSpecialLink.csv"
	tgStringsCSV        = "TelegramStrings.csv"
)

type apiProxyResourceMgr struct {
	logger *zap.Logger
	dm     *gdm.DirMonitor
	mtr    *metrics

	tgReply                map[string]map[string]map[string]*mpb.TGReplyRsc
	tgCmdInlineKeyBoardMap map[string][][]*mpb.TGInlineKeyboardRsc
	tgGameShortNameMap     map[string]*mpb.TGGameRsc
	tgSpecialLinkMap       map[string]*mpb.TGSpecialLinkRsc
	tgStringMap            map[string]*mpb.TGStringsRsc
}

func newAPIProxyResourceMgr(logger *zap.Logger, mtr *metrics) (*apiProxyResourceMgr, error) {
	rMgr := &apiProxyResourceMgr{
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

func (rm *apiProxyResourceMgr) load() error {
	var err error
	err = rm.dm.BindAndExec(tgReplyCSV, rm.loadTGReply)
	if err != nil {
		fmt.Println(1)
		return err
	}
	err = rm.dm.BindAndExec(tgInlineKeyBoardCSV, rm.loadTGInlineKeyboard)
	if err != nil {
		fmt.Println(2)
		return err
	}
	err = rm.dm.BindAndExec(tgGameCSV, rm.loadTGGame)
	if err != nil {
		fmt.Println(3)
		return err
	}
	err = rm.dm.BindAndExec(tgSpecialLinkCSV, rm.loadTGSpecialLink)
	if err != nil {
		fmt.Println(3)
		return err
	}
	err = rm.dm.BindAndExec(tgStringsCSV, rm.loadTGStrings)
	if err != nil {
		fmt.Println(3)
		return err
	}
	return nil
}

func (rm *apiProxyResourceMgr) watch() error {
	return rm.dm.StartWatch()
}

func (rm *apiProxyResourceMgr) loadTGInlineKeyboard(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[string][][]*mpb.TGInlineKeyboardRsc)
	for _, data := range datas {
		node := &mpb.TGInlineKeyboardRsc{
			Cmd:          data["cmd"],
			Switch:       gcsv.StrToBool(data["switch"]),
			Row:          gcsv.StrToUint32(data["row"]),
			Text:         data["text"],
			Url:          data["url"],
			CallbackData: data["callbackdata"],
			CallbackGame: gcsv.StrToBool(data["callbackgame"]),
		}

		l := m[node.Cmd]
		var ok bool
		for i, ll := range l {
			if ll[0].Row == node.Row {
				ok = true
				ll = append(ll, node)
				l[i] = ll
				break
			}
		}
		if !ok {
			l = append(l, []*mpb.TGInlineKeyboardRsc{node})
		}

		m[node.Cmd] = l
		rm.logger.Debug("loadTGInlineKeyboard read:", zap.Any("row", node))
	}

	rm.tgCmdInlineKeyBoardMap = m
	rm.logger.Debug("loadTGInlineKeyboard read finish:", zap.Any("rows", m))

	return nil
}

func (rm *apiProxyResourceMgr) getTGInlineKeyBoard(cmd string, lan string) *mpb.TGInlineKeyboardMarkup {
	ret := &mpb.TGInlineKeyboardMarkup{}
	l := rm.tgCmdInlineKeyBoardMap[cmd]
	for _, ll := range l {
		row := make([]*mpb.TGInlineKeyboardButton, 0, 1)
		for _, v := range ll {
			if !v.Switch {
				continue
			}
			row = append(row, &mpb.TGInlineKeyboardButton{
				Text:         rm.getKeyString(v.Text, lan),
				Url:          v.Url,
				CallbackData: v.CallbackData,
				CallbackGame: gutil.If(v.CallbackGame, &mpb.TGCallbackGame{}, nil),
			})
		}
		if len(row) > 0 {
			ret.InlineKeyboard = append(ret.InlineKeyboard, row)
		}
	}
	return ret
}

func (rm *apiProxyResourceMgr) loadTGReply(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[string]map[string]map[string]*mpb.TGReplyRsc)
	for _, data := range datas {
		node := &mpb.TGReplyRsc{
			Cmd:           data["cmd"],
			Type:          data["type"],
			SubType:       data["subtype"],
			Method:        data["method"],
			Text:          data["text"],
			Photo:         data["photo"],
			GameShortName: data["gameshortname"],
			Url:           data["url"],
		}
		sm, ok := m[node.Cmd]
		if !ok {
			sm = make(map[string]map[string]*mpb.TGReplyRsc)
			m[node.Cmd] = sm
		}
		ssm, ok := sm[node.Type]
		if !ok {
			ssm = make(map[string]*mpb.TGReplyRsc)
			sm[node.Type] = ssm
		}
		ssm[node.SubType] = node

		rm.logger.Debug("loadTGReply read:", zap.Any("row", node))
	}

	rm.tgReply = m
	rm.logger.Debug("loadTGReply read finish:", zap.Any("rows", m))
	return nil
}

func (rm *apiProxyResourceMgr) getTGReplyRsc(cmd string, tp, subTp string) *mpb.TGReplyRsc {
	return rm.tgReply[cmd][tp][subTp]
}

func (rm *apiProxyResourceMgr) getEmailAddrs() []*mpb.EmailAddrRsc {
	return nil
}

func (rm *apiProxyResourceMgr) loadTGGame(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[string]*mpb.TGGameRsc)
	for _, data := range datas {
		node := &mpb.TGGameRsc{
			GameName:      data["gamename"],
			GameShortName: data["gameshortname"],
			GameUrl:       data["gameurl"],
		}

		m[node.GameShortName] = node
		rm.logger.Debug("loadTGGame read:", zap.Any("row", node))
	}

	rm.tgGameShortNameMap = m
	rm.logger.Debug("loadTGGame read finish:", zap.Any("rows", m))
	return nil
}

func (rm *apiProxyResourceMgr) getTGGameRscByGameShortName(name string) *mpb.TGGameRsc {
	return rm.tgGameShortNameMap[name]
}

func (rm *apiProxyResourceMgr) loadTGSpecialLink(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[string]*mpb.TGSpecialLinkRsc)
	for _, data := range datas {
		node := &mpb.TGSpecialLinkRsc{
			Cmd:     data["cmd"],
			GameBot: data["gamebot"],
			Comment: data["comment"],
		}

		m[node.Cmd] = node
		rm.logger.Debug("loadTGSpecialLink read:", zap.Any("row", node))
	}

	rm.tgSpecialLinkMap = m
	rm.logger.Debug("loadTGSpecialLink read finish:", zap.Any("rows", m))
	return nil
}

func (rm *apiProxyResourceMgr) getTGSpecialRsc(cmd string) *mpb.TGSpecialLinkRsc {
	return rm.tgSpecialLinkMap[cmd]
}

func (rm *apiProxyResourceMgr) loadTGStrings(csvPath string) error {
	datas, err := gcsv.ReadCSV2Array(csvPath)
	if err != nil {
		rm.logger.Error(fmt.Sprintf("load %s failed: %s", csvPath, err.Error()))
		return err
	}
	m := make(map[string]*mpb.TGStringsRsc)
	for _, data := range datas {
		node := &mpb.TGStringsRsc{
			Key:    data["key"],
			ZhHans: data["zh_hans"],
			ZhHant: data["zh_hant"],
			En:     strings.ReplaceAll(data["en"], "`", ","),
		}

		m[node.Key] = node
		rm.logger.Debug("loadTGStrings read:", zap.Any("row", node))
	}

	rm.tgStringMap = m
	rm.logger.Debug("loadTGStrings read finish:", zap.Any("rows", m))
	return nil
}

func (rm *apiProxyResourceMgr) getKeyString(key string, lan string) string {
	rsc, ok := rm.tgStringMap[key]
	if !ok {
		return ""
	}
	rm.logger.Debug("lang", zap.String("lan", lan))
	switch lan {
	case com.LAN_ZH_HANS:
		return rsc.ZhHans
	case com.LAN_ZH_HANT:
		return rsc.ZhHant
	case com.LAN_EN:
		return rsc.En
	default:
		return rsc.En
	}
	return ""
}
