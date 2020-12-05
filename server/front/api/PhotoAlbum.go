package api

import (
	"gin-vue-admin/front/service"
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)


// @Tags PhotoAlbum
// @Summary 点击类型查询写真
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page 	query int true "page" default(1)
// @Param pageSize 	query int true "pageSize" default(10)
// @Param type	query string true "type" default(1)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /photoAlbum/getPhotoAlbumListByType [get]
func GetPhotoAlbumListByType(c *gin.Context) {
	if err, list, total,page,pageSize := service.GetPhotoAlbumListByType(c); err != nil {
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

// @Tags PhotoAlbum
// @Summary 分页获取PhotoAlbum列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param page 	query int true "page" default(1)
// @Param pageSize 	query int true "pageSize" default(10)
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /photoAlbum/getPhotoAlbumList [get]
func GetPhotoAlbumList(c *gin.Context) {
	var pageInfo request.PhotoAlbumSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := service.GetPhotoAlbumInfoList(pageInfo); err != nil {
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
