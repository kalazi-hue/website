package router

import (
	"gin-vue-admin/front/api"
	"github.com/gin-gonic/gin"
)

func FrontPhotoAlbumRouter(Router *gin.RouterGroup) {
	PhotoAlbumRouter := Router.Group("photoAlbum")
	{
		PhotoAlbumRouter.GET("getPhotoAlbumListByType", api.GetPhotoAlbumListByType)        // 根据type获取PhotoAlbum
		PhotoAlbumRouter.GET("getPhotoAlbumList", api.GetPhotoAlbumList)  // 获取PhotoAlbum列表
	}
}
