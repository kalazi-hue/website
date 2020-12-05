import service from '@/utils/request'

// @Tags PhotoAlbum
// @Summary 创建PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PhotoAlbum true "创建PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /photoAlbum/createPhotoAlbum [post]
export const createPhotoAlbum = (data) => {
     return service({
         url: "/photoAlbum/createPhotoAlbum",
         method: 'post',
         data
     })
 }


// @Tags PhotoAlbum
// @Summary 删除PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PhotoAlbum true "删除PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /photoAlbum/deletePhotoAlbum [delete]
 export const deletePhotoAlbum = (data) => {
     return service({
         url: "/photoAlbum/deletePhotoAlbum",
         method: 'delete',
         data
     })
 }

// @Tags PhotoAlbum
// @Summary 删除PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /photoAlbum/deletePhotoAlbum [delete]
 export const deletePhotoAlbumByIds = (data) => {
     return service({
         url: "/photoAlbum/deletePhotoAlbumByIds",
         method: 'delete',
         data
     })
 }

// @Tags PhotoAlbum
// @Summary 更新PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PhotoAlbum true "更新PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /photoAlbum/updatePhotoAlbum [put]
 export const updatePhotoAlbum = (data) => {
     return service({
         url: "/photoAlbum/updatePhotoAlbum",
         method: 'put',
         data
     })
 }


// @Tags PhotoAlbum
// @Summary 用id查询PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PhotoAlbum true "用id查询PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /photoAlbum/findPhotoAlbum [get]
 export const findPhotoAlbum = (params) => {
     return service({
         url: "/photoAlbum/findPhotoAlbum",
         method: 'get',
         params
     })
 }


// @Tags PhotoAlbum
// @Summary 分页获取PhotoAlbum列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取PhotoAlbum列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /photoAlbum/getPhotoAlbumList [get]
 export const getPhotoAlbumList = (params) => {
     return service({
         url: "/photoAlbum/getPhotoAlbumList",
         method: 'get',
         params
     })
 }