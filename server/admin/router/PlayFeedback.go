package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitPlayFeedbackRouter(Router *gin.RouterGroup) {
	PlayFeedbackRouter := Router.Group("playFeedback").Use(middleware.JWTAuth())
	{
		PlayFeedbackRouter.DELETE("deletePlayFeedback", api.DeletePlayFeedback) // 删除PlayFeedback
		PlayFeedbackRouter.DELETE("deletePlayFeedbackByIds", api.DeletePlayFeedbackByIds) // 批量删除PlayFeedback
		//PlayFeedbackRouter.GET("findPlayFeedback", api.FindPlayFeedback)        // 根据ID获取PlayFeedback
		PlayFeedbackRouter.GET("getPlayFeedbackList", api.GetPlayFeedbackList)  // 获取PlayFeedback列表
	}
}
