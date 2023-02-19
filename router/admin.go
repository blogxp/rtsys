// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package router

import (
	"net/http"
	"rtsys/admin/controller/common"
	"rtsys/admin/controller/demo"
	"rtsys/admin/controller/system"
	"rtsys/admin/lib"
	"rtsys/admin/middleware"
	"rtsys/utils/core"

	"github.com/gin-gonic/gin"
)

func (router *Router) Admin(engine *gin.Engine) {
	//加载管理端全局数据
	lib.AdminGlobalSetting(engine)

	//后端前缀
	adminPrefix := core.G_CONFIG.Route.Admin
	if adminPrefix == "" {
		adminPrefix = "/"
	}
	//路由分组 中间件
	var group *gin.RouterGroup
	group = engine.Group(adminPrefix, middleware.LoginAdminMiddleWare(), middleware.AuthAdminMiddleWare())

	//配置路由
	indexC := common.NewIndexController()
	//自动跳转到后台
	if core.G_CONFIG.Route.AutoAdmin {
		engine.GET(adminPrefix, func(ctx *gin.Context) {
			ctx.Redirect(http.StatusFound, indexC.ParseUrl("index", true))
		})
	}
	//common 分组
	indexC.Route(engine, group)
	common.NewPublicController().Route(engine, group)
	common.NewCommonController().Route(engine, group)

	//system分组
	system.NewNoticeController().Route(engine, group)
	system.NewLoginLogController().Route(engine, group)
	system.NewAdminController().Route(engine, group)
	system.NewMenuController().Route(engine, group)
	system.NewConfigController().Route(engine, group)
	system.NewStructController().Route(engine, group)
	system.NewPositionController().Route(engine, group)
	system.NewRoleController().Route(engine, group)

	//demo分组
	demo.NewGenController().Route(engine, group)
	demo.NewBuildController().Route(engine, group)
	demo.NewMediaController().Route(engine, group)
	demo.NewTestInfoController().Route(engine, group)

	//__new_router__
}
