<template>
	<div>
	  <div class="ad-page">
	  	<a :href="img.playUrl" v-for="(img, index) in dataList" :key="index" target="_blank" :title="img.description" v-if="img.status">
	    	<img v-lazy="img.url" >  		
	  	</a>
	  	
      <!-- 数据加载前骨架占位 -->
	  	<a href="" v-if="dataList.length <= 0""  v-for="(img, index) in 6" :key="index" class="placeholder"></a>
	  </div>
	  <adData @adTopData="getData" />		
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
			dataList: []
    }
  },
	methods: {
		getData (data) {
			this.dataList = data
		}
	}
}
</script>

<style scoped lang="scss">
@import '@/assets/css/ad.scss';
</style>