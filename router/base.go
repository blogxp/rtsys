// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package router

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
}

func LoadRouter( engine *gin.Engine) {
	engine.Static("/static/", "./static/")
	engine.Static("/uploads/", "./static/uploads/") //上传文件
	engine.StaticFile("/404.html", "./static/404.html")
	engine.StaticFile("/favicon.ico", "./static/favicon.ico")

	router := &Router{}

	router.Api(engine) //接口路由
	router.Admin(engine) //后端路由

}
