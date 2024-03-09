<template>
  <!-- <div :class="is_cri ? 'bossdetail shake' : 'bossdetail'"> -->
  <div class="bossdetail">
    <div class="headbox">
      <div class="bag">
        <img src="../../assets/bag/bag.png" />
      </div>
    </div>
    <!-- weaponlist -->
    <div class="boss">
      <div class="bossCav" ref="bossCav"></div>
      <!-- <transition name="fade-dmg">
        <div class="text-stroke bossDmg" :class="is_cri ? 'criDmg':''" v-if="showBossDmg" v-html="bossDmg"></div>
      </transition> -->
      <div class="detail" ref="bossObj">
        <div class="detpic">
          <img alt="NFT" :class="hit ? 'gameover' : ''" :src="this.bossData.bossPic" />
        </div>
        <!-- <transition name="fade-coin-0">
          <div class="bossRewardsOut" v-if="showBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-coin-1">
          <div class="bossRewardsOut" v-if="showBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-coin-2">
          <div class="bossRewardsOut" v-if="showBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-coin-3">
          <div class="bossRewardsOut" v-if="showBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-coin-4">
          <div class="bossRewardsOut" v-if="showBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-coin-5">
          <div class="bossRewardsOut" v-if="showBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-recover-coin-0">
          <div class="bossRewardsRecover-0" v-if="recoverBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-recover-coin-1">
          <div class="bossRewardsRecover-1" v-if="recoverBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-recover-coin-2">
          <div class="bossRewardsRecover-2" v-if="recoverBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-recover-coin-3">
          <div class="bossRewardsRecover-3" v-if="recoverBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-recover-coin-4">
          <div class="bossRewardsRecover-4" v-if="recoverBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition>
        <transition name="fade-recover-coin-5">
          <div class="bossRewardsRecover-5" v-if="recoverBossCoin">
            <img class="bossRewardsFlipy" src="@/assets/items/10.png" />
          </div>
        </transition> -->
        <div class="mygame">
          <div class="progress">
            <div class="val text-stroke">{{ progressBoss }}%</div>
            <div class="fail" :style="{ width: progressBoss + '%' }"></div>
          </div>
        </div>
      </div>
    </div>
    <div class="player" ref="playerObj">
      <div class="myname">
        <div class="user">
          <div class="pic"><img alt="NFT" src="~@/assets/sculpture.png" /></div>
          <div class="text text-stroke">
            <div v-html="nickname" class="name"></div>
          </div>
        </div>
      </div>
      <div class="tagame">
        <div class="progress">
          <div class="val text-stroke">{{ progressIsta }}%</div>
          <div class="victory" :style="{ width: progressIsta + '%' }"></div>
        </div>
      </div>

      <div class="pkshow" ref="pkbosstxtbody">
        <div class="itemdata onsd text-stroke" v-for="(item, index) in temppk" :key="index">
          <span class="my" v-if="item.showTitle && item.status === 0" v-html="nickname"></span>
          <span class="ta" v-if="item.showTitle && item.status === 1">{{bossName}}</span>
          {{ item.text }}
        </div>
      </div>

      <!-- pop start -->
      <div class="popbox fightbosspop" v-if="istip">
        <div class="bgs"></div>
        <!-- type 1-->
        <div class="detailbody" v-if="isType === 1">
          <div class="head">
            <span class="text-stroke">Obtained Rewards</span>
          </div>
          <div class="middle">
            <div class="fallen-rwd-notice text-stroke" v-if="this.fallenButRwd">
            Since you caused lots of damage to the BOSS, you obtained some rewards.
            </div>
            <div class="type">
              <div class="pic" v-for="items in awardsItems">
                <img alt="NFT" :src="items.itemPic" />
                <div class="lv text-stroke" v-if="items.item_id == 91000001">{{items.num / 10000}}</div>
                <div class="lv text-stroke" v-else>{{items.num}}</div>
              </div>
            </div>
          </div>
          <div class="foots">
            <div v-if="this.hasHiddenBoss" class="btnss f30" @click="onPopBox(3)"><span class="text-stroke">Confirm</span>
            </div>
            <div v-else class="btnss f30" @click="onPopClose(1)"><span class="text-stroke">Confirm</span></div>
          </div>
        </div>

        <!-- type 2-->
        <div class="detailbody" v-if="isType === 2">
          <div class="head">
            <!-- Obtaining Rewards
          <span class="icon"></span> -->
          </div>
          <div class="middle">
            <div class="type">
              <span class="text-stroke">Challenge failed</span>
            </div>
          </div>
          <div class="foots">
            <div class="btnss f30" @click="onPopClose(1)"><span class="text-stroke">Confirm</span></div>
          </div>
          <!-- <div class="close-btn" @click="istip=false"></div> -->
        </div>

        <!-- type 3-->
        <div class="hiddenbossbody" v-if="isType === 3">
          <div class="head"></div>
          <div class="middle">
            <div class="hideboss">
              <div class="bosinfo">
                <!-- <img src="~@/assets/boss/1003.png" /> -->
                <img alt="NFT" :src="this.hiddenBossPic" />
              </div>
              <div class="hiddenbosstitle text-stroke">Hidden BOSS Appeared !!!</div>
              <!-- <div class="hiddenbosstxt" v-html="this.hiddenBossLink"></div> -->
              <div class="bosstxt text-stroke">
                Please click the button to copy a code, and send it to the <span>FishillteGroup</span> chat or
                FishillteGame robot to join forces with the others in defeating the BOSS.
              </div>
            </div>
          </div>
          <div class="foots">
            <div class="btnss f30" @click="copy()"><span class="text-stroke">Copy</span></div>
          </div>
          <div class="close-btn" @click="onPopClose(1)"></div>
        </div>
      </div>
    </div>
    <!-- pop end -->

    <div class="tips-div">
      <Tip ref="tip"></Tip>
    </div>
  </div>
</template>
<script>

// import { mapActions } from 'vuex'
import boss from '@/share/boss.json'
import faildata from './faildata.json'
import windata from './windata.json'
import cache from '@/utils/cache'
import { fight, getShareHBossLink } from '@/api/fight'
import { FallItem } from '@/animation/fall'
import { Damage } from '@/animation/damageFade'

export default {
  data() {
    return {
      dpr: 1,
      istip: false,
      isType: 0,
      bossItem: 1,
      //temp
      temppk: [],
      len: 1,
      progressBoss: 100,
      progressIsta: 100,
      hit: false,
      is_cri: false,
      nickname: '',
      energy: 0,
      awardsItems: [],
      hasHiddenBoss: false,
      hiddenBossId: 0,
      hiddenBossUid: 0,
      hiddenBossPic: '',
      hiddenBossLink: '',
      bossData: {},
      bossId: 1,
      bossName:'BOSS',
      showBossDmg: false,
      bossDmg: '',
      showBossCoin: false,
      recoverBossCoin: false,
      showPlayerDmg: false,
      playerDmg: '',
      isWin: 0,
      fallenButRwd: false,

      fallItems:[]
    }
  },
  computed: {
    textLines() {
      return this.text.split('\n');
    },
  },
  components: {},
  created() {
    this.getUserInfo()
  },
  updated(){
    this.scrollToBottom()
  },
  activated() {
    this.dpr = window.devicePixelRatio || 1
    this.$nextTick(() => {
      //初始化
      this.init()
      //init boss信息
      console.log("start pk....")
      const querys = this.$route.query || {}
      if(querys && querys.id){
        this.bossId =+ querys.id || 1
        let allBossData = boss
        this.bossData = allBossData[this.bossId] || {}
        this.bossData.bossPic = require('@/assets/boss/' + this.bossData.Image + '.png')
        this.bossItem =+ this.bossData.Class + 1
        this.bossName = querys.name
        console.log("this fight boss ",this.bossData)
      }else{
        this.$refs.tip.openTip('Param Err')
      }

      if (process.env.NODE_ENV === 'development') {
        //test
        //this.attacker()
        this.fight()
      } else {
        //production
        this.fight()
        //this.attacker()
      }
    })

  },
  methods: {
    init() {
      this.progressBoss = 100
      this.progressIsta = 100
      this.temppk = []
      this.len = 1
      this.hit = false
      this.istip = false
      this.isType = 0
      this.hasHiddenBoss = false
      this.awardsItems = []
      this.bossDmg = ''
      this.playerDmg = ''
      this.isWin = 0
      this.fallenButRwd = false
      this.fallItems = []

      this.clearBossCav()
    },
    scrollToBottom(){
      this.$nextTick(()=>{
        let body = this.$refs.pkbosstxtbody
        body.scrollTo({ top: body.scrollHeight, behavior: 'smooth' })
      })
    },
    fight() {
      let bossId = +this.bossId // 每个系列第一个的id 是10001 11001 12001 13001 14001 15001
      console.log('-----------bossId', bossId)
      fight({ boss_id: bossId }).then(response => {
        let itemData = response
        console.log('fight boss response', response)
        if(!itemData || !itemData.details || itemData.details.length <= 0){
          this.$refs.tip.openTip('Fight Boss error')
          return
        }
        // 刷新buff缓存
        if(itemData.buff_cards){
          cache.setSession('buffCard', itemData.buff_cards)
        }else{
          cache.removeSession('buffCard')
        }
        if(itemData.buff_card_status){
          cache.setSession('buffCardStatus', itemData.buff_card_status)
        }else{
          cache.removeSession('buffCardStatus')
        }
        //是否触发隐藏boss
        let hiddenBoss = itemData.hidden_boss || {}
        if (hiddenBoss && hiddenBoss.boss_id) {
          this.findHiddenBoss(hiddenBoss)
        }
        //是否扣体力
        if (itemData.win) {
          this.useEnergy(itemData.energy, itemData.energy_recover_at)
          if(itemData.details && itemData.details.length > 0){
            let _atkBoss = itemData.details[itemData.details.length - 1].be_attacker ? true : false
            if(_atkBoss){
              if(!itemData.details[itemData.details.length - 1].hp_after || itemData.details[itemData.details.length - 1].hp_after === 0){
                // 刷新session 进度
                this.freshSeesionNode()
                this.isWin = 1
              }
            }
          }
        }
        // console.log(itemData)
        //设置获取奖励
        if (itemData.win) {
          let { awards } = itemData || {}
          console.log(awards)
          this.awardsItems = awards.add_items || []
          let _materials = JSON.parse(cache.getSession('bagItems')) || {}
          if (this.awardsItems.length > 0) {
            for (let items of this.awardsItems) {
              items.itemPic = require('@/assets/items/' + items.item_id + '.png')
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
        }

        let drawArr = this.calDrawMaterial(itemData.details)

        console.log('---玩家赢：win:true, 玩家输：win:false', '此局玩家赢：', itemData.win)
        //判断谁先开 0 玩家  1 boss
        let isboss = false
        if(itemData.details[0].attacker_type && itemData.details[0].attacker_type === 1){
          isboss = true
        }else{
          if(itemData.details[0].be_attacker && itemData.details[0].be_attacker === 0){
            isboss = true
          }
        }
        // console.log(itemData.details.length, '---', itemData.details[0].attacker_type, isboss)
        this.temppk.push({
          status: isboss ? 1 : 0,
          text: ' starts the attack first...',
          showTitle: true
        })

        let bossInitHp = itemData.boss_hp || 0
        let playerInitHp = itemData.player_hp || 0
        let drawItemsNum = 0
        let time = setInterval(() => {
          this.is_cri = false
          this.len++
          let item = itemData.details[this.len - 2]
          let atkBossNow = item.be_attacker ? true : false
          let atkPlayerNow = item.attacker_type ? true : false
          //console.log('====', atkBossNow, atkPlayerNow)
          console.log(atkPlayerNow ? "BOSS->玩家:" : "玩家->BOSS:", '被攻击前：', item.hp_before, '被攻击后：', item.hp_after, ' 造成伤害', item.dmg ,item.is_miss ? ' Miss' : '', item.is_cri ? ' 暴击':'')

          //设置双方血条
          if (atkBossNow) {
            this.bossDmg = item.is_miss ? 0 : item.dmg * -1
            if(item.is_cri){
              this.is_cri = true;
              let times = Math.floor(Math.random() * 2) + 1
              this.drawHp(180,80,this.bossDmg,2,times)
            }else{
              this.drawHp(180,80,this.bossDmg,1,1)
            }
            if (!item.is_miss) {
              this.drawMaterial(drawArr[drawItemsNum].id, drawArr[drawItemsNum].material)
              if(drawArr[drawItemsNum].stealAdd > 0){
                let stealNum = drawArr[drawItemsNum].stealAdd
                setTimeout(() => {
                  this.drawMaterial(drawArr[drawItemsNum].id, stealNum)
                }, 20)
              }
              if(drawArr[drawItemsNum].zpa > 0){
                let zpaNum = drawArr[drawItemsNum].zpa
                setTimeout(() => {
                  this.drawMaterial('91000001', zpaNum)
                }, 30)
              }
              drawItemsNum ++;
            }
            let bossHP = item.hp_after || 0
            this.progressBoss = bossHP == bossInitHp ? 100 : Math.ceil(bossHP / bossInitHp * 100)
          }
          if (atkPlayerNow) {
            let playerHP = item.hp_after || 0
            this.progressIsta = playerHP == playerInitHp ? 100 : Math.ceil(playerHP / playerInitHp * 100)
          }
          //boss击中 闪白
          if (atkBossNow) {
            item.is_miss ? this.hit = false : this.hit = true;
          } else {
            this.hit = false;
          }
          //boss 1, 玩家0

          let text = '';
          if (item.is_miss) {
            text = 'attacked, but it was missed.'
          } else if (item.is_cri) {
            if(atkPlayerNow){
              text = 'attacked, and caused '+item.dmg+' critical hit damage.'
              //打击震动效果text = 'delivered a full-force strike, and caused '+item.dmg+' critical hit damage.'
              this.attackPlayer(5)
            }
            if(atkBossNow){
              text = 'delivered a full-force strike, and caused '+item.dmg+' critical hit damage.'
              //打击震动效果
              this.attackBoss(5)
            }
          } else {
            text = 'caused ' + item.dmg + ' damage.'
            //打击震动效果
            atkPlayerNow ? this.attackPlayer(1) : this.attackBoss(1)
          }

          this.temppk.push({
            status: atkPlayerNow ? 1 : 0,
            text,
            showTitle : true
          })

          // 本次伤害后有一方血量归0
          if (!item.hp_after || item.hp_after === undefined || item.hp_after === 0) {
            text = atkPlayerNow ? 'has fallen.' : 'was defeated.'
            this.temppk.push({
              status: atkPlayerNow ? 0 : 1, // 被击败对象需要反一下
              text : text,
              showTitle : true
            })
          }

          //退出战斗
          if (this.len > itemData.details.length) {
            if(itemData.win){
              if(atkPlayerNow){
                if(!item.hp_after || item.hp_after === 0){
                  text = 'Fight failed, since you caused a lot of damage to the BOSS, you will still obtain some rewards.'
                  this.fallenButRwd = true
                }
              }else{
                if(!item.hp_after || item.hp_after === 0){
                  text = 'Battle victory, you will obtain all rewards'
                }
              }
            }else{
              text = 'Fight failed...'
            }
            this.temppk.push({
              status: atkPlayerNow ? 1 : 0,
              text: text,
              showTitle : false
            })
            this.hit = false
            //延时弹窗
            setTimeout(() => {
              if(itemData.win){
                this.fadeMaterial()
              }
              setTimeout(() => {
                this.clearBossCav()
                itemData.win ? this.onPopBox(1) : this.onPopBox(2)
              }, 1200)//等回收金币动画
            }, 1500)//600+800 600最后一滴血 800掉金币
            //结束定时器
            clearInterval(time)
          }
        }, 1000)
      })
    },
    attacker() {
      // let itemData = faildata
      console.log(this.thisBoss,'=====this fight boss')
      let itemData = windata
      if(!itemData || !itemData.details || itemData.details.length <= 0){
        this.$refs.tip.openTip('Fight Boss error')
        return
      }
      //是否触发隐藏boss，演示用
      // let hiddenBoss = {
      //   "boss_id": 20001,
      //   "boss_uuid": 160557906445468673,
      //   "hp": 10000,
      //   "finder": 10000008,
      //   "expire_at": 1703952053
      // }
      // if (hiddenBoss && hiddenBoss.boss_id) {
      //   this.findHiddenBoss(hiddenBoss)
      // }
      //是否扣体力
      if (itemData.win) {
        this.useEnergy(itemData.energy, itemData.energy_recover_at)
        if(itemData.details && itemData.details.length > 0){
          let _atkBoss = itemData.details[itemData.details.length - 1].be_attacker ? true : false
          if(_atkBoss){
            if(!itemData.details[itemData.details.length - 1].hp_after || itemData.details[itemData.details.length - 1].hp_after === 0){
              // 刷新session 进度
              this.freshSeesionNode()
              this.isWin = 1
            }
          }
        }
      }
      // console.log(itemData)
      //设置获取奖励
      if (itemData.win) {
        let { awards } = itemData || {}
        console.log(awards)
        this.awardsItems = awards.add_items || []
        let _materials = JSON.parse(cache.getSession('bagItems')) || {}
        if (this.awardsItems.length > 0) {
          for (let items of this.awardsItems) {
            items.itemPic = require('@/assets/items/' + items.item_id + '.png')
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
      }

      let drawArr = this.calDrawMaterial(itemData.details)

      console.log('---玩家赢：win:true, 玩家输：win:false', '此局玩家赢：', itemData.win)
      //判断谁先开 0 玩家  1 boss
      let isboss = false
      if(itemData.details[0].attacker_type && itemData.details[0].attacker_type === 1){
        isboss = true
      }else{
        if(itemData.details[0].be_attacker && itemData.details[0].be_attacker === 0){
          isboss = true
        }
      }
      // console.log(itemData.details.length, '---', itemData.details[0].attacker_type, isboss)
      this.temppk.push({
        status: isboss ? 1 : 0,
        text: ' starts the attack first...',
        showTitle: true
      })

      let bossInitHp = itemData.boss_hp || 0
      let playerInitHp = itemData.player_hp || 0
      let drawItemsNum = 0
      let time = setInterval(() => {
        this.is_cri = false
        this.len++
        let item = itemData.details[this.len - 2]
        let atkBossNow = item.be_attacker ? true : false
        let atkPlayerNow = item.attacker_type ? true : false
        //console.log('====', atkBossNow, atkPlayerNow)
        console.log(atkPlayerNow ? "BOSS->玩家:" : "玩家->BOSS:", '被攻击前：', item.hp_before, '被攻击后：', item.hp_after, ' 造成伤害', item.dmg ,item.is_miss ? ' Miss' : '', item.is_cri ? ' 暴击':'')

        //设置双方血条
        if (atkBossNow) {
          this.bossDmg = item.is_miss ? 0 : item.dmg * -1
          if(item.is_cri){
            this.is_cri = true;
            let times = Math.floor(Math.random() * 2) + 1
            this.drawHp(180,80,this.bossDmg,2,times)
          }else{
            this.drawHp(180,80,this.bossDmg,1,1)
          }
          //this.showBossDmgValue()
          if (!item.is_miss) {
            this.drawMaterial(drawArr[drawItemsNum].id, drawArr[drawItemsNum].material)
            if(drawArr[drawItemsNum].stealAdd > 0){
              let stealNum = drawArr[drawItemsNum].stealAdd
              setTimeout(() => {
                this.drawMaterial(drawArr[drawItemsNum].id, stealNum)
              }, 20)
            }
            if(drawArr[drawItemsNum].zpa > 0){
              let zpaNum = drawArr[drawItemsNum].zpa
              setTimeout(() => {
                this.drawMaterial('91000001', zpaNum)
              }, 30)
            }
            drawItemsNum ++;
          }
          let bossHP = item.hp_after || 0
          this.progressBoss = bossHP == bossInitHp ? 100 : Math.ceil(bossHP / bossInitHp * 100)
        }
        if (atkPlayerNow) {
          let playerHP = item.hp_after || 0
          this.progressIsta = playerHP == playerInitHp ? 100 : Math.ceil(playerHP / playerInitHp * 100)
        }
        //boss击中 闪白
        if (atkBossNow) {
          item.is_miss ? this.hit = false : this.hit = true;
        } else {
          this.hit = false;
        }
        //boss 1, 玩家0

        let text = '';
        if (item.is_miss) {
          text = 'attacked, but it was missed.'
        } else if (item.is_cri) {
          if(atkPlayerNow){
            text = 'attacked, and caused '+item.dmg+' critical hit damage.'
            //打击震动效果text = 'delivered a full-force strike, and caused '+item.dmg+' critical hit damage.'
            //this.attackPlayer(5)
          }
          if(atkBossNow){
            text = 'delivered a full-force strike, and caused '+item.dmg+' critical hit damage.'
            //打击震动效果
            //this.attackBoss(5)
          }
        } else {
          text = 'caused ' + item.dmg + ' damage.'
          //打击震动效果
          //atkPlayerNow ? this.attackPlayer(1) : this.attackBoss(1)
        }

        this.temppk.push({
          status: atkPlayerNow ? 1 : 0,
          text,
          showTitle : true
        })

        // 本次伤害后有一方血量归0
        if (!item.hp_after || item.hp_after === undefined || item.hp_after === 0) {
          text = atkPlayerNow ? 'has fallen.' : 'was defeated.'
          this.temppk.push({
            status: atkPlayerNow ? 0 : 1, // 被击败对象需要反一下
            text : text,
            showTitle : true
          })
        }

        //退出战斗
        if (this.len > itemData.details.length) {
          if(itemData.win){
            if(atkPlayerNow){
              if(!item.hp_after || item.hp_after === 0){
                text = 'Fight failed, since you caused a lot of damage to the BOSS, you will still obtain some rewards.'
                this.fallenButRwd = true
              }
            }else{
              if(!item.hp_after || item.hp_after === 0){
                text = 'Battle victory, you will obtain all rewards'
              }
            }
          }else{
            text = 'Fight failed...'
          }
          this.temppk.push({
            status: atkPlayerNow ? 1 : 0,
            text: text,
            showTitle : false
          })
          this.hit = false
          //延时弹窗
          setTimeout(() => {
            if(itemData.win){
              //this.showRecoverBossCoin()
              this.fadeMaterial()
            }
            setTimeout(() => {
              //this.clearBossCav()
              //itemData.win ? this.onPopBox(1) : this.onPopBox(2)
            }, 1200)//等回收金币动画
          }, 1500)//600+800 600最后一滴血 800掉金币
          //结束定时器
          clearInterval(time)
        }
      }, 1000)
    },
    calDrawMaterial(details){
      let drawArr = []
      // 计算要动画的材料数量， 用BOSS 等级代替显示材料的系数
      let maxMaterialToDraw = this.bossData.Level > 10 ?10 : this.bossData.Level
      // let maxMaterialToDraw = 200
      let mId = 0
      for (let itemId in this.bossData.Awards) {
        // 此处目前固定小于100的ID认为是升级升星材料，且一个BOSS只有一种，后续要增加材料种类后，需优化
        if(itemId < 100){
          mId = itemId
        }
      }
      let maxZpaToDraw = 0
      if(this.bossData.ManaAwardsFinal){
        maxZpaToDraw = parseInt(this.bossData.ManaAwardsFinal) > 5 ? 5 : parseInt(this.bossData.ManaAwardsFinal)
      }
      // 计算击中BOSS次数
      let attackNums = 0
      if(details && details.length > 0){
        for(let atkdetail  of details){
          if(atkdetail.be_attacker === 1 || atkdetail.attacker_type === 0){
            if(!atkdetail.is_miss){
              attackNums ++
            }
          }
        }
        if(attackNums == 1){
          // 控制单次数量不超过3
          maxMaterialToDraw = maxMaterialToDraw > 3 ? 3 : maxMaterialToDraw
        }
        console.log('maxToDraw zpa',maxZpaToDraw,' material ', maxMaterialToDraw)
        // console.log('atk nums', attackNums)
        for(let atkdetail  of details){
          if(!atkdetail.is_miss && attackNums > 0){
            let _draw = {}
            _draw.id = mId
            _draw.material = 0
            _draw.stealAdd = 0
            _draw.zpa = 0
            let criAdd = 1
            let avg = Math.ceil(maxMaterialToDraw / attackNums)
            if(avg < 1){
              avg = 1
            }
            if(atkdetail.is_cri){
              criAdd = 1.5
            }
            if(attackNums == 1){
              // console.log('final', avg)
               _draw.material = avg
               if(maxZpaToDraw > 0){
                 _draw.zpa = maxZpaToDraw
               }
            }else{
              // console.log('avg', avg)
              _draw.material = Math.round(Math.random() * avg * criAdd) + Math.round(avg / 3)
              if(_draw.material > 3){
                _draw.material = 3
              }
              if(maxZpaToDraw > 0){
                _draw.zpa = Math.round(Math.random() * maxZpaToDraw / 2)
                maxZpaToDraw = maxZpaToDraw - _draw.zpa
              }
              // console.log('time', _draw.material)
            }
            if(atkdetail.trigger_dodge_steal){
              _draw.stealAdd = Math.round(_draw.material * 0.3) < 1 ? 1 : Math.round(_draw.material * 0.3)
            }
            maxMaterialToDraw = maxMaterialToDraw - _draw.material - _draw.stealAdd
            if(maxMaterialToDraw < 0){
              maxMaterialToDraw = 0
            }
            // console.log('remaining', maxMaterialToDraw)
            attackNums--
            drawArr.push(_draw)
          }
        }
        console.log('drawArr==',drawArr)
      }
      return drawArr
    },
    attackBoss(times = 1) {
      console.log("do boss shake")
      let timer = null, num = 0
      let divBoss = this.$refs.bossObj
      clearInterval(timer)  //实现画面震动效果
      timer = setInterval(function () {
        num++             //num控制画面震动时间，点击一下后，每隔30ms震动一次，震动10次
        if (num == 10) {
          clearInterval(timer)
          num = 0
          divBoss.style.left = 0
          divBoss.style.top = 0
          return
        }
        divBoss.style.left = times * (Math.random() * -5 + 5) + 'px'
        divBoss.style.top = times * (Math.random() * -5 + 5) + 'px'
      }, 30)
    },
    attackPlayer(times = 1) {
      console.log("do shake")
      let timer = null, num = 0
      let divPlayer = this.$refs.playerObj
      clearInterval(timer)  //实现画面震动效果
      timer = setInterval(function () {
        num++             //num控制画面震动时间，点击一下后，每隔30ms震动一次，震动10次
        if (num == 10) {
          clearInterval(timer)
          num = 0
          divPlayer.style.left = 0
          divPlayer.style.top = 0
          return
        }
        divPlayer.style.left = times * (Math.random() * -5 + 5) + 'px'
        divPlayer.style.top = times * (Math.random() * -5 + 5) + 'px'
      }, 30)
    },
    //打开弹窗
    onPopBox(type) {
      this.isType = parseInt(type)
      this.istip = true
    },
    onPopClose(closeType) {
      this.istip = false
      if (closeType == 3) {
        this.onPopBox(3)
      } else {
        // location.href = '/boss'
        this.$router.push({ path: '/boss', query: { 'id': this.bossId || '', 'win': this.isWin } })
      }
    },
    getUserInfo() {
      this.nickname = cache.getSession('nickName')
      this.energy = cache.getSession('energy')
    },
    useEnergy(useEnergy, recoverAt) {
      this.energy = useEnergy || 0
      cache.setSession('energy', this.energy)
      cache.setSession('energy_update_at', recoverAt || 0)
    },
    findHiddenBoss(hiddenBoss) {
      this.hasHiddenBoss = true
      this.hiddenBossId = hiddenBoss.boss_id//1003
      this.hiddenBossUid = hiddenBoss.boss_uuid
      let picId = 0
      // 20001~20045：1002
      // 21001~21045：1102
      // 22001~22045：1202
      // 23001~23045：1302
      // 24001~24045：1402
      // 25001~25045：1502
      // 30001~30045：1003
      // 31001~31045：1103
      // 32001~32045：1203
      // 33001~33045：1303
      // 34001~34045：1403
      // 35001~35045：1503
      if (this.hiddenBossId >= 20001 && this.hiddenBossId <= 20045) {
        picId = 1002
      }
      if (this.hiddenBossId >= 21001 && this.hiddenBossId <= 21045) {
        picId = 1102
      }
      if (this.hiddenBossId >= 22001 && this.hiddenBossId <= 22045) {
        picId = 1202
      }
      if (this.hiddenBossId >= 23001 && this.hiddenBossId <= 23045) {
        picId = 1302
      }
      if (this.hiddenBossId >= 24001 && this.hiddenBossId <= 24045) {
        picId = 1402
      }
      if (this.hiddenBossId >= 25001 && this.hiddenBossId <= 25045) {
        picId = 1502
      }
      if (this.hiddenBossId >= 30001 && this.hiddenBossId <= 30045) {
        picId = 1003
      }
      if (this.hiddenBossId >= 31001 && this.hiddenBossId <= 31045) {
        picId = 1103
      }
      if (this.hiddenBossId >= 32001 && this.hiddenBossId <= 32045) {
        picId = 1203
      }
      if (this.hiddenBossId >= 33001 && this.hiddenBossId <= 33045) {
        picId = 1303
      }
      if (this.hiddenBossId >= 34001 && this.hiddenBossId <= 34045) {
        picId = 1403
      }
      if (this.hiddenBossId >= 35001 && this.hiddenBossId <= 35045) {
        picId = 1503
      }
      if (picId) {
        this.hiddenBossPic = require('@/assets/boss/' + picId + '.png') || ''
      }
      getShareHBossLink({ 'boss_uuid': this.hiddenBossUid }).then(response => {
        this.hiddenBossLink = response.share_link || ''
      })
    },
    copy() {
      console.log(this.hiddenBossLink)
      navigator.clipboard.writeText(this.hiddenBossLink)
      this.$refs.tip.openTip('Copy success')
    },
    freshSeesionNode() {
      console.log("挑战的BOSS系列", this.bossData.Class)
      console.log("挑战的BOSS", this.bossId)
      let nodes = JSON.parse(cache.getSession('bossNode')) || {}
      if(nodes){
        let id = nodes.find(obj => obj.class == this.bossData.Class).bossId
        if(parseInt(id) < parseInt(this.bossId)){
          nodes.find(obj => obj.class == this.bossData.Class).bossId = this.bossId
          cache.setSession('bossNode', nodes);
        }
      }
    },
    showBossDmgValue() {
      this.showBossDmg = true
      setTimeout(() => {
        this.showBossDmg = false
      }, 100)
    },
    showPlayerDmgValue() {
      this.showPlayerDmg = true
      setTimeout(() => {
        this.showPlayerDmg = false
      }, 100)
    },
    showBossCoinOut() {
      this.showBossCoin = true
      setTimeout(() => {
        this.showBossCoin = false
      }, 100)
    },
    showRecoverBossCoin() {
      this.recoverBossCoin = true
      setTimeout(() => {
        this.recoverBossCoin = false
      }, 100)
    },

    // canvas动画
    clearBossCav(){
      let childNode = this.$refs.bossCav.childNodes
      for(let i = childNode.length - 1; i >= 0; i--){
        this.$refs.bossCav.removeChild(childNode[i])
      }
    },
    drawHp(x,y,dmg,type,showTimes){
      let _dmg = parseInt(dmg)
      let dmgArr = []
      if(showTimes > 1){
        let _total = 0
        for(let i = 0; i < showTimes; i++){
          let d = 0
          if(i == showTimes - 1){
            d = _dmg - _total
          }else{
            d = Math.floor(_dmg / showTimes)
          }
          dmgArr.push(d)
          _total += d
        }
      }else{
        dmgArr.push(_dmg)
      }

      for (let i = 0; i < showTimes; i++) {
        let x1 = 0
        let y1 = 0
        if(type == 2){
          // 暴击触发随机位置,取10的整数
          // x  110~260
          // y  x在140~230之间时，y 70~ 100  否则 y 80~200
          x1 = Math.floor(Math.random() * (150 / 10)) * 10 + 110;
          if(x1 >= 140 || x1 <= 230){
            y1 = Math.floor(Math.random() * (40 / 10)) * 10 + 70;
          }else{
            y1 = Math.floor(Math.random() * (130 / 10)) * 10 + 80;
          }
        }
        let t = new Date().getTime().toString()
        t = parseInt(t.substring(t.length - 6, t.length)) + i
        const canvas = document.createElement('canvas');
        canvas.style.width = "390px";
        canvas.style.height = "300px";
        canvas.className  = 'cav'
        canvas.width = 390 * this.dpr;
        canvas.height = 300 * this.dpr;
        canvas.style.zIndex = t
        this.$refs.bossCav.appendChild(canvas);
        const ctx = canvas.getContext('2d');
        ctx.scale(this.dpr, this.dpr)
        this.$nextTick(() => {
          setTimeout(() => {
            if(type == 2){
              let obj = new Damage(x1, y1, dmgArr[i], ctx, canvas)
              obj.fadeCriDmgToBoss()
            }else{
              let obj = new Damage(x, y, dmgArr[i], ctx, canvas)
              obj.fadeNormalDmgToBoss()
              // console.log(obj)
            }
          },200*i)
        });
      }
    },
    drawMaterial(itemId,num){
      for (let i = 0; i < num; i++) {
        let t = new Date().getTime().toString()
        t = parseInt(t.substring(t.length - 6, t.length)) + i
        const canvas = document.createElement('canvas');
        canvas.style.width = "390px";
        canvas.style.height = "300px";
        canvas.className  = 'cav'
        canvas.width = 390 * this.dpr;
        canvas.height = 300 * this.dpr;
        canvas.style.zIndex = t
        this.$refs.bossCav.appendChild(canvas);
        const ctx = canvas.getContext('2d');
        ctx.scale(this.dpr, this.dpr)
        this.$nextTick(() => {
          setTimeout(() => {
            let obj = new FallItem(180,180,itemId,ctx,canvas,this.dpr)
            this.fallItems.push(obj)
            obj.fall()
          },20*i)
        });
      }
    },
    fadeMaterial(){
      if(this.fallItems.length > 0){
        for(let i = this.fallItems.length - 1; i >= 0; i--){
          this.fallItems[i].fade()
          this.fallItems.pop()
        }
      }
    }
  }
}
</script>

<style lang='stylus' src="./index.css" />
