package router

import (
	"gin-vue-admin/admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

func InitAuthorityRouter(Router *gin.RouterGroup) {
	AuthorityRouter := Router.Group("authority").
		Use(middleware.JWTAuth())
	{
		AuthorityRouter.POST("createAuthority", api.CreateAuthority)   // 创建角色
		AuthorityRouter.POST("deleteAuthority", api.DeleteAuthority)   // 删除角色
		AuthorityRouter.PUT("updateAuthority", api.UpdateAuthority)    // 更新角色
		AuthorityRouter.POST("copyAuthority", api.CopyAuthority)       // 更新角色
		AuthorityRouter.POST("getAuthorityList", api.GetAuthorityList) // 获取角色列表
		AuthorityRouter.POST("setDataAuthority", api.SetDataAuthority) // 设置角色资源权限
	}
}
