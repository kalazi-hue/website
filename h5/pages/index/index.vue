<template>
	<view>
		
		<lunbo /><!-- 轮播组件 -->
		
		<adBanner :adTop="true" /><!--顶部横幅 -->
		
		<video-list-home :showLatestList="true" /> <!--最近更新 -->
		
		<video-list-home :showRecommendList="true" /> <!--推荐 -->
		
		<adBanner :adBottom="true" /><!--底部横幅 -->
	
		<notice /><!--公告弹窗 -->
		
		<!--左边栏 -->
		<unidrawer ref="draw" :mask="mask" :width="drawWid" :mode='mode'>
			<image src="/static/images/icons/ico-close.png" 
				@click="closeDrawer()" 
				mode="widthFix" 
				style="width: 60rpx;height: 60rpx;position: absolute;left: 160px;top: 30rpx;z-index: 99;">
			</image>
			<leftMenu />
		</unidrawer>
	</view>
</template>

<script>
	import lunbo from '@/components/ad/lunbo.vue';
	import adBanner from '@/components/ad/adBanner.vue';
	import videoListHome from '@/components/videoListHome';
	import notice from '@/components/notice';
	import unidrawer from '@/components/uniapp/uni-drawer.vue';
	import leftMenu from '@/components/leftMenu.vue';
  export default {
    components: {
	  lunbo,
	  adBanner,
	  videoListHome,
	  notice,
	  unidrawer,
	  leftMenu
    },
    data() {
		return {
			mask: true,
			left:'left',
			drawWid:200,
			maskClick:true,
			mode:'left'
		}
    },
    methods: {
		openDrawer () {
			this.$refs.draw.open();
		},
		closeDrawer () {
			this.$refs.draw.close();		
		}
    },
	onNavigationBarButtonTap(e) { // 监听原生标题栏按钮点击事件
		console.log("点击了自定义按钮", e.index);
		if (e.index === 0) { // 顶部左边菜单
			 this.openDrawer()
		} else if (e.index === 1) {  // 顶部右边搜索
			uni.navigateTo({
				url: '/pages/movie/search'
			});
		}
	},
    onPullDownRefresh() {
        setTimeout(function () {
            uni.stopPullDownRefresh();
        }, 1000);
    }
  }
</script>
