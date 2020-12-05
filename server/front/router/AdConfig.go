package router

import (
	"gin-vue-admin/front/api"
	"github.com/gin-gonic/gin"
)

func FrontAdConfigRouter(Router *gin.RouterGroup) {
	AdConfigRouter := Router.Group("adConfig")
	{
		AdConfigRouter.GET("getAdConfigList", api.GetAdConfigList)  // 获取AdConfig列表
	}
}
