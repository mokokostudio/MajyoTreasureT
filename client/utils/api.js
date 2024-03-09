
import { request } from '@/utils/request'
const BASE_URL = process.env.VUE_APP_Server

export default {
  // 获取配置列表
  getConfig: (params) => {
    return request(`${BASE_URL}/xcx/share/share.json`, `GET`, params);
  }

  // 获取钱包流水
  //   getUserWalletList: (params) => {
  //     return request(`${BASE_URL}/admin/getUserWalletList`, {
  //       method: 'GET',
  //       params: params,
  //     });
  //   }
}



// request('post', 'classID')