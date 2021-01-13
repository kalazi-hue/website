<template>
  <div>
    <cpHeader />
    <div class="main" :class="{'pb200': fixFooter}">
      <lunbo /> <!-- 轮播图组件 -->

      <adTop /> <!-- 顶部横幅组件 -->
      
      <keep-alive>
        <router-view class="main"></router-view>
      </keep-alive>

      <adBottom /> <!-- 底部横幅组件 -->

    </div>

    <adLeftRight /> <!-- 两边广告组件 -->

    <cpFooter @closeOrder="msgGot" />

    <div class="backTop" @click="backtop" v-if="isShowBtn">
      <i class="el-icon-arrow-up"></i>
    </div>
    
  </div>
</template>

<script>
import cpHeader from '@/components/Header'
import cpFooter from '@/components/Footer'
import lunbo from '@/components/common/lunbo'
import adTop from '@/components/common/adTop'
import adBottom from '@/components/common/adBottom'
import adLeftRight from '@/components/common/adLeftRight'
export default {
  components: {
    cpHeader,
    cpFooter,
    lunbo,
    adTop,
    adBottom,
    adLeftRight
  },
  data () {
    return {
      scrollTop: 0, // 滚动条的高度
      isShowBtn: false,
      fixFooter: true
    }
  },
  mounted() {
    window.addEventListener('scroll', this.handleScroll, true)
  },
  created () {
  },
  destroyed () {
    window.removeEventListener('scroll', this.handleScroll, true)
  },
  methods: {
    msgGot (msg) {
      msg ? this.fixFooter = false : this.fixFooter = true
    },
    sortId(a, b){  
      return a.sort - b.sort  
    },
    handleScroll () { // 保存滚动值，这是兼容的写法
      this.scrollTop = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop
      const windowHeight = document.documentElement.clientHeight || document.body.clientHeight
      const scrollHeight = document.documentElement.scrollHeight||document.body.scrollHeight
      // console.log(this.scrollTop)
    },
    backtop () {
      const that = this
      let timer = setInterval(() => {
        let ispeed = Math.floor(-that.scrollTop / 5)
        document.documentElement.scrollTop = document.body.scrollTop = that.scrollTop + ispeed
        if (that.scrollTop === 0) {
          clearInterval(timer)
        }
      }, 16)
    }
  }
}
</script>
<style lang="scss">
@import './assets/css/common.css';
body {
    padding-top: 80px;
    width: 100%;
    min-height: 1000px;
    background: url('./assets/images/bg-body.jpg') center top no-repeat;
    background-color: #333;
    background-position: 0 0, 0 100%;
    background-size: 100%;
}

.main {
  width: 100%;
  max-width: 1200px;
  margin: 0 auto;  
  transition: all ease .5s;
}
.isShowFooter .main,
.isShowFooter {
    max-width: 100%!important;
}
.backTop {
  position: fixed;
  right: 50px;
  bottom: 50px;
  z-index: 10;
  background-color: rgba(0,0,0,.6);
  color: #fff;    
  width: 50px;
  height: 50px;
  line-height: 50px;
  text-align: center;
  border-radius: 4px;
  font-size: 18px;
  cursor: pointer;
}

.pb200 {
  padding-bottom: 200px;
}
</style>
