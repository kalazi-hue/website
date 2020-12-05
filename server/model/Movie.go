// 自动生成模板Movie
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type Movie struct {
      global.GVA_MODEL
      Title  string `json:"title" form:"title" gorm:"column:title;comment:片名;type:varchar(200);size:200;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:影片简介;type:varchar(200);size:200;"`
      Keyword  string `json:"keyword" form:"keyword" gorm:"column:keyword;comment:关键词;type:varchar(200);size:200;"`
      Cover  string `json:"cover" form:"cover" gorm:"column:cover;comment:封面图;type:varchar(200);size:200;"`
      PlayTime  string `json:"playTime" form:"playTime" gorm:"column:play_time;comment:影片时长;type:varchar(200);size:200;"`
      PlayUrl  string `json:"playUrl" form:"playUrl" gorm:"column:play_url;comment:播放链接;type:varchar(200);size:200;"`
      DownUrl  string `json:"downUrl" form:"downUrl" gorm:"column:down_url;comment:下载链接;type:varchar(200);size:200;"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:影片类型;type:bigint;size:19;"`
      Tags  string `json:"tags" form:"tags" gorm:"column:tags;comment:标签;type:varchar(200);size:200;"`
      Star  int `json:"star" form:"star" gorm:"column:star;comment:点赞数;type:bigint;size:19;"`
      PlayCount  int `json:"playCount" form:"playCount" gorm:"column:play_count;comment:观影次数;type:bigint;size:19;"`
      IsTop  *bool `json:"isTop" form:"isTop" gorm:"column:is_top;comment:是否置顶;type:tinyint;"`
      IsRecommend  *bool `json:"isRecommend" form:"isRecommend" gorm:"column:is_recommend;comment:是否推荐;type:tinyint;"`
      ShelfTime  time.Time `json:"shelfTime" form:"shelfTime" gorm:"column:shelf_time;comment:上架时间;type:datetime;"`
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:是否上架;type:tinyint;"`
}


func (Movie) TableName() string {
  return "movie"
}
