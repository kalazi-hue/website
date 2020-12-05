package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemRouter(Router *gin.RouterGroup) {
	SystemRouter := Router.Group("system").Use(middleware.JWTAuth(), middleware.CasbinHandler())
	{
		SystemRouter.POST("getSystemConfig", api.GetSystemConfig) // 获取配置文件内容
		SystemRouter.POST("setSystemConfig", api.SetSystemConfig) // 设置配置文件内容
		SystemRouter.POST("getServerInfo", api.GetServerInfo)     // 获取服务器信息
	}
}
