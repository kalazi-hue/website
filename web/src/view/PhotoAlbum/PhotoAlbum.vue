<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">          
            <el-form-item label="是否置顶" prop="isTop">
            <el-select v-model="searchInfo.isTop" clear placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>     
            <el-form-item label="是否上架" prop="status">
            <el-select v-model="searchInfo.status" clear placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>   
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增photoAlbum</el-button>
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
    <el-table-column type="selection" width="40"></el-table-column>

    <el-table-column prop="title" label="名称" width="120" height="80" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.title}}</p>
        </template>
    </el-table-column>
    
    <el-table-column prop="cover" label="封面图" sortable width="80">
        <template slot-scope="scope">
            <img :src="scope.row.cover" :title="scope.row.cover" alt style="background-color: #ccc;width: 60px;height: 40px" />
        </template>
    </el-table-column>
 
    <el-table-column prop="createdAt" label="创建日期" width="120" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.createdAt|formatDate}}</p>
        </template>
    </el-table-column>

    <el-table-column prop="description" label="影集简介" width="120" height="80" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.description}}</p>
        </template>
    </el-table-column>

    <el-table-column prop="content" label="影集内容" width="120" height="80" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.content}}</p>
        </template>
    </el-table-column>
    
    <el-table-column prop="keyword" label="关键词" width="120" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.keyword}}</p>
        </template>
    </el-table-column>

    <el-table-column label="点击量" prop="reads" ></el-table-column> 
    
   <el-table-column label="是否置顶" prop="isTop" >
      <template slot-scope="scope">
          <el-switch
            v-model="scope.row.isTop"
            active-color="#F56C6C"
            inactive-color="#cab3b3"
            @change="changeSwitch($event,scope.row, 'isTop')"/>
      </template>
    </el-table-column>
    
    
    <el-table-column label="是否上架" prop="status" >
      <template slot-scope="scope">
          <el-switch
            v-model="scope.row.status"
            active-color="#F56C6C"
            inactive-color="#cab3b3"
            @change="changeSwitch($event,scope.row, 'isShelf')"/>
      </template>
    </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updatePhotoAlbum('',scope.row,'','')" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deletePhotoAlbum(scope.row)">确定</el-button>
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
         <el-form-item label="名称:">
            <el-input v-model="formData.title" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="封面:">
            <el-input v-model="formData.cover" clearable placeholder="请输入" ></el-input>
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

       <el-form-item label="点击量:"><el-input v-model.number="formData.reads" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="是否置顶:">
            <el-switch active-color="#13ce66" inactive-color="#cab3b3" active-text="是" inactive-text="否" v-model="formData.isTop" clearable ></el-switch>
       </el-form-item>
       
       
         <el-form-item label="是否上架:">
            <el-switch active-color="#13ce66" inactive-color="#cab3b3" active-text="是" inactive-text="否" v-model="formData.status" clearable ></el-switch>
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
    createPhotoAlbum,
    deletePhotoAlbum,
    deletePhotoAlbumByIds,
    updatePhotoAlbum,
    findPhotoAlbum,
    getPhotoAlbumList
} from "@/api/PhotoAlbum";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "PhotoAlbum",
  mixins: [infoList],
  data() {
    return {
      listApi: getPhotoAlbumList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            title:"",
            cover:"",
            description:"",
            content:"",
            keyword:"",
            isTop:false,
            reads:0,
            status:false,
            
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
    changeSwitch (e, row, type) { // 列表中操作switch按钮
      this.updatePhotoAlbum(e, row, type, 'isSwitch');
    },
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10           
        if (this.searchInfo.isTop==""){
          this.searchInfo.isTop=null
        }         
        if (this.searchInfo.status==""){
          this.searchInfo.status=null
        }      
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
        const res = await deletePhotoAlbumByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updatePhotoAlbum(e, row, type, isSwitch) {
      const res = await findPhotoAlbum({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.rephotoAlbum;
        if (isSwitch === 'isSwitch') { // 表格列Switch按钮直接切换
          this.dialogFormVisible = false;
          if (type === 'isTop') {
            this.formData.isTop = e
          } else if (type === 'isRecommend') {
            this.formData.isRecommend = e
          } else if (type === 'isShelf') {
            this.formData.status = e
          }
          this.enterDialog(isSwitch)
        } else {
          this.dialogFormVisible = true;          
        }
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          title:"",
          cover:"",
          description:"",
          content:"",
          keyword:"",
          isTop:false,
          reads:0,
          status:false,
          
      };
    },
    async deletePhotoAlbum(row) {
      this.visible = false;
      const res = await deletePhotoAlbum({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog(isSwitch) {
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
        this.closeDialog();
        if (isSwitch !== 'isSwitch') { // 表格列Switch按钮直接切换
          this.getTableData();
        }
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