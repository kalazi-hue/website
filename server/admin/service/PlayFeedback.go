package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePlayFeedback
//@description: 创建PlayFeedback记录
//@param: playFeedback model.PlayFeedback
//@return: err error

func CreatePlayFeedback(playFeedback model.PlayFeedback) (err error) {
	err = global.GVA_DB.Create(&playFeedback).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePlayFeedback
//@description: 删除PlayFeedback记录
//@param: playFeedback model.PlayFeedback
//@return: err error

func DeletePlayFeedback(playFeedback model.PlayFeedback) (err error) {
	err = global.GVA_DB.Delete(playFeedback).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePlayFeedbackByIds
//@description: 批量删除PlayFeedback记录
//@param: ids request.IdsReq
//@return: err error

func DeletePlayFeedbackByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PlayFeedback{},"id in ?",ids.Ids).Error
	return err
}


//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPlayFeedback
//@description: 根据id获取PlayFeedback记录
//@param: id uint
//@return: err error, playFeedback model.PlayFeedback

func GetPlayFeedback(id uint) (err error, playFeedback model.PlayFeedback) {
	err = global.GVA_DB.Where("id = ?", id).First(&playFeedback).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPlayFeedbackInfoList
//@description: 分页获取PlayFeedback记录
//@param: info request.PlayFeedbackSearch
//@return: err error, list interface{}, total int64

func GetPlayFeedbackInfoList(info request.PlayFeedbackSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PlayFeedback{})
    var playFeedbacks []model.PlayFeedback
    // 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&playFeedbacks).Error
	return err, playFeedbacks, total
}