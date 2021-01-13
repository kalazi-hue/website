<template>
	<div>
	  <div class="ad-page">
	  	<a :href="img.playUrl" v-for="(img, index) in dataList" :key="index" target="_blank" :title="img.description" v-if="img.status">
	    	<img v-lazy="img.url" >  		
	  	</a>
	  </div>
	  <adData @adBottomData="getData" />		
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
.ad-page {
	margin-bottom: 100px;
}
</style>