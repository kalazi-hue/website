<template>
  <div class="slider">
    <swiper :options="swiperOption">
      <swiper-slide  v-for="(item,index) in dataList" :key="index">
        <a :href="item.playUrl" target="_blank">
          <img v-lazy="item.url" alt="" />          
        </a>
      </swiper-slide>
      <div class="swiper-pagination" slot="pagination"></div>
     </swiper>
    <div class="swiper-button-prev"></div><!--左箭头。如果放置在swiper-container外面，需要自定义样式。-->
    <div class="swiper-button-next"></div><!--右箭头。如果放置在swiper-container外面，需要自定义样式。-->
    <adData @lunboData="getData" /> 
  </div>
</template>

<script>
import { swiper, swiperSlide } from 'vue-awesome-swiper'
import 'swiper/dist/css/swiper.css'
import adData from '@/components/common/adData'
export default {
  components: {
    swiper,
    swiperSlide,
    adData
  },
  data () {
    const that = this
    return {
      dataList: [],
      swiperOption: {
        autoplay:1500,
        speed: 1000,
        lazy: {
          loadPrevNext: true
        },
        loadOnTransitionStart: true,
        watchSlidesVisibility: true, //解决了swiper前面的一个不能提前加载的bug
        autoplayDisableOnInteraction : false,
        loop:true,
        centeredSlides : true,
        slidesPerView:1.2,
        pagination : '.swiper-pagination',
        paginationClickable:true,
        prevButton:'.swiper-button-prev',
        nextButton:'.swiper-button-next'
      }
    }
  },
  created () {
  },
  methods: {
    getData (data) {
      this.dataList = data
    }
  }
}
</script>
<style scoped lang="scss">
.slider {
  margin: 20px auto;
  position: relative;
  z-index: 2;
  .swiper-container {
    height: 300px;  
    border-radius: 14px;
  }
  .swiper-slide {
    -webkit-transition: transform 1.0s;
    -moz-transition: transform 1.0s;
    -ms-transition: transform 1.0s;
    -o-transition: transform 1.0s;
    -webkit-transform: scale(0.95);
    transform: scale(0.95);
    border-radius: 14px;
    overflow: hidden;
  }
  .swiper-button-prev, .swiper-button-next {
    color: #eb0076 !important;
    background: rgb(255, 255, 255, .5);
    width: 80px;
    height: 80px;
    border-radius: 50%;
    margin-top: -40px;
    opacity: .7;
  }

  .swiper-button-prev:hover, .swiper-button-next:hover {
    background: rgb(255, 255, 255, .9);
    opacity: 1;
  }

  .swiper-slide-active,.swiper-slide-duplicate-active {
    -webkit-transform: scale(1);
    transform: scale(1);
  }

  .swiper-slide a {
    display: block;
    border-radius: 14px;
  }

  .swiper-slide-active a:after {
    background: none;
  }

  .swiper-slide-prev,
  .swiper-slide-next {
    position: relative;
  }
  .swiper-button-prev {
    left: 47px;
  } 
  .swiper-button-next {
    right: 47px;
  }

  .swiper-slide-next:after,
  .swiper-slide-prev:after {
    content: "";
    width: 100px;
    height: 100%;
    background: rgb(255, 255, 255, .5);
    display: block;
    position: absolute;
    top: 0;
  }

  .swiper-slide-prev:after {
    right: 0;
  }

  .swiper-slide-next:after {
    left: 0;
  }

  .swiper-slide img {
    width: 100%;
    height: 100%;
    border-radius: 14px;
    display: block;
    max-height: 300px;
  }

  .swiper-pagination {
    bottom: 10px;
  }

}
</style>

<style>
  .swiper-pagination-bullet {
    background: #fff !important;
    margin-left: 4px;
    margin-right: 4px;
    width: 8px;
    height: 8px;
    opacity: 1;
    transition: all ease .5s;
  }

  .swiper-pagination-bullet-active {
    width: 20px;
    height: 8px;
    background: #FFF;
    border-radius: 30px !important;
    background: #ce35e4 !important;
  }
</style>