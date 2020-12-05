package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)


//@description: 分页获取AdConfig记录


func GetAdConfigInfoList(info request.AdConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AdConfig{})
    var adConfigs []model.AdConfig
    // 如果有条件搜索 下方会自动创建搜索语句
	db = db.Where("`status` = ?",1)
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adConfigs).Error
	return err, adConfigs, total
}