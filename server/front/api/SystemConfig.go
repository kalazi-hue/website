package api

import (
	"gin-vue-admin/global"
    "gin-vue-admin/model/request"
    "gin-vue-admin/model/response"
    "gin-vue-admin/front/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

// @Tags SystemConfig
// @Summary 分页获取SystemConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page 	query int true "page" default(1)
// @Param pageSize 	query int true "pageSize" default(10)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemConfig/getSystemConfigList [get]
func GetSystemConfigList(c *gin.Context) {
	var pageInfo request.SystemConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetSystemConfigInfoList(pageInfo); err != nil {
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
