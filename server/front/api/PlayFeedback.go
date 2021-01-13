package api

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/response"
    "gin-vue-admin/front/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags PlayFeedback
// @Summary 创建PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayFeedback true "创建PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playFeedback/createPlayFeedback [post]
func CreatePlayFeedback(c *gin.Context) {
	var playFeedback model.PlayFeedback
	_ = c.ShouldBindJSON(&playFeedback)
	if err := service.CreatePlayFeedback(playFeedback); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}
