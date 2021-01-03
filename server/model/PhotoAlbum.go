// 自动生成模板PhotoAlbum
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type PhotoAlbum struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:写真名称;type:varchar(200);size:200;"`
      Cover  string `json:"cover" form:"cover" gorm:"column:cover;comment:封面;type:varchar(200);size:200;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:简介;type:varchar(200);size:200;"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:类型;type:varchar(50);size:50;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`
      Keyword  string `json:"keyword" form:"keyword" gorm:"column:keyword;comment:关键词;type:varchar(200);size:200;"`
      Is_top  *bool `json:"isTop" form:"isTop" gorm:"column:is_top;comment:is_top;type:tinyint;"`
      Reads  int `json:"reads" form:"reads" gorm:"column:reads;comment:点击量;type:bigint;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:是否上架;type:tinyint;"`
}


func (PhotoAlbum) TableName() string {
  return "photo_album"
}
