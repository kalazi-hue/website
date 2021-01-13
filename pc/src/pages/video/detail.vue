<template>
  <div>
    <div id="dplayer" class="play-root" ></div>
    <div class="videoinfo">
      <ul class="bottom">
        <li><i class="ico-view"></i>{{ videoDetail.playCount }}次观看</li>
        <li><i class="ico-view"></i>影片时长{{ videoDetail.playTime }}</li>
        <li @click="addLike"><i class="ico-like"></i>{{ videoDetail.star }}个点赞</li>
        <li v-clipboard:error="onError" v-clipboard:copy="shareHref" v-clipboard:success="onCopy"><i class="ico-share"></i>分享给好友</li>
        <li @click="dialogVisible = true"><i class="ico-flag"></i>报错反馈</li>
      </ul>
      <h4 class="title">{{ videoDetail.title }}</h4>
      <h5 class="title">视频简介：{{ videoDetail.description }}</h5>
      <h6 class="title">上架时间：{{ videoDetail.createdAt }}</h6>
      <ul class="tags">
        <li v-for="item in filterTags(videoDetail.tags)" @click="search(item)" v-if="item">{{ item }}</li>
      </ul>
      <input type="text" v-text="shareHref" :value="shareHref">
    </div>

    <!-- 报错反馈弹窗 -->
    <el-dialog
      title="提示"
      :visible.sync="dialogVisible"
      width="30%">
      <div>
        <el-radio-group v-model="errorName" @change="radioChange" >
          <el-radio label="轻微卡顿" border>轻微卡顿</el-radio>
          <el-radio label="严重卡顿" border>严重卡顿</el-radio>
          <el-radio label="无法播放" border>无法播放</el-radio>
        </el-radio-group>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取 消</el-button>
        <el-button type="primary" @click="submitFeedback">确 定</el-button>
      </span>
    </el-dialog>

  </div>

</template>

<script>
import DPlayer from 'dplayer'
import { API } from '@/api'
import axios from 'axios'
export default {
  data() {
    return {
      flag: false,
      flagReport: false,
      dp: null,
      shareHref: '',
      dialogVisible: false,
      videoDetail: [],
      errorName: ''
    }
  },
  mounted() {
    if (JSON.parse(sessionStorage.getItem('videoDetail'))) { // 获取列表页缓存的播放页数据
      this.videoDetail = JSON.parse(sessionStorage.getItem('videoDetail'))
    } else {
      this.$message('未找到片源，请返回列表页');
    }
    console.log(this.videoDetail)
    const dp = new DPlayer({
      container: document.getElementById('dplayer'),
      autoplay: true,
      screenshot: true,
      hotkey: true,
      theme: '#ff40dd',
      loop: true,
      lang: 'zh-cn',
      preload: 'auto',
      logo: require('../../assets/images/logo.png'),
      volume: 0.7,
      video: {
        url: this.videoDetail.playUrl,
        type: 'auto',
      },
    })
    this.shareHref = window.location.href
  },
  methods:{
    radioChange(val){
      console.log(val)
      this.errorName = val
    },
    submitFeedback () {
      if (this.flagReport) {
        this.$message({
          showClose: true,
          type: 'warning',
          duration: 2000,
          message: '不能重复反馈'
        })
        return false
      }
      if (!this.errorName) {
        this.$message({
          showClose: true,
          type: 'warning',
          duration: 2000,
          message: '请选择原因'
        })
        return false        
      }
      let params = {
        createdAt: this.videoDetail.createdAt,
        movieId: this.videoDetail.ID,
        playUrl: this.videoDetail.playUrl,
        title: this.videoDetail.title,
        type: this.errorName,
        updatedAt: this.videoDetail.updatedAt
      }
      API.createPlayFeedback(params).then(res => {
        this.$message({
          showClose: true,
          type: 'success',
          duration: 2000,
          message: '反馈成功！'
        })
        this.flagReport = true
        this.dialogVisible = false
        setTimeout(() => {
          this.flagReport = false
        },20000)
      })
    },
    addLike () {
      if (this.flag) {
        this.$message({
          showClose: true,
          type: 'warning',
          duration: 2000,
          message: '不能重复点赞'
        })
        return false
      }
      let params = {
        id: this.videoDetail.ID
      }
      API.movieApproval(params).then(res => {
        this.$message({
          showClose: true,
          type: 'success',
          duration: 2000,
          message: '点赞👍 +1'
        })
        this.videoDetail.star++
        this.flag = true
        setTimeout(() => {
          this.flag = false
        },20000)
      })
    },
    onPlay() {
      this.dp.play()
    },
    onCopy (e) { // 复制成功
      this.$message({
        showClose: true,
        type: 'success',
        duration: 2000,
        message: '复制成功,请去分享~'
      })
    },
    onError (e) { // 复制失败
      this.$message({
        showClose: true,
        type: 'error',
        duration: 1200,
        message: '复制失败'
      })
    },
    filterTags (val) {
      if (val) {
        let tagArr = val.split(',')
        sessionStorage.removeItem("tagLike")
        sessionStorage.setItem('tagLike', val)
        return tagArr
      }
    },
    search (name) {
      this.$router.push({ path: '/pages/video/search', query: { keywords: name } })
    }
  }
}
</script>

<style scoped lang="scss">
  .play-root {
    width: 100%;
    height: 660px;
    background-color: #000;
    margin: 20px auto 0;
  }
  .videoinfo {
    color: #ccc;
    .bottom {
      padding: 15px 15px 25px 15px;
      display: flex;
      justify-content: space-between;
      li {
        font-size: 14px;
        cursor: pointer;
        display: flex;        
        align-items: center;
        .ico-share {
          background: url('../../assets/images/icons/ico-share.png') no-repeat;
          background-size: cover;
          width: 18px;
          height: 19px;        
          margin-right: 10px;  
        }
        .ico-flag {
          background: url('../../assets/images/icons/ico-flag.png') no-repeat;
          background-size: cover;
          width: 18px;
          height: 19px;        
          margin-right: 10px;  
        }
        &:hover {
          color: #fff;
        }
      }
    }
    h4 {
      color: #fff;
      font-size: 26px;
      font-weight: normal;
    }
    h5,h6 {
      font-size: 16px;
      font-weight: normal;
      margin: 15px 0;
    }
    .tags {
      margin: 10px 0 20px 0;
      display: flex;
      flex-wrap: wrap;
      li {
        list-style: none;
        padding: 0 10px;
        min-width: 60px;
        height: 26px;
        line-height: 24px;
        border: 1px solid #d03fe5;
        border-radius: 0;
        font-size: 14px;
        text-align: center;
        margin: 0 10px 10px 0;
        cursor: pointer;
        transition: all ease 0.5s;
        background: rgba(193, 4, 99, 0.2);
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        border-radius: 30px;
        &:hover {
          background-color: #d03fe5;
        }
      }
    }
  }
</style>
<style lang="scss">
  .dplayer-logo {
    max-width: 120px;
    right: 20px;
    left: auto;
  }
  .dplayer-controller .dplayer-bar-wrap .dplayer-bar .dplayer-loaded, 
  .dplayer-controller .dplayer-bar-wrap .dplayer-bar .dplayer-played,
  .dplayer-controller .dplayer-bar-wrap .dplayer-bar,
  .dplayer-controller .dplayer-bar-wrap {
    height: 10px !important;
  }
  .dplayer-controller .dplayer-bar-wrap .dplayer-bar .dplayer-played .dplayer-thumb {
    margin-top: -5px !important;
    height: 20px !important;
    width: 20px !important;
  }
</style>
