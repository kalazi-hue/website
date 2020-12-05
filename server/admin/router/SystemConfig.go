package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSystemConfigRouter(Router *gin.RouterGroup) {
	SystemConfigRouter := Router.Group("systemConfig").Use(middleware.JWTAuth())
	{
		SystemConfigRouter.POST("createSystemConfig", api.CreateSystemConfig)   // 新建SystemConfig
		SystemConfigRouter.DELETE("deleteSystemConfig", api.DeleteSystemConfig) // 删除SystemConfig
		SystemConfigRouter.DELETE("deleteSystemConfigByIds", api.DeleteSystemConfigByIds) // 批量删除SystemConfig
		SystemConfigRouter.PUT("updateSystemConfig", api.UpdateSystemConfig)    // 更新SystemConfig
		SystemConfigRouter.GET("findSystemConfig", api.FindSystemConfig)        // 根据ID获取SystemConfig
		SystemConfigRouter.GET("getSystemConfigList", api.GetSystemConfigList)  // 获取SystemConfig列表
	}
}
