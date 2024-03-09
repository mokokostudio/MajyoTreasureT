import Vue from 'vue'
import Router from 'vue-router'
import App from '../App'

Vue.use(Router)

// Router.prototype.go = function() {
//     this.isBack = true
//     window.history.go(-1)
// }

// 路由懒加载
const loading = () => import('@/page/loading/index')
const mint = () => import('@/page/mint/index')
const boss = () => import('@/page/boss/index')
const bossDetail = () => import('@/page/bossDetail/index')
const pvp = () => import('@/page/pvp/index')
const pvpfight = () => import('@/page/pvpfight/index')
const market = () => import('@/page/market/index')

export default new Router({
    // mode: "history",
    // history:createWebHashHistory(),
    mode: "history",
    // base: process.env.BASE_URL,

    routes: [
        {
            path: '',
            redirect: '/loading',
            component: App,
            children: [
                // index页面
                {
                    name: 'mint',
                    path: '/mint',
                    component: mint
                },
                //boss
                {
                    name: 'boss',
                    path: '/boss',
                    component: boss
                },
                {
                    name: 'pk',
                    path: '/pk',
                    component: bossDetail
                },
                {
                    name: 'loading',
                    path: '/loading',
                    component: loading
                },
                {
                    name: 'pvp',
                    path: '/pvp',
                    component: pvp
                },
                {
                    name: 'pvpfight',
                    path: '/pvpfight',
                    component: pvpfight
                },
                {
                    name: 'market',
                    path: '/market',
                    component: market
                }
                // // detail页面
                // {
                //     name: 'detail',
                //     path: '/detail',
                //     component: detail
                // },
                // // search页面
                // {
                //     name: 'search',
                //     path: '/search',
                //     component: search
                // }
            ]
        }
    ]
})
