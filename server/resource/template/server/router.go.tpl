package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func Init{{.StructName}}Router(Router *gin.RouterGroup) {
	{{.StructName}}Router := Router.Group("{{.Abbreviation}}").Use(middleware.JWTAuth())
	{
		{{.StructName}}Router.POST("create{{.StructName}}", api.Create{{.StructName}})   // 新建{{.StructName}}
		{{.StructName}}Router.DELETE("delete{{.StructName}}", api.Delete{{.StructName}}) // 删除{{.StructName}}
		{{.StructName}}Router.DELETE("delete{{.StructName}}ByIds", api.Delete{{.StructName}}ByIds) // 批量删除{{.StructName}}
		{{.StructName}}Router.PUT("update{{.StructName}}", api.Update{{.StructName}})    // 更新{{.StructName}}
		{{.StructName}}Router.GET("find{{.StructName}}", api.Find{{.StructName}})        // 根据ID获取{{.StructName}}
		{{.StructName}}Router.GET("get{{.StructName}}List", api.Get{{.StructName}}List)  // 获取{{.StructName}}列表
	}
}
