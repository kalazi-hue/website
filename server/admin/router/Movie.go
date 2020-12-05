package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitMovieRouter(Router *gin.RouterGroup) {
	MovieRouter := Router.Group("movie").Use(middleware.JWTAuth())
	{
		MovieRouter.POST("createMovie", api.CreateMovie)   // 新建Movie
		MovieRouter.DELETE("deleteMovie", api.DeleteMovie) // 删除Movie
		MovieRouter.DELETE("deleteMovieByIds", api.DeleteMovieByIds) // 批量删除Movie
		MovieRouter.PUT("updateMovie", api.UpdateMovie)    // 更新Movie
		MovieRouter.GET("findMovie", api.FindMovie)        // 根据ID获取Movie
		MovieRouter.GET("getMovieList", api.GetMovieList)  // 获取Movie列表
	}
}
