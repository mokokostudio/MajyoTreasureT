import request from '@/utils/requestApi'

export function login(data) {
    console.log('======check', process.env.NODE_ENV)
    if (process.env.NODE_ENV === 'development') {
        return request({
            url: '/LoginTest',
            method: 'POST',
            data: JSON.stringify({"user_id": 10000008})
        })
    } else {
        return request({
            url: '/LoginByToken',
            method: 'POST',
            data: JSON.stringify(data)
        })
    }
}

// export function login(data) {
//     return request({
//         url: '/LoginByToken',
//         method: 'POST',
//         data: JSON.stringify(data)
//     })
// }
