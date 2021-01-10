package api

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/admin/service"
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

// @Tags PlayFeedback
// @Summary 删除PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayFeedback true "删除PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playFeedback/deletePlayFeedback [delete]
func DeletePlayFeedback(c *gin.Context) {
	var playFeedback model.PlayFeedback
	_ = c.ShouldBindJSON(&playFeedback)
	if err := service.DeletePlayFeedback(playFeedback); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PlayFeedback
// @Summary 批量删除PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /playFeedback/deletePlayFeedbackByIds [delete]
func DeletePlayFeedbackByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePlayFeedbackByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}



// @Tags PlayFeedback
// @Summary 用id查询PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayFeedback true "用id查询PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /playFeedback/findPlayFeedback [get]
func FindPlayFeedback(c *gin.Context) {
	var playFeedback model.PlayFeedback
	_ = c.ShouldBindQuery(&playFeedback)
	if err, replayFeedback := service.GetPlayFeedback(playFeedback.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"replayFeedback": replayFeedback}, c)
	}
}

// @Tags PlayFeedback
// @Summary 分页获取PlayFeedback列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PlayFeedbackSearch true "分页获取PlayFeedback列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playFeedback/getPlayFeedbackList [get]
func GetPlayFeedbackList(c *gin.Context) {
	var pageInfo request.PlayFeedbackSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPlayFeedbackInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败", zap.Any("err", err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
