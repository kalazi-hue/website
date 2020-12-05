<template>
 <div class="big">
         <div class="card">
             <div class="car-left">
                 <div>
                    <span class="card-img"> <img :src="userInfo.headerImg"  alt="" > </span>   
                    <div class="text"><h4>早安，管理员， 请开始您一天的工作吧！</h4>
                        <p class="tips-text">
                            <i class="el-icon-sunny el-icon"></i>
                            <span>今日晴，0℃ - 10℃，天气寒冷，注意添加衣物。</span>
                        </p>
                    </div>
                 </div>
             </div>

         </div>


 </div>

</template>

<script>
import { mapGetters } from 'vuex'
export default {

   computed: {
    ...mapGetters('user', ['userInfo'])
  },

    methods:{
      toTarget(name){
          this.$router.push({name})
      }
    },
}
</script>

<style lang="scss" scoped>
    .big{
        margin:100px 0 0 0;
        padding-top: 0;
        background-color: rgb(243,243,243);
        padding-top: 15px;
        .top{
            width: 100%;
            height: 360px;
            margin-top: 20px;
            overflow: hidden;
            .chart-container{
                position: relative;
                width: 100%;
                height: 100%;
                padding: 20px;
                background-color: #fff;
            }
        }
        .mid{
            width: 100%;
            height: 380px;
            .chart-wrapper {
                height: 340px;
                background: #fff;
                padding: 16px 16px 0;
                margin-bottom: 32px;
            }
        }
        .bottom{
            width: 100%;
            height: 300px;
            // margin: 20px 0;
            .el-row{
                margin-right: 4px !important;
            }
            .chart-player{
                width: 100%;
                height: 270px;
                padding: 10px;
                background-color: #fff;
            }
        }
    }

</style>
