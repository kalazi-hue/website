<template>
  <div class="photo-list">
    <h2 class="title"><i></i> {{ title }} </h2>
    <div class="container-water-fall">
      <waterfall :col="col" :data="data" @loadmore="loadmore">
        <template>
          <div class="cell-item" v-for="(item,index) in data" :key="index" @click="goPlayer(item.content)"  v-if="item.status">
            <img v-if="item.cover" :lazy-src="item.cover" alt="加载错误">
            <div class="item-body">
              <div class="item-desc">{{item.title}}</div>
              <div class="item-footer">
                <div class="name">{{item.updatedAt}}</div>
              </div>
            </div>
          </div>
        </template>
      </waterfall>
      <p v-if="nomore" class="nomore">没有更多了~</p>
      <!-- <loading :show="loading"/> -->
    </div>
  </div>
</template>

<script>
import { API } from '@/api'
export default {
  props: {
    isHome: Boolean
  },
  data() {
    return {
      nomore:false,
      data: [],
      col: 5,
      title: '美女写真',
      typeid: 1,
      photoList: [],
      totalPage: 0,
      page: 1,
      pageSize: 16
    };
  },
  computed: {
    itemWidth() {
      return 133 * 0.5 * (document.documentElement.clientWidth / 355);
    },
    gutterWidth() {
      return 10 * 0.5 * (document.documentElement.clientWidth / 355);
    }
  },
  mounted() {
    this.loadmore()
  },
  methods: {
    loadmore (type) {
      if (this.isHome && this.page> 1) { // 首页引入列表不需要加载更多
        this.nomore = false
        return false
      }
      let params = {
        page: this.page,
        pageSize: this.pageSize
      }
      API.getPhotoAlbumList(params).then(res => {
        if (res.list.length === 0) {
          this.nomore = true
          return false
        }
        this.data = this.data.concat(res.list);
        this.pageSize = res.pageSize
        this.totalPage = res.total
        this.page++
      })
    },
    goPlayer (content) {
      this.$router.push({
        path: `/pages/pic/detail?content=${content}`
      })
    },
    backtop () {
      document.documentElement.scrollTop = document.body.scrollTop = 0
    }
  }
}
</script>

<style scoped lang="scss">
  .photo-list {
    .title {
      i {
        background: url('../assets/images/icons/ico-photo.png') no-repeat;
        background-size: cover;
        width: 61px;
        height: 40px;
        margin: 0 10px 0 0;
      }
    }
    .nomore {
      margin: 20px auto;
      color: #868080;
      text-align: center;
    }
  }
</style>
<style  lang="less" scoped>
@import '@/assets/css/photo.scss';
</style>
