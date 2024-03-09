package gameservice

import (
	"sync"
	"time"

	"gitlab.com/morbackend/mor_services/mpb"
)

const (
	cacheTime5Mins = 5 * 60
)

type gameCacheMgr struct {
	pvpHisMux      sync.RWMutex
	pvpHisUpdateAt int64
	pvpHisCache    []*mpb.PVPHistory
}

func newGameCacheMgr() *gameCacheMgr {
	return &gameCacheMgr{}
}

func (gcm *gameCacheMgr) getPVPHistoryCache() ([]*mpb.PVPHistory, bool) {
	gcm.pvpHisMux.RLock()
	defer gcm.pvpHisMux.RUnlock()
	if gcm.pvpHisUpdateAt < time.Now().Unix()-cacheTime5Mins {
		return nil, false
	}
	return gcm.pvpHisCache, true
}

func (gcm *gameCacheMgr) setPVPHistoryCache(list []*mpb.PVPHistory) {
	gcm.pvpHisMux.Lock()
	defer gcm.pvpHisMux.Unlock()
	gcm.pvpHisCache = list
	gcm.pvpHisUpdateAt = time.Now().Unix()
}
