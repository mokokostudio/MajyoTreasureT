<template>
    <div class="bfoot">
      <div :class="menuIndex===1?'item fons':'item'" @click="onMenu(1)"><img alt="NFT" src="~@/assets/boss/home-icon.png" /></div>
      <div :class="menuIndex===2?'item fons':'item'"  @click="onMenu(2)"><img alt="NFT" src="~@/assets/boss/nft-icon.png" /></div>
      <div :class="menuIndex===3?'item fons':'item'"  @click="onMenu(3)"><img :class="noMarket ? 'gray' : ''" alt="NFT" src="~@/assets/boss/market-icon.png" /></div>
      <div :class="menuIndex===4?'item fons':'item'" @click="onMenu(4)"><img :class="noPK ? 'gray':''" alt="NFT" src="~@/assets/boss/pvp-icon.png" /></div>

      <div class="popbox footbox" v-if="istip">
        <div class="bgs"></div>
        <!-- type 1-->
        <div class="popbody pkpop" v-if="isType === 1">
          <div class="head">
            <span class="text-stroke">Unlock conditions</span>
          </div>
          <div class="middle">
            <div class="content text-stroke">
              <div>- Equip 3 NFT weapon</div>
              <div>- Total NFT level > 30</div>
              <div>- Staking 20 Ton</div>
            </div>
          </div>
          <div class="foots" style="display: flex;">
            <div class="btn f30" style="margin-top: 30px;" @click="onPopClose()"><span class="text-stroke">Got it</span></div>
          </div>
          <div class="close-btn" @click="onPopClose()"></div>
        </div>
      </div>
      <div class="tips-div">
        <Tip ref="tip"></Tip>
      </div>
    </div>
</template>


<script>

export default {
    props: [],
    data() {
    return {
      menuIndex: 1,
      istip: false,
      isType: 0,
      noPK: false,
      noMarket: false
    }
  },
  activated() {
    if(this.$route.path==='/boss'){
      this.menuIndex = 1
    }
    if(this.$route.path==='/mint'){
      this.menuIndex = 2
    }
    if(this.$route.path==='/market'){
      this.menuIndex = 3
    }
    if(this.$route.path==='/pvp'){
      this.menuIndex = 4
    }
  },
  mounted() {

    if(this.$route.path==='/boss'){
      this.menuIndex = 1
    }

    if(this.$route.path==='/mint'){
      this.menuIndex = 2
    }
    if(this.$route.path==='/market'){
      this.menuIndex = 3
    }
    if(this.$route.path==='/pvp'){
      this.menuIndex = 4
    }

  },
  methods: {
      onMenu(i){
          let tempurl = ['/boss','/mint','/market','/pvp']
          if(i===1){
            this.menuIndex = i;
            this.$router.push({ path: tempurl[0] })
          } else if(i===2){
            this.menuIndex = i;
            this.$router.push({ path: tempurl[1] })
          }else if(i===3){
            // this.$refs.tip.openTip('Marketplace is building...', 1500)
            this.menuIndex = i;
            this.$router.push({ path: tempurl[2] })
          }else if(i===4){
            // this.openPK()
            this.menuIndex = i;
            this.$router.push({ path: tempurl[3] })
          }
      },
      onPopBox(type) {
        console.log(type, '====type')
        this.istip = true
        this.isType = parseInt(type)
      },
      onPopClose() {
        this.isType = 0
        this.istip = false
      },
      openPK(){
        this.onPopBox(1)
      }
  }
}
</script>

<style lang='stylus' src="./index.css" />
