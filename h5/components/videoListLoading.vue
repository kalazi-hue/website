<template>
	<view class="video-list">
		<view v-for="(newsItem, index) in 6" :key="index" class="list-item" :class="{'list-item-even': index % 2 === 0}">
			<view class="list-top">
				<image src="" mode="scaleToFill" class="list-img"></image>
				<view class="shadow"></view>
				<text class="list-date"></text>
			</view>
			<view class="list-title list-title-loading"></view>
			<view class="list-bottom list-bottom-loading">
			</view>
		</view>
	</view>  
</template>

<style scoped>
@import '@/static/css/videoList.css';	
.list-bottom-loading {
	height: 40rpx;
}
.list-title-loading {
	margin-top: 0;
}
.page-title {
    margin-left: 30rpx;
}
</style>
