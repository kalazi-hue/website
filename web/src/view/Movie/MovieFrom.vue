<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="片名:">
                <el-input v-model="formData.title" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="影片简介:">
                <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="封面图:">
                <el-input v-model="formData.cover" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="影片时长:">
                <el-input v-model="formData.playTime" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="播放链接:">
                <el-input v-model="formData.playUrl" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="下载链接:">
                <el-input v-model="formData.downUrl" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="影片类型:"><el-input v-model.number="formData.type" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="标签:">
                <el-input v-model="formData.tags" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="点赞数:"><el-input v-model.number="formData.star" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="观影次数:"><el-input v-model.number="formData.playCount" clearable placeholder="请输入"></el-input>
          </el-form-item>
           
             <el-form-item label="是否置顶:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isTop" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="是否推荐:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isRecommend" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="上架时间:">
                  <el-date-picker type="date" placeholder="选择日期" v-model="formData.shelfTime" clearable></el-date-picker>
           </el-form-item>
           
             <el-form-item label="是否上架:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
          </el-form-item>
           <el-form-item>
           <el-button @click="save" type="primary">保存</el-button>
           <el-button @click="back" type="primary">返回</el-button>
           </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
    createMovie,
    updateMovie,
    findMovie
} from "@/api/Movie";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "Movie",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            title:"",
            description:"",
            cover:"",
            playTime:"",
            playUrl:"",
            downUrl:"",
            type:0,
            tags:"",
            star:0,
            playCount:0,
            isTop:false,
            isRecommend:false,
            shelfTime:new Date(),
            status:false,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createMovie(this.formData);
          break;
        case "update":
          res = await updateMovie(this.formData);
          break;
        default:
          res = await createMovie(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
      }
    },
    back(){
        this.$router.go(-1)
    }
  },
  async created() {
   // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if(this.$route.query.id){
    const res = await findMovie({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.removie
       this.type == "update"
     }
    }else{
     this.type == "create"
   }
  
}
};
</script>

<style>
</style>