<template>
	<view class="tags-page">
		<view class="page-title">
			<image src="/static/images/icons/ico-hot.png" mode="widthFix" class="page-title-icons" style="width: 38rpx;"></image>
			<text class="title-txt">热门标签</text>
		</view>
		<view class="tag-list">
			<text class="tag" v-for="item in tagList" @click="toSearch(item.name)" :key="item.ID" v-if="tagList.length > 0 && item.status"> #{{ item.name }}</text>
			<text class="tag tag-loading" v-for="item in 1" v-else></text>
		</view>
	</view>
</template>

<script>
import api from '@/api'
export default {
  data () {
    return {
		tagList: [],
		searchText: ''
    }
  },
  created () {
	this.fetchTags()
  },
  methods: {
	toSearch (name) {
		uni.navigateTo({
			url: '/pages/video/search?tag=' + name
		});
	},
	// 获取标签列表
	fetchTags () {
		let params = {
			page: 1,
			pageSize: 15
		}
		uni.request({
			method: 'GET',
			url: `${ api.configApi }/tag/getTagList`,
			data: params,
			success: (res) => {
				this.tagList = res.data.data.list
			}
		});
	}
  }
}
</script>

<style scoped>
.tags-page {
	margin: 20rpx 30rpx;
	width: 690rpx;
}
.page-title {
	width: 300rpx;
	margin: 0 0 30rpx 0;
	padding: 0;
	align-items: center;
    flex-direction: row;
}
.page-title-icons {
	margin: 0 20rpx 0 0;
}
.label {
	color: #fff;
	margin: 0 0 15px 0;
	/* #ifndef APP-NVUE */
	display: flex;
	/* #endif */
	font-size: 24px;
	align-items: center;
}
.ico-hottag {
	width: 20px;
	height: 23px;
	margin-right: 10px;
}
.tag-list {
	/* #ifndef APP-NVUE */
	display: flex;
	/* #endif */
	justify-content: flex-start;
    flex-direction: row;
	flex-wrap: wrap;
    margin: 0;
	padding: 0;
}
.tag {
	border-width: 1px;
	border-style: solid;
	border-color: #d03fe5;
	font-size: 26rpx;
	text-align: center;
	margin: 0 10rpx 20rpx 0;
    lines: 1;
    text-overflow: ellipsis;
    background-color: #603d65;
	border-radius: 5px;
	/* #ifndef APP-NVUE */
	padding: 8rpx 10rpx;
	overflow: hidden;
	list-style: none;
	transition: all ease 0.5s;
	text-overflow: ellipsis;
	white-space: nowrap;
	/* #endif */
	
	/* #ifdef APP-NVUE */
	padding: 15rpx 10rpx;
	/* #endif */
}
.tag-loading {
	width: 130rpx;
	height: 50rpx;
    border-color: #5b3c60;
}
.tags-btn {
  font-size: 14px;
  color: #666;
  text-align: center;
}
</style>
