import request from '@/utils/requestApi'

export function fight(data) {
    return request({
        url: '/Fight',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function getShareHBossLink(data) {
    return request({
        url: '/GenerateShareHBossLink',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

// 刷BUFF
export function getRandomBuff(data) {
    return request({
        url: '/RandomBuffCards',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function choseBuff(data) {
    return request({
        url: '/ChoseBuffCard',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

// PVP业务
export function getRankList(data) {
    return request({
        url: '/GetPVPRanks',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function getPVPInfo(data) {
    return request({
        url: '/GetPVPInfo',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function getChallengeTarget(data) {
    return request({
        url: '/GetPVPChallengeTargets',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function pvpFight(data) {
    return request({
        url: '/FightPVP',
        method: 'POST',
        data: JSON.stringify(data)
    })
}

export function pvpHis(data) {
    return request({
        url: '/GetPVPHistory',
        method: 'POST',
        data: JSON.stringify(data)
    })
}
