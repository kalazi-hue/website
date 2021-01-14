<script>
	export default {
		onLaunch: function() {
			// #ifdef APP-NVUE
			plus.screen.lockOrientation('portrait-primary');

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
