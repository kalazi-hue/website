<template>
	<view>
		<special-banner :banner-list="bannerList" :swiper-config="swiperConfig" v-if="bannerList.length > 0"></special-banner>
		<view class="banner-loading" v-else>
			<view class="placeholder left-loading"></view>
			<view class="placeholder mid-loading"></view>
			<view class="placeholder right-loading"></view>
		</view>
	</view>
</template>

<script>
import api from '@/api'
import specialBanner from "@/components/ad/specialBanner.vue"
export default {
    components: {
	  specialBanner
    },
	data() {
	    return {
	        bannerList: [],
	        swiperConfig: {
	            indicatorDots: true,
	            indicatorColor: 'rgba(255, 255, 255, .2)',
	            indicatorActiveColor: 'rgba(178, 120, 236, 1.0)',
	            autoplay: false,
	            interval: 3000,
	            duration: 300,
	            circular: true,
	            previousMargin: '58rpx',
	            nextMargin: '58rpx'
	        }
	    }
	},
	created () {
		this.getAdConfigList()
	},
	methods:{
		getAdConfigList () {
			let params = {}
			params.page = 1
			params.pageSize = 200	  
			uni.request({
				method: 'GET',
				url: `${ api.configApi }/adConfig/getAdConfigList`,
				data: params,
				success: (res) => {
					const adList = res.data.data.list
					let subLunboList = [],subTopList = [],subBottomList = [],subLeftList = [],subRightList = []
					for (let i = 0; i < adList.length; i++) {
						if (adList[i].position.indexOf('轮播') > -1) {
							subLunboList.push(adList[i])
						} else if (adList[i].position.indexOf('顶部') > -1) {
							subTopList.push(adList[i])
						} else if (adList[i].position.indexOf('底部') > -1) {
							subBottomList.push(adList[i])
						} else if (adList[i].position.indexOf('左飘') > -1) {
							subLeftList.push(adList[i])
						} else if (adList[i].position.indexOf('右飘') > -1) {
							subRightList.push(adList[i])
						}
					}
					// console.log(subLunboList);
					this.bannerList = subLunboList
					uni.setStorageSync('subTopList', JSON.stringify(subTopList))
					uni.setStorageSync('subBottomList', JSON.stringify(subBottomList))
				},
				fail: (err) => {
					console.log(err);
				},
				complete: (e) => {
				}
			});
		}
	}
}
</script>

<style scoped>
.banner-loading {
	/* #ifndef APP-NVUE */
	display: flex;
	/* #endif */	
	justify-content: space-around;
	padding: 30rpx 0 60rpx 0;
    position: relative;
    z-index: 2;
}
.left-loading,
.right-loading {
	width: 30rpx;
	height: 300rpx;	
	margin-top: 30rpx;
	/* #ifndef APP-NVUE */
	border-radius: 6px;
	/* #endif */
}
.mid-loading {
	width: 630rpx;
	height: 360rpx;
	z-index: 200;
	/* #ifndef APP-NVUE */
	border-radius: 6px;
	/* #endif */
}
</style>
