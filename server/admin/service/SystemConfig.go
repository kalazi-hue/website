package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateSystemConfig
//@description: 创建SystemConfig记录
//@param: systemConfig model.SystemConfig
//@return: err error

func CreateSystemConfig(systemConfig model.SystemConfig) (err error) {
	err = global.GVA_DB.Create(&systemConfig).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSystemConfig
//@description: 删除SystemConfig记录
//@param: systemConfig model.SystemConfig
//@return: err error

func DeleteSystemConfig(systemConfig model.SystemConfig) (err error) {
	err = global.GVA_DB.Delete(systemConfig).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSystemConfigByIds
//@description: 批量删除SystemConfig记录
//@param: ids request.IdsReq
//@return: err error

func DeleteSystemConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.SystemConfig{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSystemConfig
//@description: 更新SystemConfig记录
//@param: systemConfig *model.SystemConfig
//@return: err error

func UpdateSystemConfig(systemConfig *model.SystemConfig) (err error) {
	err = global.GVA_DB.Save(systemConfig).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSystemConfig
//@description: 根据id获取SystemConfig记录
//@param: id uint
//@return: err error, systemConfig model.SystemConfig

func GetSysConfig(id uint) (err error, systemConfig model.SystemConfig) {
	err = global.GVA_DB.Where("id = ?", id).First(&systemConfig).Error
	return
}

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