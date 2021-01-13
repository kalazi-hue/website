<template>
  <div class="video-list">
    <h2 class="title"><i></i>{{ title }}</h2>
    <p class="total">
      <i class="el-icon-info"></i>
      共搜索到<span>{{ videoList.length }}</span>条包含关键字<span>"{{ keywords || tag }}"</span>的影片
    </p>
    <ul class="list">
      <li v-for="(list, index) in videoList" @click="goPlayer(list)" :key="index" v-if="list.status">
        <div class="img">
          <div class="overlay"><i class="ico-play"></i></div>
          <img v-lazy="list.cover" alt="">
          <span class="date">{{ list.createdAt.split('T')[0] || '' }}</span>
        </div>
        <h5>{{ list.title }}</h5>
        <h6>
          <span class="score"><i class="ico-view"></i>{{ list.playCount }}</span>
          <span class="score"><i class="ico-like"></i>{{ list.star }}</span>
        </h6>
      </li>
    </ul>
    <p v-if="videoList.length <= 0" class="nodata"><i class="el-icon-folder-delete"></i>无数据,请更换关键字后重试</p>
    <el-pagination
    @size-change="handleSizeChange"
    @current-change="handleCurrentChange"
    :current-page="page"
    :page-size="10"
    layout="prev, pager, next, jumper"
    :total="totalPage"
    >
    </el-pagination>
  </div>
</template>

<script>
import { API } from '@/api'
export default {
  data () {
    return {
      title: '搜索结果',
      keywords: '',
      tag: '',
      videoList: [],
      totalPage: 0,
      page: 1,
      pageSize: 8,
    }
  },
  created () {
    this.getPath()
  },
  methods: {
    handleSizeChange (val) {
      console.log(val)
    },
    handleCurrentChange (val) { // 页码
      this.page = val
      this.getMovieListByKeyword()
    },
    getMovieListByKeyword (type) {
      let params = {
        page: this.page,
        pageSize: this.pageSize,
        keyword: this.keywords 
      }
      API.getMovieListByKeyword(params).then(res => {
        this.videoList = res.list
        this.pageSize = res.pageSize
        this.totalPage = res.total
      })
    },
    getMovieListByTag (type) {
      let params = {
        page: this.page,
        pageSize: this.pageSize,
        tags : this.tag 
      }
      API.getMovieListByTag(params).then(res => {
        this.videoList = res.list
        this.pageSize = res.pageSize
        this.totalPage = res.total
      })
    },
    goPlayer (item) {
      sessionStorage.setItem('videoDetail', JSON.stringify(item))
      this.$router.push({
        path: `/pages/video/detail/${item.ID}`
      })
    },
    backtop () {
      document.documentElement.scrollTop = document.body.scrollTop = 0
    },
    getPath () {
      this.keywords = this.$route.query.keywords
      this.tag = this.$route.query.tag
      if (this.keywords) {
        this.getMovieListByKeyword()        
      } else if (this.tag) {
        this.getMovieListByTag()        
      }
    }
  },
  watch: {
    '$route':'getPath'
  }
}
</script>

<style scoped lang="scss">
@import '@/assets/css/video.scss';
.video-list {
  .title {
    i {
      background: url('../../assets/images/icons/ico-novel.png') no-repeat;
      background-size: cover;
      width: 25px;
      height: 28px;
    }
  }
}
.total {
  color: #ccc;
  font-size: 18px;
  margin: 0 0 20px;
  .el-icon-info {
    font-size: 20px;
    margin-right: 5px;
  }
  span {
    color: #ce35e4;
    font-weight: bold;
    margin: 0 5px;
  }
}
.nodata {
  text-align: center;
  color: #666;
  line-height: 40px;
  padding: 50px 0;
  .el-icon-folder-delete {
    font-size: 180px;
    display: block;
    color: #444;
  }
}
</style>