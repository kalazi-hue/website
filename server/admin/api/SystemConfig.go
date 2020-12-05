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

// @Tags SystemConfig
// @Summary 创建SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemConfig true "创建SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemConfig/createSystemConfig [post]
func CreateSystemConfig(c *gin.Context) {
	var systemConfig model.SystemConfig
	_ = c.ShouldBindJSON(&systemConfig)
	if err := service.CreateSystemConfig(systemConfig); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags SystemConfig
// @Summary 删除SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemConfig true "删除SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemConfig/deleteSystemConfig [delete]
func DeleteSystemConfig(c *gin.Context) {
	var systemConfig model.SystemConfig
	_ = c.ShouldBindJSON(&systemConfig)
	if err := service.DeleteSystemConfig(systemConfig); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags SystemConfig
// @Summary 批量删除SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /systemConfig/deleteSystemConfigByIds [delete]
func DeleteSystemConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteSystemConfigByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags SystemConfig
// @Summary 更新SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemConfig true "更新SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /systemConfig/updateSystemConfig [put]
func UpdateSystemConfig(c *gin.Context) {
	var systemConfig model.SystemConfig
	_ = c.ShouldBindJSON(&systemConfig)
	if err := service.UpdateSystemConfig(&systemConfig); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags SystemConfig
// @Summary 用id查询SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemConfig true "用id查询SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /systemConfig/findSystemConfig [get]
func FindSystemConfig(c *gin.Context) {
	var systemConfig model.SystemConfig
	_ = c.ShouldBindQuery(&systemConfig)
	if err, resystemConfig := service.GetSysConfig(systemConfig.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resystemConfig": resystemConfig}, c)
	}
}

// @Tags SystemConfig
// @Summary 分页获取SystemConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.SystemConfigSearch true "分页获取SystemConfig列表"
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
