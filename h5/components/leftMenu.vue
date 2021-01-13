<template>
	<view>
		<view class="menus">
			<view class="item">
				<image src="/static/images/icons/ico-app.png" mode="widthFix" class="ico-img"></image>
				<text class="menu-title" @click="open_url(configList[0].exchangeGroup)">APP下载</text>
			</view>
			<view class="item">
				<image src="/static/images/icons/ico-group.png" mode="widthFix" class="ico-img"></image>
				<text class="menu-title" @click="open_url(configList[0].exchangeGroup)">用户交流群</text>
			</view>
			<view class="item">
				<image src="/static/images/icons/ico-github.png" mode="widthFix" class="ico-img"></image>
				<text class="menu-title" @click="open_url(configList[0].gitHub)">Github地址</text>
			</view>
			<view class="item">
				<image src="/static/images/icons/ico-email.png" mode="widthFix" class="ico-img"></image>
				<text class="menu-title" @click="open_url(configList[0].email)">商务邮箱</text>
			</view>
			<view class="item">
				<image src="/static/images/icons/ico-backup.png" mode="widthFix" class="ico-img"></image>
				<text class="menu-title" @click="open_url(configList[0].alternateAddress)">备用地址</text>
			</view>
			<!-- <view class="item">
				<image src="/static/images/icons/ico-yinsi.png" mode="widthFix" class="ico-img"></image>
				<text class="menu-title" @click="open_url('/pages/')">隐私政策</text>
			</view> -->
			<view class="item">
				<image src="/static/images/icons/ico-contact.png" mode="widthFix" class="ico-img"></image>
				<text class="menu-title" @click="open_url(1)">联系我们</text>
			</view>
			<!-- <view class="item">
				<image src="/static/images/icons/ico-faq.png" mode="widthFix" class="ico-img"></image>
				<text class="menu-title" @click="open_url('')">常见问题</text>
			</view> -->
		</view>
		
	</view>
	
</template>

<script>
export default {
	components: {
	},
    data() {
		return {
			configList: []
		}
    },
	created() {
		this.configList = uni.getStorageSync('configList')
	},
    methods: {
		open_url (url) {
			console.log(url)
			const reg = RegExp(/http/);
			if (reg.test(url)) {
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
    }
}
</script>

<style scoped>
	.menus {
		padding: 100rpx 0 0 60rpx;
	}
	.item {
		margin-top: 60rpx;	
		flex-direction: row;
		align-items: center;
	}
	.ico-img {
		width: 50rpx;
		height: 50rpx;
		margin-right: 20rpx;
		/* #ifndef APP-PLUS */
		vertical-align: -10px;
		/* #endif */
	}
</style>
