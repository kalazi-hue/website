<template>
  <div class="header">
    <div class="main">
      <h1 @click="toDtail(0)">
        <img src="../assets/images/logo.png" alt="logo" class="logo">
      </h1>
      <div class="menus">
        <ul>
          <li v-for="(item, index) in menuList" @click="toDtail(item.typeid,item.title)" :class='{active:item.title === activeTitle}' :key="index"> {{ item.title }}</li>
        </ul>
      </div>
      <div class="search">
        <el-input
          placeholder="请输入关键字搜片源"
          v-model="keywords"
          >
          <i slot="suffix" class="el-input__icon el-icon-search" @click="doSearch"></i>
        </el-input>
      </div>    
    </div>
  </div>
</template>

<script>
import { API } from '@/api'
import { menus } from '@/config/menus'
export default {
  data () {
    return {
      keywords: '',
      activeTitle: '首页',
      token: '',
      menuList: []
    }
  },
  mounted () {
  },
  created () {
    setTimeout(() => {
      this.menuList = menus.list
      if (this.$route.query.typeid) { // 视频菜单刷新选中
        this.title = this.$route.query.typeid
        this.activeTitle = this.$route.query.typeid
      }
      if (this.$route.path.indexOf('/pages/pic/') > -1) {  // 写真菜单刷新选中
        this.activeTitle = '写真'
      } else if (this.$route.path.indexOf('/pages/novel/') > -1) {  // 小说菜单刷新选中
        this.activeTitle = '小说'
      } else if (this.$route.path.indexOf('/pages/live/') > -1) {  // 直播菜单刷新选中
        this.activeTitle = '直播'
      }
    }, 500)
  },
  methods: {
    doSearch () {
      if (!this.keywords) {
        this.$notify({
          title: '提示',
          message: '请输入关键字',
          type: 'warning',
          duration: 1000
        });
        return false
      } else {
        this.$router.push({
            path: `/pages/movie/search?keywords=${this.keywords}`
        })          
      }
    },
    toDtail (id,title) {
      this.activeTitle = title
      if (title === '首页') {
        this.$router.push({
            path: `/`
        })
      } else if (title === '写真') {
        this.$router.push({
            path: `/pages/pic/index`
        })
      }  else if (title === '小说') {
        this.$router.push({
            path: `/pages/novel/list?viewid=${id}`
        })
      } else if (title === '直播') {
        this.$router.push({
            path: `/pages/live/list`
        })
      } else {
        this.$router.push({
            path: `/pages/movie/index?typeid=${title}`
        })     
      }
    }
  }
}
</script>
 
<style lang="scss" scoped>
.header {
  background-color: #141414;
  height: 80px;
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  z-index: 9;
  .main {
    display: flex;
    align-items: center;
    height: 100%;
    justify-content: space-between;
    .logo {
      height: 60px;
    }
    .menus {
      ul {
        height: 100%;
        display: flex;
        align-items: center;
        justify-content: space-between;
        li {
          height: 100%;
          line-height: 80px;
          color: #fff;
          font-size: 20px;
          padding: 0 20px;
          cursor: pointer;
          &:hover,
          &.active {
            color: #000;
            background: url('../assets/images/bg-hover.png') no-repeat center;
            background-size: contain;
          }
        }
      }
    }
    .search {
      .el-icon-search {
        margin-top: 5px;
        cursor: pointer;
        background: url('../assets/images/icons/ico-search.png') no-repeat; 
        text-indent: -9999px;         
      }
    }
  }
}

@media screen and (max-width: 1180px) {
  .header .main .menus ul li {
    font-size: 18px;
    padding: 0 10px;
  }
}
</style>

<style lang="scss">
.search {
  .el-input__inner {
    color: #ccc;
    background-color: transparent;
    border: none;
    border-bottom: 1px solid #777;
  }
  .el-input__icon {
    font-size: 26px;
    color: #b40dff;
  }
}
</style>