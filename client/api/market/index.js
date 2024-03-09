import request from '@/utils/requestApi'

export function getOrders(data) {
    return request({
        url: '/GetGoodsOrdersOnSell',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function getMyOrders(data) {
    return request({
        url: '/GetMyGoodsOrdersOnSell',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function publish(data) {
    return request({
        url: '/PublishGoodsOrder',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function purchase(data) {
    return request({
        url: '/PurchaseGoodsOrder',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function takeOff(data) {
    return request({
        url: '/TakeOffGoodsOrder',
        method: 'POST',
        data: JSON.stringify(data)
    })
}
