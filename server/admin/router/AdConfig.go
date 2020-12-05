package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAdConfigRouter(Router *gin.RouterGroup) {
	AdConfigRouter := Router.Group("adConfig").Use(middleware.JWTAuth())
	{
		AdConfigRouter.POST("createAdConfig", api.CreateAdConfig)   // 新建AdConfig
		AdConfigRouter.DELETE("deleteAdConfig", api.DeleteAdConfig) // 删除AdConfig
		AdConfigRouter.DELETE("deleteAdConfigByIds", api.DeleteAdConfigByIds) // 批量删除AdConfig
		AdConfigRouter.PUT("updateAdConfig", api.UpdateAdConfig)    // 更新AdConfig
		AdConfigRouter.GET("findAdConfig", api.FindAdConfig)        // 根据ID获取AdConfig
		AdConfigRouter.GET("getAdConfigList", api.GetAdConfigList)  // 获取AdConfig列表
	}
}
