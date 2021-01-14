<template>
	<div class="notice">
    <el-dialog
		  :title="title || '官方公告'"
		  :visible.sync="dialogVisible"
		  width="400px">
		  <span>{{ noticeContent || '暂无公告' }}</span>
		  <span slot="footer" class="dialog-footer">
		    <el-button type="primary" @click="closeDialog">确 定</el-button>
		  </span>
		</el-dialog>
    <div class="overlay" v-if="dialogVisible"></div>
	</div>
</template>

<script>
export default {
  props: {
    noticeContent: String
  },
	data () {
		return {
      dialogVisible: false,
      title: ''
		}
	},
	created () {
    if (this.noticeContent && !sessionStorage.getItem('showNotice')) {
      setTimeout(()=> {
        this.dialogVisible = true
      }, 800)
    } else {
      this.dialogVisible = false
    }
	},
  methods: {
    closeDialog () {
      this.dialogVisible = false
      sessionStorage.setItem('showNotice', 1)
    }
  }
}
</script>

<style lang="scss">
.notice {
  .el-dialog {
    background: url('../assets/images/bg-notice.png') no-repeat;
    background-size: cover;
    height: 420px;
    padding-top: 200px;
    box-shadow: none;
    .el-dialog__title {
      font-size: 26px;
      font-weight: bold;
      margin: 10px 0 0 0;
      display: block;      
    }
    .el-dialog__body {
      padding: 5px 20px;   
      color: #000;
      font-size: 15px;   
      height: 90px;
      overflow-y: auto;
    }
    .el-dialog__footer {
      padding: 0;
      text-align: center;
      .el-button {
        background: url('../assets/images/bg-notice-btn.png') no-repeat;
        background-size: cover;
        width: 182px;
        border: none;
        font-size: 20px;    
      }
    }
  }
  .overlay {
    position: fixed;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    opacity: .5;
    background: #000;
    z-index: 2000;
  }
}
</style>