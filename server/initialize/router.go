package initialize

import (
	_ "gin-vue-admin/docs"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/admin/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	Fr "gin-vue-admin/front/router"
	"gin-vue-admin/docs"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")
	docs.SwaggerInfo.Title = "啊啊啊啊"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "2.0"
	docs.SwaggerInfo.Host = "127.0.0.1:9999"
	docs.SwaggerInfo.BasePath = "/api"

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	AdminGroup := Router.Group("admin")
	router.InitUserRouter(AdminGroup)                  // 注册用户路由
	router.InitBaseRouter(AdminGroup)                  // 注册基础功能路由 不做鉴权
	router.InitMenuRouter(AdminGroup)                  // 注册menu路由
	router.InitAuthorityRouter(AdminGroup)             // 注册角色路由
	router.InitApiRouter(AdminGroup)                   // 注册功能api路由
	router.InitFileUploadAndDownloadRouter(AdminGroup) // 文件上传下载功能路由
	router.InitSimpleUploaderRouter(AdminGroup)        // 断点续传（插件版）
	router.InitWorkflowRouter(AdminGroup)              // 工作流相关路由
	router.InitCasbinRouter(AdminGroup)                // 权限相关路由
	router.InitJwtRouter(AdminGroup)                   // jwt相关路由
	router.InitSystemRouter(AdminGroup)                // system相关路由
	router.InitCustomerRouter(AdminGroup)              // 客户路由
	router.InitAutoCodeRouter(AdminGroup)              // 创建自动化代码
	router.InitSysDictionaryDetailRouter(AdminGroup)   // 字典详情管理
	router.InitSysDictionaryRouter(AdminGroup)         // 字典管理
	router.InitSysOperationRecordRouter(AdminGroup)    // 操作记录
	router.InitEmailRouter(AdminGroup)                 // 邮件相关路由

	router.InitAdConfigRouter(AdminGroup)				//广告
	router.InitMovieRouter(AdminGroup)				//影片
	router.InitTagRouter(AdminGroup)				//标签
	router.InitSystemConfigRouter(AdminGroup)				//标签
	router.InitPhotoAlbumRouter(AdminGroup)			//写真
	global.GVA_LOG.Info("admin router register success")

	ApiGroup := Router.Group("api")
	Fr.FrontAdConfigRouter(ApiGroup)
	Fr.FrontMovieRouter(ApiGroup)
	Fr.FrontTagRouter(ApiGroup)
	Fr.FrontSystemConfigRouter(ApiGroup)
	Fr.FrontPhotoAlbumRouter(ApiGroup)
	global.GVA_LOG.Info("api router register success")

	return Router
}
