<template>
	<view class="video-list-home">
		<view class="page-title">
			<text class="title-txt">{{ title }}</text>
			<image src="/static/images/icons/bg-video.png" mode="widthFix" class="page-title-icons" style="width: 90px;height: 4px;" v-if="showLatestList"></image>
			<image src="/static/images/icons/bg-hot.png" mode="widthFix" class="page-title-icons" style="width: 90px;height: 4px;" v-if="showRecommendList"></image>
		</view>
		<view class="video-list" v-if="videoList.length > 0">
			<view 
				v-for="(newsItem, index) in videoList" 
				@click="goDetail(newsItem.ID)" 
				 v-if="newsItem.status"
				:key="index" 
				class="list-item" 
				:class="{'list-item-even': 
				index % 2 === 0}">
				<view class="list-top">
					<image :src="newsItem.cover" mode="scaleToFill" class="list-img"></image>
					<view class="shadow"></view>
					<text class="list-date">{{newsItem.createdAt.split(' ')[0]}}</text>
				</view>
				<text class="list-title">{{newsItem.title}}</text>
				<view class="list-bottom">
					<view class="list-playCount">
						<image src="/static/images/icons/ico-eye.png" mode="widthFix" class="ico-eye"></image>
						<text class="text-views">{{newsItem.playCount}}次播放</text>
					</view>
					<view class="list-star">
						<image src="/static/images/icons/ico-like.png" mode="widthFix" class="ico-like"></image>
						<text class="text-views">{{newsItem.star}}</text>
					</view>
				</view>
			</view>
		</view>   
		<videoListLoading v-else /> <!-- 骨架屏 -->
	</view>
</template>

<script>
import api from '@/api'
import videoListLoading from '@/components/videoListLoading.vue';
export default {
    components: {
	  videoListLoading
    },
	props: {
		showLatestList: Boolean,
		showRecommendList: Boolean
	},
	data() {
	    return {
			title: '今日更新',
	        videoList: [],
	        params: {
				page: 1,
				pageSize: 6
	        }
	    }
	},
	created () {
		this.showLatestList ? this.getMovieList('getLatestMovieList') : ''
		this.showRecommendList ? this.getMovieList('getRecommendMovieList') : ''
		this.showRecommendList ? this.title = '热门推荐' : this.title = '今日更新'
	},
	methods: {
		getMovieList(type) {
			uni.request({
				method: 'GET',
				url: `${ api.configApi }/movie/${type}`,
				data: this.params,
				success: (res) => {
					this.videoList = res.data.data.list
				}
			});
		},
		goDetail (vid) {
			uni.navigateTo({
				url: '/pages/movie/detail?videoId=' + vid
			});
		}
	}
}
</script>

<style scoped>
@import '@/static/css/videoList.css';
</style>
