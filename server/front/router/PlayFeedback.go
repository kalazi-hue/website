package router

import (
	"gin-vue-admin/front/api"
	"github.com/gin-gonic/gin"
)

func FrontPlayFeedbackRouter(Router *gin.RouterGroup) {
	PlayFeedbackRouter := Router.Group("playFeedback")
	{
		PlayFeedbackRouter.POST("createPlayFeedback", api.CreatePlayFeedback) // 删除PlayFeedback
	}
}
