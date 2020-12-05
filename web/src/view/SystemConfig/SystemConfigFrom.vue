<template>
<div>
    <el-form :model="formData" label-position="right" label-width="80px">
             <el-form-item label="公告:">
                <el-input v-model="formData.notice" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="最新访问域名:">
                <el-input v-model="formData.lastAccessDomain" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="最新下载域名:">
                <el-input v-model="formData.lastDownDomain" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="交流群:">
                <el-input v-model="formData.exchangeGroup" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="GitHub地址:">
                <el-input v-model="formData.gitHub" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="反馈邮箱:">
                <el-input v-model="formData.email" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="备用地址:">
                <el-input v-model="formData.alternateAddress" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="隐私政策:">
                <el-input v-model="formData.privacyPolicy" clearable placeholder="请输入" ></el-input>
          </el-form-item>
           
             <el-form-item label="联系我们:">
                <el-input v-model="formData.contact" clearable placeholder="请输入" ></el-input>
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
    createSystemConfig,
    updateSystemConfig,
    findSystemConfig
} from "@/api/SystemConfig";  //  此处请自行替换地址
import infoList from "@/mixins/infoList";
export default {
  name: "SystemConfig",
  mixins: [infoList],
  data() {
    return {
      type: "",formData: {
            notice:"",
            lastAccessDomain:"",
            lastDownDomain:"",
            exchangeGroup:"",
            gitHub:"",
            email:"",
            alternateAddress:"",
            privacyPolicy:"",
            contact:"",
            
      }
    };
  },
  methods: {
    async save() {
      let res;
      switch (this.type) {
        case "create":
          res = await createSystemConfig(this.formData);
          break;
        case "update":
          res = await updateSystemConfig(this.formData);
          break;
        default:
          res = await createSystemConfig(this.formData);
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
    const res = await findSystemConfig({ ID: this.$route.query.id })
    if(res.code == 0){
       this.formData = res.data.resystemConfig
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