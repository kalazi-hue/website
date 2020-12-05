// 自动生成模板AdConfig
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type AdConfig struct {
      global.GVA_MODEL
      Position  string `json:"position" form:"position" gorm:"column:position;comment:广告位置;type:varchar(200);size:200;"`
      Url  string `json:"url" form:"url" gorm:"column:url;comment:图片链接;type:varchar(200);size:200;"`
      PlayUrl  string `json:"playUrl" form:"playUrl" gorm:"column:play_url;comment:播放链接;type:varchar(200);size:200;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:内容描述;type:varchar(200);size:200;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:启用状态;type:tinyint;"`
}


func (AdConfig) TableName() string {
  return "ad_config"
}
