package router

import (
	"gin-vue-admin/front/api"
	"github.com/gin-gonic/gin"
)

func FrontTagRouter(Router *gin.RouterGroup) {
	TagRouter := Router.Group("tag")
	{
		TagRouter.GET("getTagList", api.GetTagList)  // 获取Tag列表
	}
}
