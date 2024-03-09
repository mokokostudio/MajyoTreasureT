import request from '@/utils/requestApi'

// export function login(data) {
//     return request({
//         url: '/LoginTest',
//         method: 'POST',
//         data: JSON.stringify(data)
//     })
// }

export function mint(data) {
    return request(
        '/testMint',
        'POST',
        data
    )
}
