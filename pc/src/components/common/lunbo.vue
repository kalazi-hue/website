<template>
  <div class="container page-home">
    <div class="banner">
        <div class="swiper-wrapper">
            <swiper :options="swiperOption" ref="mySwiper">
                <swiper-slide v-for="(item,i) in dataList" :key="i" v-if="item.status">
                    <div class="cover">
                        <img :src="item.url" :alt="item.description" @click="openLunbo(item.playUrl)">
                    </div>
                </swiper-slide>
                <div class="swiper-pagination"  slot="pagination"></div>
            </swiper>
        </div>
    </div>
    <adData @lunboData="getData" />
  </div>
</template>

<script>
require('swiper/dist/css/swiper.css');
import { swiper, swiperSlide } from 'vue-awesome-swiper'
import adData from '@/components/common/adData'
export default {
  components: {
    swiper,
    swiperSlide,
    adData
  },
  computed: {
      swiper() {
          return this.$refs.mySwiper.swiper
      }
  },
  data () {
    return {
      dataList: [],
      swiperOption: {
        loop:true,
        mousewheelControl: true,// 禁止鼠标滚轮切换
        initialSlide :1,
        lazy: {
            loadPrevNext: true,
        },
        autoplay:1200,
        speed:1500,
        autoplayDisableOnInteraction : false,
        pagination : '.swiper-pagination',
        paginationClickable:true,
        watchSlidesProgress:true,
        centeredSlides: true, //设定为true时，活动块会居中，而不是默认状态下的居左。
        spaceBetween:20,
        slidesPerView: 1.2,
        loopedSlides :2,
        observer: true,
        observeParents: true  //等待过渡完毕。自动切换会在slide过渡完毕后才开始计时。
      }
    }
  },
  methods: {
    getData (data) {
      this.dataList = data
    },
    openLunbo (url) {
      window.open(url)
    }
  }
}
</script>

<style lang="less" rel="stylesheet/less">
.banner{
    padding: 20px 0;
    .swiper-slide{
      border-radius: 15px;
      .cover{
        position: relative;
        height: 300px;
        img{
          width: 100%;
          height: 100%;
          border-radius: 15px;
          cursor: pointer;
        }
      }
    }
    .swiper-slide-active {
    }
    .swiper-pagination{ //设置分页器的大小
      .swiper-pagination {
        bottom: 16px;
        right: 15px;
        left: auto;
        width: auto!important;
      }
      .swiper-pagination-bullet{
        width: 8px;
        height: 8px;
        transition: all ease .5s;
      }
      .swiper-pagination-bullet-active{
        background: #fff;
        width: 20px;
        border-radius: 15px;
      }
    }
  }
</style>