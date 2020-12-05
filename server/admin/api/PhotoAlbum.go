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

// @Tags PhotoAlbum
// @Summary 创建PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PhotoAlbum true "创建PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /photoAlbum/createPhotoAlbum [post]
func CreatePhotoAlbum(c *gin.Context) {
	var photoAlbum model.PhotoAlbum
	_ = c.ShouldBindJSON(&photoAlbum)
	if err := service.CreatePhotoAlbum(photoAlbum); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// @Tags PhotoAlbum
// @Summary 删除PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PhotoAlbum true "删除PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /photoAlbum/deletePhotoAlbum [delete]
func DeletePhotoAlbum(c *gin.Context) {
	var photoAlbum model.PhotoAlbum
	_ = c.ShouldBindJSON(&photoAlbum)
	if err := service.DeletePhotoAlbum(photoAlbum); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags PhotoAlbum
// @Summary 批量删除PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /photoAlbum/deletePhotoAlbumByIds [delete]
func DeletePhotoAlbumByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := service.DeletePhotoAlbumByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags PhotoAlbum
// @Summary 更新PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PhotoAlbum true "更新PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /photoAlbum/updatePhotoAlbum [put]
func UpdatePhotoAlbum(c *gin.Context) {
	var photoAlbum model.PhotoAlbum
	_ = c.ShouldBindJSON(&photoAlbum)
	if err := service.UpdatePhotoAlbum(&photoAlbum); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// @Tags PhotoAlbum
// @Summary 用id查询PhotoAlbum
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.PhotoAlbum true "用id查询PhotoAlbum"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /photoAlbum/findPhotoAlbum [get]
func FindPhotoAlbum(c *gin.Context) {
	var photoAlbum model.PhotoAlbum
	_ = c.ShouldBindQuery(&photoAlbum)
	if err, rephotoAlbum := service.GetPhotoAlbum(photoAlbum.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rephotoAlbum": rephotoAlbum}, c)
	}
}

// @Tags PhotoAlbum
// @Summary 分页获取PhotoAlbum列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PhotoAlbumSearch true "分页获取PhotoAlbum列表"
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
