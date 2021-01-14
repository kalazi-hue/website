import Vue from 'vue'
import Router from 'vue-router'

const index = (resolve) => import('@/pages/index').then((module) => {
  resolve(module)
})
const videoList = (resolve) => import('@/pages/movie/index').then((module) => {
  resolve(module)
})
const videoDetail = (resolve) => import('@/pages/movie/detail').then((module) => {
  resolve(module)
})
const videoSearch = (resolve) => import('@/pages/movie/search').then((module) => {
  resolve(module)
})
const novelList = (resolve) => import('@/pages/novel/list').then((module) => {
  resolve(module)
})
const novelDetail = (resolve) => import('@/pages/novel/detail').then((module) => {
  resolve(module)
})
const picList = (resolve) => import('@/pages/pic/index').then((module) => {
  resolve(module)
})
const picDetail = (resolve) => import('@/pages/pic/detail').then((module) => {
  resolve(module)
})
const liveList = (resolve) => import('@/pages/live/list').then((module) => {
  resolve(module)
})
const liveDetail = (resolve) => import('@/pages/live/detail').then((module) => {
  resolve(module)
})
const helpApp = (resolve) => import('@/pages/help/app').then((module) => {
  resolve(module)
})
const helpContact = (resolve) => import('@/pages/help/contact').then((module) => {
  resolve(module)
})
const helpFaq = (resolve) => import('@/pages/help/faq').then((module) => {
  resolve(module)
})

Vue.use(Router)

// 解决重复点击路由会出现如下报错
const originalPush = Router.prototype.push
Router.prototype.push = function push (location) {
  return originalPush.call(this, location).catch(err => err)
}


const router = new Router({
  // mode: 'history',
  routes: [
    { path: '*', redirect: '/pages/index' },
    { path: '/pages/index', name: 'index',component: index },
    { path: '/pages/movie/index', component: videoList },
    { path: '/pages/movie/detail', component: videoDetail },
    { path: '/pages/movie/search', component: videoSearch },
    { path: '/pages/pic/index', component: picList },
    { path: '/pages/pic/detail', component: picDetail },
    { path: '/pages/novel/list', component: novelList },
    { path: '/pages/novel/detail', component: novelDetail },
    { path: '/pages/live/list', component: liveList },
    { path: '/pages/live/detail', component: liveDetail },
    { path: '/pages/help/app', component: helpApp },
    { path: '/pages/help/contact', component: helpContact },
    { path: '/pages/help/faq', component: helpFaq }
  ]
})

router.beforeEach((to, from, next) => {
  if (to.meta.title) {
    document.title = to.meta.title
  }
  setTimeout(()=>{
    var _hmt = _hmt || [];
    (function() {
      var hm = document.createElement("script");
      hm.src = "";
      var s = document.getElementsByTagName("script")[0]; 
      s.parentNode.insertBefore(hm, s);
    })();
  },0)
  next()
})

export default router
