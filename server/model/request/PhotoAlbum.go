package request

import "gin-vue-admin/model"

type PhotoAlbumSearch struct{
    model.PhotoAlbum
    PageInfo
}