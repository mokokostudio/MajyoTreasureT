import request from '@/utils/requestApi'

export function getItems(data) {
    return request({
        url: '/GetItems',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function upgradeBaseEquipStar(data) {
    return request({
        url: '/UpgradeBaseEquipStar',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function upgradeBaseEquipLevel(data) {
    return request({
        url: '/UpgradeBaseEquipLevel',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

