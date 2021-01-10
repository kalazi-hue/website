package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
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
