package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitEmailRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("email").Use(middleware.JWTAuth())
	{
		UserRouter.POST("emailTest", api.EmailTest) // 发送测试邮件
	}
}
