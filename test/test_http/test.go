package test_http

import (
	"io"
	"rtsys/router"
	"rtsys/utils/core"

	"github.com/gin-gonic/gin"
)

func TestLoadEnv() *gin.Engine {
	core.TestLoad("../../")
	gin.SetMode(gin.ReleaseMode)

	gin.DefaultWriter = io.Discard
	engine := gin.New()
	//使用自定义格式的日志log

	engine.Use(gin.LoggerWithFormatter(core.CustomB5Log)).Use(gin.Recovery())

	router.LoadRouter(engine)
	return engine
}
