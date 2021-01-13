<script>
	export default {
		onLaunch: function() {
			console.log('App Launch');
			// #ifdef APP-NVUE
			plus.screen.lockOrientation('portrait-primary');

			let appid = plus.runtime.appid;
			if (appid && appid.toLocaleLowerCase() != "hbuilder") {
				uni.request({
					url: 'https://uniapp.dcloud.io/update', //检查更新的服务器地址
					data: {
						appid: plus.runtime.appid,
						version: plus.runtime.version
					},
					success: (res) => {
						console.log('success', res);
						if (res.statusCode == 200 && res.data.isUpdate) {
							let openUrl = plus.os.name === 'iOS' ? res.data.iOS : res.data.Android;
							// 提醒用户更新
							uni.showModal({
								title: '更新提示',
								content: res.data.note ? res.data.note : '是否选择更新',
								success: (showResult) => {
									if (showResult.confirm) {
										plus.runtime.openURL(openUrl);
									}
								}
							})
						}
					}
				})
			}

			// 引入字体图标，兼容NVUE
			const domModule = weex.requireModule('dom')  
			domModule.addRule('fontFace', {  
			   fontFamily: 'myfont',  
			   src: "url('/static/fonts/iconfont.ttf')"  
				// src: "url('https://at.alicdn.com/t/font_2272660_7790wz50b4g.ttf')"  
			});  
			
			// 重置底部安全区域背景
				var Color = plus.android.importClass("android.graphics.Color");    
				    plus.android.importClass("android.view.Window");    
				    var mainActivity = plus.android.runtimeMainActivity();    
				    var window_android = mainActivity.getWindow();
					window_android.setNavigationBarColor(Color.argb(255,20,20,20,));
			// #endif  
		},
		onShow: function() {
			// console.log('App Show')
		},
		onHide: function() {
			// console.log('App Hide')
		}
	}
</script>

<style>
@import './static/css/common.css';
@import './static/css/iconfont.css';
.myfont{
    font-family:"myfont" !important;
    font-size:16px;
	font-style:normal;
    /* #ifndef APP-NVUE */
    -webkit-font-smoothing: antialiased;
    -webkit-text-stroke-width: 0.2px;
    -moz-osx-font-smoothing: grayscale;
    /* #endif */
}
</style>
