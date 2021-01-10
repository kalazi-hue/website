package router

import (
	"gin-vue-admin/front/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func FrontPlayFeedbackRouter(Router *gin.RouterGroup) {
	PlayFeedbackRouter := Router.Group("playFeedback").Use(middleware.JWTAuth())
	{
		PlayFeedbackRouter.POST("createPlayFeedback", api.CreatePlayFeedback) // 删除PlayFeedback
	}
}
