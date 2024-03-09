<template>
  <div class="pvpfight-view">
    <div class="battle-ground">
      <div class="head-info text-stroke">
        <div class="my-head">
          <div class="avt">
            <img alt="NFT" src="~@/assets/sculpture.png" />
            <span>{{nickname}}</span>
          </div>
          <div class="my-hp">
            <div class="hp-progress">
              <div class="val my-val text-stroke">{{this.progressChallengerVal}}</div>
              <div class="victory" :style="{ width: progressChallenger + '%' }"></div>
            </div>
          </div>
        </div>
        <div class="tar-head">
          <div class="avt">
            <span>{{this.targetName}}</span>
            <img alt="NFT" src="~@/assets/sculpture.png" />
          </div>
          <div class="tar-hp">
            <div class="hp-progress">
              <div class="val tar-val text-stroke">{{this.progressDefenderVal}}</div>
              <div class="fail" :style="{ width: progressDefender + '%' }"></div>
            </div>
          </div>
        </div>
      </div>
      <div class="battle-body" ref="batbodyCavs"></div>
      <div class="battle-equip">
        <div class="my-nft">
          <div class="nftequip" v-for="nft in this.myNfts" :key="nft.id">
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
        <div class="tar-nft">
          <div class="nftequip">
            <img alt="NFT"  src="~@/assets/nfts/Norns4.png" />
            <div class="nftlv tarlv"><span class="text-stroke">11</span></div>
          </div>
          <div class="nftequip">
            <img alt="NFT"  src="~@/assets/nfts/Garm6.png" />
            <div class="nftlv tarlv"><span class="text-stroke">15</span></div>
          </div>
        </div>
      </div>
    </div>

    <div class="battle-info-show">
      <div class="pkshow" ref="pktxtbody">
        <div class="itemdata onsd text-stroke" v-for="(item, index) in battleTxt" :key="index">
          <span class="my" v-if="item.showTitle && item.status === 0" v-html="nickname"></span>
          <span class="tar" v-if="item.showTitle && item.status === 1">{{targetName}}</span>
          <span>{{ item.text }}</span>
        </div>
      </div>
    </div>

    <!-- pop start -->
    <div class="popbox pvpfightbox" v-if="istip">
      <div class="bgs"></div>
      <!-- type 1-->
      <div class="popbody" v-if="isType === 1">
        <div class="head">
          <span v-if="isWin === 1" class="text-stroke">Congratulations</span>
          <span v-if="isWin === 0" class="text-stroke">Challenge Failed</span>
        </div>
        <div class="middle">
          <div v-if="isWin === 1" class="notice-win text-stroke">
            <div>You won the battle</div>
            <div>and ranking becomes</div>
            <div class="num">No.{{this.targetRank}}</div>
          </div>
          <div v-if="isWin === 0" class="notice-win text-stroke">
            <div>Your ranking is still No.{{this.myRank}}</div>
            <div class="suggest">·Enhance equipment to defeat opponents</div>
          </div>
        </div>
        <div class="foots">
          <div class="btn f30" @click="onPopClose(1)"><span class="text-stroke">Confirm</span></div>
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

import cache from '@/utils/cache'
import jskit from '@/utils/index'
import tmpdata from './tmpdata.json'
import { pvpFight } from '@/api/fight'
import { MainCharacter } from '@/animation/mainCharacter'
import { Damage } from '@/animation/damageFade'

export default {
  data() {
    return {
      enemyskillCav: null,
      myskillCav:null,

      istip: false,
      isType: 0,

      targetRank: 0,
      targetUid: 0,
      targetName:'',
      myRank:0,

      myNfts:{},
      tarNfts:{},

      bossItem: 1,

      battleTxt: [],
      len: 1,
      progressDefender: 100,
      progressChallenger: 100,
      progressDefenderVal: "--",
      progressChallengerVal: "--",
      hit: false,
      is_cri: false,
      nickname: '',

      showBossDmg: false,
      defenderDmgTxt: '',
      showBossCoin: false,
      recoverBossCoin: false,
      isWin: 0,

      // 战斗
      dpr: 1,
      battleFinish: false,
      myCharacter:{},
      enemyCharacter:{},
      skillAni : {}
    }
  },
  computed: {
    textLines() {
      return this.text.split('\n');
    },
  },
  components: {},
  created() {
  },
  activated() {
    this.dpr = window.devicePixelRatio || 1
    this.$nextTick(() => {
      console.log('dpr==',this.dpr)
      //初始化
      this.init()
      this.initCavs()
      this.initNfts()

      //init boss信息
      console.log("start pk....")
      const querys = this.$route.query || {}

      if(querys && querys.rank && querys.uid){
        this.targetRank = querys.rank
        this.targetUid = querys.uid
        this.targetName = querys.tname || 'Defender'
      }else{
        this.$refs.tip.openTip('Param Err')
      }

      this.pvpFight()
      //this.testFight()

      this.createBattleAni()
    })
  },
  updated(){
    this.scrollToBottom()
  },
  methods: {
    init() {
      this.targetRank = 0
      this.targetUid = 0
      this.progressDefender = 100
      this.progressChallenger = 100
      this.progressDefenderVal = '--'
      this.progressChallengerVal = '--'
      this.battleTxt = []
      this.len = 1
      this.hit = false
      this.istip = false
      this.isType = 0
      this.defenderDmgTxt = ''
      this.isWin = 0

      this.nickname = cache.getSession('nickName') || 'Not found'
    },
    initCavs(){
      this.battleFinish = false

      this.myCharacter = new MainCharacter(1, this.dpr, this.$refs.batbodyCavs)
      console.log('my character ',this.myCharacter)

      this.enemyCharacter = new MainCharacter(2, this.dpr, this.$refs.batbodyCavs)
      console.log('enemy character ',this.enemyCharacter)

    },
    clearCav(){
      let childNode = this.$refs.batbodyCavs.childNodes
      for(let i = childNode.length - 1; i >= 0; i--){
        this.$refs.batbodyCavs.removeChild(childNode[i])
      }
    },
    initNfts(){
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
      if(nft3){
        nftArr.push(nft3)
      }
      this.myNfts = nftArr
    },
    scrollToBottom(){
      this.$nextTick(()=>{
        let body = this.$refs.pktxtbody
        body.scrollTo({ top: body.scrollHeight, behavior: 'smooth' })
      })
    },
    pvpFight() {
      let data = {
        target_id : parseInt(this.targetUid),
        target_rank : parseInt(this.targetRank)
      }
      console.log('-----------data', data)
      pvpFight(data).then(response => {
        console.log('pvp fight response', response)
        // let fightData = tmpdata
        let fightData = response
        console.log(fightData)
        if(fightData && fightData.details && fightData.details.length > 0){
          let mana =+ fightData.mana || 0
          cache.setSession('zpa', mana)
          this.$store.commit("GLOBAL_ZPA", mana)
          this.progressDefenderVal = fightData.defender_hp + ''
          this.progressChallengerVal = fightData.challenger_hp + ''
          if(fightData.win){
            this.isWin = 1
            this.targetRank = fightData.new_rank
          }else{
            this.isWin = 0
            this.myRank = fightData.old_rank
          }
          console.log('此局我方赢：', fightData.win)
          //判断谁先开
          let isme = false
          if(fightData.details[0].attacker_type && fightData.details[0].attacker_type === 2){
            // type 2 攻击者   3  防守者
            isme = true
          }else{
            if(fightData.details[0].be_attacker && fightData.details[0].be_attacker === 3){
              isme = true
            }
          }
          // console.log("isme===",isme)
          this.battleTxt.push({
            status: !isme ? 1 : 0,
            text: ' starts the attack first...',
            showTitle: true
          })

          let challengeInitHp = fightData.challenger_hp || 0
          let defenderInitHp = fightData.defender_hp || 0
          let time = setInterval(() => {
            if(this.myCharacter.status == 2 && this.enemyCharacter.status == 2){
              this.len ++
              let item = fightData.details[this.len - 2]
              // console.log(item)
              let atkDefenderNow = false
              if(item.be_attacker && item.be_attacker === 3){
                atkDefenderNow = true
              }
              let atkChallengerNow = false
              if(item.attacker_type && item.attacker_type === 3){
                atkChallengerNow = true
              }
              // console.log('====攻击防守者', atkDefenderNow, "===攻击挑战者",atkChallengerNow)
              // console.log(atkDefenderNow ? "我方 -> 对方:" : "对方 -> 我方:", '被攻击前：', item.hp_before, '被攻击后：', item.hp_after, ' 造成伤害', item.dmg ,item.is_miss ? ' Miss' : '', item.is_cri ? ' 暴击':'')

              // 设置双方血条
              if (atkDefenderNow) {
                let defenderHP = item.hp_after || 0
                this.enemyCharacter.hpLeft = defenderHP
                if(item.is_cri){
                  this.createMySkill()
                  setTimeout(() => {
                    let dmg = item.is_miss ? 0 : item.dmg
                    this.drawDmg(false, dmg, 2, 3)
                    this.progressDefender = defenderHP == defenderInitHp ? 100 : Math.ceil(defenderHP / defenderInitHp * 100)
                    this.progressDefenderVal = defenderHP + ''
                  }, 1000)
                }else{
                  this.attackAni(1)
                  setTimeout(() => {
                    let dmg = item.is_miss ? 0 : item.dmg
                    this.drawDmg(false, dmg, 1, 1)
                    this.progressDefender = defenderHP == defenderInitHp ? 100 : Math.ceil(defenderHP / defenderInitHp * 100)
                    this.progressDefenderVal = defenderHP + ''
                  }, 400)
                }
              }
              if (atkChallengerNow) {
                let challengerHP = item.hp_after || 0
                this.myCharacter.hpLeft = challengerHP
                if(item.is_cri){
                  this.createEnemySkill()
                  setTimeout(() => {
                    let dmg = item.is_miss ? 0 : item.dmg
                    this.drawDmg(true, dmg, 2, 5)
                    this.progressChallenger = challengerHP ==  challengeInitHp ? 100 : Math.ceil(challengerHP / challengeInitHp * 100)
                    this.progressChallengerVal = challengerHP + ''
                  }, 1000)
                }else{
                  this.attackAni(2)
                  setTimeout(() => {
                    let dmg = item.is_miss ? 0 : item.dmg
                    this.drawDmg(true, dmg, 1, 1)
                    this.progressChallenger = challengerHP ==  challengeInitHp ? 100 : Math.ceil(challengerHP / challengeInitHp * 100)
                    this.progressChallengerVal = challengerHP + ''
                  }, 400)
                }
              }

              let text = '';
              if (item.is_miss) {
                text = 'attacked, but it was missed.'
              } else if (item.is_cri) {
                text = 'cast a skill, and caused '+item.dmg+' critical hit damage.'
              } else {
                text = 'caused ' + item.dmg + ' damage.'
              }
              this.battleTxt.push({
                status: atkChallengerNow ? 1 : 0,
                text : text,
                showTitle : true
              })

              // 本次伤害后有一方血量归0
              if (!item.hp_after || item.hp_after === undefined || item.hp_after === 0) {
                text = 'was defeated.'
                this.battleTxt.push({
                  status: atkChallengerNow ? 0 : 1, // 需要反一下，被击败的是对方
                  text : text,
                  showTitle : true
                })
              }

              // 退出战斗
              if (this.len > fightData.details.length) {
                if(this.isWin == 1){
                  text = 'Battle victory, you defeated ' + this.targetName
                }else{
                  text = 'Challenge failed...'
                }
                this.battleTxt.push({
                  status: atkChallengerNow ? 1 : 0,
                  text: text,
                  showTitle: false
                })
                // this.hit = false
                //延时弹窗
                setTimeout(() => {
                  // 延时之后进入结算弹窗
                  this.onPopBox(1)
                }, 1500)
                //结束定时器
                clearInterval(time)
              }
            }
          }, 100)

        }else{
          this.$refs.tip.openTip('Fight error')
        }
      })
    },
    testFight() {
      let data = {
        target_id : parseInt(this.targetUid),
        target_rank : parseInt(this.targetRank)
      }
      console.log('-----------data', data)
      let fightData = tmpdata
      console.log(fightData)
      if(fightData && fightData.details && fightData.details.length > 0){
        let mana =+ fightData.mana || 0
        cache.setSession('zpa', mana)
        this.$store.commit("GLOBAL_ZPA", mana)
        this.progressDefenderVal = fightData.defender_hp + ''
        this.progressChallengerVal = fightData.challenger_hp + ''
        if(fightData.win){
          this.isWin = 1
          this.targetRank = fightData.new_rank
        }else{
          this.isWin = 0
          this.myRank = fightData.old_rank
        }
        console.log('此局我方赢：', fightData.win)
        //判断谁先开
        let isme = false
        if(fightData.details[0].attacker_type && fightData.details[0].attacker_type === 2){
          // type 2 攻击者   3  防守者
          isme = true
        }else{
          if(fightData.details[0].be_attacker && fightData.details[0].be_attacker === 3){
            isme = true
          }
        }
        // console.log("isme===",isme)
        this.battleTxt.push({
          status: !isme ? 1 : 0,
          text: ' starts the attack first...',
          showTitle: true
        })

        let challengeInitHp = fightData.challenger_hp || 0
        let defenderInitHp = fightData.defender_hp || 0
        let time = setInterval(() => {
          if(this.myCharacter.status == 2 && this.enemyCharacter.status == 2){
            this.len ++
            let item = fightData.details[this.len - 2]
            // console.log(item)
            let atkDefenderNow = false
            if(item.be_attacker && item.be_attacker === 3){
              atkDefenderNow = true
            }
            let atkChallengerNow = false
            if(item.attacker_type && item.attacker_type === 3){
              atkChallengerNow = true
            }
            // console.log('====攻击防守者', atkDefenderNow, "===攻击挑战者",atkChallengerNow)
            // console.log(atkDefenderNow ? "我方 -> 对方:" : "对方 -> 我方:", '被攻击前：', item.hp_before, '被攻击后：', item.hp_after, ' 造成伤害', item.dmg ,item.is_miss ? ' Miss' : '', item.is_cri ? ' 暴击':'')

            // 设置双方血条
            if (atkDefenderNow) {
              let defenderHP = item.hp_after || 0
              this.enemyCharacter.hpLeft = defenderHP
              if(item.is_cri){
                this.createMySkill()
                setTimeout(() => {
                  let dmg = item.is_miss ? 0 : item.dmg
                  this.drawDmg(false, dmg, 2, 3)
                  this.progressDefender = defenderHP == defenderInitHp ? 100 : Math.ceil(defenderHP / defenderInitHp * 100)
                  this.progressDefenderVal = defenderHP + ''
                }, 1000)
              }else{
                this.attackAni(1)
                setTimeout(() => {
                  let dmg = item.is_miss ? 0 : item.dmg
                  this.drawDmg(false, dmg, 1, 1)
                  this.progressDefender = defenderHP == defenderInitHp ? 100 : Math.ceil(defenderHP / defenderInitHp * 100)
                  this.progressDefenderVal = defenderHP + ''
                }, 400)
              }
            }
            if (atkChallengerNow) {
              let challengerHP = item.hp_after || 0
              this.myCharacter.hpLeft = challengerHP
              if(item.is_cri){
                this.createEnemySkill()
                setTimeout(() => {
                  let dmg = item.is_miss ? 0 : item.dmg
                  this.drawDmg(true, dmg, 2, 5)
                  this.progressChallenger = challengerHP ==  challengeInitHp ? 100 : Math.ceil(challengerHP / challengeInitHp * 100)
                  this.progressChallengerVal = challengerHP + ''
                }, 1000)
              }else{
                this.attackAni(2)
                setTimeout(() => {
                  let dmg = item.is_miss ? 0 : item.dmg
                  this.drawDmg(true, dmg, 1, 1)
                  this.progressChallenger = challengerHP ==  challengeInitHp ? 100 : Math.ceil(challengerHP / challengeInitHp * 100)
                  this.progressChallengerVal = challengerHP + ''
                }, 400)
              }

            }

            let text = '';
            if (item.is_miss) {
              text = 'attacked, but it was missed.'
            } else if (item.is_cri) {
              // 播暴击振屏
              //this.is_cri = true;
              //text = 'delivered a full-force strike, and caused '+item.dmg+' critical hit damage.'
              text = 'cast a skill, and caused '+item.dmg+' critical hit damage.'
            } else {
              text = 'caused ' + item.dmg + ' damage.'
            }
            this.battleTxt.push({
              status: atkChallengerNow ? 1 : 0,
              text : text,
              showTitle : true
            })

            // 本次伤害后有一方血量归0
            if (!item.hp_after || item.hp_after === undefined || item.hp_after === 0) {
              text = 'was defeated.'
              this.battleTxt.push({
                status: atkChallengerNow ? 0 : 1, // 需要反一下，被击败的是对方
                text : text,
                showTitle : true
              })
            }

            // 退出战斗
            if (this.len > fightData.details.length) {
              if(this.isWin == 1){
                text = 'Battle victory, you defeated ' + this.targetName
              }else{
                text = 'Challenge failed...'
              }
              this.battleTxt.push({
                status: atkChallengerNow ? 1 : 0,
                text: text,
                showTitle: false
              })
              // this.hit = false
              //延时弹窗
              setTimeout(() => {
                // 延时之后进入结算弹窗
                this.onPopBox(1)
              }, 1500)
              //结束定时器
              clearInterval(time)
            }
          }
        }, 100)

      }else{
        this.$refs.tip.openTip('Fight error')
      }
    },
    //打开弹窗
    onPopBox(type) {
      this.isType = parseInt(type)
      this.istip = true
    },
    onPopClose(closeType) {
      this.istip = false
      this.isType = 0
      if (closeType == 1) {
        this.battleFinish = true
        this.clearCav()
        this.$router.push({ path: '/pvp'})
      }
    },

    // 战斗canvas
    createAniCavs(element, i){
      let t = new Date().getTime().toString()
      t = parseInt(t.substring(t.length - 6, t.length)) + i
      element.style.width = "380px";
      element.style.height = "230px";
      element.className  = 'battlecav'
      element.width = 380 * this.dpr;
      element.height = 230 * this.dpr;
      element.style.zIndex = t
      this.$refs.batbodyCavs.appendChild(element);
      const ctx = element.getContext('2d');
      ctx.scale(this.dpr, this.dpr)
      return ctx
    },
    createBattleAni(){
      this.myCharacter.move(20)
      this.enemyCharacter.move(20)
      //this.myCharacter.stand()
      // setTimeout(() => {
      //   this.myCharacter.attack(60)
      //   this.enemyCharacter.attack(60)
      // },1200)
    },
    attackAni(type){
      if(type === 1){
        this.myCharacter.attack(60,this.enemyCharacter)
      }else{
        this.enemyCharacter.attack(60,this.myCharacter)
      }
    },
    createMySkill(){
      let sk = Math.round(Math.random() * 5)
      let code = ''
      if(sk == 0){
        code = '1000'
      }else if(sk == 1){
        code = '1001'
      }else if(sk == 2){
        code = '1002'
      }else if(sk == 3){
        code = '1003'
      }else if(sk == 4){
        code = '1004'
      }else if(sk == 5){
        code = '1005'
      }
      this.myCharacter.Skill(code, this.enemyCharacter)
    },
    createEnemySkill(){
      let sk = Math.round(Math.random() * 5)
      let code = ''
      if(sk == 0){
        code = '1006'
      }else if(sk == 1){
        code = '1007'
      }else if(sk == 2){
        code = '1008'
      }else if(sk == 3){
        code = '1009'
      }else if(sk == 4){
        code = '1010'
      }else if(sk == 5){
        code = '1011'
      }
      this.enemyCharacter.Skill(code, this.myCharacter)
    },
    drawDmg(atkMy,dmgTxt,type,showTimes){
      let _dmg = parseInt(dmgTxt)
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
      let cht = atkMy ? this.myCharacter : this.enemyCharacter
      let dir = atkMy ? -1 : 1
      for (let i = 0; i < showTimes; i++) {
        const canvas = document.createElement('canvas');
        const ctx = this.createAniCavs(canvas, i)
        this.$nextTick(() => {
          setTimeout(() => {
            let x = cht.xStand + cht.width / 2
            let y = cht.yStand
            let obj = new Damage(x, y, dmgArr[i], ctx, canvas)
            obj.moveDirection(dir)
            if(type === 1){
              obj.fadeNormalDmgToPVP()
            }else{
              obj.fadeCriDmgToPVP()
            }
          },200*i)
        });
      }
    }

  }
}
</script>

<style lang='stylus' src="./index.css" />
