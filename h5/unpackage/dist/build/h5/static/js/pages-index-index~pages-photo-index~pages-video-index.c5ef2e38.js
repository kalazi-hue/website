(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-index-index~pages-photo-index~pages-video-index"],{"0736":function(t,n,e){"use strict";e.r(n);var a=e("173e"),o=e("9615");for(var i in o)"default"!==i&&function(t){e.d(n,t,(function(){return o[t]}))}(i);e("7551e"),e("df7c");var r,c=e("f0c5"),l=Object(c["a"])(o["default"],a["b"],a["c"],!1,null,"daad558c",null,!1,a["a"],r);n["default"]=l.exports},"0cba":function(t,n,e){"use strict";var a=e("27bb");e.d(n,"a",(function(){return a["a"]})),e.d(n,"b",(function(){return a["b"]})),e.d(n,"c",(function(){return a["c"]}))},1148:function(t,n,e){"use strict";var a=e("a691"),o=e("1d80");t.exports="".repeat||function(t){var n=String(o(this)),e="",i=a(t);if(i<0||i==1/0)throw RangeError("Wrong number of repetitions");for(;i>0;(i>>>=1)&&(n+=n))1&i&&(e+=n);return e}},"173e":function(t,n,e){"use strict";var a;e.d(n,"b",(function(){return o})),e.d(n,"c",(function(){return i})),e.d(n,"a",(function(){return a}));var o=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("v-uni-view",{staticClass:"hx-navbar",style:{width:t.screenWidth+"px"}},[e("v-uni-view",{staticClass:"hx-navbar__content",class:{"hx-navbar__fixed":t.conf.fixed,"hx-navbar__shadow":t.conf.shadow,"hx-navbar__border":t.conf.border},style:{width:t.screenWidth+"px","background-color":t.bgIsLine?"":t.backgroundColorRgba,"background-image":t.bgIsLine?t.backgroundColorRgba:"",height:t.navbarHeight,color:t.txtColor}},[t.backgroundImage?e("v-uni-view",{staticClass:"hx-navbar__content__imgctn",style:{opacity:t.bgImgTransparent[1]}},[e("v-uni-image",{staticClass:"hx-navbar__imgctn__img",style:{height:t.conf.height+t.statusBarHeight+"px",width:t.screenWidth+"px"},attrs:{src:t.backgroundImageEnd,mode:"scaleToFill"}})],1):t._e(),t.backgroundImage?e("v-uni-view",{staticClass:"hx-navbar__content__imgctn",style:{opacity:t.bgImgTransparent[0]}},[e("v-uni-image",{staticClass:"hx-navbar__imgctn__img",style:{height:t.conf.height+t.statusBarHeight+"px",width:t.screenWidth+"px"},attrs:{src:t.backgroundImage,mode:"scaleToFill"}})],1):t._e(),e("v-uni-view",{staticClass:"hx-navbar__content__main",style:{height:t.conf.height+"px",fontSize:t.conf.fontSize}},[t.conf.back?e("v-uni-view",{staticClass:"hx-navbar__content__main_back",on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.onBack.apply(void 0,arguments)}}},[e("v-uni-text",{staticClass:"hxicon hx-navbar__icon",class:{"hx-navbar__icontran":0==t.bgTransparent},style:{color:t.txtColor}},[t._v("")]),t.conf.backTxt?e("v-uni-text",{staticClass:"hx_font_size"},[t._v(t._s(t.conf.backTxt))]):t._e()],1):t._e(),t.conf.leftSlot?e("v-uni-view",{staticClass:"hx-navbar__content__main_left",staticStyle:{padding:"0",margin:"0 13px 0 0"}},[t.conf.leftSlotSwitch&&0==t.slotSwitchOpacity?t._t("leftSwitch"):t._t("left")],2):t._e(),t.conf.leftButton?e("v-uni-view",{staticClass:"hx-navbar__content__main_left",style:{padding:t.conf.back||t.conf.leftSlot?"0":"0 9px 0 13px;"}},[t._l(t.conf.leftButton,(function(n,a){return[e("v-uni-view",{key:a+"_0",staticClass:"hx-navbar__content__main_left_btn",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.onClickBtn(n)}}},[n.position&&"left"==n.position?e("v-uni-text",{staticClass:"hx-navbar__icon",class:t.conf.font,style:{color:n.color?n.color:t.txtColor,marginRight:n.txt?"2px":"0"}},[t._v(t._s(t.iconHandle(n.icon)))]):t._e(),n.txt?e("v-uni-text",{staticClass:"hx-navbar__content__main_lefticon_txt hx_font_size",style:{color:n.color?n.color:t.txtColor}},[t._v(t._s(n.txt))]):t._e(),n.position&&"left"==n.position?t._e():e("v-uni-text",{staticClass:"hx-navbar__icon",class:t.conf.font,style:{color:n.color?n.color:t.txtColor,marginLeft:n.txt?"2px":"0"}},[t._v(t._s(t.iconHandle(n.icon)))])],1)]}))],2):t._e(),t.conf.search||t.conf.maxSlot?t._e():e("v-uni-view",{staticClass:"hx-navbar__content__main_center"},[e("v-uni-view",{staticClass:"hx-navbar__content__main_center_flex"}),e("v-uni-view",{staticClass:"hx-navbar__content__main_center_txt"},[!t.conf.centerSlot&&t.title?e("v-uni-text",{staticClass:"hx_font_size hx_text_overflow"},[t._v(t._s(t.title))]):t._e(),t.conf.centerSlotSwitch&&0==t.slotSwitchOpacity?t._t("centerSwitch"):t._t("center")],2),e("v-uni-view",{staticClass:"hx-navbar__content__main_center_flex"})],1),!t.conf.search&&t.conf.maxSlot?e("v-uni-view",{staticClass:"hx-navbar__content__main_center",staticStyle:{overflow:"hidden"}},[t.conf.maxSlotSwitch&&0==t.slotSwitchOpacity?t._t("maxSwitch"):t._t("max")],2):t._e(),t.conf.search?e("v-uni-view",{staticClass:"hx-navbar__content__main_search",style:{"padding-left":t.conf.leftButton?"9px":"13px","padding-right":t.conf.rightButton?"9px":"13px"},on:{click:function(n){arguments[0]=n=t.$handleEvent(n),t.searchClick.apply(void 0,arguments)}}},[e("v-uni-text",{staticClass:"hxicon hx-navbar__content__main_search_hxicon",staticStyle:{color:"#dbdbdb"}},[t._v("")]),e("v-uni-input",{staticClass:"hx-navbar__content__main_search_input hx_font_size",attrs:{type:"text","placeholder-style":"color:#dbdbdb;","confirm-type":"search",value:t.conf.search.value,placeholder:t.conf.search.placeholder?t.conf.search.placeholder:"请输入搜索内容",disabled:t.conf.search.disabled},on:{confirm:function(n){arguments[0]=n=t.$handleEvent(n),t.searchConfirm.apply(void 0,arguments)}}})],1):t._e(),t.conf.search?t._e():e("v-uni-view",{staticStyle:{flex:"1"}}),t.conf.rightButton?e("v-uni-view",{staticClass:"hx-navbar__content__main_right",style:{padding:t.conf.rightSlot?"0":"0 13px 0 9px;"}},[t._l(t.conf.rightButton,(function(n,a){return[e("v-uni-view",{key:a+"_0",staticClass:"hx-navbar__content__main_right_btn",on:{click:function(e){arguments[0]=e=t.$handleEvent(e),t.onClickBtn(n)}}},[n.position&&"left"==n.position?e("v-uni-text",{staticClass:"hxicon hx-navbar__content__main_right_icon hx-navbar__icon",class:t.conf.font,style:{color:n.color?n.color:"",marginRight:n.txt?"2px":"0"}},[t._v(t._s(t.iconHandle(n.icon)))]):t._e(),n.txt?e("v-uni-text",{staticClass:"hx-navbar__content__main_right_txt hx_font_size",style:{color:n.color?n.color:t.txtColor}},[t._v(t._s(n.txt))]):t._e(),n.position&&"left"==n.position?t._e():e("v-uni-text",{staticClass:"hxicon hx-navbar__content__main_right_icon hx-navbar__icon",class:t.conf.font,style:{color:n.color?n.color:"",marginLeft:n.txt?"2px":"0"}},[t._v(t._s(t.iconHandle(n.icon)))])],1)]}))],2):t._e(),t.conf.rightSlot?e("v-uni-view",{staticClass:"hx-navbar__content__main_right",staticStyle:{padding:"0",margin:"0"}},[t.conf.rightSlotSwitch&&0==t.slotSwitchOpacity?t._t("rightSwitch"):t._t("right")],2):t._e()],1)],1),t.conf.barPlaceholder&&t.conf.fixed?e("v-uni-view",[t.conf.statusBar?e("v-uni-view",{style:{height:t.statusBarHeight+"px"}}):t._e(),e("v-uni-view",{style:{height:t.conf.height+"px"}})],1):t._e()],1)},i=[]},"20cb":function(t,n,e){"use strict";e.r(n);var a=e("0cba"),o=e("c749");for(var i in o)"default"!==i&&function(t){e.d(n,t,(function(){return o[t]}))}(i);e("3cee");var r,c=e("f0c5"),l=Object(c["a"])(o["default"],a["b"],a["c"],!1,null,"a9640046",null,!1,a["a"],r);n["default"]=l.exports},"256e":function(t,n,e){var a=e("9cb0");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var o=e("4f06").default;o("56e7f3e0",a,!0,{sourceMap:!1,shadowMode:!1})},"27bb":function(t,n,e){"use strict";(function(t){var a;e.d(n,"b",(function(){return o})),e.d(n,"c",(function(){return i})),e.d(n,"a",(function(){return a}));try{a={hxNavbar:e("0736").default}}catch(r){if(-1===r.message.indexOf("Cannot find module")||-1===r.message.indexOf(".vue"))throw r;t.error(r.message),t.error("1. 排查组件名称拼写是否正确"),t.error("2. 排查组件是否符合 easycom 规范，文档：https://uniapp.dcloud.net.cn/collocation/pages?id=easycom"),t.error("3. 若组件不符合 easycom 规范，需手动引入，并在 components 中注册该组件")}var o=function(){var t=this,n=t.$createElement,e=t._self._c||n;return e("v-uni-view",[e("hx-navbar",{ref:"hxnb",attrs:{config:t.config}}),e("v-uni-view",{staticClass:"title"},[t._v("22222222222")])],1)},i=[]}).call(this,e("5a52")["default"])},"2eae":function(t,n,e){var a=e("833c");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var o=e("4f06").default;o("7bc9fee7",a,!0,{sourceMap:!1,shadowMode:!1})},"3cee":function(t,n,e){"use strict";var a=e("2eae"),o=e.n(a);o.a},"408a":function(t,n,e){var a=e("c6b6");t.exports=function(t){if("number"!=typeof t&&"Number"!=a(t))throw TypeError("Incorrect invocation");return+t}},"6afe":function(t,n,e){var a=e("9454");"string"===typeof a&&(a=[[t.i,a,""]]),a.locals&&(t.exports=a.locals);var o=e("4f06").default;o("231fdb67",a,!0,{sourceMap:!1,shadowMode:!1})},"7551e":function(t,n,e){"use strict";var a=e("6afe"),o=e.n(a);o.a},7999:function(t,n,e){"use strict";var a=e("4ea4");Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var o=a(e("0736")),i={components:{hxNavbar:o.default},data:function(){return{config:{title:"",back:!1,color:"#ffffff",leftButton:[{key:"menu",icon:"",position:"left"}],rightButton:[{key:"search",icon:"",position:"right"}]},skeletonShow:!0,current:2,isFixed:!1,mode:"left",mask:!0,maskClick:!0,zIndex:"10001",maskZIndex:"10000",visible:!1}},onLoad:function(){var t=this;setTimeout((function(){t.skeletonShow=!1}),1800)},onPageScroll:function(t){this.$refs.hxnb.pageScroll(t)},methods:{onClickBtn:function(t){uni.showToast({title:"key为 ".concat(t.key," 的按钮"),icon:"none",duration:1300})},activeTab:function(t){uni.navigateTo({url:t.pagePath})},toPath:function(t){uni.navigateTo({url:t})}}};n.default=i},"833c":function(t,n,e){var a=e("24fb");n=a(!1),n.push([t.i,".tui-content-box[data-v-a9640046]{\ndisplay:-webkit-box;display:-webkit-flex;display:flex;\n-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row}.ico-menu[data-v-a9640046]{width:20px;height:18px;margin:5px 0 0 10px;background-image:url(/static/images/icons/ico-menus.png);\ndisplay:-webkit-box;display:-webkit-flex;display:flex;background-repeat:no-repeat;background-size:contain\n}.logo[data-v-a9640046]{width:130px;height:30px;background-image:url(/static/images/logo@3x.png);\nmargin:5px auto 0;display:-webkit-box;display:-webkit-flex;display:flex;background-repeat:no-repeat;background-size:contain\n}.ico-search[data-v-a9640046]{width:20px;height:21px;margin:5px 10px 0 0;background-image:url(/static/images/icons/ico-search.png);\ndisplay:-webkit-box;display:-webkit-flex;display:flex;background-repeat:no-repeat;background-size:contain\n}",""]),t.exports=n},9454:function(t,n,e){var a=e("24fb");n=a(!1),n.push([t.i,".tui-content-box[data-v-daad558c]{\ndisplay:-webkit-box;display:-webkit-flex;display:flex;\n-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row}.hx-navbar__content__main_left_btn[data-v-daad558c],\n.ico-menu[data-v-daad558c]{width:20px;height:18px!important;margin:5px 0 0 10px;background-image:url(/static/images/icons/ico-menus.png);\ndisplay:-webkit-box;display:-webkit-flex;display:flex;background-repeat:no-repeat;background-size:contain;text-indent:-9999px\n}.hx-navbar__content__main_center_txt[data-v-daad558c],\n.logo[data-v-daad558c]{width:130px;height:30px!important;background-image:url(/static/images/logo@3x.png);\nmargin:8px auto 0;display:-webkit-box;display:-webkit-flex;display:flex;background-repeat:no-repeat;background-size:contain;background-position:50%;text-indent:-9999px\n}.hx-navbar__content__main_right_btn[data-v-daad558c],\n.ico-search[data-v-daad558c]{width:20px;height:21px!important;margin:5px 10px 0 0;background-image:url(/static/images/icons/ico-search.png);\ndisplay:-webkit-box;display:-webkit-flex;display:flex;background-repeat:no-repeat;background-size:contain;text-indent:-9999px\n}",""]),t.exports=n},9615:function(t,n,e){"use strict";e.r(n);var a=e("b91a"),o=e.n(a);for(var i in a)"default"!==i&&function(t){e.d(n,t,(function(){return a[t]}))}(i);n["default"]=o.a},"9cb0":function(t,n,e){var a=e("24fb");n=a(!1),n.push([t.i,'@charset "UTF-8";\r\n/**\r\n * 这里是uni-app内置的常用样式变量\r\n *\r\n * uni-app 官方扩展插件及插件市场（https://ext.dcloud.net.cn）上很多三方插件均使用了这些样式变量\r\n * 如果你是插件开发者，建议你使用scss预处理，并在插件代码中直接使用这些变量（无需 import 这个文件），方便用户通过搭积木的方式开发整体风格一致的App\r\n *\r\n */\r\n/**\r\n * 如果你是App开发者（插件使用者），你可以通过修改这些变量来定制自己的插件主题，实现自定义主题功能\r\n *\r\n * 如果你的项目同样使用了scss预处理，你也可以直接在你的 scss 代码中使用如下变量，同时无需 import 这个文件\r\n */\r\n/* 颜色变量 */\r\n/* 行为相关颜色 */\r\n/* 文字基本颜色 */\r\n/* 背景颜色 */\r\n/* 边框颜色 */\r\n/* 尺寸变量 */\r\n/* 文字尺寸 */\r\n/* 图片尺寸 */\r\n/* Border Radius */\r\n/* 水平间距 */\r\n/* 垂直间距 */\r\n/* 透明度 */\r\n/* 文章场景相关 */.hxicon[data-v-daad558c]{text-decoration:none}.hd[data-v-daad558c]{overflow:hidden}[class*="hx-navbar__"][data-v-daad558c]{display:-webkit-box;display:-webkit-flex;display:flex}.hx-navbar[data-v-daad558c]{position:relative;padding-top:0;overflow:hidden;-webkit-box-flex:1;-webkit-flex:1;flex:1}.hx-navbar__stretch[data-v-daad558c]{-webkit-box-align:stretch;-webkit-align-items:stretch;align-items:stretch}.hx-navbar__status[data-v-daad558c]{position:relative;z-index:3}.hx-navbar__icon[data-v-daad558c]{position:relative;top:1px}.hx-navbar__icontran[data-v-daad558c]{border-radius:100%!important;background-color:rgba(0,0,0,.5)!important;width:30px!important;height:30px!important;line-height:30px!important;text-align:center!important;color:#fff!important}.hx-navbar__btntran[data-v-daad558c]{padding:0!important;margin-left:6px!important;border-radius:100%!important;background-color:rgba(0,0,0,.5)!important;width:30px!important;height:30px!important;line-height:30px!important;text-align:center!important;color:#fff!important}.hx-navbar__content[data-v-daad558c]{position:relative;-webkit-box-flex:1;-webkit-flex:1;flex:1;z-index:1;width:100%;-webkit-box-orient:vertical;-webkit-box-direction:normal;-webkit-flex-direction:column;flex-direction:column}.hx-navbar__content__imgctn[data-v-daad558c]{position:absolute;left:0;top:0;right:0;bottom:0;z-index:2}.hx-navbar__content__main[data-v-daad558c]{position:relative;z-index:3;-webkit-box-flex:1;-webkit-flex:1;flex:1;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-align:center;-webkit-align-items:center;align-items:center;padding:0}.hx-navbar__content__main_back[data-v-daad558c]{-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-align:center;-webkit-align-items:center;align-items:center;height:100%;padding:0 13px;margin:0;position:relative;z-index:2;line-height:44px}.hx-navbar__content__main_search[data-v-daad558c]{position:relative;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-flex:1;-webkit-flex:1;flex:1;height:100%}.hx-navbar__content__main_search_hxicon[data-v-daad558c]{position:absolute;left:24px;font-size:18px;color:#dbdbdb}.hx-navbar__content__main_search_input[data-v-daad558c]{height:30px;-webkit-box-flex:1;-webkit-flex:1;flex:1;border-radius:40px;background-color:rgba(32,32,32,.35);padding:0 16px 0 36px}.hx-navbar__content__main_right[data-v-daad558c]{position:relative;z-index:2;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-align:center;-webkit-align-items:center;align-items:center;height:100%;padding:0 13px 0 9px;margin:0}.hx-navbar__content__main_right_txt[data-v-daad558c]{max-width:60px;lines:1;text-overflow:ellipsis}.hx-navbar__content__main_right_btn[data-v-daad558c]{-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;align-items:center;height:100%;padding:0 5px}.hx-navbar__content__main_right_icon[data-v-daad558c]{-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;text-align:center}.hx-navbar__content__main_center[data-v-daad558c]{-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-flex:1;-webkit-flex:1;flex:1;-webkit-box-align:stretch;-webkit-align-items:stretch;align-items:stretch;position:absolute;top:0;left:0;right:0;bottom:0;z-index:0;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center}.hx-navbar__content__main_center_txt[data-v-daad558c]{-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-flex:46;-webkit-flex:46;flex:46;display:block;white-space:nowrap;overflow:hidden;height:44px;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;-webkit-box-align:center;-webkit-align-items:center;align-items:center;text-align:center;text-overflow:ellipsis;line-height:44px}.hx-navbar__content__main_center_flex[data-v-daad558c]{-webkit-box-flex:32;-webkit-flex:32;flex:32}.hx-navbar__content__main_left[data-v-daad558c]{position:relative;z-index:1;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-align:center;-webkit-align-items:center;align-items:center;height:100%;padding:0 9px 0 13px;margin:0}.hx-navbar__content__main_left_txt[data-v-daad558c]{width:60px;lines:1;text-overflow:ellipsis}.hx-navbar__content__main_left_btn[data-v-daad558c]{-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-align:center;-webkit-align-items:center;align-items:center;height:100%;padding:0 5px}.hx-navbar__fixed[data-v-daad558c]{position:fixed;top:0;z-index:99}.hx-navbar__shadow[data-v-daad558c]{box-shadow:0 %?2?% %?12?% #ccc}.hx-navbar__border[data-v-daad558c]:after{position:absolute;z-index:3;bottom:0;left:0;right:0;height:1px;-webkit-transform:scaleY(.5);transform:scaleY(.5);background-color:#efefef}.hx-navbar__icon[data-v-daad558c]{top:0!important;font-size:20px}.hx_font_size[data-v-daad558c]{font-size:18px}.hx_text_overflow[data-v-daad558c]{display:block;white-space:nowrap;overflow:hidden;text-overflow:ellipsis;text-align:center}',""]),t.exports=n},b680:function(t,n,e){"use strict";var a=e("23e7"),o=e("a691"),i=e("408a"),r=e("1148"),c=e("d039"),l=1..toFixed,s=Math.floor,d=function(t,n,e){return 0===n?e:n%2===1?d(t,n-1,e*t):d(t*t,n/2,e)},f=function(t){var n=0,e=t;while(e>=4096)n+=12,e/=4096;while(e>=2)n+=1,e/=2;return n},b=l&&("0.000"!==8e-5.toFixed(3)||"1"!==.9.toFixed(0)||"1.25"!==1.255.toFixed(2)||"1000000000000000128"!==(0xde0b6b3a7640080).toFixed(0))||!c((function(){l.call({})}));a({target:"Number",proto:!0,forced:b},{toFixed:function(t){var n,e,a,c,l=i(this),b=o(t),u=[0,0,0,0,0,0],x="",g="0",h=function(t,n){var e=-1,a=n;while(++e<6)a+=t*u[e],u[e]=a%1e7,a=s(a/1e7)},_=function(t){var n=6,e=0;while(--n>=0)e+=u[n],u[n]=s(e/t),e=e%t*1e7},p=function(){var t=6,n="";while(--t>=0)if(""!==n||0===t||0!==u[t]){var e=String(u[t]);n=""===n?e:n+r.call("0",7-e.length)+e}return n};if(b<0||b>20)throw RangeError("Incorrect fraction digits");if(l!=l)return"NaN";if(l<=-1e21||l>=1e21)return String(l);if(l<0&&(x="-",l=-l),l>1e-21)if(n=f(l*d(2,69,1))-69,e=n<0?l*d(2,-n,1):l/d(2,n,1),e*=4503599627370496,n=52-n,n>0){h(0,e),a=b;while(a>=7)h(1e7,0),a-=7;h(d(10,a,1),0),a=n-1;while(a>=23)_(1<<23),a-=23;_(1<<a),h(1,1),_(2),g=p()}else h(0,e),h(1<<-n,0),g=p()+r.call("0",b);return b>0?(c=g.length,g=x+(c<=b?"0."+r.call("0",b-c)+g:g.slice(0,c-b)+"."+g.slice(c-b))):g=x+g,g}})},b91a:function(t,n,e){"use strict";e("99af"),e("fb6a"),e("a9e3"),e("b680"),e("d3b7"),e("acd8"),e("e25e"),e("ac1f"),e("25f0"),e("5319"),e("1276"),Object.defineProperty(n,"__esModule",{value:!0}),n.default=void 0;var a={name:"hxNavbar",components:{},data:function(){return{title:"",backgroundColorRgba:"",backgroundColorRgb:"rgb(222,222,222)",backgroundImage:null,backgroundImageEnd:null,txtColor:"#333333",bgArr:[],colorArr:[],statusBarBackground:"",bgTransparent:1,bgImgTransparent:[1,1],jnWidth:0,bgIsLine:!1,slotSwitchOpacity:1,conf:{title:"",height:44,fixed:!0,statusBar:!0,statusBarFontColor:"#000000",statusBarBackground:null,font:"hxicon",fontSize:"18px",color:"#333333",backgroundColor:[1,"#ffffff"],backgroundImg:null,backgroundColorLinearDeg:"to right",slideHeight:100,slideBackgroundColor:null,back:!0,backTxt:null,backTabPage:null,backPage:null,leftButton:null,rightButton:null,search:null,shadow:!1,border:!1,barPlaceholder:!0,slotSwitch:0,rightSlot:!1,rightSlotSwitch:!1}}},props:{config:{type:Object,default:function(){return{}}}},computed:{statusBarHeight:function(){return uni.getSystemInfoSync().statusBarHeight},navbarHeight:function(){var t=this;return uni.getSystemInfoSync().statusBarHeight+t.conf.height+"px"},screenWidth:function(){return uni.getSystemInfoSync().screenWidth}},created:function(){var t=this;t.init()},mounted:function(){},watch:{},methods:{iconHandle:function(t){return t=t.replace(/(&#x|;)/g,""),unescape("%u"+t)},init:function(){var t=this;t.conf=Object.assign(t.conf,t.config),""!=t.conf.title&&uni.setNavigationBarTitle({title:t.conf.title}),t.conf.title&&("object"==typeof t.conf.title&&2==t.conf.title.length?t.title=t.conf.title[0]:t.title=t.conf.title),t.conf.statusBarBackground&&("object"==typeof t.conf.statusBarBackground&&2==t.conf.statusBarBackground.length?t.statusBarBackground=t.conf.statusBarBackground[0]:t.statusBarBackground=t.conf.statusBarBackground),t.conf.statusBarFontColor&&uni.setNavigationBarColor({frontColor:"object"==typeof t.conf.statusBarFontColor?t.conf.statusBarFontColor[0]:t.conf.statusBarFontColor,backgroundColor:"#000000"}),""!=t.conf.color&&("object"==typeof t.conf.color&&2==t.conf.color.length?(t.txtColor=t.conf.color[0],t.colorArr=t.gradientColor(t.conf.color[0],t.conf.color[1],t.conf.slideHeight)):t.txtColor=t.conf.color)},onBack:function(){var t=this;t.conf.backTabPage?uni.switchTab({url:t.conf.backTabPage}):t.conf.backPage?uni.redirectTo({url:t.conf.backPage}):getCurrentPages().length>1?uni.navigateBack():history.back()},onClickBtn:function(t){this.$emit("clickBtn",t)},pageScroll:function(t){var n=this,e=parseFloat(t.scrollTop.toFixed(2)),a=n.conf.slideHeight,o=Math.round(e);if(o>0?o-=1:o=0,n.conf.title&&"object"==typeof n.conf.title&&2==n.conf.title.length&&(n.title=e<=a?n.conf.title[0]:n.conf.title[1]),n.conf.color&&"object"==typeof n.conf.color&&2==n.conf.color.length){var i=n.colorArr,r=o<=i.length-1?i[o]:i[i.length-1];n.txtColor=r}n.slotSwitchOpacity=n.transHandle(e,a,1,0),n.conf.statusBarFontColor&&"object"==typeof n.conf.statusBarFontColor&&2==n.conf.statusBarFontColor.length&&(e<=a?uni.setNavigationBarColor({frontColor:n.conf.statusBarFontColor[0],backgroundColor:"#ffffff"}):uni.setNavigationBarColor({frontColor:n.conf.statusBarFontColor[1],backgroundColor:"#ffffff"})),n.conf.statusBarBackground&&"object"==typeof n.conf.statusBarBackground&&2==n.conf.statusBarBackground.length&&(n.statusBarBackground=e<=a?n.conf.statusBarBackground[0]:n.conf.statusBarBackground[1]);var c=n.conf.backgroundImg;if(n.conf.backgroundImg&&"object"==typeof n.conf.backgroundImg&&c.length>1){var l=c.length;n.bgTransparent=l<=3?1:n.transHandle(e,a,c[2],c[3]),n.bgImgTransparent[0]=l<=3?1:n.transHandle(e,a,c[2],0),n.bgImgTransparent[1]=c[3];e<=a?c[0]:c[1],n.backgroundImage=n.bgImgStringHandle(c[0]),n.backgroundImageEnd=n.bgImgStringHandle(c[1])}if(n.conf.slideBackgroundColor){var s=n.bgArr,d=[];for(var f in n.bgArr){var b=o<=s[f].length-1?s[f][o]:s[f][s[f].length-1];d.push(b.replace(/(?:\(|\)|rgb|RGB)*/g,"").split(","))}var u=n.conf.backgroundColor[0],x=n.conf.slideBackgroundColor[0],g=x;if(e<=a){var h=Math.abs(x-u),_=parseFloat(h/a).toFixed(4),p=parseFloat(_*e).toFixed(2);g=u>x?u-p:u+p,g=parseFloat(g).toFixed(2)}var v=n.conf.slideBackgroundColor[1],w="";if("object"==typeof v&&v.length>1){w="linear-gradient("+n.conf.backgroundColorLinearDeg+",";var k=d.length;for(var f in d){var m=d[f];w+="rgba("+m[0]+","+m[1]+","+m[2]+","+g+")",k!=1*f+1&&(w+=",")}w+=")"}else w="rgba("+d[0][0]+","+d[0][1]+","+d[0][2]+","+g+")";n.bgTransparent=g,n.backgroundColorRgba=w}},transHandle:function(t,n,e,a){var o=a;if(t<=n){var i=Math.abs(a-e),r=parseFloat(i/n).toFixed(4),c=parseFloat(r*t).toFixed(2);o=e>a?e-c:e+c,o=parseFloat(o).toFixed(2)}return o},gradientColor:function(t,n,e){for(var a=this,o=a.colorRgb(t),i=o[0],r=o[1],c=o[2],l=a.colorRgb(n),s=l[0],d=l[1],f=l[2],b=(s-i)/e,u=(d-r)/e,x=(f-c)/e,g=[],h=0;h<e;h++){var _="rgb("+parseInt(b*h+i)+","+parseInt(u*h+r)+","+parseInt(x*h+c)+")",p=a.colorHex(_);g.push(p)}return g},colorRgb:function(t){var n=/^#([0-9a-fA-f]{3}|[0-9a-fA-f]{6})$/;t=t.toLowerCase();if(t&&n.test(t)){if(4===t.length){for(var e="#",a=1;a<4;a+=1)e+=t.slice(a,a+1).concat(t.slice(a,a+1));t=e}var o=[];for(a=1;a<7;a+=2)o.push(parseInt("0x"+t.slice(a,a+2)));return o}return t},colorHex:function(t){var n=t,e=/^#([0-9a-fA-f]{3}|[0-9a-fA-f]{6})$/;if(/^(rgb|RGB)/.test(n)){for(var a=n.replace(/(?:(|)|rgb|RGB)*/g,"").split(","),o="#",i=0;i<a.length;i++){var r=Number(a[i]).toString(16);r=r<10?"0"+r:r,"0"===r&&(r+=r),o+=r}return 7!==o.length&&(o=n),o}if(!e.test(n))return n;var c=n.replace(/#/,"").split("");if(6===c.length)return n;if(3===c.length){var l="#";for(i=0;i<c.length;i+=1)l+=c[i]+c[i];return l}},setBgColor:function(t,n){var e=this;if("object"==typeof t&&t.length>0){var a="linear-gradient("+e.conf.backgroundColorLinearDeg+",",o=null;for(var i in t){o=t[i];var r=e.colorRgb(o);a+="rgba(".concat(r[0],", ").concat(r[1],", ").concat(r[2],", ").concat(n,")"),t.length!=1*i+1&&(a+=",")}a+=")",e.backgroundColorRgba=a}else{var c=t,l=e.colorRgb(c);e.backgroundColorRgba="rgba(".concat(l[0],", ").concat(l[1],", ").concat(l[2],", ").concat(n,")")}},bgImgStringHandle:function(t){return t},searchConfirm:function(t){this.$emit("searchConfirm",t.detail)},searchClick:function(t){this.$emit("searchClick",!0)}}};n.default=a},c749:function(t,n,e){"use strict";e.r(n);var a=e("7999"),o=e.n(a);for(var i in a)"default"!==i&&function(t){e.d(n,t,(function(){return a[t]}))}(i);n["default"]=o.a},df7c:function(t,n,e){"use strict";var a=e("256e"),o=e.n(a);o.a}}]);