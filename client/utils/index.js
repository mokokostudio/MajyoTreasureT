export default {
  //Query url value
  quv(name) {
    let url = location.href; // 获取完整的URL
    let theRequest = {};
    if (url.indexOf('?') != -1) {
      let str = url.substr(url.indexOf('?') + 1); // 获取'?'后面的字符串
      let strs = str.split('&'); // 将查询参数按'&'进行分割
      for (let i = 0; i < strs.length; i++) {
        // 将每个参数名和参数值存储到对象中
        theRequest[strs[i].split('=')[0]] = unescape(strs[i].split('=')[1]);
      }
    }
    return theRequest
  },

  localStorageSet(name, data) {
    localStorage.setItem(name, JSON.stringify(data))
  },
  localStorageGet(key) {
    return JSON.parse(localStorage.getItem(key)) || ''
  },
  showTooltip(show) {
    console.log(123321)
  }
}