package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/gin-gonic/gin"
	"strconv"
)


//@description: 根据关键词获取写真
func GetPhotoAlbumListByType(c *gin.Context) (err error, list interface{}, total int64, page int, pageSize int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	typeId := c.DefaultQuery("type", "1")
	limit := pageSize
	offset := pageSize * (page - 1)

	// 创建db
	db := global.GVA_DB.Model(&model.Movie{})
	var movies []model.Movie
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`status` = ? and `type` =?", 1, typeId)

	err = db.Count(&total).Error
	err = db.Order("shelf_time desc, id desc").Limit(limit).Offset(offset).Find(&movies).Error
	return err, movies, total, page, pageSize
}


//@description: 分页获取PhotoAlbum记录


func GetPhotoAlbumInfoList(info request.PhotoAlbumSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PhotoAlbum{})
    var photoAlbums []model.PhotoAlbum
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Is_top != nil {
        db = db.Where("`is_top` = ?",info.Is_top)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&photoAlbums).Error
	return err, photoAlbums, total
}