package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"github.com/gin-gonic/gin"
	"strconv"
)



//@description: 分页获取Movie记录

func GetMovieInfoList(info request.MovieSearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

    // 创建db
	db := global.GVA_DB.Model(&model.Movie{})
    var movies []model.Movie
    // 如果有条件搜索 下方会自动创建搜索语句

    if info.Type != "" {
        db = db.Where("`type` = ?",info.Type)
    }
    if info.IsRecommend != nil {
        db = db.Where("`is_recommend` = ?",info.IsRecommend)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&movies).Error
	return err, movies, total
}

//@description: 最近更新
func GetLatestMovieInfoList(info request.MovieSearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&model.Movie{})
	var movies []model.Movie
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`status` = ?",1)

	err = db.Count(&total).Error
	err = db.Order("shelf_time desc").Limit(limit).Offset(offset).Find(&movies).Error
	return err, movies, total
}


//@description: 热门推荐
func GetRecommendMovieList(info request.MovieSearch) (err error, list interface{}, total int64) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&model.Movie{})
	var movies []model.Movie
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`status` = ? and `is_recommend` = ?", 1,1)

	err = db.Count(&total).Error
	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&movies).Error
	return err, movies, total
}


//@description: 首页->国产、日韩分类列表
func GetMovieListByTypeId(c *gin.Context) (err error, list interface{}, total int64, page int, pageSize int) {
	var typeId string
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	typeId = c.DefaultQuery("typeId", "")
	limit := pageSize
	offset := pageSize * (page - 1)

	// 创建db
	db := global.GVA_DB.Model(&model.Movie{})
	var movies []model.Movie
	// 如果有条件搜索 下方会自动创建搜索语句
	if typeId == "" {
		db = db.Where("`status` = ?", 1)
	} else {
		db = db.Where("`status` = ? and `type` = ?", 1, typeId)
	}

	err = db.Count(&total).Error
	err = db.Order("shelf_time desc, id desc").Limit(limit).Offset(offset).Find(&movies).Error
	return err, movies, total, page, pageSize
}


//@description: 根据标签获取影片
func GetMovieListByTag(c *gin.Context) (err error, list interface{}, total int64, page int, pageSize int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	tags := c.DefaultQuery("tags", "t")
	limit := pageSize
	offset := pageSize * (page - 1)

	// 创建db
	db := global.GVA_DB.Model(&model.Movie{})
	var movies []model.Movie
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`status` = ? and `tags` like ?", 1, "%"+ tags +"%")

	err = db.Count(&total).Error
	err = db.Order("shelf_time desc, id desc").Limit(limit).Offset(offset).Find(&movies).Error
	return err, movies, total, page, pageSize
}


//@description: 根据关键词获取影片
func GetMovieListByKeyword(c *gin.Context) (err error, list interface{}, total int64, page int, pageSize int) {
	page, _ = strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ = strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	keyword := c.DefaultQuery("keyword", "t")
	limit := pageSize
	offset := pageSize * (page - 1)

	// 创建db
	db := global.GVA_DB.Model(&model.Movie{})
	var movies []model.Movie
	// 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`status` = ? and `keyword` like ?", 1, "%"+ keyword +"%")

	err = db.Count(&total).Error
	err = db.Order("shelf_time desc, id desc").Limit(limit).Offset(offset).Find(&movies).Error
	return err, movies, total, page, pageSize
}
