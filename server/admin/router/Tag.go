package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitTagRouter(Router *gin.RouterGroup) {
	TagRouter := Router.Group("tag").Use(middleware.JWTAuth())
	{
		TagRouter.POST("createTag", api.CreateTag)   // 新建Tag
		TagRouter.DELETE("deleteTag", api.DeleteTag) // 删除Tag
		TagRouter.DELETE("deleteTagByIds", api.DeleteTagByIds) // 批量删除Tag
		TagRouter.PUT("updateTag", api.UpdateTag)    // 更新Tag
		TagRouter.GET("findTag", api.FindTag)        // 根据ID获取Tag
		TagRouter.GET("getTagList", api.GetTagList)  // 获取Tag列表
	}
}
