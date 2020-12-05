package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAdConfig
//@description: 创建AdConfig记录
//@param: adConfig model.AdConfig
//@return: err error

func CreateAdConfig(adConfig model.AdConfig) (err error) {
	err = global.GVA_DB.Create(&adConfig).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAdConfig
//@description: 删除AdConfig记录
//@param: adConfig model.AdConfig
//@return: err error

func DeleteAdConfig(adConfig model.AdConfig) (err error) {
	err = global.GVA_DB.Delete(adConfig).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAdConfigByIds
//@description: 批量删除AdConfig记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAdConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.AdConfig{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAdConfig
//@description: 更新AdConfig记录
//@param: adConfig *model.AdConfig
//@return: err error

func UpdateAdConfig(adConfig *model.AdConfig) (err error) {
	err = global.GVA_DB.Save(adConfig).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAdConfig
//@description: 根据id获取AdConfig记录
//@param: id uint
//@return: err error, adConfig model.AdConfig

func GetAdConfig(id uint) (err error, adConfig model.AdConfig) {
	err = global.GVA_DB.Where("id = ?", id).First(&adConfig).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAdConfigInfoList
//@description: 分页获取AdConfig记录
//@param: info request.AdConfigSearch
//@return: err error, list interface{}, total int64

func GetAdConfigInfoList(info request.AdConfigSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.AdConfig{})
    var adConfigs []model.AdConfig
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Position != "" {
        db = db.Where("`position` = ?",info.Position)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&adConfigs).Error
	return err, adConfigs, total
}