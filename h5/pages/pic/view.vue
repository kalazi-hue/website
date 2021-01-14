<template>
	<view class="view-page">
			<view class="view-info">
				<view class="view-img" @tap="previewImage">
					<image :src="photoDetails.cover" mode="aspectFit" class="uni-image" fade-show="true"></image>
					<view class="tips icon">点击浏览全部（共{{ imgArr.length }}张）</view>
				</view>
				<view class="info">写真名称：{{ photoDetails.title }}</view>
				<view class="info">图片数量：{{ imgArr.length }} 张</view>
				<view class="info">发行时间：{{ photoDetails.createdAt }}</view>
			</view>
		</view>
</template>

<script>
	var t1 = new Date().getTime()
	import api from '@/api'
	export default {
		components: {
		},
		data() {
			return {
				total: 0,
				loading: true,
				photoDetails: [],
				imgArr: []
			}
		},
		onLoad(event) {
			// 目前在某些平台参数会被主动 decode，暂时这样处理。
			try {
				this.photoDetails = JSON.parse(decodeURIComponent(query));
			} catch (error) {
				this.photoDetails = JSON.parse(event.query);
			}
			let newImgArr = []
			newImgArr = this.photoDetails.content.split(",")
			for (let i in newImgArr) {
			    let newItem = newImgArr[i].replace(/\s+/g,"")
				this.imgArr.push(newItem)
			}
			this.previewImage()
		},
		methods: {
			previewImage: function(res) {    
				uni.previewImage({  
					current: 0,
					urls: this.imgArr,
					indicator: "number", // 图片指示器样式，可取值："default" - 底部圆点指示器； "number" - 顶部数字指示器； "none" - 不显示指示器。
					longPressActions: {
						itemColor: '#fff'
					}
				});  
			}
		}
	}
</script>

<style scoped>
.uni-image {
	width: 750upx;
    background-color: #000;
}
.tips{
	color: #fff;
	font-size: 26upx;
	margin: 10upx 0;
	text-align: center;
}
.view-title {
	color: #fff;
	font-size: 36upx;
	text-align: left;
	padding: 40upx 40upx 10upx;
}
.info {
	color: #999;
	font-size: 32upx;
	text-align: left;	
	padding: 10upx 40upx;
	flex-direction: row;
    align-items: flex-start;
}
</style>
