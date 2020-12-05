// 自动生成模板SystemConfig
package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type SystemConfig struct {
      global.GVA_MODEL
      Notice  string `json:"notice" form:"notice" gorm:"column:notice;comment:公告1;type:text;"`
      Notice2  string `json:"notice2" form:"notice2" gorm:"column:notice2;comment:公告2;type:text;"`
      LastAccessDomain  string `json:"lastAccessDomain" form:"lastAccessDomain" gorm:"column:last_access_domain;comment:最新访问域名;type:varchar(200);size:200;"`
      LastDownDomain  string `json:"lastDownDomain" form:"lastDownDomain" gorm:"column:last_down_domain;comment:最新下载域名;type:varchar(200);size:200;"`
      ExchangeGroup  string `json:"exchangeGroup" form:"exchangeGroup" gorm:"column:exchange_group;comment:交流群;type:varchar(100);size:100;"`
      GitHub  string `json:"gitHub" form:"gitHub" gorm:"column:git_hub;comment:GitHub地址;type:varchar(200);size:200;"`
      Email  string `json:"email" form:"email" gorm:"column:email;comment:反馈邮箱;type:varchar(200);size:200;"`
      AlternateAddress  string `json:"alternateAddress" form:"alternateAddress" gorm:"column:alternate_address;comment:备用地址;type:varchar(200);size:200;"`
      PrivacyPolicy  string `json:"privacyPolicy" form:"privacyPolicy" gorm:"column:privacy_policy;comment:隐私政策;type:text;"`
      Contact  string `json:"contact" form:"contact" gorm:"column:contact;comment:联系我们;type:varchar(200);size:200;"`
}


func (SystemConfig) TableName() string {
  return "system_config"
}
