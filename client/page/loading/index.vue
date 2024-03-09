<template>
  <div class="loading-view">
    <div class="bak"></div>
    <div class="bgtxt text-stroke">
      <div ref="txtDiv1"></div>
      <div ref="txtDiv2"></div>
      <div ref="txtDiv3"></div>
    </div>
    <div v-if="showProgress" class="loading-div text-stroke">
      <div class="loading-uni">
        <div>Loading</div>
        <div class="loading-text text-stroke">
        	 <span class="dot"></span>
        </div>
      </div>
      <div class="loading">
        <div class="progress">
          <div class="val text-stroke">{{loadingVal}}%</div>
          <div class="doing" :style="{ width: loadingVal + '%' }"></div>
        </div>
      </div>
    </div>
    <div v-if="!showProgress">
      <div class="starttxt text-stroke"><span>Click the button to start game.</span></div>
      <div class="gamestart"><span @click="gotoBoss()"></span></div>
    </div>
    <div class="tips-div">
      <Tip ref="tip"></Tip>
    </div>
  </div>
</template>
<script>
import cache from '@/utils/cache'
import { getItems } from '@/api/items'

export default {
  props: [],
  data() {
    return {
      showProgress: false,
      loginFinished: false,
      loginSuccess: false,
      loadingVal: 0,
      count:0,
      percent:0
    }
  },
  created() {
    this.handleLogin()
  },
  mounted() {
    this.showTxtAni()
    this.showLoadingProgress()
    //console.log("loading===", this.$store.state)

    // this.$store.commit('SET_ZPA', 10000)
    // console.log("zpa====",this.$store.state.zpa)
    this.percent = 0
    this.count = 0
    let files = require.context("@/assets",true,/\.*\.png|gif$/).keys()
    //console.log(files)
    for(let file of files){
      if(file.length > 2){
        let img = new Image()
        img.src = require('@/assets/'+file.substring(2, file.length))
        img.onload = (() => {
          this.count++
          // 计算图片加载的百分数，绑定到percent变量
          let percentNum = Math.floor(this.count / files.length * 100)
          this.percent = percentNum
          //console.log(this.percent)
        })
      }
    }



  },
  methods: {
    handleLogin() {
      const querys = this.$route.query || {}
      // console.log('querys==========', querys)
      let token = cache.getSession('token') || ''
      if (!token) {
        let data = { "token": querys.token || '' }
        // console.log('----data', data)
        this.$store.dispatch('user/login', data)
          .then(() => {
            console.log('=========loginSuccess', data)
            this.initItems()
          })
          .catch(() => {
            console.log('=========loginError', data)
            this.loginFinished = true
            this.loginSuccess = false
            // this.$refs.tip.openTip('123')
          })
      }else{
        this.initItems()
      }
    },
    initItems(){
      getItems().then(response => {
        if(response){
          console.log("get items success==", response)
          if(response.items){
            cache.setSession('bagItems', response.items)
          }
          if(response.base_equips){
            cache.setSession('baseEquip', response.base_equips)
          }
          if(response.base_equip_upgrade_infos){
            cache.setSession('baseUpInfo', response.base_equip_upgrade_infos)
          }
          if(response.mana){
            cache.setSession('zpa', response.mana)
          }
          this.loginFinished = true
          this.loginSuccess = true
        }else{
          this.$refs.tip.openTip('Get items error')
        }
      })
    },
    gotoBoss(){
      this.$router.push({ path: '/boss'})
    },
    showTxtAni(){
      let txt1 = "Awaken in a realm unknown, a mere blink ago, you traversed dimensions, summoned as the savior to this extraordinary world.";
      let txt2 = "Legend has it that 300 years prior, five formidable witches wreaked havoc, their powerful artifacts reducing civilization to fragments. Under their dark reign, humanity yearned for salvation, day after dreary day.";
      let txt3 = "Now, wielding your own mystical artifact, you stand as the chosen one. You're the savior, destined to confront the sorceresses and seize the endless treasures they guard.";
      let txtArr1 = txt1.split(" ")
      let txtArr2 = txt2.split(" ")
      let txtArr3 = txt3.split(" ")
      let txtDiv1 = this.$refs.txtDiv1
      let txtDiv2 = this.$refs.txtDiv2
      let txtDiv3 = this.$refs.txtDiv3
      let len1 = txtArr1.length
      let len2 = txtArr2.length
      let len3 = txtArr3.length
      let idx1 = 0
      let idx2 = 0
      let idx3 = 0
      let handle = setInterval(() => {
        // console.log("length:",len1,len2,len3)
        if(idx1 < len1){
          let span = document.createElement("span")
          span.textContent = txtArr1[idx1]
          txtDiv1.appendChild(span)
          idx1++;
        }else{
          len1 = -1
        }
        if(len1 < 0){
          if(idx2 < len2){
            let span = document.createElement("span")
            span.textContent = txtArr2[idx2]
            txtDiv2.appendChild(span)
            idx2++;
          }else{
            len2 = -1
          }
        }
        if(len1 < 0 && len2 < 0){
          if(idx3 < len3){
            let span = document.createElement("span")
            span.textContent = txtArr3[idx3]
            txtDiv3.appendChild(span)
            idx3++;
          }else{
            len3 = -1
          }
        }
        if(len1 < 0 && len2 < 0 && len3 < 0){
          clearInterval(handle)
        }
      },200)
    },
    showLoadingProgress(){
      if(!this.showProgress){
        this.showProgress = true
      }
      let perAdd = 1
      let tempVal = 0
      let handle = setInterval(() => {
        // console.log(this.loginFinished)
        if(this.loginFinished && this.percent == 100){
          perAdd = 100
        }
        if(tempVal < 100){
          if(this.loadingVal + perAdd > 100){
            this.loadingVal = 100
            tempVal = 100
          }else{
            this.loadingVal = this.loadingVal + perAdd
            tempVal = tempVal + perAdd
          }
        }else{
          tempVal = 999
        }
        // console.log("num", this.loadingVal, tempVal)
        if(tempVal == 999){
          if(this.loginFinished){
            clearInterval(handle)
            if(this.loginSuccess){
              this.showProgress = false
            }
            this.loginFinished = false
          }
        }
      },40)
    }
  }
}
</script>

<style lang='stylus' src="./index.css" />
