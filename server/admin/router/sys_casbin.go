package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitCasbinRouter(Router *gin.RouterGroup) {
	CasbinRouter := Router.Group("casbin").
		Use(middleware.JWTAuth()).
		Use(middleware.OperationRecord())
	{
		CasbinRouter.POST("updateCasbin", api.UpdateCasbin)
		CasbinRouter.POST("getPolicyPathByAuthorityId", api.GetPolicyPathByAuthorityId)
	}
}
