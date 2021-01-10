// 自动生成模板PlayFeedback
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type PlayFeedback struct {
      global.GVA_MODEL
      MovieId  int `json:"movieId" form:"movieId" gorm:"column:movie_id;comment:;type:bigint;size:19;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:;type:varchar(200);size:200;"`
      PlayUrl  string `json:"playUrl" form:"playUrl" gorm:"column:play_url;comment:;type:varchar(200);size:200;"`
      Type  string `json:"type" form:"type" gorm:"column:type;comment:;type:varchar(50);size:50;"`
}


func (PlayFeedback) TableName() string {
  return "play_feedback"
}
