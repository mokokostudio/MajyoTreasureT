import Vue from 'vue'
import axios from 'axios'
import qs from 'qs'

Vue.prototype.$http = axios

export var request = async( url = '',method = 'POST', data = {}) => {
    let result
    method = method.toUpperCase()
    
    if (method === 'GET') {
        await axios.get(url, { params: data })
        .then(res => {
            result = res.data
        })
    } else if (method === 'POST') {
        await axios.post(url, qs.stringify(data))
        .then(res => {
            result = res.data
        })
    }
    return result
}

export var requestApi = async( url = '',method = 'POST', data = {}) => {
    let result
    method = method.toUpperCase()
    // console.log(1111111, url, method, data)
    // url = 'WWW.BAIDU.COM:8888' + url
    url = process.env.VUE_APP_BASE_API + url
    
    if (method === 'GET') {
        await axios.get(url, { params: data })
        .then(res => {
            result = res.data
        })
    } else if (method === 'POST') {
        await axios.post(url, data)
        .then(res => {
            result = res.data
        })
    }
    return result
}
