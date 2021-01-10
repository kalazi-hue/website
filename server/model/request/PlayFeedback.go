package request

import "gin-vue-admin/model"

type PlayFeedbackSearch struct{
    model.PlayFeedback
    PageInfo
}