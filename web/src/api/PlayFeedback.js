import service from '@/utils/request'

// @Tags PlayFeedback
// @Summary 创建PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayFeedback true "创建PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playFeedback/createPlayFeedback [post]
export const createPlayFeedback = (data) => {
     return service({
         url: "/playFeedback/createPlayFeedback",
         method: 'post',
         data
     })
 }


// @Tags PlayFeedback
// @Summary 删除PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayFeedback true "删除PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playFeedback/deletePlayFeedback [delete]
 export const deletePlayFeedback = (data) => {
     return service({
         url: "/playFeedback/deletePlayFeedback",
         method: 'delete',
         data
     })
 }

// @Tags PlayFeedback
// @Summary 删除PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /playFeedback/deletePlayFeedback [delete]
 export const deletePlayFeedbackByIds = (data) => {
     return service({
         url: "/playFeedback/deletePlayFeedbackByIds",
         method: 'delete',
         data
     })
 }

// @Tags PlayFeedback
// @Summary 更新PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayFeedback true "更新PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /playFeedback/updatePlayFeedback [put]
 export const updatePlayFeedback = (data) => {
     return service({
         url: "/playFeedback/updatePlayFeedback",
         method: 'put',
         data
     })
 }


// @Tags PlayFeedback
// @Summary 用id查询PlayFeedback
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PlayFeedback true "用id查询PlayFeedback"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /playFeedback/findPlayFeedback [get]
 export const findPlayFeedback = (params) => {
     return service({
         url: "/playFeedback/findPlayFeedback",
         method: 'get',
         params
     })
 }


// @Tags PlayFeedback
// @Summary 分页获取PlayFeedback列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PlayFeedback列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /playFeedback/getPlayFeedbackList [get]
 export const getPlayFeedbackList = (params) => {
     return service({
         url: "/playFeedback/getPlayFeedbackList",
         method: 'get',
         params
     })
 }