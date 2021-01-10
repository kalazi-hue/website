package api

import (
	"gin-vue-admin/front/service"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Movie
// @Summary 首页->「最新更新」模块
// @accept application/json
// @Produce application/json
// @Param page 	query int true "page" default(1)
// @Param pageSize 	query int true "pageSize" default(10)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/getLatestMovieList [get]
func GetLatestMovieList(c *gin.Context) {
	var pageInfo request.MovieSearch
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetLatestMovieInfoList(pageInfo); err != nil {
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

// @Tags Movie
// @Summary 首页->「热门推荐」模块
// @accept application/json
// @Produce application/json
// @Param page 	query int true "page" default(1)
// @Param pageSize 	query int true "pageSize" default(10)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/getRecommendMovieList [get]
func GetRecommendMovieList(c *gin.Context) {
	var pageInfo request.MovieSearch
	if pageInfo.PageSize == 0 {
		pageInfo.PageSize = 10
	}
	if pageInfo.Page == 0 {
		pageInfo.Page = 1
	}
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetRecommendMovieList(pageInfo); err != nil {
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

// @Tags Movie
// @Summary 首页->顶部->「国产、日韩」模块
// @accept application/json
// @Produce application/json
// @Param page 	query int true "page" default(1)
// @Param pageSize 	query int true "pageSize" default(10)
// @Param typeId query string true "分类id" default(1)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/getMovieListByTypeId [get]
func GetMovieListByTypeId(c *gin.Context) {

	if err, list, total,page,pageSize := service.GetMovieListByTypeId(c); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, "获取成功", c)
	}
}


// @Tags Movie
// @Summary 影片详情页面，点击tag获取影片
// @accept application/json
// @Produce application/json
// @Param page 	query int true "page" default(1)
// @Param pageSize 	query int true "pageSize" default(10)
// @Param tags query string true "标签名称" default(t)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/getMovieListByTag [get]
func GetMovieListByTag(c *gin.Context) {
	if err, list, total,page,pageSize := service.GetMovieListByTag(c); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, "获取成功", c)
	}
}


// @Tags Movie
// @Summary 首页，关键词搜索影片
// @accept application/json
// @Produce application/json
// @Param page 	query int true "page" default(1)
// @Param pageSize 	query int true "pageSize" default(10)
// @Param keyword query string true "关键词" default(t)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/getMovieListByKeyword [get]
func GetMovieListByKeyword(c *gin.Context) {
	if err, list, total,page,pageSize := service.GetMovieListByKeyword(c); err != nil {
		global.GVA_LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, "获取成功", c)
	}
}


// @Tags Movie
// @Summary 影片点赞
// @accept application/json
// @Produce application/json
// @Param id query int true "影片id" default(1000)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"操作成功"}"
// @Router /movie/movieApproval [post]
func MovieApproval(c *gin.Context) {
	response.Ok(c)
}