const getters = {
    size: state => state.app.size,
    account: state => state.user.account,
    userId: state => state.user.userId,
    nickName: state => state.user.nickName,
    lan: state => state.user.lan,
    energy: state => state.user.energy,
    energy_update_at: state => state.user.energy_update_at,
    token: state => state.user.token,
}

const mutations = {
  SET_ENG:(state,name) => {
    getters.energy = name
  }
}

export default getters
