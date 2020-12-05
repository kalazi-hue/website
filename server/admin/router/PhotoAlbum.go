package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPhotoAlbumRouter(Router *gin.RouterGroup) {
	PhotoAlbumRouter := Router.Group("photoAlbum").Use(middleware.JWTAuth())
	{
		PhotoAlbumRouter.POST("createPhotoAlbum", api.CreatePhotoAlbum)   // 新建PhotoAlbum
		PhotoAlbumRouter.DELETE("deletePhotoAlbum", api.DeletePhotoAlbum) // 删除PhotoAlbum
		PhotoAlbumRouter.DELETE("deletePhotoAlbumByIds", api.DeletePhotoAlbumByIds) // 批量删除PhotoAlbum
		PhotoAlbumRouter.PUT("updatePhotoAlbum", api.UpdatePhotoAlbum)    // 更新PhotoAlbum
		PhotoAlbumRouter.GET("findPhotoAlbum", api.FindPhotoAlbum)        // 根据ID获取PhotoAlbum
		PhotoAlbumRouter.GET("getPhotoAlbumList", api.GetPhotoAlbumList)  // 获取PhotoAlbum列表
	}
}
