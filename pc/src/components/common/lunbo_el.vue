<template>
  <el-carousel :interval="2000" :autoplay='false' type="card" height="300px" @click="change">
    <el-carousel-item v-for="(item,index) in dataList" :key="index">
      <a :href="item.playUrl" target="_blank">
        <img v-lazy="item.url" alt="" />          
      </a>
    </el-carousel-item>
    <adData @lunboData="getData" /> 
  </el-carousel>
</template>

<script>
import adData from '@/components/common/adData'
export default {
  components: {
    adData
  },
  data () {
    return {
      dataList: []
    }
  },
  created () {
  },
  methods: {
    getData (data) {
      this.dataList = data
    },
    change(e) {
      console.log(e)
    }
  }
}
</script>

<style>
  .el-carousel__item {
    width: 1000px!important;
    border-radius: 14px;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
    border-radius: 14px;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
    border-radius: 14px;
  }
  .el-carousel__item--card.is-active {
    transform: translateX(100px) scale(1)!important;
  }
</style>