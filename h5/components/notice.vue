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
			noticeContent: '',
			androidUrl: null
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
					if (uni.getSystemInfoSync().platform === 'android') {
						if (that.configList[0].notice2.indexOf('.apk') > -1) { // 安卓apk存在
							that.androidUrl = that.configList[0].notice2
							//#ifndef H5
							that.AndroidCheckUpdate(that.androidUrl) // 检测更新
							//#endif
						}
					}
				}
			});
		},
		AndroidCheckUpdate (androidUrl) {  // 安卓提醒用户更新 
			let newVersion = 0
			newVersion = androidUrl.split('_')[1].split('.apk')[0]
			plus.runtime.getProperty(plus.runtime.appid, function(wgtinfo){
				if (newVersion != wgtinfo.version) {
					uni.showModal({
						title: '发现新版本',
						confirmText: '本地更新',
						cancelText: '浏览器下载',
						content: '目前版本过低，请更新至最新版本 v' + newVersion,
						success: function (res) {
							if (res.confirm) {
								_this.doUpData()
							} else if (res.cancel) {
								plus.runtime.openURL(androidUrl);
							}
						}
					});
				}
			});
		},
		doUpData() { // 执行更新下载
			uni.showLoading({
				title: '更新中,请等待……'
			})
			uni.downloadFile({//执行下载
				url: this.androidUrl, //下载地址
				success: downloadResult => {//下载成功
					uni.hideLoading();
					if (downloadResult.statusCode == 200) {
						uni.showModal({
							title: '',
							content: '更新成功，确定现在重启吗？',
							confirmText: '重启',
							confirmColor: '#EE8F57',
							success: function(res) {
								if (res.confirm == true) {
									plus.runtime.install(//安装
										downloadResult.tempFilePath, {
											force: true
										},
										function(res) {
											utils.showToast('更新成功，重启中');
											plus.runtime.restart();
										}
									);
								}
							}
						});
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
