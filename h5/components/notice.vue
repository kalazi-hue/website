<template>
	<view class="notice">
		<uni-popup ref="popup" type="dialog">
			<image src="/static/images/bg-notice.png" mode="aspectFill" class="bg-notice"></image>
			<view class="notice-content">
				<view class="notice-msg">
					{{ noticeContent }}
				</view>
				<image src="/static/images/bg-notice-btn.png" mode="aspectFill" class="btn-notice" @tap="close"></image>
			</view>
		</uni-popup>
	</view>
</template>

<script>
import api from '@/api'
import uniPopup from '@/components/uniapp/uni-popup/uni-popup.vue'
import uniPopupMessage from '@/components/uniapp/uni-popup/uni-popup-message.vue'
import uniPopupDialog from '@/components/uniapp/uni-popup/uni-popup-dialog.vue'
export default {
	components: {
		uniPopup,
		uniPopupMessage,
		uniPopupDialog
	},
  data() {
		return {
			configList: [],
			noticeContent: ''
		}
  },
	mounted () {
		this.getSystemConfigList()
	},
  methods: {
		open(){
			 this.$refs.popup.open()
		},
		close(){
			 this.$refs.popup.close()
		},
		getSystemConfigList(type) {
			const that = this
			uni.request({
				method: 'GET',
				url: `${ api.configApi }/systemConfig/getSystemConfigList`,
				data: this.params,
				success: (res) => {
					this.configList = res.data.data.list
					uni.setStorageSync('configList', this.configList)
					if (that.configList[0].notice !== '') {
						that.open()
						that.noticeContent = that.configList[0].notice
					} else {
						that.close()						
					}
				}
			});
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
  }
}
</script>

<style scoped>
	.bg-notice {
		height: 650rpx;
		width: 620rpx;
		position: relative;
	}
	.notice-content {
		height: 320rpx;
		position: absolute;
		bottom: 10rpx;
		left: 0;
	}
	.notice-msg {
		width: 520rpx;
		height: 190rpx;		
		padding: 30rpx 50rpx 0 50rpx;
		font-size: 28rpx;
	}
	.btn-notice {
		height: 70rpx;
		width: 290rpx;
		left: 160rpx;
	}
</style>
