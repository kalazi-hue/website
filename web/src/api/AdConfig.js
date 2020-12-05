import service from '@/utils/request'

// @Tags AdConfig
// @Summary 创建AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdConfig true "创建AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adConfig/createAdConfig [post]
export const createAdConfig = (data) => {
     return service({
         url: "/adConfig/createAdConfig",
         method: 'post',
         data
     })
 }


// @Tags AdConfig
// @Summary 删除AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdConfig true "删除AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adConfig/deleteAdConfig [delete]
 export const deleteAdConfig = (data) => {
     return service({
         url: "/adConfig/deleteAdConfig",
         method: 'delete',
         data
     })
 }

// @Tags AdConfig
// @Summary 删除AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /adConfig/deleteAdConfig [delete]
 export const deleteAdConfigByIds = (data) => {
     return service({
         url: "/adConfig/deleteAdConfigByIds",
         method: 'delete',
         data
     })
 }

// @Tags AdConfig
// @Summary 更新AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdConfig true "更新AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /adConfig/updateAdConfig [put]
 export const updateAdConfig = (data) => {
     return service({
         url: "/adConfig/updateAdConfig",
         method: 'put',
         data
     })
 }


// @Tags AdConfig
// @Summary 用id查询AdConfig
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.AdConfig true "用id查询AdConfig"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /adConfig/findAdConfig [get]
 export const findAdConfig = (params) => {
     return service({
         url: "/adConfig/findAdConfig",
         method: 'get',
         params
     })
 }


// @Tags AdConfig
// @Summary 分页获取AdConfig列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取AdConfig列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /adConfig/getAdConfigList [get]
 export const getAdConfigList = (params) => {
     return service({
         url: "/adConfig/getAdConfigList",
         method: 'get',
         params
     })
 }