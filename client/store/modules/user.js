import cache from '@/utils/cache'
import { login } from '@/api/login'

export default {
    namespaced: true,   //命名空间
    state: {
        account : '',
        userId: '',
        nickName: 'userPlayer',
        lan: '',
        energy: 0,
        energy_update_at: 0,
        token: '',
        defeatHistory: null,
        bossNode: [
          {class:0,bossId:0},
          {class:1,bossId:0},
          {class:2,bossId:0},
          {class:3,bossId:0},
          {class:4,bossId:0},
          {class:5,bossId:0}
        ]
    },
    mutations: {
        SET_ACCOUNT: (state, account) => {
            cache.setSession('account', account)
            state.account = account
        },
        SET_USERID: (state, userId) => {
            cache.setSession('userId', userId)
            state.userId = userId
        },
        SET_NICKNAME: (state, nickName) => {
            cache.setSession('nickName', nickName)
            state.nickName = nickName
        },
        SET_LAN: (state, lan) => {
            cache.setSession('lan', lan)
            state.lan = lan
        },
        SET_ENERGY: (state, energy) => {
            cache.setSession('energy', energy)
            state.energy = energy
        },
        SET_ENERGYUPAT: (state, energy_update_at) => {
            cache.setSession('energy_update_at', energy_update_at)
            state.energy_update_at = energy_update_at
        },
        SET_TOKEN: (state, token) => {
            cache.setSession('token', token)
            state.token = token
        },
        SET_DEFEAT_HISTORY: (state, defeatHistory) => {
          if(defeatHistory && defeatHistory.list && defeatHistory.list.length > 0){
            cache.setSession('defeatHistory', 1);
            state.defeatHistory = 1;
          }else{
            cache.setSession('defeatHistory', 0);
            state.defeatHistory = 0;
          }
        },
        SET_BOSS_HISTORY_DETAIL: (state, defeatHistory) => {
          if(defeatHistory && defeatHistory.list && defeatHistory.list.length > 0){
            console.log("初始化进度",defeatHistory.list.length)
            for(let hisObj of defeatHistory.list){
              if(!hisObj.boss_class || hisObj.boss_class == ''){
                state.bossNode.find(obj => obj.class === 0).bossId = hisObj.boss_id
              }else{
                state.bossNode.find(obj => obj.class === hisObj.boss_class).bossId = hisObj.boss_id
              }
            }
          }
          cache.setSession('bossNode', state.bossNode);
        },
        SET_BUFF_CARD: (state, buffCard) => {
          if(buffCard && buffCard.length > 0){
            cache.setSession('buffCard', buffCard);
          }else{
            cache.removeSession('buffCard');
          }
        },
        SET_BUFF_CARD_STATUS: (state, buffCardStatus) => {
          if(buffCardStatus){
            cache.setSession('buffCardStatus', buffCardStatus);
          }else{
            cache.removeSession('buffCardStatus');
          }
        }
    },
    actions: {
        login({ commit }, data) {
            // console.log(data)
            // data = {"user_id": 10000008}
            return new Promise((resolve, reject) => {
                login(data).then(response => {
                    console.log('response', response)
                    const { account } = response
                    commit('SET_ACCOUNT', account.account)
                    commit('SET_USERID', account.user_id)
                    commit('SET_NICKNAME', account.nickname)
                    commit('SET_LAN', account.lan)
                    commit('SET_ENERGY', response.energy)
                    commit('SET_ENERGYUPAT', response.energy_update_at)
                    commit('SET_TOKEN', response.token)
                    commit('SET_BOSS_HISTORY_DETAIL', response.boss_defeat_history)
                    commit('SET_DEFEAT_HISTORY', response.boss_defeat_history)
                    commit('SET_BUFF_CARD', response.buff_cards)
                    commit('SET_BUFF_CARD_STATUS', response.buff_card_status)
                    // setToken(response.token)
                    resolve(response)
                }).catch(error => {
                  console.log(error)
                  reject(error)
                })
            })
        },
    }
}
