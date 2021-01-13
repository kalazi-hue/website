<template>
  <div class="slider">
    <div class="swiper-container" >
      <div class="swiper-wrapper" >
        <div class="swiper-slide" ><img src="http://www.1990tu.com/i/20201206105926tm7.png" /></div>
        <div class="swiper-slide" ><img src="http://www.1990tu.com/i/20201206105904pyl.png" /></div>
        <div class="swiper-slide" ><img src="http://www.1990tu.com/i/20201206105845mzv.png" /></div>
      </div>
    </div>
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
  data(){
    return {
      }
  },
  mounted(){
    this.$nextTick(()=>{
      new Swiper('.swiper-container', {
        loop: true, // 循环模式选项
        loopedSlides: 3,
      })
    })
  }
}
</script>

<style lang="scss" >
.slider {
  margin: 20px auto;
  position: relative;
  z-index: 2;
  overflow: hidden;
  .swiper-container {
    overflow: visible!important;
    width: 1000px;
    height: 300px;
    border-radius: 15px;
      .swiper-slide {
      width: 1000px;
      border-radius: 15px;      
    }
    img {
      width: 100%;
      height: 100%;
      border-radius: 15px;
    }
  }
  .swiper-container .swiper-wrapper .swiper-slide-prev,
  .swiper-container .swiper-wrapper .swiper-slide-next{margin-top: 10px; height: 280px!important;}
  .swiper-container .swiper-wrapper .swiper-slide-next img{ height: 280px!important;}
  .swiper-container .swiper-wrapper .swiper-slide-prev img{ height: 280px!important;}
  .swiper-container .swiper-wrapper .swiper-slide-active{ width: 1000px;margin-right: 20px;margin-left: 20px}
}
</style>