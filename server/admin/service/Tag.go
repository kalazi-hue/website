package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTag
//@description: 创建Tag记录
//@param: tag model.Tag
//@return: err error

func CreateTag(tag model.Tag) (err error) {
	err = global.GVA_DB.Create(&tag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTag
//@description: 删除Tag记录
//@param: tag model.Tag
//@return: err error

func DeleteTag(tag model.Tag) (err error) {
	err = global.GVA_DB.Delete(tag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteTagByIds
//@description: 批量删除Tag记录
//@param: ids request.IdsReq
//@return: err error

func DeleteTagByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.Tag{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTag
//@description: 更新Tag记录
//@param: tag *model.Tag
//@return: err error

func UpdateTag(tag *model.Tag) (err error) {
	err = global.GVA_DB.Save(tag).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTag
//@description: 根据id获取Tag记录
//@param: id uint
//@return: err error, tag model.Tag

func GetTag(id uint) (err error, tag model.Tag) {
	err = global.GVA_DB.Where("id = ?", id).First(&tag).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTagInfoList
//@description: 分页获取Tag记录
//@param: info request.TagSearch
//@return: err error, list interface{}, total int64

func GetTagInfoList(info request.TagSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.Tag{})
    var tags []model.Tag
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("`name` LIKE ?","%"+ info.Name+"%")
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&tags).Error
	return err, tags, total
}