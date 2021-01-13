import { fetch, post, patch, put } from './request.js'
 
/**
* api接口的统一出口
*/

export const API = {
	// 分页获取SystemConfig列表
  getSystemConfigList: params => { return fetch('/systemConfig/getSystemConfigList', params).then(res => res.data)},
	// 获取广告列表
  getAdConfigList: params => { return fetch('/adConfig/getAdConfigList', params).then(res => res.data)},
	// 分页获取Tag列表
 	getTagList: params => { return fetch('/tag/getTagList', params).then(res => res.data)},
	// 首页->「最新更新」模块
  getLatestMovieList: params => { return fetch('/movie/getLatestMovieList', params).then(res => res.data)},
	// 获取推荐影片列表
  getRecommendMovieList: params => { return fetch('/movie/getRecommendMovieList', params).then(res => res.data)},
	// 影片详情页面，点击tag获取影片
  getMovieListByTag: params => { return fetch('/movie/getMovieListByTag', params).then(res => res.data)},
	// 首页->顶部->「国产、日韩」模块
  getMovieListByTypeId: params => { return fetch('/movie/getMovieListByTypeId', params).then(res => res.data)},
	// 首页，关键词搜索影片
  getMovieListByKeyword: params => { return fetch('/movie/getMovieListByKeyword', params).then(res => res.data)},
	// 点赞
  movieApproval: params => { return post('/movie/movieApproval', params).then(res => res.data)},
	// 分页获取写真列表
 	getPhotoAlbumList: params => { return fetch('/photoAlbum/getPhotoAlbumList', params).then(res => res.data)},
	// 点击类型查询写真
 	getPhotoAlbumListByType: params => { return fetch('/photoAlbum/getPhotoAlbumListByType', params).then(res => res.data)},
	// 反馈
 	createPlayFeedback: params => { return post('/playFeedback/createPlayFeedback', params).then(res => res.data)}
}

export default API