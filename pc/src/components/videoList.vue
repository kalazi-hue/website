<template>
	<div class="video-list">
    <h2 class="title"><i></i>{{ title }}</h2>
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

      <!-- 数据加载前骨架占位 -->
      <li v-for="(list, index) in 16" v-if="videoList.length <= 0" :key="index">
        <div class="img placeholder"></div>
        <h5 class="placeholder"></h5>
        <h6 class="placeholder"></h6>
      </li>
    </ul>
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
import { menus } from '@/config/menus'
export default {
	props: {
    isHome: Boolean,
    showNewList: Boolean,
    showRecommendList: Boolean
	},
	data () {
		return {
      title: '国产视频',
      typeid: 1,
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
      this.getMovieListByTypeId()
    },
    getMovieListByTypeId (type) {
      console.log(this.$route.path)
      let params = {
        page: this.page,
        pageSize: this.pageSize,
        typeId: this.typeid 
      }
      API.getMovieListByTypeId(params).then(res => {
        this.videoList = res.list
        this.pageSize = res.pageSize
        this.totalPage = res.total
        // console.log(this.totalPage)
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
      if (this.$route.path.indexOf('/pages/video/') > -1) {
        this.typeid = this.$route.query.typeid
        this.getMovieListByTypeId()
        this.getTitle()     
      }
    },
    getTitle () {
      let _this = this
      menus.list.forEach(item=> {
        if (Number(_this.$route.query.typeid) === Number(item.typeid)) {
          _this.title = item.title
        }
      })   
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
      background: url('../assets/images/icons/ico-video.png') no-repeat;
      background-size: cover;
      width: 28px;
      height: 28px;
    }
  }
}
</style>