package request

import "gin-vue-admin/model"

type TagSearch struct{
    model.Tag
    PageInfo
}