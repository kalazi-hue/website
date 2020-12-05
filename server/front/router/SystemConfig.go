package router

import (
	"gin-vue-admin/front/api"
	"github.com/gin-gonic/gin"
)

func FrontSystemConfigRouter(Router *gin.RouterGroup) {
	SystemConfigRouter := Router.Group("systemConfig")
	{
		SystemConfigRouter.GET("getSystemConfigList", api.GetSystemConfigList)  // 获取SystemConfig列表
	}
}
