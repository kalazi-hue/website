<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="名称:">
                <el-input v-model="formData.title" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="封面:">
                <el-input v-model="formData.reads" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="简介:">
                <el-input v-model="formData.description" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="内容:">
                <el-input v-model="formData.content" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="关键词:">
                <el-input v-model="formData.keyword" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="是否置顶:">
                <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.isTop" clearable ></el-switch>
          </el-form-item>
           
             <el-form-item label="点击量:"><el-input v-model.number="formData.reads" clearable placeholder="请输入"></el-input>
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
    createPhotoAlbum,
    updatePhotoAlbum,
    findPhotoAlbum
} from "@/api/PhotoAlbum";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "PhotoAlbum",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            title:"",
            reads:"",
            description:"",
            content:"",
            keyword:"",
            isTop:false,
            reads:0,
            status:false,
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createPhotoAlbum(this.formData);
          break;
        case "update":
          res = await updatePhotoAlbum(this.formData);
          break;
        default:
          res = await createPhotoAlbum(this.formData);
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
    const res = await findPhotoAlbum({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.rephotoAlbum
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