<template>
  <div class="home-list">
    <div class="title-video">
      <h2>{{ title }}</h2> 
      <ul class="list-title">
        <li v-for="(list, index) in videoList.slice(10)" @click="goPlayer(list.ID)" :key="index" v-if="list.status">
            {{ list.title }}
        </li>
      </ul>
      <span @click="goPath(path)" class="more">更多<i class="el-icon-d-arrow-right"></i></span>
    </div>

    <div class="list-box">
      <div class="list-swiper">
        <swiper :options="swiperOption" class="swiper-wrap"  ref="mySwiper">
          <swiper-slide v-for="(list, index) in videoListSwiper" :key="index" v-if="list.status">
            <div class="img" @click="goPlayer(list.ID)">
              <div class="overlay"><i class="ico-play"></i></div>
              <span class="playtime">{{ list.play_time }}</span>
              <img v-lazy="list.cover" alt="">
              <span class="title">{{ list.title }}</span>
            </div>
          </swiper-slide>
          <div class="swiper-pagination" slot="pagination"></div>
        </swiper>
      </div>

      <ul class="list">
        <li v-for="(list, index) in videoList" v-show="index > 4 && index < 11" :key="index" v-if="list.status">
          <div class="img" @click="goPlayer(list.ID)">
              <div class="overlay"><i class="ico-play"></i></div>
              <span class="playtime">{{ list.play_time }}</span>
              <img v-lazy="list.cover" alt="">
              <span class="title">{{ list.title }}</span>
            </div>
        </li>

        <!-- 数据加载前骨架占位 -->
        <li v-for="(list, index) in 6" v-if="videoList.length <= 0" :key="index">
          <div class="img placeholder"></div>
        </li>

      </ul>      
    </div><!-- list-box end -->
  </div>
</template>

<script>
import { API } from '@/api';
import { swiper, swiperSlide } from 'vue-awesome-swiper'
import 'swiper/dist/css/swiper.css'
export default {
  components: {
    swiper,
    swiperSlide
  },
	props: {
    showLatestList: Boolean,
    showRecommendList: Boolean
	},
	data () {
		return {
      title: '今日更新',
      path: '',
      videoList: [],
      videoListSwiper: [],
      swiperOption: {
        autoplay:1200,
        speed:1000,
        autoplayDisableOnInteraction : false,
        loop:true,
        centeredSlides : true,
        slidesPerView:1,
        pagination : '.swiper-pagination',
        paginationClickable:true
      }
		}
	},
	created () {
    this.showLatestList ? this.getLatestMovieList('getLatestMovieList') : ''
    this.showRecommendList ? this.getRecommendMovieList('getRecommendMovieList') : ''
    this.showRecommendList ? this.title = '热门推荐' : this.title = '今日更新'
    this.showRecommendList ? this.path = `/pages/movie/index?typeid=${1}` : this.path = `/pages/movie/index?typeid=${2}`
	},
  methods: {
    getLatestMovieList (type) {
      let params = {
        page: 1,
        pageSize: 20
      }
      API.getLatestMovieList(params).then(res => {
        let videoList2 = JSON.parse(JSON.stringify(res.list)); 
        this.videoList = res.list
        this.videoListSwiper = videoList2.splice(0, 4)
      })
    },
    getRecommendMovieList (type) {
      let params = {
        page: 1,
        pageSize: 20
      }
      API.getRecommendMovieList(params).then(res => {
        let videoList2 = JSON.parse(JSON.stringify(res.list)); 
        this.videoList = res.list
        this.videoListSwiper = videoList2.splice(0, 4)
      })
    },
    goPath (path) {
      this.$router.push({
        path: this.path
      })
    },
    goPlayer (vid) {
      this.$router.push({
        path: `/pages/movie/detail?videoId=${vid}`
      })
    },
    backtop () {
      document.documentElement.scrollTop = document.body.scrollTop = 0
    }
  }
}
</script>

<style scoped lang="scss">
.home-list {
  margin-bottom: 20px;
  &:nth-child(1) {
    h2 {
      background: url('../assets/images/icons/bg-update.png') no-repeat center bottom;     
    }
  }
  &:nth-child(2) {
    h2 {
      background: url('../assets/images/icons/bg-hot.png') no-repeat center bottom;     
    }
  }
  &:nth-child(3) {
    h2 {
      background: url('../assets/images/icons/bg-photo.png') no-repeat center bottom;     
    }
  }
  .title-video {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    h2 {
      width: 10%;
      padding-bottom: 10px;
    }
    .more {
      display: flex;
      align-items: center;
      justify-content: flex-end;
      width: 5%;
      padding: 0 0 0 10px;
      cursor: pointer;
        &:hover {
          color: #f90;
        }
    }
    .list-title {
      width: 80%;
      display: flex;
      align-items: flex-end;
      white-space: nowrap;
      overflow: hidden;
      text-overflow: ellipsis;
      flex-wrap: nowrap;
      border-bottom: 1px solid #8e8383;
      padding-bottom: 10px;
      position: relative;
      li {
        padding: 0 10px;
        font-size: 15px;
        border-right: 2px solid #acacac;
        cursor: pointer;
        &:last-child {
          border-right: none;
        }
        &:hover {
          color: #fff!important;
        }
      }
    }
  }
}
.home-list {
   h2 {
    font-weight: normal;
    a {
      text-decoration: none;
      font-size: 30px;
      color: #fff;
    }
    .list-title {
    }
  }
  .list-box {
    display: flex;
    margin-top: 20px;
    img {
      width: 100%;
      height: 100%;
      border-radius: 10px;
    }
    .swiper-container {
      width: 480px;
      height: 314px;
      margin: 0;
      border-radius: 10px;
      background-image: linear-gradient(90deg, #222222 25%, #151212 37%, #222222 63%);
      .title {
        width: 100%;
        height: 75px !important;
        font-size: 20px!important;
        padding: 30px 10px 0 15px !important;
      }
      .playtime {
        width: 100%;
        height: 75px !important;
        font-size: 16px!important;
        padding: 15px 20px 0 15px !important;
      }
      .swiper-pagination {
        bottom: 16px;
        right: 15px;
        left: auto;
        width: auto!important;
      }
    }
    .list {
      flex-wrap: wrap;
      display: flex;
      width: 720px;
      margin-bottom: -15px;
      li {
        border-radius: 10px;
        width: 31%;
        margin: 0 0 15px 15px;
      }
    }
    .img {
      position: relative;
      overflow: hidden;
      border-radius: 10px;
      height: 100%;
      .overlay {
        text-indent: 15px;
        position: absolute;
        top: 0;
        left: 0;
        cursor: pointer;
        z-index: 1;
        display: flex;
        align-items: center;
        justify-content: center;
        height: 98%;
        width: 100%;
        opacity: 0;
        filter: alpha(opacity=0);
        color: #fff;
        -webkit-backface-visibility: hidden;
        backface-visibility: hidden;
        transition: all .25s ease-out;
        opacity: 0;
        filter: alpha(opacity=0);
        &:hover {
          background: rgba(0,0,0,.2);
          filter: progid:DXImageTransform.Microsoft.gradient(startColorstr=#33000000,endColorstr=#33000000);
          opacity: 1;
          filter: alpha(opacity=100);
          transform: scale(1);
          text-align: center;
          i {
            transform: scale(1);  
            opacity: 1;
            filter: alpha(opacity=100);
            transform: scale(1);
          }
        }
        i  {
          transform: scale(2);
          -webkit-backface-visibility: hidden;
          backface-visibility: hidden;
          transition: all .25s ease-out;
          opacity: 0;
          filter: alpha(opacity=0);
        }
        .el-icon-video-play {
          font-size: 50px;
          font-weight: normal;
          margin-top: -20px;
        }
      }
      .playtime {
        position: absolute;
        width: 100%;
        top: 0;
        right: 0;
        color: #fff;
        font-size: 12px;
        padding: 5px 15px 5px 5px;
        height: 35px;
        background: url(../assets/images/shadow-top.png) center top no-repeat;
        background-size: 100%;
        text-align: right;
      }
      .title {
        text-align: left;
        display: block;
        width: 100%;
        position: absolute;
        bottom: 0;
        left: 0;
        padding: 10px 10px 0 10px;
        height: 35px;
        background: url(../assets/images/shadow-bottom.png) center bottom no-repeat;
        background-size: 100%;
        font-size: 12px;
        color: #fff;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
      }
    }
  }
}


.home-list:nth-child(1) .list-title li:nth-child(6),
.home-list:nth-child(2) .list-title li:nth-child(5),
.home-list:nth-child(3) .list-title li:nth-child(1) {
    color: #00fda2;
 }
.home-list:nth-child(1) .list-title li:nth-child(4),
.home-list:nth-child(2) .list-title li:nth-child(3),
.home-list:nth-child(3) .list-title li:nth-child(2) {
    color: #9f41c9;
 }
.home-list:nth-child(1) .list-title li:nth-child(2),
.home-list:nth-child(2) .list-title li:nth-child(1),
.home-list:nth-child(3) .list-title li:nth-child(3) {
    color: #1dd2ff;
 }
.home-list:nth-child(1) .list-title li:nth-child(5),
.home-list:nth-child(2) .list-title li:nth-child(6),
.home-list:nth-child(3) .list-title li:nth-child(4) {
    color: #ca9706;
 }
.home-list:nth-child(1) .list-title li:nth-child(3),
.home-list:nth-child(2) .list-title li:nth-child(4),
.home-list:nth-child(3) .list-title li:nth-child(5) {
    color: #d9457f;
 }
.home-list:nth-child(1) .list-title li:nth-child(1),
.home-list:nth-child(2) .list-title li:nth-child(2),
.home-list:nth-child(3) .list-title li:nth-child(6) {
    color: #f90;
 }
</style>