package router

import (
	"gin-vue-admin/front/api"
	"github.com/gin-gonic/gin"
)

func FrontMovieRouter(Router *gin.RouterGroup) {
	MovieRouter := Router.Group("movie")
	{
		//MovieRouter.GET("getMovieList", api.GetMovieList)  // 获取Movie列表
		MovieRouter.GET("getLatestMovieList", api.GetLatestMovieList)  // 首页->最近更新
		MovieRouter.GET("getRecommendMovieList", api.GetRecommendMovieList)  // 首页->热门推荐
		MovieRouter.GET("getMovieListByTypeId", api.GetMovieListByTypeId)  // 首页->顶部分类->
		MovieRouter.GET("getMovieListByTag", api.GetMovieListByTag)  //根据tag获取影片
		MovieRouter.GET("getMovieListByKeyword", api.GetMovieListByKeyword)  //首页->关键词搜索影片
		MovieRouter.POST("movieApproval", api.MovieApproval)//影片点赞
		MovieRouter.GET("movieInfoUpdate", api.MovieInfoUpdate)//切片后信息更新
	}
}
