<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">        
<!--        <el-form-item>-->
<!--          <el-button @click="onSubmit" type="primary">查询</el-button>-->
<!--        </el-form-item>-->
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

    
    <el-table-column label="影片ID" prop="movieId" width="80"></el-table-column>

    <el-table-column label="影片标题" width="260" show-overflow-tooltip>
        <template slot-scope="scope">
            <p>{{scope.row.title}}</p>
        </template>
    </el-table-column>

    <el-table-column label="播放链接" show-overflow-tooltip>
        <template slot-scope="scope">
            <p>{{scope.row.playUrl}}</p>
        </template>
    </el-table-column>
    
    <el-table-column label="反馈类型" prop="type" width="100"></el-table-column>
    
    <el-table-column label="创建日期"  width="100">
             <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
        </el-table-column>
      <el-table-column label="操作" width="160">
        <template slot-scope="scope">
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deletePlayFeedback(scope.row)">确定</el-button>
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
         <el-form-item label="movieId字段:"><el-input v-model.number="formData.movieId" clearable placeholder="请输入"></el-input>
      </el-form-item>
       
         <el-form-item label="title字段:">
            <el-input v-model="formData.title" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="playUrl字段:">
            <el-input v-model="formData.playUrl" clearable placeholder="请输入" ></el-input>
      </el-form-item>
       
         <el-form-item label="type字段:">
            <el-input v-model="formData.type" clearable placeholder="请输入" ></el-input>
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
    createPlayFeedback,
    deletePlayFeedback,
    deletePlayFeedbackByIds,
    updatePlayFeedback,
    findPlayFeedback,
    getPlayFeedbackList
} from "@/api/PlayFeedback";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/date";
import infoList from "@/mixins/infoList";
export default {
  name: "PlayFeedback",
  mixins: [infoList],
  data() {
    return {
      listApi: getPlayFeedbackList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
            movieId:0,
            title:"",
            playUrl:"",
            type:"",
            
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
        const res = await deletePlayFeedbackByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updatePlayFeedback(row) {
      const res = await findPlayFeedback({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.replayFeedback;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
          movieId:0,
          title:"",
          playUrl:"",
          type:"",
          
      };
    },
    async deletePlayFeedback(row) {
      this.visible = false;
      const res = await deletePlayFeedback({ ID: row.ID });
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
          res = await createPlayFeedback(this.formData);
          break;
        case "update":
          res = await updatePlayFeedback(this.formData);
          break;
        default:
          res = await createPlayFeedback(this.formData);
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