<template>
	<view class="ad-banner" :class="{'ad-banner-bottom' : adBottom === true }">
		<!-- 
		设备类型 equipment :1:PC,2:H5 
		位置 location 1:顶部,2:底部,3:左边,4:右边
		-->
		<view class="list">
			<view v-for="(item, index) in adList" :key="index" class="item" v-if="item.status">
				<image :src="item.url" :alt="item.description" :title="item.description" @tap="open_url(item.playUrl, index)" mode="widthFix" fade-show="true" class="uni-image"></image>
			</view>
		</view>
	</view>
</template>
<script>
let _self
import api from '@/api'
export default {
	props:{
		adTop:{
			type: Boolean,
			default: false
		},
		adBottom:{
			type: Boolean,
			default: false
		}
	},
	data () {
		return {
			adList: []
		}
	},
	methods: {
		compare(property){
		    return function(a,b){
		        var value1 = a[property];
		        var value2 = b[property];
		        return value1 - value2;
		    }
		},
		open_url (url) {
			if (url.indexOf('http') > -1) {
			   //#ifdef H5
					window.open(url)
			   //#endif
			   
			   //#ifndef H5
				   plus.runtime.openURL(url); 
			   //#endif				
			} else {
				uni.navigateTo({
					url: url
				});
			}
		}
	},
	created () {
		setTimeout(() => {
			let value = null
			if (this.adTop) {
				value = uni.getStorageSync('subTopList');
				this.adList = value
			} else {
				value = uni.getStorageSync('subBottomList');
				this.adList = value
			}
		}, 1500)
	}
}
</script>

<style scoped> 
.ad-banner-bottom {
    padding-bottom: 10px;	
}
.list {
	flex-wrap: nowrap;
	/* #ifndef APP-NVUE */
	display: block;
	/* #endif */
}
.item {
	padding-bottom: 2px;
	width: 100%;
	/* #ifndef APP-NVUE */
	border-radius: 0;
	/* #endif */
	position: relative;
}
.uni-image {
	width: 100%;
	height: 92upx!important;
	/* #ifndef APP-NVUE */
	border-radius: 0;
	/* #endif */
	position: relative;
	z-index: 2;
}
.uni-image {
	width: 100%;
	height: 100%;
	/* #ifndef APP-NVUE */
	max-height: 80px;
	/* #endif */
}
</style>
