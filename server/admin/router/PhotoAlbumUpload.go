package router

import (
	"gin-vue-admin/admin/api"
	"github.com/gin-gonic/gin"
)

func PhotoAlbumUploadRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("photoAlbumUpload")
	{
		ApiRouter.POST("photoAlbumUpload", api.PhotoAlbumUploaderUpload) // 上传功能
		ApiRouter.POST("photoAlbumMerge", api.PhotoAlbumUploadMerge) // 合并
		ApiRouter.GET("photoAlbumCheckMd5", api.PhotoAlbumCheckMd5)
	}
}
