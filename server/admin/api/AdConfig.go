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

// @Tags AdConfig
// @Summary 创建AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdConfig true "创建AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adConfig/createAdConfig [post]
func CreateAdConfig(c *gin.Context) {
	var adConfig model.AdConfig
	_ = c.ShouldBindJSON(&adConfig)
	if err := service.CreateAdConfig(adConfig); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags AdConfig
// @Summary 删除AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdConfig true "删除AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adConfig/deleteAdConfig [delete]
func DeleteAdConfig(c *gin.Context) {
	var adConfig model.AdConfig
	_ = c.ShouldBindJSON(&adConfig)
	if err := service.DeleteAdConfig(adConfig); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags AdConfig
// @Summary 批量删除AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /adConfig/deleteAdConfigByIds [delete]
func DeleteAdConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteAdConfigByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags AdConfig
// @Summary 更新AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdConfig true "更新AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adConfig/updateAdConfig [put]
func UpdateAdConfig(c *gin.Context) {
	var adConfig model.AdConfig
	_ = c.ShouldBindJSON(&adConfig)
	if err := service.UpdateAdConfig(&adConfig); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags AdConfig
// @Summary 用id查询AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdConfig true "用id查询AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adConfig/findAdConfig [get]
func FindAdConfig(c *gin.Context) {
	var adConfig model.AdConfig
	_ = c.ShouldBindQuery(&adConfig)
	if err, readConfig := service.GetAdConfig(adConfig.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"readConfig": readConfig}, c)
	}
}

// @Tags AdConfig
// @Summary 分页获取AdConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AdConfigSearch true "分页获取AdConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adConfig/getAdConfigList [get]
func GetAdConfigList(c *gin.Context) {
	var pageInfo request.AdConfigSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetAdConfigInfoList(pageInfo); err != nil {
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
