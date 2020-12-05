package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateMovie
//@description: 创建Movie记录
//@param: movie model.Movie
//@return: err error

func CreateMovie(movie model.Movie) (err error) {
	err = global.GVA_DB.Create(&movie).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMovie
//@description: 删除Movie记录
//@param: movie model.Movie
//@return: err error

func DeleteMovie(movie model.Movie) (err error) {
	err = global.GVA_DB.Delete(movie).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteMovieByIds
//@description: 批量删除Movie记录
//@param: ids request.IdsReq
//@return: err error

func DeleteMovieByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Movie{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateMovie
//@description: 更新Movie记录
//@param: movie *model.Movie
//@return: err error

func UpdateMovie(movie *model.Movie) (err error) {
	err = global.GVA_DB.Save(movie).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMovie
//@description: 根据id获取Movie记录
//@param: id uint
//@return: err error, movie model.Movie

func GetMovie(id uint) (err error, movie model.Movie) {
	err = global.GVA_DB.Where("id = ?", id).First(&movie).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetMovieInfoList
//@description: 分页获取Movie记录
//@param: info request.MovieSearch
//@return: err error, list interface{}, total int64

func GetMovieInfoList(info request.MovieSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Movie{})
    var movies []model.Movie
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Title != "" {
        db = db.Where("`title` LIKE ?","%"+ info.Title+"%")
    }
    if info.Tags != "" {
        db = db.Where("`tags` LIKE ?","%"+ info.Tags+"%")
    }
    if info.IsTop != nil {
        db = db.Where("`is_top` = ?",info.IsTop)
    }
    if info.IsRecommend != nil {
        db = db.Where("`is_recommend` = ?",info.IsRecommend)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&movies).Error
	return err, movies, total
}