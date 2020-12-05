package request

import "gin-vue-admin/model"

type SystemConfigSearch struct{
    model.SystemConfig
    PageInfo
}