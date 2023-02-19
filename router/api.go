package router

import (
	"rtsys/api/controller/store"
	"rtsys/api/controller/user"
	"rtsys/api/middleware"
	"rtsys/utils/core"

	"github.com/gin-gonic/gin"
)

func (router *Router) Api(engine *gin.Engine) {
	//接口前缀
	apiPrefix := core.G_CONFIG.Route.Api
	if apiPrefix == "" || apiPrefix == "/" {
		apiPrefix = "/api"
	}
	baseGroup := engine.Group(apiPrefix)

	//user模块
	groupUser := baseGroup.Group("/user")
	{
		groupUser.Use(middleware.NewUserLoginMiddleWare().Handle())
		user.NewPublicApi().Route(engine, groupUser)
	}

	//store模块
	groupStore := baseGroup.Group("/store")
	{
		groupStore.Use(middleware.NewStoreLoginMiddleWare().Handle())
		store.NewIndexApiStore().Route(engine, groupStore)
		store.NewPublicApiStore().Route(engine, groupStore)
	}

}
