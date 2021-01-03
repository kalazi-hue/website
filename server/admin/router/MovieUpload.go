package router

import (
	"gin-vue-admin/admin/api"
	"github.com/gin-gonic/gin"
)

func MovieUploadRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("movieUpload")
	{
		ApiRouter.POST("movieUpload", api.MovieUploaderUpload) // 上传功能
		ApiRouter.POST("movieMerge", api.MovieUploadMerge) // 合并
		ApiRouter.GET("movieCheckMd5", api.MovieCheckMd5)
	}
}
