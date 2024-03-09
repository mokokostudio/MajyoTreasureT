import Vue from 'vue'
// 引入axios
import axios from "axios"
import cache from '@/utils/cache'

Vue.prototype.$http = axios

const service = axios.create({
    baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
    // withCredentials: true, // send cookies when cross-domain requests
    timeout: 300000, // request timeout
    headers: {
        // "X-Custom-Header": "foobar"
        "Content-Type": "application/json"
    }
})
//   3.添加请求拦截器
// 请求拦截器
service.interceptors.request.use(
    (config) => {
        // 在发送请求之前做些什么.
        let token = cache.getSession('token') || ''
        config.headers['Authorization'] = 'Bearer ' + token
        return config
    },
    (error) => {
        // 对请求错误做些什么
        console.log('request-=======', error)
        return Promise.reject(error)
    }
);
// 4.添加响应拦截器
// 添加响应拦截器
service.interceptors.response.use(
    (response) => {
        // 2xx 范围内的状态码都会触发该函数。
        // 对响应数据做点什么
        // 手动清除 Toast
        return response.data.data // java版, 重要 重要：需调整部分pvp处的j版代码
        //return response.data
    },
    (error) => {
        // 超出 2xx 范围的状态码都会触发该函数。
        // 对响应错误做点什么
        // console.log('response==========', error.response.data.length, this)
        // let err_code = error.response.data.trim()
        // if(err_code === 'ERR_NOT_ENOUGH_MANA'){
        //   return err_code
        // }
        return Promise.reject(error)
    }
)
// 5.向外抛出实例对象
export default service
