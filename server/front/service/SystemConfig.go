package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSystemConfigInfoList
//@description: 分页获取SystemConfig记录
//@param: info request.SystemConfigSearch
//@return: err error, list interface{}, total int64

func GetSystemConfigInfoList(info request.SystemConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.SystemConfig{})
    var systemConfigs []model.SystemConfig
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&systemConfigs).Error
	return err, systemConfigs, total
}