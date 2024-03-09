// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import 'babel-polyfill'
import Vue from 'vue'
import App from './App'
import router from './router'
import store from './store'

// css
import '@/assets/css/reset.css'
import '@/assets/css/icon.css'
import '@/assets/css/transition.css'
import '@/assets/css/global.css'

// 第三方库
// import MintUI, { MessageBox, Toast, Indicator } from 'mint-ui'     // 饿了么移动端UI组件
// import 'mint-ui/lib/style.css'
import VueAwesomeSwiper from 'vue-awesome-swiper'       // swiper滑动组件
import 'swiper/dist/css/swiper.css'

// 自定义组件
// import listItem from '@/components/listItem'            // 文章列表组件
import Tip from '@/components/tips'

import '@/directives'                                   // 指令

Vue.config.productionTip = false
// Vue.use(MintUI)
// Vue.prototype.$toast = Toast
// Vue.prototype.$msgBox = MessageBox
// Vue.prototype.$indicator = Indicator
Vue.use(VueAwesomeSwiper)

// 注册全局组件
// Vue.component('list-item', listItem)
const components = {Tip}

Object.keys(components).forEach(key => {
  Vue.component(key, components[key])
})

/* eslint-disable no-new */
new Vue({
    el: '#app',
    router,
    store,
    template: '<App/>',
    components: { App }
})
