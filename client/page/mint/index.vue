<template>
  <div class="mint-view">
    <!-- headbox -->
    <headers/>
    <!-- weaponlist -->
    <div class="weaponlist">
      <div class="wbox-bg">
        <div class="wbox">
          <!-- <div v-for="item in 50" :key="item">{{ item }} </div> -->
          <div class="item" v-for="item in this.nfts.weapon" :key="item.type">
            <!-- <div :class="item === 2 || item === 4 ? 'collaction text-stroke' : 'ingame text-stroke'"
              v-text="item === 2 || item === 4 ? 'Collaction' : 'In Game'"></div> -->
            <div class="user">
              <div v-if="item.type == 1" class="pic">
                <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Garm1.png" />
                <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Garm2.png" />
                <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Garm3.png" />
                <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Garm4.png" />
                <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Garm5.png" />
                <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Garm6.png" />
              </div>
              <div v-if="item.type == 2" class="pic">
                <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Edda1.png" />
                <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Edda2.png" />
                <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Edda3.png" />
                <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Edda4.png" />
                <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Edda5.png" />
                <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Edda6.png" />
              </div>
              <div v-if="item.type == 3" class="pic">
                <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Norns1.png" />
                <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Norns2.png" />
                <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Norns3.png" />
                <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Norns4.png" />
                <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Norns5.png" />
                <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Norns6.png" />
              </div>
              <div v-if="item.type == 4" class="pic">
                <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Parvati1.png" />
                <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Parvati2.png" />
                <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Parvati3.png" />
                <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Parvati4.png" />
                <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Parvati5.png" />
                <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Parvati6.png" />
              </div>
              <div v-if="item.type == 5" class="pic">
                <img v-if="item.level == 1" alt="NFT" src="~@/assets/nfts/Tartarus1.png" />
                <img v-if="item.level > 1 && item.level <= 3" alt="NFT" src="~@/assets/nfts/Tartarus2.png" />
                <img v-if="item.level > 3 && item.level <= 6" alt="NFT" src="~@/assets/nfts/Tartarus3.png" />
                <img v-if="item.level > 6 && item.level <= 9" alt="NFT" src="~@/assets/nfts/Tartarus4.png" />
                <img v-if="item.level > 9 && item.level <= 12" alt="NFT" src="~@/assets/nfts/Tartarus5.png" />
                <img v-if="item.level > 12" alt="NFT" src="~@/assets/nfts/Tartarus6.png" />
              </div>
              <div class="name">#Test</div>
            </div>
            <div class="text text-stroke">
              <div class="tls">{{item.name}}</div>
              <div class="atr">
                <div>
                  <span>{{item.main_atr_1}}: {{item.main_atr_1_val / 100}}%</span>
                  <span>{{item.main_atr_2}}: {{item.main_atr_2_val / 100}}%</span>
                </div>
                <div>
                  <span v-if="item.add_atr_1_val > 0">{{item.add_atr_1}}: {{item.add_atr_1_val / 100}}%</span>
                  <span v-if="item.add_atr_2_val > 0">{{item.add_atr_2}}: {{item.add_atr_2_val / 100}}%</span>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

    </div>
    <!-- foot  @click="count++" {{ count }}-->
    <div class="foot">
      <div class="bxicon"><img alt="NFT" src="~@/assets/NFTChest.png" /></div>
      <div class="text text-stroke">Majyo Weapon</div>
      <div class="des">0 left</div>
      <div class="btn" @click="mintGo()"><span class="text-stroke">Mint</span></div>
    </div>

    <div class="tips-div">
      <Tip ref="tip"></Tip>
    </div>
    <foots/>

    <div class="popbox mintbox" v-if="istip">
      <div class="bgs"></div>

      <!-- type 1-->
      <div class="popbody" v-if="isType === 1">
        <div class="head">
          武器名
          <span class="icon"></span>
        </div>
        <div class="middle">
          <div class="type">
            <div class="pic"><img alt="NFT" src="~@/assets/nfts/Garm1.png" /></div>
            <div class="name">#1231231</div>
            <div class="ingame">In Game</div>

            <div class="atr">
              <span>ATK：5%</span>
              <span>CRT：5%</span>
              <span>NP：5%</span>
              <span>ATK：8%</span>
            </div>

          </div>
        </div>
        <div class="foots">
          <div class="btn">Claim</div>
        </div>
        <div class="close-btn" @click="istip = false"></div>
      </div>
      <!-- type 2-->
      <div class="popbody" v-if="isType === 2">
        <div class="head">
          武器名
          <span class="icon"></span>
        </div>
        <div class="middle">
          <div class="type">
            <div class="pic"><img alt="NFT" src="~@/assets/nfts/Garm1.png" /></div>
            <div class="name">#1231231</div>
            <div class="collaction">Collaction</div>
            <div class="atr">
              <span>ATK：5%</span>
              <span>CRT：5%</span>
              <span>NP：5%</span>
              <span>ATK：8%</span>
            </div>

          </div>
        </div>
        <div class="foots">
          <div class="btn f30">Charge Game</div>
        </div>
        <div class="close-btn" @click="istip = false"></div>
      </div>

      <!-- type 3-->
      <div class="popbody" v-if="isType === 3">
        <div class="head">
          Mint NFT
          <p class="tips">Mint for one random weapon</p>
          <span class="icon"></span>
        </div>
        <div class="middle">
          <div class="type">
            <div class="listpic">
              <div class="pic"><img alt="NFT" src="~@/assets/nfts/Garm1.png" /></div>
              <div class="pic"><img alt="NFT" src="~@/assets/nfts/Garm1.png" /></div>
              <div class="pic"><img alt="NFT" src="~@/assets/nfts/Garm1.png" /></div>
              <div class="pic"><img alt="NFT" src="~@/assets/nfts/Garm1.png" /></div>
              <div class="pic"><img alt="NFT" src="~@/assets/nfts/Garm1.png" /></div>
            </div>
            <p class="tips">Mint for two random attributes with</p>
            <p class="tips">randomizad values</p>
            <div class="blockhe">
              <div class="atr">
                <span>HP：5%</span>
                <span>ATK：5%</span>
                <span>HIK：5%</span>
                <span>SPD：5%</span>
                <span>CRT：5%</span>
              </div>
            </div>
            <div class="btxt">
              <p>GAS FEE:<span>TON 0.0101</span></p>
              <p>COST:0.1u≈<span>0.1256</span></p>
            </div>
          </div>
        </div>
        <div class="foots">
          <div class="btn f30" @click="onPopBox(4)">Mint</div>
        </div>
        <div class="close-btn" @click="istip = false"></div>
      </div>

      <!-- type 4-->
      <div class="popbody" v-if="isType === 4">
        <div class="head">
          Obtaining Rewards
          <span class="icon"></span>
        </div>
        <div class="middle">
          <div class="type">
            <div class="pic"><img alt="NFT" src="~@/assets/nfts/Garm1.png" /></div>
            <div class="name">#1231231</div>
            <div class="collaction">Collaction</div>
            <div class="atr">
              <span>ATK：5%</span>
              <span>CRT：5%</span>
            </div>
          </div>
        </div>
        <div class="foots">
          <div class="btn f30">Certitude</div>
        </div>
        <div class="close-btn" @click="istip = false"></div>
      </div>

    </div>
  </div>
</template>
<script>

// import { mapActions } from 'vuex'
import api from '@/utils/api'
import foots from '@/components/foots'
import headers from '@/components/header'
import nftsInit from '@/share/nftsTmp.json'
import jskit from '@/utils/index'
import { mint } from '@/api/index'
// import cache from '@/utils/cache'
// npm i @tonconnect/sdk
// https://ton-connect.github.io/sdk/modules/_tonconnect_sdk.html
// https://www.bilibili.com/video/BV1SM4y1t715/?spm_id_from=333.337.search-card.all.click&vd_source=a8c2aafa30345ac167cc14310c07f2dc
// https://www.nordicsemi.com/Products/Development-tools/nRF-Command-Line-Tools
// http://www.igiveyou.cn/?post=29
// https://pan.baidu.com/s/1TweXsQPKAKWretlG-iu5iQ?pwd=9p6m

// nRF Command Line Tools ->https://www.nordicsemi.com/Products/Development-tools/nRF-Command-Line-Tools
// JLink->https://pan.baidu.com/s/1TweXsQPKAKWretlG-iu5iQ?pwd=9p6m

// nRF connect for Destok ->https://www.nordicsemi.com/Products/Development-tools/nRF-Connect-for-Desktop/Download#infotabs
// nRF Connect SDK->

export default {
  data() {
    return {
      istip: false,
      isType: 0,
      nfts:{}
    }
  },
  components: {foots, headers},
  mounted() {
    this.nfts = jskit.localStorageGet('nftsTmp')
    if(!this.nfts){
      this.nfts = nftsInit
      jskit.localStorageSet('nftsTmp', this.nfts)
    }

  },
  activated(){

  },
  methods: {
    onPopBox(type) {
      this.istip = true;
      this.isType = parseInt(type);

      console.log(this.isType, 'this.isType')
    },
    mintGo() {
      // mint({}).then(response => {

      // })
      this.$refs.tip.openTip('Coming soon')
    }
  }
}
</script>

<style lang='stylus' src="./index.css" />
