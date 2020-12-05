package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@description: 分页获取Tag记录

func GetTagInfoList(info request.TagSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Tag{})
    var tags []model.Tag
    // 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`status` = ?",1)
	err = db.Count(&total).Error
	err = db.Order("sort desc").Limit(limit).Offset(offset).Find(&tags).Error
	return err, tags, total
}