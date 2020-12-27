<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="片名">
          <el-input placeholder="搜索条件" v-model="searchInfo.title"></el-input>
        </el-form-item>                
        <el-form-item label="标签">
          <el-input placeholder="搜索条件" v-model="searchInfo.tags"></el-input>
        </el-form-item>        
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
            <el-form-item label="是否推荐" prop="isRecommend">
            <el-select v-model="searchInfo.isRecommend" clear placeholder="请选择">
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

    <el-table-column prop="title" label="片名" width="120" height="80" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.title}}</p>
        </template>
    </el-table-column>

    <el-table-column prop="description" label="影片简介" width="120" height="80" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.description}}</p>
        </template>
    </el-table-column>
    
    <el-table-column prop="cover" label="封面图" sortable width="80">
        <template slot-scope="scope">
            <img :src="scope.row.cover" :title="scope.row.cover" alt style="width: 60px;height: 40px" />
        </template>
    </el-table-column>
    
    <el-table-column label="影片时长" prop="playTime" ></el-table-column> 
    
    <el-table-column prop="playUrl" label="播放链接" width="120" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.playUrl}}</p>
        </template>
    </el-table-column>

    <el-table-column prop="downUrl" label="下载链接" width="120" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.downUrl}}</p>
        </template>
    </el-table-column>

    <el-table-column label="影片类型" prop="type" ></el-table-column> 

    <el-table-column prop="tags" label="标签" width="120" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.tags}}</p>
        </template>
    </el-table-column>
    
    <el-table-column label="点赞数" prop="star" ></el-table-column> 
    
    <el-table-column label="观影次数" prop="playCount" ></el-table-column> 
    
    <el-table-column label="是否置顶" prop="isTop" >
         <template slot-scope="scope">{{scope.row.isTop|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="是否推荐" prop="isRecommend" >
         <template slot-scope="scope">{{scope.row.isRecommend|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column prop="shelfTime" label="上架时间" width="120" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.shelfTime}}</p>
        </template>
    </el-table-column>
    
    <el-table-column label="是否上架" prop="status" >
         <template slot-scope="scope">{{scope.row.status|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column prop="createdAt" label="创建日期" width="120" show-overflow-tooltip>
        <template slot-scope="scope">
            <p class="">{{scope.row.createdAt|formatDate}}</p>
        </template>
    </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-button class="table-button" @click="updateMovie(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteMovie(scope.row)">确定</el-button>
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
    createMovie,
    deleteMovie,
    deleteMovieByIds,
    updateMovie,
    findMovie,
    getMovieList
} from "@/api/Movie";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "Movie",
  mixins: [infoList],
  data() {
    return {
      listApi: getMovieList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
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
        if (this.searchInfo.isTop==""){
          this.searchInfo.isTop=null
        }        
        if (this.searchInfo.isRecommend==""){
          this.searchInfo.isRecommend=null
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
        const res = await deleteMovieByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateMovie(row) {
      const res = await findMovie({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.removie;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
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
          
      };
    },
    async deleteMovie(row) {
      this.visible = false;
      const res = await deleteMovie({ ID: row.ID });
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