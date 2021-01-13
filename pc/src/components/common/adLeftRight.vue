<template>
	<div>	
	  <div class="ad-side" v-if="isSHowSide">
	  	<div class="left">
	  		<i class="el-icon-circle-close" @click="close"></i>
		  	<a :href="img.playUrl" v-for="(img, index) in leftList" target="_blank" :title="img.description" v-if="img.status">
		    	<img v-lazy="img.url" >  		
		  	</a>	  		

	      <!-- 数据加载前骨架占位 -->
		  	<!-- <a href="" v-if="leftList.length <= 0" ></a> -->
	  	</div>
	  	<div class="right">
	  		<i class="el-icon-circle-close" @click="close"></i>
		  	<a :href="item.playUrl" v-for="(item, i) in rightList" target="_blank" :title="item.description" v-if="img.status">
		    	<img v-lazy="item.url" >  		
		  	</a> 		
	      <!-- 数据加载前骨架占位 -->
		  	<!-- <a href="" v-if="rightList.length <= 0" ></a> -->
	  	</div>
	  </div>
	  <adData @adLeftData="getDataLeft" />		
	  <adData @adRightData="getDataRight" />		
  </div>
</template>

<script>
import adData from '@/components/common/adData'
export default {
	components: {
		adData
	},
  data () {
    return {
    	isSHowSide: false,
			leftList: [],
			rightList: []
    }
  },
	methods: {
		getDataLeft (data) {
			if (data.length > 0) {
				this.leftList = data
				// this.isSHowSide = true 
			}
		},
		getDataRight (data) {
			this.rightList = data
		},
		close () {
			this.isSHowSide = false
		}
	}
}
</script>
 
<style scoped lang="scss">
@import '@/assets/css/ad.scss';
</style>