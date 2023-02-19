// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

//////      简单搭建的操作登录判断结构体   嵌套BaseApi   实际控制器可使用该结构体取代BaseApi        ///////

package store

import (
	"rtsys/api/lib"
	"rtsys/api/services"
	"rtsys/utils/tool"

	"github.com/gin-gonic/gin"
)

type LoginCheck struct {
	NoLoginActions []string //不需要登录判断的方法数组
	lib.BaseApi
}

// Handle 登录判断，未登录直接返回json
func (l *LoginCheck) Handle(ctx *gin.Context, action string) bool {
	if tool.InArray(action, l.NoLoginActions) {
		return true
	}
	appToken := services.GetApiLoginInfo(ctx)
	if appToken == nil {
		l.Error(ctx, "登录失效，请重新登录", nil, lib.ANoLogin)
		return false
	}
	return true
}
