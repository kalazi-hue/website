import service from '@/utils/request'

// @Tags SystemConfig
// @Summary 创建SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemConfig true "创建SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemConfig/createSystemConfig [post]
export const createSystemConfig = (data) => {
     return service({
         url: "/systemConfig/createSystemConfig",
         method: 'post',
         data
     })
 }


// @Tags SystemConfig
// @Summary 删除SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemConfig true "删除SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemConfig/deleteSystemConfig [delete]
 export const deleteSystemConfig = (data) => {
     return service({
         url: "/systemConfig/deleteSystemConfig",
         method: 'delete',
         data
     })
 }

// @Tags SystemConfig
// @Summary 删除SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemConfig/deleteSystemConfig [delete]
 export const deleteSystemConfigByIds = (data) => {
     return service({
         url: "/systemConfig/deleteSystemConfigByIds",
         method: 'delete',
         data
     })
 }

// @Tags SystemConfig
// @Summary 更新SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemConfig true "更新SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /systemConfig/updateSystemConfig [put]
 export const updateSystemConfig = (data) => {
     return service({
         url: "/systemConfig/updateSystemConfig",
         method: 'put',
         data
     })
 }


// @Tags SystemConfig
// @Summary 用id查询SystemConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SystemConfig true "用id查询SystemConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /systemConfig/findSystemConfig [get]
 export const findSystemConfig = (params) => {
     return service({
         url: "/systemConfig/findSystemConfig",
         method: 'get',
         params
     })
 }


// @Tags SystemConfig
// @Summary 分页获取SystemConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取SystemConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemConfig/getSystemConfigList [get]
 export const getSystemConfigList = (params) => {
     return service({
         url: "/systemConfig/getSystemConfigList",
         method: 'get',
         params
     })
 }