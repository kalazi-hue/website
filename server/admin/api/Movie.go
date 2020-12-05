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

// @Tags Movie
// @Summary 创建Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Movie true "创建Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/createMovie [post]
func CreateMovie(c *gin.Context) {
	var movie model.Movie
	_ = c.ShouldBindJSON(&movie)
	if err := service.CreateMovie(movie); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags Movie
// @Summary 删除Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Movie true "删除Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /movie/deleteMovie [delete]
func DeleteMovie(c *gin.Context) {
	var movie model.Movie
	_ = c.ShouldBindJSON(&movie)
	if err := service.DeleteMovie(movie); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Movie
// @Summary 批量删除Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /movie/deleteMovieByIds [delete]
func DeleteMovieByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeleteMovieByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Movie
// @Summary 更新Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Movie true "更新Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /movie/updateMovie [put]
func UpdateMovie(c *gin.Context) {
	var movie model.Movie
	_ = c.ShouldBindJSON(&movie)
	if err := service.UpdateMovie(&movie); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags Movie
// @Summary 用id查询Movie
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Movie true "用id查询Movie"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /movie/findMovie [get]
func FindMovie(c *gin.Context) {
	var movie model.Movie
	_ = c.ShouldBindQuery(&movie)
	if err, removie := service.GetMovie(movie.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"removie": removie}, c)
	}
}

// @Tags Movie
// @Summary 分页获取Movie列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.MovieSearch true "分页获取Movie列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /movie/getMovieList [get]
func GetMovieList(c *gin.Context) {
	var pageInfo request.MovieSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetMovieInfoList(pageInfo); err != nil {
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
