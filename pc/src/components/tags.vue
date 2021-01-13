<template>
  <div class="tags-page">
    <label><i class="ico-hottag"></i>热门标签</label>
    <ul>
      <li v-for="item in tagList" @click="search(item.name)" :class="{'active':tagname == item.name}" v-if="item.status"> {{ item.name }}</li>
    </ul>
  </div>
</template>

<script>
import { API } from '@/api';
export default {
  data () {
    return {
      tagList: {},
      searchText: '',
      tagname: this.$route.query.tagname || this.$route.query.keywords
    }
  },
  created () {
    if (JSON.parse(sessionStorage.getItem('tagList'))) {
      this.tagList = JSON.parse(sessionStorage.getItem('tagList'))
    } else {
      this.fetchTags()
    }
  },
  methods: {
    search (name) {
      this.$router.push({ path: '/pages/video/search', query: { tag: name } })
    },
    // 获取标签列表
    fetchTags () {
      let params = {
        page: 1,
        pageSize: 33
      }
      API.getTagList(params).then(res => {
        this.tagList = res.list
        sessionStorage.setItem('tagList', JSON.stringify(this.tagList))
      })
    }
  }
}
</script>

<style scoped lang="scss">
  .tags-page {
    margin: 20px auto 0;
    max-width: 1200px;
    label {
      color: #fff;
      margin: 0 0 15px 0;
      display: flex;
      font-size: 24px;
      align-items: center;
      .ico-hottag {
        background: url('../assets/images/icons/ico-hot.png') no-repeat;
        background-size: cover;
        width: 20px;
        height: 23px;
        margin-right: 10px;
      }
    }
    ul {
      display: flex;
      justify-content: space-between;
      flex-wrap: wrap;
      margin-right: -15px;
      &:after{
        flex: auto;
        content: "";
      }
      li {
        list-style: none;
        width: 85px;
        padding: 0 5px;
        height: 30px;
        line-height: 28px;
        border: 1px solid #d03fe5;
        border-radius: 0;
        color: #fff;
        font-size: 14px;
        text-align: center;
        margin: 0 15px 15px 0;
        cursor: pointer;
        transition: all ease 0.5s;
        background: rgb(193, 4, 99, .2);
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        border-radius: 30px;
        &.active,
        &:hover {
          background-color: #d03fe5;
        }
      }      
    }
  }
  .tags-btn {
    font-size: 14px;
    color: #666;
    text-align: center;
    p {
      position: relative;
      cursor: pointer;
      &:before,
      &:after {
        cursor: initial;
        top: 12px;
        width: 49.5%;
        content: '';
        background-color:#666;
        height:1px;
        position: absolute;
      }
      &:before {
        left: 0;
      }
      &:after {
        right: 0;        
      }
      i {
        font-size: 16px;
      } 
      i.el-icon-arrow-down {
        margin-top: 7px;
      }
    }
  }

  // 白色主题
  .lightTheme {
    .tags-page li {
      border: 1px solid #ccc;
      color: #333;
    }
    label {
      color: #ff0000;
    }
  }
</style>
