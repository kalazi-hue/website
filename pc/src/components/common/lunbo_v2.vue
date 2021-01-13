<template>
  <div class="slider">
      <div class="swiper-container">
          <div class="swiper-wrapper">
              <div class="swiper-slide swiper-slide-center none-effect"><a href="#"><img src="http://www.1990tu.com/i/20201206105845mzv.png" ></a></div>
              <div class="swiper-slide"><a href="#"><img src="http://www.1990tu.com/i/20201206105926tm7.png" ></a></div>
              <div class="swiper-slide"><a href="#"><img src="http://www.1990tu.com/i/20201206105904pyl.png" ></a></div>
          </div>
          
      </div>
      <div class="swiper-pagination"></div>
      <div>
        <div class="swiper-button-prev"><i class="el-icon-arrow-left"></i></div>
        <div class="swiper-button-next"><i class="el-icon-arrow-right"></i></div>
      </div>
    <adData @lunboData="getData" /> 
  </div>
</template>

<script>
import '../../assets/js/swiper/swiper.min.js'
import '../../assets/js/swiper/swiper.min.css'
import adData from '@/components/common/adData'
export default {
  components: {
    adData
  },
  data () {
    const that = this
    return {
      dataList: []
    }
  },
  mounted () {
    var swiper = new Swiper('.swiper-container',{
      // autoplay:1500,
      speed:1000,
      autoplayDisableOnInteraction : false,
      loop:true,
      centeredSlides : true,
      slidesPerView:1.2,
      // spaceBetween: 30,
      pagination : '.swiper-pagination',
      paginationClickable:true,
      prevButton:'.swiper-button-prev',
      nextButton:'.swiper-button-next'
    });
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
}
</style>

<style>

.swiper-container {
  width: 100%;
  height: 300px;
}
 
.swiper-slide {
  -webkit-transition: transform 1.0s;
  -moz-transition: transform 1.0s;
  -ms-transition: transform 1.0s;
  -o-transition: transform 1.0s;
  -webkit-transform: scale(0.95);
  transform: scale(0.95);
  background: url('../../assets/images/banner-default.jpg') no-repeat;
}
 
.swiper-slide-active,.swiper-slide-duplicate-active {
  -webkit-transform: scale(1);
  transform: scale(1);
}
 

.none-effect {
  -webkit-transition: none;
  -moz-transition: none;
  -ms-transition: none;
  -o-transition: none;
}

.swiper-slide a {
  display: block;
  border-radius: 14px;
}
 
.swiper-slide-active a:after {
  background: none;
}
 

.swiper-slide img {
  width: 100%;
  height: 100%;
  border-radius: 14px;
  display: block;
}
.swiper-pagination {
  position: relative;
  margin-bottom: 30px;
}

.swiper-pagination-bullet {
  background: #00a0e9;
  margin-left: 4px;
  margin-right: 4px;
  width: 17px;
  height: 17px;
  opacity: 1;
  margin-bottom: 4px;
}

.swiper-pagination-bullet-active {
  width: 13px;
  height: 13px;
  background: #FFF;
  border: 6px solid #00a0e9;
  margin-bottom: 0;
}

.slider .swiper-button-next,
.slider .swiper-button-prev {
    opacity: .7;
    position: absolute;
    top: 50%;
    width: 60px;
    height: 60px;
    line-height: 60px;
    margin-top: -40px;
    text-align: center;
    display: block;
    border-radius: 50%;
    background: rgb(255, 255, 255, .5);
    background-image: none!important;
}
.slider .swiper-button-next i,
.slider .swiper-button-prev i {
    font-size: 60px;
    color: #cd0067;
}

.slider .swiper-button-prev:hover, 
.slider .swiper-button-next:hover {
    background: rgb(255, 255, 255, .9);
    opacity: 1;
  }
.slider .swiper-button-prev {
  left: 55px;
} 
.slider .swiper-button-next {
  right: 55px;
}

.slider .swiper-pagination {
  position: relative;
  bottom: 30px;
}
.slider .swiper-pagination-bullet {
  background: #fff;
  margin-left: 4px;
  margin-right: 4px;
  width: 8px;
  height: 8px;
  opacity: 1;
  transition: all ease .5s;
}

.slider .swiper-pagination-bullet-active {
  width: 30px;
  height: 8px;
  background: #FFF;
  border-radius: 30px;
  background: #ce35e4;
}
</style>