<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                  
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增SystemConfig</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>


      <el-table-column label="公告1" prop="notice" ></el-table-column>
      <el-table-column label="公告2" prop="notice2" ></el-table-column>

      <el-table-column label="最新访问域名" prop="lastAccessDomain" ></el-table-column>
    
    <el-table-column label="最新下载域名" prop="lastDownDomain" ></el-table-column> 
    
    <el-table-column label="交流群" prop="exchangeGroup" ></el-table-column> 
    
    <el-table-column label="GitHub地址" prop="gitHub" ></el-table-column> 
    
    <el-table-column label="反馈邮箱" prop="email" ></el-table-column> 
    
    <el-table-column label="备用地址" prop="alternateAddress" ></el-table-column> 
    
    <el-table-column label="隐私政策" prop="privacyPolicy" ></el-table-column> 
    
    <el-table-column label="联系我们" prop="contact" ></el-table-column> 
    
    <el-table-column label="创建日期" >
         <template slot-scope="scope">{{scope.row.createdAt|formatDate}}</template>
    </el-table-column>
        
      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateSystemConfig(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteSystemConfig(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      <el-form :model="formData" label-position="right" label-width="80px">
         <el-form-item label="公告1:">
            <el-input v-model="formData.notice" clearable placeholder="请输入" ></el-input>
      </el-form-item>
          <el-form-item label="公告2:">
            <el-input v-model="formData.notice2" clearable placeholder="请输入" ></el-input>
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
       </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createSystemConfig,
    deleteSystemConfig,
    deleteSystemConfigByIds,
    updateSystemConfig,
    findSystemConfig,
    getSystemConfigList
} from "@/api/SystemConfig";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "SystemConfig",
  mixins: [infoList],
  data() {
    return {
      listApi: getSystemConfigList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            notice:"",
            notice2:"",
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
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10             
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        if(this.multipleSelection.length == 0){
          this.$message({
            type: 'warning',
            message: '请选择要删除的数据'
          })
          return
        }
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteSystemConfigByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateSystemConfig(row) {
      const res = await findSystemConfig({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.resystemConfig;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        notice:"",
        notice2:"",
          lastAccessDomain:"",
          lastDownDomain:"",
          exchangeGroup:"",
          gitHub:"",
          email:"",
          alternateAddress:"",
          privacyPolicy:"",
          contact:"",
          
      };
    },
    async deleteSystemConfig(row) {
      this.visible = false;
      const res = await deleteSystemConfig({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
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
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();
  
}
};
</script>

<style>
</style>