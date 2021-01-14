// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router/router'
import VueLazyload from 'vue-lazyload'
/* eslint-disable */
import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
import('babel-polyfill') // 兼容低版本IE
import waterfall from 'vue-waterfall2' 
import VueClipboard  from 'vue-clipboard2'
import store from './store'
import VueAwesomeSwiper from 'vue-awesome-swiper'
import 'swiper/dist/css/swiper.css'

Vue.use(VueAwesomeSwiper)

Vue.use(VueClipboard)
Vue.use(waterfall)
Vue.prototype.$store = store

Vue.use(ElementUI)
Vue.prototype.baseURL = process.env.API_ROOT
Vue.use(VueLazyload, {
  error: require('./assets/images/loading.gif'),
  loading: require('./assets/images/loading.gif'),
  preLoad: 1, // 预加载高度比例
})

new Vue({
  el: "#app",
  router: router,
  render: h => {
    return h(App);
  }
});

router.beforeEach((to, from, next) => {
  if (to.path !== from.path) {
    document.body.scrollTop = 0
    document.documentElement.scrollTop = 0
    window.pageYOffset = 0    
  }
  next()
})

// 百度统计
router.afterEach( ( to, from, next ) => {
  setTimeout(()=>{
    var _hmt = _hmt || [];
    (function() {
        document.getElementById('baidu_tj') && document.getElementById('baidu_tj').remove();
        var hm = document.createElement("script");
        hm.src = "https://hm.baidu.com/hm.js?1111111111111111";
        hm.id = "baidu_tj"
        var s = document.getElementsByTagName("script")[0];
        s.parentNode.insertBefore(hm, s);
    })();
  },0);
});