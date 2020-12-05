import service from '@/utils/request'

// @Tags Movie
// @Summary 创建Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Movie true "创建Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/createMovie [post]
export const createMovie = (data) => {
     return service({
         url: "/movie/createMovie",
         method: 'post',
         data
     })
 }


// @Tags Movie
// @Summary 删除Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Movie true "删除Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /movie/deleteMovie [delete]
 export const deleteMovie = (data) => {
     return service({
         url: "/movie/deleteMovie",
         method: 'delete',
         data
     })
 }

// @Tags Movie
// @Summary 删除Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /movie/deleteMovie [delete]
 export const deleteMovieByIds = (data) => {
     return service({
         url: "/movie/deleteMovieByIds",
         method: 'delete',
         data
     })
 }

// @Tags Movie
// @Summary 更新Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Movie true "更新Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /movie/updateMovie [put]
 export const updateMovie = (data) => {
     return service({
         url: "/movie/updateMovie",
         method: 'put',
         data
     })
 }


// @Tags Movie
// @Summary 用id查询Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Movie true "用id查询Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /movie/findMovie [get]
 export const findMovie = (params) => {
     return service({
         url: "/movie/findMovie",
         method: 'get',
         params
     })
 }


// @Tags Movie
// @Summary 分页获取Movie列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取Movie列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/getMovieList [get]
 export const getMovieList = (params) => {
     return service({
         url: "/movie/getMovieList",
         method: 'get',
         params
     })
 }