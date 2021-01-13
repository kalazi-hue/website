import Vue from 'vue'
import App from './App'
import VueClipboard from 'vue-clipboard2'
 
Vue.use(VueClipboard)
Vue.config.productionTip = false
Vue.prototype.$host = 'https://unidemo.dcloud.net.cn/';
	
/* #ifdef H5 */
window.Hls = require('hls.js');
/* #endif */

App.mpType = 'app'

const app = new Vue({
  ...App
})
app.$mount()
