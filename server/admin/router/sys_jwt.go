package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitJwtRouter(Router *gin.RouterGroup) {
	ApiRouter := Router.Group("jwt").Use(middleware.JWTAuth())
	{
		ApiRouter.POST("jsonInBlacklist", api.JsonInBlacklist) // jwt加入黑名单
	}
}
