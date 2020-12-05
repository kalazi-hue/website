package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreatePhotoAlbum
//@description: 创建PhotoAlbum记录
//@param: photoAlbum model.PhotoAlbum
//@return: err error

func CreatePhotoAlbum(photoAlbum model.PhotoAlbum) (err error) {
	err = global.GVA_DB.Create(&photoAlbum).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePhotoAlbum
//@description: 删除PhotoAlbum记录
//@param: photoAlbum model.PhotoAlbum
//@return: err error

func DeletePhotoAlbum(photoAlbum model.PhotoAlbum) (err error) {
	err = global.GVA_DB.Delete(photoAlbum).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeletePhotoAlbumByIds
//@description: 批量删除PhotoAlbum记录
//@param: ids request.IdsReq
//@return: err error

func DeletePhotoAlbumByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]model.PhotoAlbum{},"id in ?",ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdatePhotoAlbum
//@description: 更新PhotoAlbum记录
//@param: photoAlbum *model.PhotoAlbum
//@return: err error

func UpdatePhotoAlbum(photoAlbum *model.PhotoAlbum) (err error) {
	err = global.GVA_DB.Save(photoAlbum).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPhotoAlbum
//@description: 根据id获取PhotoAlbum记录
//@param: id uint
//@return: err error, photoAlbum model.PhotoAlbum

func GetPhotoAlbum(id uint) (err error, photoAlbum model.PhotoAlbum) {
	err = global.GVA_DB.Where("id = ?", id).First(&photoAlbum).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetPhotoAlbumInfoList
//@description: 分页获取PhotoAlbum记录
//@param: info request.PhotoAlbumSearch
//@return: err error, list interface{}, total int64

func GetPhotoAlbumInfoList(info request.PhotoAlbumSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&model.PhotoAlbum{})
    var photoAlbums []model.PhotoAlbum
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Is_top != nil {
        db = db.Where("`is_top` = ?",info.Is_top)
    }
    if info.Status != nil {
        db = db.Where("`status` = ?",info.Status)
    }
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&photoAlbums).Error
	return err, photoAlbums, total
}