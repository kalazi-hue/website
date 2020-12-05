package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitSysDictionaryRouter(Router *gin.RouterGroup) {
	SysDictionaryRouter := Router.Group("sysDictionary").
		Use(middleware.JWTAuth()).
		Use(middleware.OperationRecord())
	{
		SysDictionaryRouter.POST("createSysDictionary", api.CreateSysDictionary)   // 新建SysDictionary
		SysDictionaryRouter.DELETE("deleteSysDictionary", api.DeleteSysDictionary) // 删除SysDictionary
		SysDictionaryRouter.PUT("updateSysDictionary", api.UpdateSysDictionary)    // 更新SysDictionary
		SysDictionaryRouter.GET("findSysDictionary", api.FindSysDictionary)        // 根据ID获取SysDictionary
		SysDictionaryRouter.GET("getSysDictionaryList", api.GetSysDictionaryList)  // 获取SysDictionary列表
	}
}
