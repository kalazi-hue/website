package request

import "gin-vue-admin/model"

type AdConfigSearch struct{
    model.AdConfig
    PageInfo
}