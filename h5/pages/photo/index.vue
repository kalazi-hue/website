<template>
  <view class="photo-page">
	  
	<view class="page-title">
		<text class="title-txt">写真图片</text>
		<image src="/static/images/icons/bg-video.png" mode="widthFix" class="page-title-icons" style="width: 90px;height: 4px;"></image>
	</view>
			  
	<!--瀑布流列表组件 -->
	<WaterfallsFlow :wfList='dataList' @itemTap="itemTap" />
	
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
import api from '@/api'
import unidrawer from '@/components/uniapp/uni-drawer.vue';
import leftMenu from '@/components/leftMenu.vue';
import WaterfallsFlow from '@/components/uniapp/WaterfallsFlow/WaterfallsFlow.vue'

  export default {
    components: {
		unidrawer,
		leftMenu,
		WaterfallsFlow
    },
    data() {
		return {
			mask: true,
			left: 'left',
			drawWid: 200,
			mode: 'left',
			requestParams: {
				page: 1,
				pageSize: 10
			},
			dataList:[],
			requiredData:[
				{
					image:'https://cdn.pixabay.com/photo/2015/10/21/03/44/beach-998788_960_720.jpg',
					avatar:'https://cdn.pixabay.com/user/2015/10/12/02-06-28-605_250x250.jpg',
					nickName:'rFonz'
				},
				{
					image:'https://cdn.pixabay.com/photo/2020/06/08/22/10/couple-5276245_960_720.jpg',
					avatar:'https://cdn.pixabay.com/user/2020/04/02/05-34-12-636_250x250.jpg',
					nickName:'VlacK_HammeR'
				},
				{
					image:'https://cdn.pixabay.com/photo/2020/06/11/21/25/imp-5288348_960_720.jpg',
					avatar:'https://cdn.pixabay.com/user/2020/06/10/08-34-47-445_250x250.jpg',
					nickName:'Alexas_Fotos'
				},
				{
					image:'https://cdn.pixabay.com/photo/2020/06/10/01/32/cumulonimbus-5280578_960_720.jpg',
					avatar:'https://pixabay.com/static/img/profile_image_dummy.svg',
					nickName:'phil454'
				},
				{
					image:'https://cdn.pixabay.com/photo/2020/06/06/04/16/table-5265190_960_720.jpg',
					avatar:'https://cdn.pixabay.com/user/2020/06/13/19-29-33-253_250x250.jpg',
					nickName:'NickyPe'
				},
				{
					image:'https://cdn.pixabay.com/photo/2016/03/26/13/09/notebook-1280538_960_720.jpg',
					avatar:'https://cdn.pixabay.com/user/2017/08/03/10-16-32-389_250x250.png',
					nickName:'freephotocc'
				},
				{
					image:'https://cdn.pixabay.com/photo/2020/06/08/05/28/vortex-5273023_960_720.png',
					avatar:'https://cdn.pixabay.com/user/2015/12/02/23-35-18-266_250x250.png',
					nickName:'GDJ'
				},
				{
					image:'https://cdn.pixabay.com/photo/2020/06/03/16/30/wave-5255593_960_720.jpg',
					avatar:'https://cdn.pixabay.com/user/2020/05/23/12-55-39-456_250x250.jpg',
					nickName:'pasja1000'
				},
				{
					image:'https://cdn.pixabay.com/photo/2020/05/28/08/00/girl-5230306_960_720.jpg',
					avatar:'https://cdn.pixabay.com/user/2020/01/29/07-13-31-807_250x250.jpg',
					nickName:'Syaibatulhamdi'
				},
				{
					image:'https://cdn.pixabay.com/photo/2020/06/10/09/19/cake-5281948_960_720.jpg',
					avatar:'https://cdn.pixabay.com/user/2018/01/25/18-40-42-65_250x250.jpg',
					nickName:'matthiasboeckel'
				}
			]
		}
    },
	onLoad() {
		// 模拟请求数据
		this.loadData()
	},
	onReachBottom(){
		// 模拟触底刷新
		// setTimeout(()=>{
		// 	this.requiredData.push(...this.dataList);
		// },1000)
		this.loadData()
	},
	onPullDownRefresh(){
		// 模拟上拉刷新
		// const newList = this.requiredData.reverse();
		// this.dataList = newList;
		this.requestParams.page = 1
		this.loadData()
		uni.stopPullDownRefresh();
	},
    methods: {
		openDrawer () {
			this.$refs.draw.open();
		},
		closeDrawer () {
			this.$refs.draw.close();		
		},
		itemTap(item){
		},
		loadData(refresh) {
			let data_List = []
			uni.request({
				method: 'GET',
				url: `${ api.configApi }/photoAlbum/getPhotoAlbumList`,
				data: this.requestParams,
				success: (res) => {
					const data = res.data.data.list
					if (data.length <= 0) {
						return false
					}
					this.dataList = this.dataList.concat(data);
					this.requestParams.page ++
				}
			});
		}
    },
	onNavigationBarButtonTap(e) { // 监听原生标题栏按钮点击事件
		let that = this
		if (e.index === 0) { // 顶部左边菜单
			 that.openDrawer()
		} else if (e.index === 1) {  // 顶部右边搜索
			uni.navigateTo({
				url: '/pages/video/search'
			});
		}
	}
  }
</script>

<style>
@import '@/static/css/tabbar.css';	
</style>

<style lang="scss" scoped>
	.page{
		min-height: 100vh;
	}
</style>