<template>
  <footer class="footer" :class="{'fixed-Footer': fixFooter}">
    <div class="main">
      <i class="el-icon-close" @click="closeFix" v-if="fixFooter"></i>
      <div class="latest-domain">
        <label>最新网址</label>
        <input type="text" :value="configList.lastAccessDomain" id="copyObj">
        <label @click="CopyUrl()">复制分享</label>
      </div>
      <ul class="nav-bottom">
        <li @click="toDtail(configList.lastDownDomain)"><img src="../assets/images/icons/ico-app.png">下载APP</li>
        <li @click="toDtail(configList.gitHub)"><img src="../assets/images/icons/ico-github.png">Github地址</li>
        <li @click="toDtail(configList.email)"><img src="../assets/images/icons/ico-email.png">邮箱</li>
        <li @click="toDtail(configList.alternateAddress)"><img src="../assets/images/icons/ico-backup.png">备用地址</li>
        <li @click="toDtail(configList.exchangeGroup)"><img src="../assets/images/icons/ico-group.png">用户交流群</li>
        <li @click="toDtail('/pages/help/faq')"><img src="../assets/images/icons/ico-yinsi.png">隐私政策</li>
        <li @click="toDtail('/pages/help/faq')"><img src="../assets/images/icons/ico-faq.png">常见问题</li>
        <li @click="toDtail('/pages/help/contact')"><img src="../assets/images/icons/ico-contact.png">联系我们</li>
      </ul>
    </div>

    <notice :noticeContent="configList.notice" /> <!-- 公告弹窗 -->

  </footer>
</template>

<script>
import { API } from '@/api'
import notice from '@/components/notice'
export default {
  components: {
    notice
  },
  data () {
    return {
      fixFooter: true,
      adList: [],
      configList: [],
    }
  },
  created () {
    if (JSON.parse(sessionStorage.getItem('adList'))) {
      this.adList = JSON.parse(sessionStorage.getItem('adList'))
      this.$store.commit('editAdList', this.adList)
    } else {
      this.getAdConfigList()
    }
    if (JSON.parse(sessionStorage.getItem('configList'))) {
      this.configList = JSON.parse(sessionStorage.getItem('configList'))
    } else {
      this.getSystemConfigList()
    }
  },
  methods: {
    getAdConfigList () {
      let params = {
        page: 1,
        pageSize: 100
      }
      API.getAdConfigList(params).then(res => {
        this.adList = res.list
        sessionStorage.setItem('adList', JSON.stringify(this.adList))
        this.$store.commit('editAdList', this.adList)
      })
    },
    getSystemConfigList () {
      let params = {
        page: 1,
        pageSize: 100
      }
      API.getSystemConfigList(params).then(res => {
        this.configList = res.list[0]
        sessionStorage.setItem('configList', JSON.stringify(this.configList))
      })
    },
    toDtail (path) {
      if (path.indexOf('/pages/') > -1) {
        this.$router.push({
            path: path
        })
      } else {
        window.open(path)
      }
    },
    CopyUrl(){
      let url = document.querySelector('#copyObj');
      url.select(); // 选择对象
      document.execCommand("Copy");
      this.$message({
        message: '复制成功，请粘贴分享',
        type: 'success'
      });
    },
    closeFix () {
      this.fixFooter = false
      this.$emit('closeOrder', 'yes')
    }
  }
}
</script>

<style scoped lang="scss">
  .footer {
    background-color: #141414;
    text-align: center;
    transition: all ease .5s;
    opacity: 0.92;
    .main {
      padding: 20px 0;
      position: relative;
      .el-icon-close {
        position: absolute;
        top: 10px;
        right: 0;
        font-size: 30px;
        color: #cd0067;
        font-weight: bold;
        cursor: pointer;
      }
      .latest-domain {
        font-size: 26px;
        width: 700px;
        margin: 0 auto;
        display: flex;
        justify-content: space-evenly;
        background-color: #cd0067;
        height: 45px;
        line-height: 45px;
        border-radius: 30px;
        input {
          outline: none;
          font-size: 22px;
          text-align: center;
          width: 360px;
          border: none;
          background-color: #fff;
          color: #000;
        }
        label {
          cursor: pointer;        
        }
      }
      .nav-bottom {
        margin: 20px 0 5px 0;
        display: flex;
        justify-content: space-between;   
        li {
          cursor: pointer;
          transition: all ease .2s;
          &:hover {
            margin-top: -5px;
          }
          img {
            display: block;
            margin: 0 auto 10px;
          }
        }   
      }
    }
  }
  .fixed-Footer {
    position: fixed;
    width: 100%;
    bottom: 0;
    left: 0;
    z-index: 6666;
  }
</style>