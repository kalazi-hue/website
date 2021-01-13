<template>
	<view class="wf-item-page" @click="goDetail(item)">
		<image :src="item.cover" mode="widthFix" class="item-img" />
		<view class="item-info">
			<view class="info-title">{{item.title}}</view>
		</view>
	</view>
</template>

<script>
	export default {
		props:{
			item:{
				type:Object,
				require:true
			}
		},
		methods: {
			goDetail(item) {
				console.log(JSON.stringify(item))
				uni.navigateTo({
					url: '/pages/photo/view?query=' +JSON.stringify(item)
				});
			}
		}
	}
</script>

<style lang="scss" scoped>
	.wf-item-page{
		overflow: hidden;
	}
	
	.item-img{
		width: 100%;
	}
	
	.item-info{
		display: flex;
		align-items: center;
		padding: 5px;
	}
	
	.info-title{
		font-size: 12px;
		color: #fff;
	}
</style>
