(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["pages-photo-detail~pages-photo-index"],{"06f0":function(e,t,i){"use strict";var a=i("f328"),n=i.n(a);n.a},"182a":function(e,t,i){"use strict";i.r(t);var a=i("40f1"),n=i.n(a);for(var o in a)"default"!==o&&function(e){i.d(t,e,(function(){return a[e]}))}(o);t["default"]=n.a},"40f1":function(e,t,i){"use strict";Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var a={props:{newsItem:{type:Object,default:function(e){return{}}}},methods:{click:function(){this.$emit("click")},close:function(e){e.stopPropagation(),this.$emit("close")}}};t.default=a},"43c8":function(e,t,i){"use strict";var a=i("4ea4");i("99af"),i("d81d"),i("a434"),i("a9e3"),i("d3b7"),i("ac1f"),i("25f0"),i("5319"),Object.defineProperty(t,"__esModule",{value:!0}),t.default=void 0;var n=a(i("ade3")),o=i("3a05"),r=a(i("572e")),s=a(i("a8e5")),l=a(i("1466")),c=a(i("4d13")),d=a(i("d5b9")),f=a(i("c6ae")),u={components:{uniList:r.default,uniCell:s.default,uniRefresh:l.default,uniLoadMore:c.default,noData:d.default,newsItem:f.default},props:{nid:{type:[Number,String],default:""}},data:function(){var e;return e={dataList:[],navigateFlag:!1,pulling:!1,refreshing:!1,refreshFlag:!1,refreshText:"",isLoading:!1,loadingText:"加载中...",isNoData:!1},(0,n.default)(e,"pulling",!1),(0,n.default)(e,"angle",0),(0,n.default)(e,"loadingMoreText",{contentdown:"",contentrefresh:"",contentnomore:""}),(0,n.default)(e,"refreshIcon","data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAADAAAAAwCAMAAABg3Am1AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAB5QTFRFcHBw3Nzct7e39vb2ycnJioqK7e3tpqam29vb////D8oK7wAAAAp0Uk5T////////////ALLMLM8AAABxSURBVHja7JVBDoAgDASrjqj//7CJBi90iyYeOHTPMwmFZrHjYyyFYYUy1bwUZqtJIYVxhf1a6u0R7iUvWsCcrEtwJHp8MwMdvh2amHduiZD3rpWId9+BgPd7Cc2LIkPyqvlQvKxKBJ//Qwq/CacAAwDUv0a0YuKhzgAAAABJRU5ErkJggg=="),e},created:function(){var e=this;this.pullTimer=null,this.requestParams={columnId:this.nid,minId:0,pageSize:10,column:"id,post_id,title,author_name,cover,published_at,comments_count"},this._isWidescreen=!1;var t=uni.createMediaQueryObserver(this);t.observe({minWidth:768},(function(t){e._isWidescreen=t}))},methods:{loadData:function(e){var t=this;if(!this.isLoading){this.isLoading=!0,this.isNoData=!1,this.requestParams.time=(new Date).getTime()+"";new Date;uni.request({url:"https://unidemo.dcloud.net.cn/api/news",data:this.requestParams,success:function(i){new Date;var a=i.data;t.isNoData=a.length<=0;var n=a.map((function(e){return{id:t.newGuid()+e.id,newsid:e.id,article_type:1,datetime:(0,o.friendlyDate)(new Date(e.published_at.replace(/\-/g,"/")).getTime()),title:e.title,image_url:e.cover,source:e.author_name,comment_count:e.comments_count,post_id:e.post_id}}));e?(t.dataList=n,t.requestParams.minId=0):(t.dataList=t.dataList.concat(n),t.requestParams.minId=a[a.length-1].id),t.dataList.length>0&&t._isWidescreen&&t.dataList.length<=10&&t.goDetail(t.dataList[0])},fail:function(e){0==t.dataList.length&&(t.isNoData=!0)},complete:function(i){t.isLoading=!1,e&&(t.refreshing=!1,t.refreshFlag=!1,t.refreshText="已刷新",t.pullTimer&&clearTimeout(t.pullTimer),t.pullTimer=setTimeout((function(){t.pulling=!1}),1e3))}})}},loadMore:function(e){this.loadData()},clear:function(){this.dataList.length=0,this.requestParams.minId=0},goDetail:function(e){this._isWidescreen?uni.$emit("updateDetail",{detail:encodeURIComponent(JSON.stringify(e))}):uni.navigateTo({url:"/pages/detail/detail?query="+encodeURIComponent(JSON.stringify(e))})},closeItem:function(e){var t=this;uni.showModal({content:"不感兴趣？",success:function(i){i.confirm&&t.dataList.splice(e,1)}})},refreshData:function(){this.isLoading||(this.pulling=!0,this.refreshing=!0,this.refreshText="正在刷新...",this.loadData(!0))},onrefresh:function(e){this.refreshData()},onpullingdown:function(e){this.refreshing||(this.pulling=!1,Math.abs(e.pullingDistance)>Math.abs(e.viewHeight)?(this.refreshFlag=!0,this.refreshText="释放立即刷新"):(this.refreshFlag=!1,this.refreshText="下拉可以刷新"))},newGuid:function(){var e=function(){return(65536*(1+Math.random())|0).toString(16).substring(1)};return(e()+e()+"-"+e()+"-4"+e().substr(0,3)+"-"+e()+"-"+e()+e()+e()).toUpperCase()}}};t.default=u},"47dc":function(e,t,i){"use strict";var a;i.d(t,"b",(function(){return n})),i.d(t,"c",(function(){return o})),i.d(t,"a",(function(){return a}));var n=function(){var e=this,t=e.$createElement,i=e._self._c||t;return e.newsItem.title?i("v-uni-view",{staticClass:"media-item view",on:{click:function(t){arguments[0]=t=e.$handleEvent(t),e.click.apply(void 0,arguments)}}},[i("v-uni-view",{staticClass:"view",style:{flexDirection:1===e.newsItem.article_type||2===e.newsItem.article_type?2===e.newsItem.article_type?"row":"row-reverse":"column"}},[i("v-uni-text",{staticClass:"media-title",class:{"media-title2":1===e.newsItem.article_type||2===e.newsItem.article_type}},[e._v(e._s(e.newsItem.title))]),e.newsItem.image_list||e.newsItem.image_url?i("v-uni-view",{staticClass:"image-section flex-row",class:{"image-section-right":2===e.newsItem.article_type,"image-section-left":1===e.newsItem.article_type}},[e.newsItem.image_url?i("v-uni-image",{staticClass:"image-list1",class:{"image-list2":1===e.newsItem.article_type||2===e.newsItem.article_type},attrs:{"fade-show":!1,src:e.newsItem.image_url}}):e._e(),e._l(e.newsItem.image_list,(function(t,a){return e.newsItem.image_list?i("v-uni-image",{key:a,staticClass:"image-list3",attrs:{"fade-show":!1,src:t.url}}):e._e()}))],2):e._e()],1),i("v-uni-view",{staticClass:"media-foot flex-row"},[i("v-uni-view",{staticClass:"media-info flex-row"},[i("v-uni-text",{staticClass:"info-text"},[e._v(e._s(e.newsItem.source))]),i("v-uni-text",{staticClass:"info-text"},[e._v(e._s(e.newsItem.comment_count)+"条评论")]),i("v-uni-text",{staticClass:"info-text"},[e._v(e._s(e.newsItem.datetime))])],1),i("v-uni-view",{staticClass:"close-view",on:{click:function(t){t.stopPropagation(),arguments[0]=t=e.$handleEvent(t),e.close.apply(void 0,arguments)}}},[i("v-uni-view",{staticClass:"close-l close-h"}),i("v-uni-view",{staticClass:"close-l close-v"})],1)],1),i("v-uni-view",{staticClass:"media-item-line",staticStyle:{position:"absolute"}})],1):e._e()},o=[]},"66cc":function(e,t,i){"use strict";var a;i.d(t,"b",(function(){return n})),i.d(t,"c",(function(){return o})),i.d(t,"a",(function(){return a}));var n=function(){var e=this,t=e.$createElement,i=e._self._c||t;return i("v-uni-view",{staticClass:"page-news"},[i("uni-list",{staticClass:"listview",attrs:{enableBackToTop:!0,"scroll-y":!0},on:{scrolltolower:function(t){arguments[0]=t=e.$handleEvent(t),e.loadMore()}}},[i("uni-refresh",{staticClass:"refresh",attrs:{display:e.refreshing?"show":"hide"},on:{refresh:function(t){arguments[0]=t=e.$handleEvent(t),e.onrefresh.apply(void 0,arguments)},pullingdown:function(t){arguments[0]=t=e.$handleEvent(t),e.onpullingdown.apply(void 0,arguments)}}},[i("div",{staticClass:"refresh-view"},[i("v-uni-image",{staticClass:"refresh-icon",class:{"refresh-icon-active":e.refreshFlag},style:{width:e.refreshing||e.pulling?0:"32px"},attrs:{src:e.refreshIcon}}),e.refreshing?i("uni-load-more",{staticClass:"loading-icon",attrs:{status:"loading",contentText:e.loadingMoreText}}):e._e(),i("v-uni-text",{staticClass:"loading-text"},[e._v(e._s(e.refreshText))])],1)]),e._l(e.dataList,(function(t,a){return i("uni-cell",{key:t.id},[i("news-item",{attrs:{newsItem:t},on:{close:function(t){arguments[0]=t=e.$handleEvent(t),e.closeItem(a)},click:function(i){arguments[0]=i=e.$handleEvent(i),e.goDetail(t)}}})],1)})),e.isLoading||e.dataList.length>4?i("uni-cell",[i("v-uni-view",{staticClass:"loading-more"},[i("v-uni-text",{staticClass:"loading-more-text"},[e._v(e._s(e.loadingText))])],1)],1):e._e()],2),e.isNoData?i("no-data",{staticClass:"no-data",on:{retry:function(t){arguments[0]=t=e.$handleEvent(t),e.loadMore.apply(void 0,arguments)}}}):e._e()],1)},o=[]},"7bb5":function(e,t,i){"use strict";var a=i("c17f"),n=i.n(a);n.a},"9f8d":function(e,t,i){"use strict";i.r(t);var a=i("66cc"),n=i("b829");for(var o in n)"default"!==o&&function(e){i.d(t,e,(function(){return n[e]}))}(o);i("7bb5");var r,s=i("f0c5"),l=Object(s["a"])(n["default"],a["b"],a["c"],!1,null,"5874cbaa",null,!1,a["a"],r);t["default"]=l.exports},b829:function(e,t,i){"use strict";i.r(t);var a=i("43c8"),n=i.n(a);for(var o in a)"default"!==o&&function(e){i.d(t,e,(function(){return a[e]}))}(o);t["default"]=n.a},c17f:function(e,t,i){var a=i("d7b7");"string"===typeof a&&(a=[[e.i,a,""]]),a.locals&&(e.exports=a.locals);var n=i("4f06").default;n("73fa63bc",a,!0,{sourceMap:!1,shadowMode:!1})},c6ae:function(e,t,i){"use strict";i.r(t);var a=i("47dc"),n=i("182a");for(var o in n)"default"!==o&&function(e){i.d(t,e,(function(){return n[e]}))}(o);i("06f0");var r,s=i("f0c5"),l=Object(s["a"])(n["default"],a["b"],a["c"],!1,null,"e2ac4524",null,!1,a["a"],r);t["default"]=l.exports},d23d:function(e,t,i){var a=i("24fb");t=a(!1),t.push([e.i,".view[data-v-e2ac4524]{-webkit-box-orient:vertical;-webkit-box-direction:normal;-webkit-flex-direction:column;flex-direction:column}.flex-row[data-v-e2ac4524]{-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row}.flex-col[data-v-e2ac4524]{-webkit-box-orient:vertical;-webkit-box-direction:normal;-webkit-flex-direction:column;flex-direction:column}.list-cell[data-v-e2ac4524]{padding:0 15px}.uni-list-cell-hover[data-v-e2ac4524]{background-color:#eee}.media-item[data-v-e2ac4524]{position:relative;-webkit-box-flex:1;-webkit-flex:1;flex:1;-webkit-box-orient:vertical;-webkit-box-direction:normal;-webkit-flex-direction:column;flex-direction:column;padding:10px 15px 10px 15px}.media-item-line[data-v-e2ac4524]{position:absolute;left:15px;right:15px;bottom:0;height:1px;background-color:#ebebeb}.media-image-right[data-v-e2ac4524]{-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row}.media-image-left[data-v-e2ac4524]{-webkit-box-orient:horizontal;-webkit-box-direction:reverse;-webkit-flex-direction:row-reverse;flex-direction:row-reverse}.media-title[data-v-e2ac4524]{-webkit-box-flex:1;-webkit-flex:1;flex:1}.media-title[data-v-e2ac4524]{lines:3;text-overflow:ellipsis;font-size:15px;color:#555}.media-title2[data-v-e2ac4524]{-webkit-box-flex:1;-webkit-flex:1;flex:1;margin-top:3px;line-height:20px}.image-section[data-v-e2ac4524]{margin-top:10px;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between}.image-section-right[data-v-e2ac4524]{margin-top:0;margin-left:5px;width:112px;height:73px}.image-section-left[data-v-e2ac4524]{margin-top:0;margin-right:5px;width:112px;height:73px}.image-list1[data-v-e2ac4524]{height:240px}.image-list2[data-v-e2ac4524]{width:112px;height:73px}.image-list3[data-v-e2ac4524]{width:112px;height:73px}.media-info[data-v-e2ac4524]{-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-align:center;-webkit-align-items:center;align-items:center}.info-text[data-v-e2ac4524]{margin-right:10px;color:#999;font-size:12px}.media-foot[data-v-e2ac4524]{margin-top:12px;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-pack:justify;-webkit-justify-content:space-between;justify-content:space-between}.close-view[data-v-e2ac4524]{position:relative;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;width:20px;height:15px;line-height:15px;border-width:%?1?%;border-style:solid;border-color:#aaa;border-radius:4px;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;text-align:center}.close-l[data-v-e2ac4524]{position:absolute;width:9px;height:1px;background-color:#aaa}.close-h[data-v-e2ac4524]{-webkit-transform:rotate(45deg);transform:rotate(45deg)}.close-v[data-v-e2ac4524]{-webkit-transform:rotate(-45deg);transform:rotate(-45deg)}",""]),e.exports=t},d7b7:function(e,t,i){var a=i("24fb");t=a(!1),t.push([e.i,".no-data[data-v-5874cbaa]{-webkit-box-flex:1;-webkit-flex:1;flex:1;position:absolute;left:0;top:0;right:0;bottom:0;z-index:10}.page-news[data-v-5874cbaa]{-webkit-box-flex:1;-webkit-flex:1;flex:1;-webkit-box-orient:vertical;-webkit-box-direction:normal;-webkit-flex-direction:column;flex-direction:column;position:absolute;left:0;top:0;right:0;bottom:0}.listview[data-v-5874cbaa]{position:absolute;left:0;top:0;right:0;bottom:0;\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;-webkit-box-orient:vertical;-webkit-box-direction:normal;-webkit-flex-direction:column;flex-direction:column;\r\n\r\nflex-direction:column\n}.refresh[data-v-5874cbaa]{-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center}.refresh-view[data-v-5874cbaa]{\r\ndisplay:-webkit-box;display:-webkit-flex;display:flex;\r\nwidth:%?750?%;height:64px;-webkit-box-orient:horizontal;-webkit-box-direction:normal;-webkit-flex-direction:row;flex-direction:row;-webkit-flex-wrap:nowrap;flex-wrap:nowrap;-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center}.refresh-icon[data-v-5874cbaa]{width:32px;height:32px;-webkit-transition-duration:.5s;transition-duration:.5s;-webkit-transition-property:-webkit-transform;transition-property:-webkit-transform;transition-property:transform;transition-property:transform,-webkit-transform;-webkit-transform:rotate(0deg);transform:rotate(0deg);-webkit-transform-origin:15px 15px;transform-origin:15px 15px}.refresh-icon-active[data-v-5874cbaa]{-webkit-transform:rotate(180deg);transform:rotate(180deg)}.loading-icon[data-v-5874cbaa]{width:28px;height:28px;margin-right:5px;color:grey}.loading-text[data-v-5874cbaa]{margin-left:2px;font-size:16px;color:#999}.loading-more[data-v-5874cbaa]{-webkit-box-align:center;-webkit-align-items:center;align-items:center;-webkit-box-pack:center;-webkit-justify-content:center;justify-content:center;padding-top:14px;padding-bottom:14px;text-align:center}.loading-more-text[data-v-5874cbaa]{font-size:%?28?%;color:#999}",""]),e.exports=t},f328:function(e,t,i){var a=i("d23d");"string"===typeof a&&(a=[[e.i,a,""]]),a.locals&&(e.exports=a.locals);var n=i("4f06").default;n("7c2ce2f5",a,!0,{sourceMap:!1,shadowMode:!1})}}]);