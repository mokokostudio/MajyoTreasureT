<template>
  <!-- headbox -->
  <div class="headbox">
    <div class="user">
      <div class="userinfo">
        <div class="pic"><img alt="NFT" src="~@/assets/sculpture.png" /></div>
        <div class="text">
          <div class="name text-stroke">
            <span v-html="nickname"></span>
          </div>
          <!-- <div class="name">Ls voidborne</div> -->
          <div class="accountInfo">
            <div class="ether"><span class="text-stroke">{{(this.$store.state.zpa / 10000).toFixed(4)}}</span></div>
            <div class="des"><span class="text-stroke">0.00</span></div>
            <div class="energy" @click="showEnergyTip()">
              <span class="text-stroke">{{energy}}</span>
              <span class="text-stroke">/</span>
              <span class="text-stroke">200</span>
            </div>
          </div>
        </div>
      </div>
     <div class="head-bag" @click="onShowBag()">
        <img src="../../assets/bag/bag.png"/>
      </div>
    </div>
    <div class="energyTip" v-if="showenergytip">
      <div class="energyTipContent">Stamina recovers by 1 point every 8 minutes.</div>
    </div>

    <div class="bagbox" v-if="isshowbag">
      <div class="bagcontent">
        <div class="bagitem">
          <img src="../../assets/items/10.png"/>
          <span class="text-stroke">X&nbsp;{{items.item10}}</span>
        </div>
        <div class="bagitem">
          <img src="../../assets/items/11.png"/>
          <span class="text-stroke">X&nbsp;{{items.item11}}</span>
        </div>
        <div class="bagitem">
          <img src="../../assets/items/12.png"/>
          <span class="text-stroke">X&nbsp;{{items.item12}}</span>
        </div>
        <div class="bagitem">
          <img src="../../assets/items/13.png"/>
          <span class="text-stroke">X&nbsp;{{items.item13}}</span>
        </div>
        <div class="bagitem">
          <img src="../../assets/items/14.png"/>
          <span class="text-stroke">X&nbsp;{{items.item14}}</span>
        </div>
        <div class="bagitem">
          <img src="../../assets/items/15.png"/>
          <span class="text-stroke">X&nbsp;{{items.item15}}</span>
        </div>
      </div>
    </div>

    <div class="tips-div">
      <Tip ref="tip"></Tip>
    </div>
  </div>
</template>


<script>
import cache from '@/utils/cache'

export default {
  props: [],
  data() {
    return {
      nickname: '',
      energy: 0,
      zpa:0,
      isshowbag: false,
      showenergytip: false,
      items: {
        item10:0,
        item11:0,
        item12:0,
        item13:0,
        item14:0,
        item15:0
      },
      handle:{}
    }
  },
  activated() {
    this.getUserInfo()

    clearInterval(this.handle)
    this.handle = setInterval(() => {
      console.log(this.energy,"===local eng")
      if(this.energy < 200){
        this.energy ++;
        cache.setSession('energy', this.energy)
      }else{
        clearInterval(this.handle)
      }
    },480000)
  },
  created(){

  },
  mounted() {
    // console.log( this.$router,'====',this.$route.path)
    if (this.$route.path === '/boss') {
      this.menuIndex = 1
    }

    if (this.$route.path === '/index') {
      this.menuIndex = 2
    }

  },
  methods: {
    getUserInfo() {
      this.nickname = cache.getSession('nickName') || 'Player'
      this.energy =+ cache.getSession('energy') || 0
      let _zpa =+ cache.getSession('zpa') || 0
      this.$store.commit("GLOBAL_ZPA", _zpa)
      if(this.$store.state.zpa != 0){
        this.zpa = (this.$store.state.zpa / 10000).toFixed(4)
      }

      //console.log("store zpa=", this.$store.state.zpa)
    },
    onShowBag() {
      let _items = JSON.parse(cache.getSession('bagItems')) || {}
      let _sbxData = JSON.parse(cache.getSession('baseEquip')) || {}
      let _sbxUpData = JSON.parse(cache.getSession('baseUpInfo')) || {}
      console.log(_items)
      console.log(_sbxData)
      console.log(_sbxUpData)
      if (Object.keys(_items).length != 0) {
        for (let item of _items) {
          if (item.item_id == 10) {
            this.items.item10 = item.num
          } else if (item.item_id == 11) {
            this.items.item11 = item.num
          } else if (item.item_id == 12) {
            this.items.item12 = item.num
          } else if (item.item_id == 13) {
            this.items.item13 = item.num
          } else if (item.item_id == 14) {
            this.items.item14 = item.num
          } else if (item.item_id == 15) {
            this.items.item15 = item.num
          }
        }
      }
      if(!this.isshowbag){
        this.isshowbag = true;
      }else{
        this.isshowbag = false;
      }
    },
    showEnergyTip(){
      this.showenergytip = true;
      setTimeout(() => {
        this.showenergytip = false;
      }, 1.5 * 1000)
    }
  }
}
</script>

<style lang='stylus' src="./index.css" />
