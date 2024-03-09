<template>
  <div class="market-view">
    <!-- headbox -->
    <headers/>

    <div v-if="this.isBuy" class="goods-list">
      <div v-if="orderData.length === undefined || orderData.length <= 0" class="find-notice">
        <span v-if="showLoading" class="text-stroke">freshing the goods<span class="dot"></span></span>
        <span v-if="!showLoading" class="text-stroke">No goods</span>
      </div>
      <div v-if="orderData && orderData.length > 0" class="goods" v-for="order in orderData" :key="order.order_uuid" @click="selectOnSaleOrder(order.order_uuid)">
        <div class="goods-item">
          <div class="goods-icon">
            <img v-if="order.goods.item_id == 10" src="~@/assets/items/10.png" />
            <img v-if="order.goods.item_id == 11" src="~@/assets/items/11.png" />
            <img v-if="order.goods.item_id == 12" src="~@/assets/items/12.png" />
            <img v-if="order.goods.item_id == 13" src="~@/assets/items/13.png" />
            <img v-if="order.goods.item_id == 14" src="~@/assets/items/14.png" />
            <img v-if="order.goods.item_id == 15" src="~@/assets/items/15.png" />
          </div>
          <div class="goods-info text-stroke"><span>Items</span><span class="goods-num">&nbsp;x&nbsp;{{order.goods.num}}</span></div>
        </div>
        <div class="goods-sale-info">
          <div v-if="order.status == 0" class="goods-state text-stroke">{{order.countTime}}s</div>
          <div class="goods-price">
            <img src="~@/assets/ETHERicon.png" />
            <span class="text-stroke">{{(order.price / 10000).toFixed(4)}}</span>
          </div>
        </div>
      </div>
      <div class="more" v-if="showMoreBtn"><div class="more-btn" @click="showMore()"><span class="text-stroke">More</span></div></div>
    </div>

    <div v-if="!this.isBuy" class="my-goods-list">
      <div class="my-listed-items">
        <div v-if="myOrderData && myOrderData.length > 0" class="goods" v-for="order in myOrderData" :key="order.order_uuid" @click="selectMyOrder(order.order_uuid)">
          <div v-if="order.status == 2" class="goods-repoint"></div>
          <div class="goods-item">
            <div class="goods-icon">
              <img v-if="order.goods.item_id == 10" src="~@/assets/items/10.png" />
              <img v-if="order.goods.item_id == 11" src="~@/assets/items/11.png" />
              <img v-if="order.goods.item_id == 12" src="~@/assets/items/12.png" />
              <img v-if="order.goods.item_id == 13" src="~@/assets/items/13.png" />
              <img v-if="order.goods.item_id == 14" src="~@/assets/items/14.png" />
              <img v-if="order.goods.item_id == 15" src="~@/assets/items/15.png" />
            </div>
            <div class="goods-info text-stroke"><span>Items</span><span class="goods-num">&nbsp;x&nbsp;{{order.goods.num}}</span></div>
          </div>
          <div class="goods-sale-info">
            <div v-if="order.status == 2" class="goods-state text-stroke">Sold out</div>
            <div class="goods-price">
              <img src="~@/assets/ETHERicon.png" />
              <span class="text-stroke">{{(order.price / 10000).toFixed(4)}}</span>
            </div>
          </div>
        </div>
        <div v-if="leftBucket > 0" class="goods" v-for="index in leftBucket"></div>
      </div>
      <div class="list-window">
        <div class="list-body">
          <div class="list-item">
            <div>
              <img v-if="this.selectItem == 10" src="../../assets/items/10.png"/>
              <img v-if="this.selectItem == 11" src="../../assets/items/11.png"/>
              <img v-if="this.selectItem == 12" src="../../assets/items/12.png"/>
              <img v-if="this.selectItem == 13" src="../../assets/items/13.png"/>
              <img v-if="this.selectItem == 14" src="../../assets/items/14.png"/>
              <img v-if="this.selectItem == 15" src="../../assets/items/15.png"/>
            </div>
          </div>
          <div class="list-info">
            <div class="list-num">
              <span class="text-stroke">Number:&nbsp;</span>
              <input type="number" class="input-placeholder" v-model="formData.num" placeholder="Number of inputs">
              <span class="text-stroke">&nbsp;Left:{{this.selectLeftNum}}</span>
            </div>
            <div class="list-price">
              <img src="~@/assets/ETHERicon.png" />
              <span class="text-stroke">Total Price:&nbsp;</span>
              <input type="number" class="input-placeholder" v-model="formData.price" placeholder="Enter total price" @change="calFee()">
              <span class="text-stroke">&nbsp;Need > 0.01</span>
            </div>
          </div>
          <div class="list-deal">
            <div class="list-fee text-stroke">
              <span>Handling fee:&nbsp;</span>
              <img src="~@/assets/ETHERicon.png" />
              <span class="fee">{{this.handleFee}}</span>
            </div>
            <div class="list-btn">
              <div v-if="!listProtect" @click="saleItems()"><span class="text-stroke">List</span></div>
              <div v-if="listProtect"><span class="text-stroke">List</span></div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="goods-filter">
      <div class="fliter-items">
        <div @click="selectitemFilter(10)">
          <img v-if="this.selectItem == 10" class="fliter-choiced" src="../../assets/icon-border.png"/>
          <img src="../../assets/items/10.png"/>
        </div>
        <div @click="selectitemFilter(11)">
          <img v-if="this.selectItem == 11" class="fliter-choiced" src="../../assets/icon-border.png"/>
          <img src="../../assets/items/11.png"/>
        </div>
        <div @click="selectitemFilter(12)">
          <img v-if="this.selectItem == 12" class="fliter-choiced" src="../../assets/icon-border.png"/>
          <img src="../../assets/items/12.png"/>
        </div>
        <div @click="selectitemFilter(13)">
          <img v-if="this.selectItem == 13" class="fliter-choiced" src="../../assets/icon-border.png"/>
          <img src="../../assets/items/13.png"/>
        </div>
        <div @click="selectitemFilter(14)">
          <img v-if="this.selectItem == 14" class="fliter-choiced" src="../../assets/icon-border.png"/>
          <img src="../../assets/items/14.png"/>
        </div>
        <div @click="selectitemFilter(15)">
          <img v-if="this.selectItem == 15" class="fliter-choiced" src="../../assets/icon-border.png"/>
          <img src="../../assets/items/15.png"/>
        </div>
      </div>
      <div class="goods-dotype">
        <div :class="this.isBuy ? 'g-btn-unable':'g-btn'" @click="changeTab(1)"><span class="text-stroke">Buy</span></div>
        <div :class="this.isBuy ? 'g-btn':'g-btn-unable'" @click="changeTab(2)"><span class="text-stroke">Sell</span></div>
      </div>
    </div>

    <div class="tips-div">
      <Tip ref="tip"></Tip>
    </div>
    <foots/>

    <div class="popbox marketbox" v-if="istip">
      <div class="bgs"></div>
      <!-- type 1-->
      <div class="popbody" v-if="isType === 1">
        <div class="head">
          <span class="text-stroke">My Order</span>
        </div>
        <div class="middle">
          <div class="order-item">
            <div v-if="this.selectOrder" class="order-item-info">
              <img v-if="this.selectOrder.goods.item_id == 10" src="../../assets/items/10.png"/>
              <img v-if="this.selectOrder.goods.item_id == 11" src="../../assets/items/11.png"/>
              <img v-if="this.selectOrder.goods.item_id == 12" src="../../assets/items/12.png"/>
              <img v-if="this.selectOrder.goods.item_id == 13" src="../../assets/items/13.png"/>
              <img v-if="this.selectOrder.goods.item_id == 14" src="../../assets/items/14.png"/>
              <img v-if="this.selectOrder.goods.item_id == 15" src="../../assets/items/15.png"/>
            </div>
          </div>
          <div class="order-sale-info">
            <span class="text-stroke">Num: {{this.selectOrder.goods.num}}</span>
          </div>
          <div class="order-sale-info">
            <span class="text-stroke">Price:<img src="~@/assets/ETHERicon.png" />{{(this.selectOrder.price / 10000).toFixed(4)}}</span>
          </div>
        </div>
        <div class="foots">
          <div v-if="!delistProtect" class="btn f30 noani" @click="delist()"><span class="text-stroke">Delist</span></div>
          <div v-if="delistProtect" class="btn f30 noani"><span class="text-stroke">Delist</span></div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>

      <div class="popbody" v-if="isType === 2">
        <div class="head">
          <span class="text-stroke">Order</span>
        </div>
        <div class="middle">
          <div class="order-item">
            <div v-if="this.selectOrder" class="order-item-info">
              <img v-if="this.selectOrder.goods.item_id == 10" src="../../assets/items/10.png"/>
              <img v-if="this.selectOrder.goods.item_id == 11" src="../../assets/items/11.png"/>
              <img v-if="this.selectOrder.goods.item_id == 12" src="../../assets/items/12.png"/>
              <img v-if="this.selectOrder.goods.item_id == 13" src="../../assets/items/13.png"/>
              <img v-if="this.selectOrder.goods.item_id == 14" src="../../assets/items/14.png"/>
              <img v-if="this.selectOrder.goods.item_id == 15" src="../../assets/items/15.png"/>
            </div>
          </div>
          <div class="order-sale-info">
            <span class="text-stroke">Num: {{this.selectOrder.goods.num}}</span>
          </div>
          <div class="order-sale-info">
            <span class="text-stroke">Price:<img src="~@/assets/ETHERicon.png" />{{(this.selectOrder.price / 10000).toFixed(4)}}</span>
          </div>
        </div>
        <div class="foots">
          <div v-if="!purchaseProtect" class="btn f30" @click="purchase()"><span class="text-stroke">Purchase</span></div>
          <div v-if="purchaseProtect" class="btn f30"><span class="text-stroke">Purchase</span></div>
        </div>
        <div class="close-btn" @click="onPopClose()"></div>
      </div>

    </div>

  </div>
</template>
<script>

import api from '@/utils/api'
import foots from '@/components/foots'
import headers from '@/components/header'
import cache from '@/utils/cache'
import jskit from '@/utils/index'
import { getOrders, getMyOrders, publish, purchase, takeOff } from '@/api/market'

export default {
  data() {
    return {
      istip: false,
      isType: 0,
      isBuy: true,
      selectItem: 10,
      selectLeftNum: 0,
      formData:{},
      orderData:{},
      myOrderData:{},
      selectOrder:{},
      page:0,
      leftBucket:8,
      showMoreBtn: false,
      showLoading: true,
      handleFee: 0,
      countHandle:null,

      purchaseProtect: false,
      listProtect: false,
      delistProtect: false
    }
  },
  components: {foots, headers},
  mounted() {

  },
  activated(){
    this.page = 0
    this.getOrders(10, this.page)

    this.countHandle = setInterval(() => {
      this.showOrderTimeCount()
    },1000)
  },
  beforeRouteLeave(to, from, next) {
    clearInterval(this.countHandle)
    next()
  },
  methods: {
    onPopBox(type) {
      this.istip = true;
      this.isType = parseInt(type);

      console.log(this.isType, 'this.isType')
    },
    onPopClose(){
      this.istip = false
      this.isType = 0
    },
    selectitemFilter(index){
      if(this.selectItem != index){
        this.selectItem = index
        this.page = 0
        this.getOrders(index, this.page)
        this.calLeftMaterial(index)
      }
    },
    changeTab(index){
      this.selectOrder = {}
      if(index === 1){
        if(!this.isBuy){
          this.isBuy = true
          this.page = 0
          this.getOrders(this.selectItem, this.page)
        }
      }else{
        if(this.isBuy){
          this.isBuy = false
          this.calLeftMaterial(this.selectItem)
          this.getMyOrders()
        }
      }
    },
    showMore(){
      this.page++
      this.getOrders(this.selectItem, this.page)
    },
    calLeftMaterial(index){
      this.selectLeftNum = 0
      let _items = JSON.parse(cache.getSession('bagItems')) || {}
      if (Object.keys(_items).length != 0) {
        for (let item of _items) {
          if (item.item_id == index) {
            this.selectLeftNum = item.num
          }
        }
      }
    },
    showOrderTimeCount(){
      if(this.orderData && this.orderData.length > 0){
        let orderArr = []
        let curTime = Math.floor(Date.now() / 1000)
        for(let i = 0; i < this.orderData.length; i++){
          if(this.orderData[i].status == 0){
            let _time = this.orderData[i].open_at - curTime
            this.orderData[i].countTime = _time
            if(_time > 60){
              this.orderData[i].status == 1
            }
          }
          orderArr.push(this.orderData[i])
        }
       this.orderData = orderArr
      }
    },
    getOrders(itemId, page){
      this.showMoreBtn = false
      this.showLoading = true
      this.selectOrder = {}
      let orderArr = []
      if(page > 0){
        for(let i = 0; i < this.orderData.length; i++){
          orderArr.push(this.orderData[i])
        }
      }else{
        this.orderData = {}
      }
      let data = {item_id : itemId, page_num : page}
      getOrders(data).then(response => {
        console.log("get orders", response)
        if(response){
          let curPageTotal = 20 * (page + 1)
          if(response.order_cnt >= curPageTotal){
            this.showMoreBtn = true
          }
          if(response.orders && response.orders.length > 0){
            for(let i = 0; i < response.orders.length; i++){
              orderArr.push(response.orders[i])
            }
            this.orderData = orderArr
          }else{
            this.showLoading = false
          }
          console.log('orders', this.orderData)
        }
      })
    },
    getMyOrders(){
      let orderArr = []
      this.myOrderData = {}
      getMyOrders({}).then(response => {
        if(response){
          if(response.orders && response.orders.length > 0){
            for(let i = 0; i < response.orders.length; i++){
              orderArr.push(response.orders[i])
            }
            this.myOrderData = orderArr
            this.leftBucket = parseInt(8 - this.myOrderData.length)
            console.log("my orders", this.myOrderData, 'left bucket', this.leftBucket )
          }else{
            this.leftBucket = 8
          }
        }else{
          this.leftBucket = 8
        }
      })
    },
    calFee(){
      if(this.selectItem == 10){
        this.handleFee = (this.formData.price * 0.03).toFixed(4)
      }else{
        this.handleFee = (this.formData.price * 0.01).toFixed(4)
      }
    },
    saleItems(){
      if(this.leftBucket <= 0){
        this.$refs.tip.openTip('Listed items exceed limit')
        return
      }
      if(this.formData){
        if(this.formData.num === undefined){
          this.$refs.tip.openTip('Please enter number')
          return
        }else{
          let num = parseInt(this.formData.num)
          if(isNaN(num) || num <= 0){
            this.$refs.tip.openTip('Please enter correct number')
            return
          }else{
            this.formData.num = num
            if(num > this.selectLeftNum){
              this.$refs.tip.openTip('Exceeded remaining quantity')
              return
            }
          }
        }
        this.handleFee = 0
        if(this.formData.price === undefined){
          this.$refs.tip.openTip('Please enter price')
          return
        }else{
          let price = parseFloat(this.formData.price)
          if(isNaN(price)){
            this.$refs.tip.openTip('Please enter correct price')
            return
          }else{
            price = price.toFixed(4)
            this.formData.price = price
            if(price <= 0.01){
              this.$refs.tip.openTip('Price too low')
              return
            }
          }
        }
        this.calFee()
        this.doSale()
      }
    },
    doSale(){
      let _price = parseInt(this.formData.price * 10000)
      let _num = this.formData.num
      let data = {item_id : this.selectItem, num : _num, price : _price}
      this.listProtect = true
      publish(data).then(response => {
        console.log("sale orders", response)
        if(response){
          this.$refs.tip.openTip('List success')
          this.getMyOrders()
          this.formData.num = null
          this.formData.price = null
          this.handleFee = 0
          let _materials = JSON.parse(cache.getSession('bagItems')) || {}
          if(Object.keys(_materials).length != 0){
            let _left = parseInt(_materials.find(obj => obj.item_id === this.selectItem).num - _num)
            _materials.find(obj => obj.item_id === this.selectItem).num = _left
            cache.setSession('bagItems', _materials)
            this.selectLeftNum = _left
          }
        }
        this.listProtect = false
      })
    },
    selectOnSaleOrder(uuid){
      this.selectOrder = this.orderData.find(obj => obj.order_uuid === uuid)
      if(this.selectOrder && this.selectOrder.status == 1){
        this.onPopBox(2)
      }
    },
    selectMyOrder(uuid){
      this.selectOrder = this.myOrderData.find(obj => obj.order_uuid === uuid)
      if(this.selectOrder && this.selectOrder.status == 1){
        this.onPopBox(1)
      }
    },
    delist(){
      if(this.selectOrder && this.selectOrder.order_uuid){
        this.delistProtect = true
        takeOff({order_uuid : this.selectOrder.order_uuid}).then(response => {
          console.log("delist orders", response)
          if(response){
            let _materials = JSON.parse(cache.getSession('bagItems')) || {}
            if(Object.keys(_materials).length != 0){
              let _left = parseInt(_materials.find(obj => obj.item_id === this.selectOrder.goods.item_id).num + this.selectOrder.goods.num)
              _materials.find(obj => obj.item_id === this.selectOrder.goods.item_id).num = _left
              cache.setSession('bagItems', _materials)
              this.calLeftMaterial(this.selectItem)
            }
            this.getMyOrders()
            this.selectOrder = {}
            this.onPopClose()
            this.$refs.tip.openTip('Delist success')
          }
          this.delistProtect = false
        })
      }
    },
    purchase(){
      if(this.selectOrder && this.selectOrder.order_uuid){
        this.purchaseProtect = true
        purchase({order_uuid : this.selectOrder.order_uuid}).then(response =>{
          if(response){
            if(response === 'ERR_NOT_ENOUGH_MANA'){
              this.$refs.tip.openTip('ZPA not enough')
            }else{
              let _materials = JSON.parse(cache.getSession('bagItems')) || {}
              if(Object.keys(_materials).length != 0){
                let _left = parseInt(_materials.find(obj => obj.item_id === this.selectOrder.goods.item_id).num + this.selectOrder.goods.num)
                _materials.find(obj => obj.item_id === this.selectOrder.goods.item_id).num = _left
                cache.setSession('bagItems', _materials)
                this.calLeftMaterial(this.selectItem)
              }
              let zpa =+ response.mana_left || 0
              cache.setSession('zpa', zpa)
              this.$store.commit("GLOBAL_ZPA", zpa)
            
              let orderArr = []
              for(let obj of this.orderData){
                if(obj.order_uuid != this.selectOrder.order_uuid){
                  orderArr.push(obj)
                }
              }
              this.orderData = orderArr
              if(!this.orderData || this.orderData.length == 0){
                this.showLoading = false
              }
              this.selectOrder = {}
              this.onPopClose()
            }
          }
          this.purchaseProtect = false
        })
      }
    }

  }
}
</script>

<style lang='stylus' src="./index.css" />
