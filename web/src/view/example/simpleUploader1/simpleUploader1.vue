<template>
  <uploader
      :options="options"
      :file-status-text="statusText"
      :autoStart="false"
      @file-added="fileAdded"
      @file-progress="onFileProgress"
      @file-success="onFileSuccess"
      @file-error="onFileError"
      class="uploader-example"
  >
    <uploader-unsupport></uploader-unsupport>

    <el-alert
        style="margin-top: 0"
        title="第一步：请填写影片信息"
        type="success"
        show-icon
        :closable="false"
        effect="dark">
    </el-alert>

    <div class="img-detail">
      <el-form :model="formData" ref="formData" label-position="right" label-width="80px">
        <el-form-item label="*影片类型">
          <el-select v-model="formData.type" placeholder="请选择类型">
            <el-option label="国产" value="国产"></el-option>
            <el-option label="日韩" value="日韩"></el-option>
            <el-option label="欧美" value="欧美"></el-option>
            <el-option label="动漫" value="动漫"></el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="*影片名称:">
          <el-input v-model="formData.title" clearable placeholder="片名必填" ></el-input>
        </el-form-item>

        <el-form-item label="影片描述:">
          <el-input v-model="formData.desc" clearable placeholder="可以不填" ></el-input>
        </el-form-item>

        <el-form-item label="影片标签:">
          <el-select v-model="inputValue" @change="changeLocationValue" multiple placeholder="可以下拉选择多个" style="width: 800px">
            <el-option
              v-for="item in tagList"
              :key="item.ID"
              :label="item.name"
              :value="item.ID">
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="是否置顶:">
          <el-switch active-color="#409EFF" inactive-color="#cab3b3" active-text="是" inactive-text="否" v-model="formData.isTop" clearable ></el-switch>
        </el-form-item>

        <el-form-item label="是否推荐:">
          <el-switch active-color="#409EFF" inactive-color="#cab3b3" active-text="是" inactive-text="否" v-model="formData.isRecommend" clearable ></el-switch>
        </el-form-item>

        <el-form-item label="启用状态:">
          <el-switch active-color="#409EFF" inactive-color="#cab3b3" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
        </el-form-item>

      </el-form>
      <el-button type="primary" @click="checkForm" style="margin: 20px auto 0;" >下一步<i class="el-icon-bottom el-icon--bottom"></i></el-button>
    </div>

    <el-alert
        v-show="isShowStep2"
        title="第二步：请上传视频封面"
        show-icon
        type="success"
        :closable="false"
        effect="dark">
    </el-alert>

    <uploader-drop v-show="isShowStep2" class="uploader-drop-step2">
      <p>拖拽封面至此或点击</p>
      <uploader-btn>选择文件</uploader-btn>
    </uploader-drop>

    <el-alert
        v-show="isShowStep3"
        title="第三步：请上传原视频（mp4格式）"
        show-icon
        type="success"
        :closable="false"
        effect="dark">
    </el-alert>

    <uploader-drop v-show="isShowStep3">
      <p>拖拽图片集至此或点击</p>
      <uploader-btn>选择文件</uploader-btn>
    </uploader-drop>


    <el-alert
        v-show="isShowStep2 || isShowStep3"
        title="上传结果列表"
        type="success"
        show-icon
        :closable="false">
    </el-alert>

    <uploader-list v-show="isShowStep2 || isShowStep3" v-if="isShowList"></uploader-list>

    <el-button type="primary" @click="resetForm('formData')" style="margin: 20px auto 0;" v-show="isShowStep3">
      再次上传<i class="el-icon-refresh-right"></i>
    </el-button>

  </uploader>
</template>

<script>
var notUploadedChunks = []; // 已经上传过的文件chunkNumber数组
var isUploaded = false; // 文件已经上传成功了
import { mapGetters } from "vuex";
import { checkFileMd5,mergeFileMd5 } from "@/api/simpleUploader1";
import { getAllTagList } from "@/api/Tag";
import SparkMD5 from "spark-md5";
const path = process.env.VUE_APP_BASE_API;
export default {
  name: "simpleUploader",
  data(){
    return{
      md5:"",
      isCover: true, // 是否封面
      isShowStep1: true,
      isShowStep2: false,
      isShowStep3: false,
      isShowList: true,
      inputVisible: false,
      inputValue: '',
      tagList: [],
      departmentNames:[],
      formData: {
        title: '',
        desc: '',
        type: '',
        status:true,
        isTop:false,
        isRecommend:false,
        tag: [],
      }
    }
  },
  computed: {
    ...mapGetters("user", ["userInfo", "token"]),
    statusText() {
      return {
        success: "成功了",
        error: "出错了",
        uploading: "上传中",
        paused: "暂停中",
        waiting: "等待中"
      };
    },
    options() {
      return {
        target: path + "/movieUpload/movieUpload",
        testChunks: false,
        simultaneousUploads: 50, // 并发上传数
        chunkSize: 2 * 1024 * 1024,
        // query: this.formData, // 自定义参数
        allowDuplicateUploads: true, // 某文件上传后允许再次上传
        headers: {
          "x-token": this.token,
          "x-user-id": this.userInfo.ID
        },
        checkChunkUploadedByResponse(chunk) {
          if (isUploaded) {
            return true; // return true 会忽略当前文件，不会再发送给后台
          } else {
            // 根据已经上传过的切片来进行忽略
            return (
                notUploadedChunks &&
                notUploadedChunks.some(
                    item => item.chunkNumber == chunk.offset + 1
                )
            );
          }
        }
      };
    }
  },
  created () {
    this.getAllTagList()
  },
  methods: {
    changeLocationValue(val){
      this.departmentNames = []
      for(let i=0;i<=val.length-1;i++){
        this.tagList.find((item)=>{
          if(item.ID == val[i]){
            this.departmentNames.push(item.name)
          }
        });
      }
      this.formData.tag = this.departmentNames
    },
    async getAllTagList(row) {
      const res = await getAllTagList();
      if (res.code == 0) {
        console.log(res.data.list)
        this.tagList = res.data.list
      }
    },
    // 上传单个文件
    fileAdded(file) {
      this.computeMD5(file); // 生成MD5
    },
    // 计算MD5值
    computeMD5(file) {
      var that = this;
      isUploaded = false; // 这个文件是否已经上传成功过
      notUploadedChunks = []; // 未成功的chunkNumber
      var fileReader = new FileReader();
      var md5 = "";

      file.pause();

      fileReader.readAsArrayBuffer(file.file);
      fileReader.onload = async function(e) {
        if (file.size != e.target.result.byteLength) {
          this.error(
              "Browser reported success but could not read the file until the end."
          );
          return false;
        }
        md5 = SparkMD5.ArrayBuffer.hash(e.target.result, false);

        file.uniqueIdentifier = md5;
        if (md5 != "") {
          const res = await checkFileMd5({ md5: md5 });
          if (res.code == 0) {
            if (res.data.isDone) {
              // 上传成功过
              isUploaded = true;
              that.$message({
                message: "该文件已经上传成功过了，秒传成功。",
                type: "success"
              });

              file.cancel();
            } else {
              // 开始上传
              isUploaded = false;
              notUploadedChunks = res.data.chunks;
              if(notUploadedChunks.length){
                file.resume();
              }
            }
          }
        }
      };
      fileReader.onerror = function() {
        this.error(
            "generater md5 时FileReader异步读取文件出错了，FileReader onerror was triggered, maybe the browser aborted due to high memory usage."
        );
        return false;
      };
    },
    // 上传进度
    onFileProgress() {},
    // 上传成功
    async onFileSuccess(rootFile, file) {
      if (!this.isShowStep3) {
        this.isCover = true
      } else {
        this.isCover = false
      }
      this.isShowStep3 = true
      await mergeFileMd5({
        md5:file.uniqueIdentifier,
        fileName:file.name,
        title:this.formData.title,
        desc:this.formData.desc,
        type:this.formData.type,
        status:this.formData.status,
        isTop:this.formData.isTop,
        isRecommend:this.formData.isRecommend,
        tag:this.formData.tag,
        isCover:this.isCover
      })
    },
    onFileError(rootFile, file, response) {
      this.$message({
        message: response,
        type: "error"
      });
    },
    handleClose(tag) {
      this.formData.tag.splice(this.formData.tag.indexOf(tag), 1);
    },
    showInput() {
      this.inputVisible = true;
      this.$nextTick(() => {
        this.$refs.saveTagInput.$refs.input.focus();
      });
    },
    handleInputConfirm() {
      let inputValue = this.inputValue;
      if (inputValue) {
        this.formData.tag.push(inputValue);
      }
      this.inputVisible = false;
      this.inputValue = '';
    },
    resetForm() {
      this.$confirm('确定第三步影片全部上传完成?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.$message({
          type: 'success',
          message: '重置成功，可以继续上传!'
        });
        this.isCover = true
        this.isShowList = false
        this.isShowStep1 = true
        this.isShowStep2 = false
        this.isShowStep3 = false
        this.formData.title = ''
        this.formData.desc = ''
        this.formData.type = ''
        this.formData.tag = []
        setTimeout(() => {
          this.isShowList = true
        }, 2000)
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消重置'
        });
      });
    },
    checkForm () {
      // 验证是否填写表单信息 title  desc  type  status  tag
      var that = this;
      if (!that.formData.type) {
        that.$message({
          message: '请填写影片类型',
          type: "error"
        });
        return false
      } else if (!that.formData.title) {
        that.$message({
          message: '请填写影片名称',
          type: "error"
        });
        return false
      } 
      if (this.isShowStep1 && this.isCover) {
        this.isShowStep2 = true
      }
    }
  }
};
</script>

<style>
.uploader .uploader-drop {
  border: 1px dashed #5006e2;
  text-align: center;
  height: 100px;
  color: #3997f9;
}
.uploader .el-alert {
  margin: 20px 0 10px;
}
.uploader .el-alert__title {
  font-size: 16px;
  font-weight: bold;
}
.uploader-example {
  width: 880px;
  padding: 15px;
  margin: 115px 15px 20px;
  font-size: 12px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
}
.uploader-example .uploader-btn {
  margin-right: 4px;
  margin-top: 20px;
  border: 1px solid #409eff;
}
.uploader .uploader-drop p {
  margin-top: 20px;
}
.uploader .uploader-drop-step2 {
  height: 50px;
}
.uploader .uploader-drop-step2 p {
  margin-top: 0;
}
.uploader-example .uploader-drop-step2 .uploader-btn {
  margin-top: 5px !important;
}
.uploader-example .uploader-list {
  max-height: 440px;
  overflow: auto;
  overflow-x: hidden;
  overflow-y: auto;
}
.uploader-example .el-form-item {
  margin-bottom: 0;
}
.el-tag + .el-tag {
  margin-left: 10px;
}
.img-detail .button-new-tag {
  margin-left: 10px;
  height: 32px;
  line-height: 30px;
  padding-top: 0!important;
  padding-bottom: 0!important;
}
.input-new-tag {
  width: 90px;
  margin-left: 10px;
  vertical-align: bottom;
}
.el-icon-bottom {
  margin: 0 0 0 10px;
}
</style>