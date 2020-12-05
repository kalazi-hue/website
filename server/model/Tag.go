// 自动生成模板Tag
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Tag struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:标签名称;type:varchar(200);size:200;"`
      Sort  int `json:"sort" form:"sort" gorm:"column:sort;comment:排序权重，值越高越靠前;type:mediumint;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:启用状态;type:tinyint;"`
}


func (Tag) TableName() string {
  return "tag"
}
