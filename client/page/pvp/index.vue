<template>
  <div class="pvp-view">
    <!-- headbox -->
    <headers/>

    <div class="prize-pool">
      <div class="prize-content">
        <div class="pool-div" @click="onPopBox(3)">
          <div class="prize-num prize-stroke"><em>x</em>{{curPrizepool}}</div>
          <div class="prize-time time-stroke">{{curSeasonDate || '---'}}</div>
        </div>
        <div class="check-time text-stroke">
          <div v-if="!this.counting">{{this.countdownTime}}</div>
          <div v-else>Settling<span class="dot"></span></div>
          <div class="history hbtn" @click="getPVPHis()">
            <span class="text-stroke">History</span>
          </div>
        </div>
      </div>
    </div>


    <div v-if="this.findingRankList" class="rank-list">
      <div class="find-notice">
        <span class="text-stroke">Refreshing the rankings<span class="dot"></span></span>
      </div>
    </div>
    <div v-else class="rank-list">
      <div class="rank" v-for="rk in this.rankList" :key="rk.rank" @click="showDetail(rk.rank, 1)">
        <div class="rank-left text-stroke">
          <div class="tl">NO.</div>
          <div class="num">{{rk.rank}}</div>
        </div>
        <div class="rank-center text-stroke">
          <div class="rk-player">{{rk.account_info.nickname}}</div>
          <div class="rk-atr">
            <div>HP:{{rk.attr.hp}}</div>
            <div>ATK:{{rk.attr.atk}}</div>
            <div class="detail">
              <img src="~@/assets/info.png" alt="detail" />
            </div>
          </div>
        </div>
        <div class="rank-right">
          <div v-if="rk.rwd && rk.rwd.length >= 1" class="rds" v-for="item in rk.rwd" :key="item.id">
            <img v-if="item.id == 10" src="~@/assets/items/10.png" />
            <img v-if="item.id == 91000001" src="~@/assets/ETHERicon.png" />
            <div v-if="item.id == 10" class="itemNum text-stroke">{{item.num}}</div>
            <div v-if="item.id == 91000001" class="itemNum text-stroke">{{item.num / 100}}%</div>
          </div>
          <div v-if="!rk.rwd || rk.rwd.length < 1" class="norwd">
            <span class="text-stroke">No rewards</span>
          </div>
        </div>
      </div>
    </div>

    <div class="my-rank">
      <div class="rank" @click="showDetail(0,0)">
        <div class="rank-left text-stroke">
          <div class="tl">NO.</div>
          <div class="num">{{myRank.rank > 0 ? myRank.rank : '--'}}</div>
        </div>
        <div class="rank-center text-stroke">
          <div class="rk-player">{{myRank.account_info.nickname}}</div>
          <div class="rk-atr">
            <div>HP:{{myRank.attr.hp}}</div>
            <div>ATK:{{myRank.attr.atk}}</div>
            <div class="detail">
              <img src="~@/assets/info.png" alt="detail" />
            </div>
          </div>
        </div>
        <div class="rank-right">
          <div v-if="myRank.rwd && myRank.rwd.length >= 1" class="rds" v-for="item in myRank.rwd" :key="item.id">
            <img v-if="item.id == 10" src="~@/assets/items/10.png" />
            <img v-if="item.id == 91000001" src="~@/assets/ETHERicon.png" />
            <div v-if="item.id == 10" class="itemNum text-stroke">{{item.num}}</div>
            <div v-if="item.id == 91000001" class="itemNum text-stroke">{{item.num / 100}}%</div>
          </div>
          <div v-if="!myRank.rwd || myRank.rwd.length < 1">
            <span class="norwd text-stroke">No rewards</span>
          </div>
        </div>
      </div>
      <div v-if="!this.counting" class="btn" @click="showTarget()"><span class="text-stroke">Challenge</span></div>
      <div v-else class="text-stroke"><span class="rest">Take a break and wait for the system to settle the reward</span></div>
    </div>

    <div class="tips-div">
      <Tip ref="tip"></Tip>
    </div>
    <foots/>

    <div class="popbox pvpbox" v-if="istip">
      <div class="bgs"></div>
      <!-- type 1-->
      <div class="popbody" v-if="isType === 1">
        <div class="head">
          <span class="text-stroke">Challenge List</span>
        </div>
        <div v-if="this.findingTarget" class="middle">
          <div class="target-notice">
            <div class="text-stroke">Looking for players<span class="dot"></span></div>
          </div>
        </div>
        <div v-else class="middle">
          <div v-if="this.targetList && this.targetList.length > 0" class="cl-list">
            <div class="cl-rank" :class="targetSelected === tar.rank ? 'choiced':''" v-for="tar in this.targetList" :key="tar.rank" @click="selectTarget(tar.rank)">
              <div class="rank-left text-stroke">
                <div class="tl">NO.</div>
                <div class="num">{{tar.rank}}</div>
              </div>
              <div class="rank-center text-stroke">
                <div class="rk-player2">{{tar.account_info.nickname}}</div>
                <div class="pvp-atr2">
                  <div>HP:{{tar.attr.hp}}</div>
                  <div>ATK:{{tar.attr.atk}}</div>
                  <div class="detail" @click="showDetail(tar.rank,2)">
                    <img src="~@/assets/info.png" alt="detail" />
                  </div>
                </div>
              </div>
              <div class="rank-right">
                <div v-if="tar.rwd && tar.rwd.length >= 1" class="rds" v-for="item in tar.rwd" :key="item.id">
                  <img v-if="item.id == 10" src="~@/assets/items/10.png" />
                  <img v-if="item.id == 91000001" src="~@/assets/ETHERicon.png" />
                  <div v-if="item.id == 10" class="itemNum text-stroke">{{item.num}}</div>
                  <div v-if="item.id == 91000001" class="itemNum text-stroke">{{item.num / 100}}%</div>
                </div>
                <div v-if="!tar.rwd || tar.rwd.length < 1" class="norwd">
                  <span class="text-stroke">No rewards</span>
                </div>
              </div>
            </div>
          </div>
          <div v-else class="target-notice">
            <div v-if="this.myRank.rank === 1" class="no1 text-stroke">You are the No.1</div>
            <div class="text-stroke">No player found to challenge.</div>
          </div>
        </div>
        <div class="foots">
          <div class="cl-cost">
            <p v-if="!this.counting && this.myRank.challenger_cnt > 0" class="notice text-stroke">Free challenges remaining today: {{this.myRank.challenger_cnt}}</p>
            <p v-if="!this.counting && this.myRank.challenger_cnt <= 0" class="text-stroke">COST:<img src="~@/assets/ETHERicon.png" /><span>{{this.pvpComsume}}</span></p>
          </div>
          <div v-if="!this.counting" class="btn f30" style="margin-top: 10px;" @click="gotoPvp()"><span class="text-stroke">Fight</span></div>
          <div v-if="!this.counting && this.showFresh" class="fresh-cl" @click="freshTarget()"><img src="~@/assets/pvp/fresh-cl.png" /></div>
          <div v-if="this.counting" class="f30"><span class="rest text-stroke">Settling...</span></div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>

      <div class="popbody" v-if="isType === 2">
        <div class="head">
          <span class="text-stroke">Rank Rewards</span>
        </div>
        <div class="middle" v-if="thisLastRwd && thisLastRwd.rank > 0">
          <div class="rk-rwd">
            <div class="pic" v-for="items in thisLastRwd.rewardsItem">
              <img v-if="items.item_id == 10" src="~@/assets/items/10.png" />
              <img v-if="items.item_id == 91000001" src="~@/assets/ETHERicon.png" />
              <div class="lv text-stroke" v-if="items.item_id == 91000001">{{items.num / 10000}}</div>
              <div class="lv text-stroke" v-if="items.item_id == 10">{{items.num}}</div>
            </div>
          </div>
          <div class="rwd-notice text-stroke">Congrats! You won the No.{{thisLastRwd.rank}} on {{thisLastRwd.date}}</div>
        </div>
        <div class="foots">
          <div class="btn f30" @click="onPopSettleReward()"><span class="text-stroke">Confirm</span></div>
        </div>
      </div>

      <div class="popbody" v-if="isType === 3">
        <div class="head">
          <span class="text-stroke">Rules</span>
        </div>
        <div class="middle">
          <div class="pool-rule">
            <div class="title text-stroke">Prize Pool</div>
            <div>1.Release a certain amount of ZPA from offical</div>
            <div>2.More BOSS defeated, richer prize pool will be</div>
            <div>3.Prize pool resets after settlement</div>
          </div>
          <div class="challenge-rule">
            <div class="title text-stroke">Challenge rule</div>
            <div>1.Challenge several times for free every day</div>
            <div>2.Defeat an opponent to earn his ranking</div>
            <div>3.System select your opponent based on your ranking</div>
          </div>
          <div class="settle-rule">
            <div class="title text-stroke">Settle</div>
            <div>1.Settled once a day</div>
            <div>2.Rewards are awarded based on ranking</div>
          </div>
        </div>
        <div class="foots">
          <div class="btn f30 noani" @click="onPopClose()"><span class="text-stroke">Got it</span></div>
        </div>
      </div>

      <div class="popbody" v-if="isType === 4">
        <div class="head">
          <span class="text-stroke">PVP History</span>
        </div>
        <div class="middle">
          <div class="c-his">
            <div class="text-stroke" v-for="his in this.pvpHis" :key="his.date">
              {{his.date}}&nbsp;&nbsp;No.{{his.rank}}&nbsp;&nbsp;&nbsp;{{his.account.nickname}}
            </div>
          </div>
        </div>
        <div class="foots">
          <div class="btn f30 noani" @click="onPopClose()"><span class="text-stroke">Confirm</span></div>
        </div>
      </div>

    </div>

    <div class="popbox pvpbox subbox" v-if="isSubTip">
      <div class="bgs"></div>
      <!-- type 1-->
      <div class="popbody playerinfo" v-if="isSubType === 1">
        <div class="head">
          <span class="text-stroke">Player Info</span>
        </div>
        <div class="middle">
          <div class="pvp-base">
            <div class="avatar">
              <div class="pic"><img alt="avatar" src="~@/assets/character.png" /></div>
            </div>
            <div class="pvp-atr">
              <div class="pvp-name text-stroke">{{ thisRank.account_info.nickname }}<span v-if="thisRank.rank && thisRank.rank > 0">No.{{thisRank.rank}}</span></div>
              <div class="dataline text-stroke">HP:&nbsp;{{ thisRank.attr.hp }}<span>+{{ Math.floor(thisRank.attr.hp * thisRank.attr.hp_add_rate / 10000) }}</span></div>
              <div class="dataline text-stroke">ATK:&nbsp;{{ thisRank.attr.atk }}<span>+{{ Math.floor(thisRank.attr.atk * thisRank.attr.atk_add_rate / 10000) }}</span></div>
              <div class="dataline text-stroke">
                HIT:&nbsp;{{ thisRank.attr.hit_rate > 0 ? thisRank.attr.hit_rate / 100 : 0 }}%
              </div>
              <div class="dataline text-stroke">
                SPD:&nbsp;{{ thisRank.attr.dodge_rate > 0 ? thisRank.attr.dodge_rate / 100 : 0 }}%
              </div>
              <div class="dataline text-stroke">
                CRT:&nbsp;{{ thisRank.attr.cri_rate > 0 ? thisRank.attr.cri_rate / 100 : 0 }}%
              </div>
              <div class="dataline text-stroke">
                C.DMG:&nbsp;{{ thisRank.attr.cri_dmg_rate > 0 ? thisRank.attr.cri_dmg_rate / 100 : 0 }}%
              </div>
            </div>
          </div>

          <div class="equips">
            <div v-if="thisRank.base_equips && thisRank.base_equips.length > 0" class="base">
              <div class="baseequip" v-for="baseItem in thisRank.base_equips" :key="baseItem.equip_type">
                <img v-if="baseItem.equip_type == 1 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary1.png" />
                <img v-if="baseItem.equip_type == 1 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary11.png" />
                <img v-if="baseItem.equip_type == 1 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary12.png" />

                <img v-if="baseItem.equip_type == 2 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary2.png" />
                <img v-if="baseItem.equip_type == 2 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary21.png" />
                <img v-if="baseItem.equip_type == 2 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary22.png" />

                <img v-if="baseItem.equip_type == 3 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary3.png" />
                <img v-if="baseItem.equip_type == 3 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary31.png" />
                <img v-if="baseItem.equip_type == 3 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary32.png" />

                <img v-if="baseItem.equip_type == 4 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary4.png" />
                <img v-if="baseItem.equip_type == 4 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary41.png" />
                <img v-if="baseItem.equip_type == 4 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary42.png" />

                <img v-if="baseItem.equip_type == 5 && baseItem.level < 3" alt="ori" src="~@/assets/boss/ordinary5.png" />
                <img v-if="baseItem.equip_type == 5 && baseItem.level >= 3 && baseItem.level < 9" alt="ori" src="~@/assets/boss/ordinary51.png" />
                <img v-if="baseItem.equip_type == 5 && baseItem.level >= 9" alt="ori" src="~@/assets/boss/ordinary52.png" />

                <div class="star"><span class="text-stroke">{{baseItem.star}}</span></div>
                <div class="lv"><span class="text-stroke">Lv.{{baseItem.level}}</span></div>
              </div>
            </div>
            <div v-else class="base">
              <span class="noinfo text-stroke">No base equipment found.</span>
            </div>
            <div v-if="thisRank.nft_equips && thisRank.nft_equips.length > 0" class="nfts">
              <div class="nftequip" v-for="nft in thisRank.nft_equips" :key="nft.id">
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 1" src="~@/assets/nfts/Garm1.png" />
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 2" src="~@/assets/nfts/Edda1.png" />
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 3" src="~@/assets/nfts/Norns1.png" />
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 4" src="~@/assets/nfts/Parvati1.png" />
                <img alt="NFT" v-if="nft.level == 1 && nft.type == 5" src="~@/assets/nfts/Tartarus1.png" />

                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm2.png" />
                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda2.png" />
                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
                <img v-if="nft.level > 1 && nft.level <= 3 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />

                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm3.png" />
                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda3.png" />
                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns3.png" />
                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
                <img v-if="nft.level > 3 && nft.level <= 6 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />

                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm4.png" />
                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda4.png" />
                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns4.png" />
                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
                <img v-if="nft.level > 6 && nft.level <= 9 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />

                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm5.png" />
                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda5.png" />
                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns5.png" />
                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
                <img v-if="nft.level > 9 && nft.level <= 12 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />

                <img v-if="nft.level > 12 && nft.type == 1" alt="NFT" src="~@/assets/nfts/Garm6.png" />
                <img v-if="nft.level > 12 && nft.type == 2" alt="NFT" src="~@/assets/nfts/Edda6.png" />
                <img v-if="nft.level > 12 && nft.type == 3" alt="NFT" src="~@/assets/nfts/Norns6.png" />
                <img v-if="nft.level > 12 && nft.type == 4" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
                <img v-if="nft.level > 12 && nft.type == 5" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />

                <div class="nftlv"><span class="text-stroke">{{nft.level}}</span></div>
              </div>
            </div>
            <div v-else class="nfts">
              <span class="noinfo text-stroke">No NFT weapon found.</span>
            </div>
          </div>
        </div>
        <div class="foots">
          <div class="btn f30 text-stroke noani" @click="onSubPopClose()"><span class="text-stroke">Close</span></div>
        </div>
        <div class="close-btn" @click="onSubPopClose()"></div>
      </div>
    </div>

  </div>
</template>
<script>

import pvpconfig from '@/share/PVPConfig.json'
import pvprwdconfig from '@/share/PVPRankRewards.json'
import api from '@/utils/api'
import foots from '@/components/foots'
import headers from '@/components/header'
import cache from '@/utils/cache'
import jskit from '@/utils/index'
import { getRankList, getPVPInfo, getChallengeTarget, pvpHis } from '@/api/fight'

export default {
  data() {
    return {
      istip: false,
      isSubTip: false,
      isType: 0,
      isSubType: 0,
      ranks:{},
      rankList:{},
      findingRankList: false,
      targetList:{},
      targetSelected: 0,
      showFresh: false,
      findingTarget: false,
      thisRank:{},
      // my rank 部分结构需要与服务器返回一致，同时叠加自己的pvp info
      // nft装备结构是临时的
      myRank:{
        rank:0,
        account_info:{},
        attr:{},
        base_equips:{},
        nft_equips:{},
        challenger_cnt: 0, // 剩余挑战次数
        update_at: 0 ,// 次数更新时间
        rwd:{}
      },
      countdownTime: '',
      counting: false,
      countingHandle:null,
      pvpMaxFreeNum:0,
      pvpComsume:0,
      seasonId: 1,
      seasonRwdGap: 6,
      seasonRwdArr:{},
      tempRwd:{},

      curSeasonDate:'',
      curPrizepool:0,
      lastRwd:{},
      thisLastRwd:{
        rank: 0,
        date: '',
        rewardsItem:[]
      },

      pvpHis:{}
    }
  },
  components: {foots, headers},
  mounted() {

  },
  activated(){
    this.initPvpConfig()
    this.initSeasonRwd()
    this.getRanks()
    this.getMyPVPInfo()
    this.initMyRank()

    this.counting = false
    this.countingHandle = setInterval(() => {
      this.showTimeCountdown()
    },1000)
  },
  beforeRouteLeave(to, from, next) {
    clearInterval(this.countingHandle)
    next()
  },
  methods: {
    initPvpConfig(){
      // console.log(pvpconfig[1])
      this.pvpMaxFreeNum = pvpconfig[1].DailyPVPFreeChallengeCnt
      this.pvpComsume = pvpconfig[1].PVPManaConsume
    },
    initSeasonRwd(){
      let rwdArr = []
      for(let i = 1; i <= this.seasonRwdGap; i++){
        rwdArr.push(pvprwdconfig[this.seasonId + '_' + i])
      }
      this.seasonRwdArr = rwdArr
      console.log(this.seasonRwdArr)
    },
    calRankRwd(rank){
      this.tempRwd = {}
      if(this.seasonRwdArr){
        let rwdItems = []
        for(let rwd of this.seasonRwdArr){
          // console.log(rwd)
          let rang = rwd.RankRange.split(':')
          let item = rwd.Rewards.split(':')
          let extra = rwd.ExtraRewardPoolRate
          if(rang && rang.length == 2){
            let min = rang[0]
            let max = rang[1]
            if(rank >= min && rank <= max){
              if(extra > 0){
                let _item = {}
                _item['id'] = 91000001
                _item['num'] = parseInt(extra)
                rwdItems.push(_item)
              }
              if(item && item.length == 2 && item[1] > 0){
                let _item = {}
                _item['id'] = parseInt(item[0])
                _item['num'] = parseInt(item[1])
                rwdItems.push(_item)
              }
            }
          }
        }
        this.tempRwd = rwdItems
        // console.log('tmp rwd', this.tempRwd.length)
      }
    },
    onPopBox(type) {
      this.istip = true;
      this.isType = parseInt(type);

      console.log(this.isType, 'this.isType')
    },
    onPopClose(){
      this.istip = false
      this.isType = 0
    },
    onSubPopBox(type) {
      this.isSubTip = true
      this.isSubType = parseInt(type)
      console.log(this.isSubType, 'this.isSubType')
    },
    onSubPopClose(){
      this.isSubTip = false
      this.isSubType = 0
    },
    onTips(txt,time) {
      this.$refs.tip.openTip(txt, time)
    },
    showDetail(rank, from){
      this.thisRank = {}
      // 如果from = 0, 查看自己
      if(from === 0){
        this.thisRank = this.myRank
      }else if(from === 1){
        // 找到rank list的玩家
        this.thisRank = this.rankList.find(obj => obj.rank === rank)
      }else if(from === 2){
        // target list 的玩家
        this.thisRank = this.targetList.find(obj => obj.rank === rank)
      }
      console.log(this.thisRank)
      if(!this.thisRank.attr.hp_add_rate){
        this.thisRank.attr.hp_add_rate = 0
      }
      if(!this.thisRank.attr.atk_add_rate){
        this.thisRank.attr.atk_add_rate = 0
      }
      this.onSubPopBox(1)
    },
    getRanks(){
      this.rankList = {}
      this.findingRankList = true
      getRankList({ page_num : 0 }).then(response => {
        console.log("rank list",response)
        if(response){
          if(response.rank_list){
            console.log(response.rank_list)
            cache.setSession('tmpRankList', response.rank_list)
            this.rankList = response.rank_list
            if(this.rankList && this.rankList.length > 0){
              let _list = []
              for(let tmp of this.rankList){
                this.calRankRwd(tmp.rank)
                tmp.rwd = this.tempRwd
                _list.push(tmp)
              }
              this.rankList = _list
            }
          }
          this.findingRankList = false
        }
      })
    },
    getMyPVPInfo(){
      this.lastRwd = {}
      getPVPInfo({}).then(response => {
        //console.log("my pvp", JSON.parse(response))
        console.log("my pvp", response)
        if(response){
          this.myRank.rank = response.rank
          this.myRank.challenger_cnt = response.challenger_cnt || 0
          this.calRankRwd(this.myRank.rank)
          this.myRank.rwd = this.tempRwd
          let _pool =+ response.pvp_mana_awards_pool || 0
          if(_pool != 0){
            this.curPrizepool = (_pool / 10000).toFixed(4)
          }
          if(response.pvp_season_date){
            let date = response.pvp_season_date.toString()
            let year = date.substring(0,4)
            let month = date.substring(4,6)
            let day = date.substring(6,8)
            this.curSeasonDate = year + "." + month + "." + day
          }
          // if(response.pvp_settle_rewards){
          //   this.lastRwd = response.pvp_settle_rewards
          //   this.onPopSettleReward()
          // }
          // j版
          if(response.settle_rewards){
            this.lastRwd = response.settle_rewards
            this.onPopSettleReward()
          }
        }
      })
    },
    freshMaterialBySettle(rewards){
      let _materials = JSON.parse(cache.getSession('bagItems')) || {}
      if (rewards.length > 0) {
        for (let items of rewards) {
          if (Object.keys(_materials).length != 0) {
            if (_materials.find(obj => obj.item_id === items.item_id)) {
              _materials.find(obj => obj.item_id === items.item_id).num += items.num
            } else {
              if (items.item_id < 100) {
                // 目前材料编号只有 10 11 12 13 14 15
                _materials.push({ item_id: items.item_id, num: items.num })
              }
            }
          }
          if(items.item_id === 91000001){
            let zpa =+ cache.getSession('zpa') || 0
            zpa += items.num
            cache.setSession('zpa', zpa)
            this.$store.commit("GLOBAL_ZPA", zpa)
          }
        }
        cache.setSession('bagItems', _materials)
      }
    },
    onPopSettleReward(){
      this.onPopClose()
      this.onSubPopClose()
      // console.log("length",this.lastRwd.length, this.lastRwd)
      if(this.lastRwd && this.lastRwd.length > 0){
        this.thisLastRwd.rank = this.lastRwd[this.lastRwd.length - 1].rank
        let _rewards = this.lastRwd[this.lastRwd.length - 1].rewards
        if(_rewards && _rewards.length > 0){
          let _date = this.lastRwd[this.lastRwd.length - 1].date.toString()
          let _year = _date.substring(0,4)
          let _month = _date.substring(4,6)
          let _day = _date.substring(6,8)
          this.thisLastRwd.date = _year + "-" + _month + "-" + _day
          this.thisLastRwd.rewardsItem = _rewards
          // 前端累加材料道具（j版后台未提前结算）
          this.freshMaterialBySettle(_rewards)

          this.lastRwd.pop()
          // 奖励弹窗
          setTimeout(() => {
            this.onPopBox(2)
          }, 300)
        }
      }
    },
    initMyRank(){
      this.myRank.attr = JSON.parse(cache.getSession('battleInfo')) || {}
      this.myRank.account_info['nickname'] = cache.getSession('nickName') || 'Not found'
      this.myRank.base_equips = JSON.parse(cache.getSession('baseEquip')) || {}
      let nft1 = jskit.localStorageGet('equip1')
      let nft2 = jskit.localStorageGet('equip2')
      let nft3 = jskit.localStorageGet('equip3')
      let nftArr = []
      if(nft1){
        nftArr.push(nft1)
      }
      if(nft2){
        nftArr.push(nft2)
      }
      if(nft2){
        nftArr.push(nft3)
      }
      this.myRank.nft_equips = nftArr
      console.log(this.myRank)
    },
    getTarget(type){
      this.targetList = {}
      this.showFresh = false
      getChallengeTarget({}).then(response => {
        console.log(response)
        if(response){
          this.targetList = response.target_list
          let firstRank = 0
          if(this.targetList && this.targetList.length > 0){
            let _list = []
            for(let tmp of this.targetList){
              this.calRankRwd(tmp.rank)
              tmp.rwd = this.tempRwd
              _list.push(tmp)
            }
            _list.sort((a, b) => a.rank - b.rank)
            this.targetList = _list
            this.selectTarget(this.targetList[0].rank)
          }
          this.showFresh = true
          this.findingTarget = false
          if(type === 1){
            this.onPopBox(1)
          }
        }
      })
    },
    showTarget(){
      this.getTarget(1)
    },
    selectTarget(rank){
      this.targetSelected = rank
    },
    freshTarget(){
      this.findingTarget = true
      this.getTarget(2)
    },
    gotoPvp(){
      let zpa =+ cache.getSession('zpa') || 0
      if(this.myRank.challenger_cnt <=0 && zpa < this.pvpComsume){
        this.$refs.tip.openTip('ZPA token is not enough!')
        return
      }
      let rank = this.targetSelected
      if(rank && rank > 0){
        let obj = this.targetList.find(obj => obj.rank === rank)
        if(obj && obj.account_info){
          let uid = this.targetList.find(obj => obj.rank === rank).account_info.user_id
          let tname = this.targetList.find(obj => obj.rank === rank).account_info.nickname
          if(uid && uid > 0){
            this.onPopClose()
            this.$router.push({ path: '/pvpfight', query: { 'rank': rank, 'uid' : uid, 'tname' : tname } })
          }else{
            this.$refs.tip.openTip('Param Error')
          }
        }else{
          this.$refs.tip.openTip('Cannot find this player')
        }
      }else{
        this.$refs.tip.openTip('Please choose a player')
      }
    },
    calTimeCount(futureTime, curTime){
      let remain = futureTime - curTime;//毫秒
      let s = parseInt(remain / 1000 % 60)//秒 余下多少秒（0-59）
      let m = parseInt(remain / 1000 / 60 % 60); //余下了多少分钟（0-59）
      let h = parseInt(remain / 1000 / 60 / 60 % 24);//余下多少时（0-23）
      s = s < 10 ? '0' + s : s;
      h = h < 10 ? '0' + h : h;
      m = m < 10 ? '0' + m : m;
      let t = parseInt(remain / 1000 / 60 / 60 / 24);
      // console.log(t," day ", h, ' hour ', m,' min ', s, ' sec')
      this.countdownTime = "End in " + h + ":" + m + ":" + s;
    },
    showTimeCountdown(){
      this.countdownTime = ''
      let curDate = new Date()
      let futureDate = new Date()
      let freePvpNumFreshDate = new Date()
      futureDate.setUTCHours(23, 0, 0, 0)
      freePvpNumFreshDate.setUTCHours(23, 59, 55, 0)
      // console.log( "curDate: ", curDate.getTime() );
      // console.log("futureDate: ", futureDate.getTime())
      let ctime = curDate.getTime()
      let ftime = futureDate.getTime()
      let fpftime = freePvpNumFreshDate.getTime()
      if(ctime > ftime){
        if(ctime > fpftime){
          this.myRank.challenger_cnt = this.pvpMaxFreeNum
        }
        this.counting = true
      }else{
        this.counting = false
        this.calTimeCount(ftime, ctime)
      }
    },
    getPVPHis(){
      this.pvpHis = {}
      pvpHis({}).then(response => {
        if(response && response.list){
          let arr = []
          for(let rank of response.list){
            let _date = rank.date.toString()
            let _year = _date.substring(0,4)
            let _month = _date.substring(4,6)
            let _day = _date.substring(6,8)
            rank.date = _year + "." + _month + "." + _day
            arr.push(rank)
          }
          this.pvpHis = arr
        }
        console.log(this.pvpHis)
        this.onPopBox(4)
      })
    }
  }
}
</script>

<style lang='stylus' src="./index.css" />
