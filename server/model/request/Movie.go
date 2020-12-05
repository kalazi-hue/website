package request

import "gin-vue-admin/model"

type MovieSearch struct{
    model.Movie
    PageInfo
}