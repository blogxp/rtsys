// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package lib

// ** 后台公共控制器 ** //

import (
	"rtsys/utils/core"
	"strings"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	Id     string //控制器名称
	Group  string //控制器分组
	IModel core.IModel
	HookAction
}

// Dispatch 路由分配器
// 生成具体访问ulr
// 调用控制器构造函数 Construct
func (ca *Controller) Dispatch(action string, prefix bool, handler func(*gin.Context)) (string, gin.HandlerFunc) {
	path := ca.ParseUrl(action, prefix)
	return path, func(ctx *gin.Context) {
		n := ca.Construct(ctx, action)
		if n {
			handler(ctx)
		}
	}
}

// Construct 构造函数
// 赋值属性及调用前置操作BeforeAction
func (ca *Controller) Construct(ctx *gin.Context, action string) bool {
	if !ca.BeforeAction(ctx, action) {
		return false
	}
	if ca.HookAction.BeforeAction != nil {
		return ca.HookAction.BeforeAction(ctx, action)
	}
	return true
}

// BeforeAction 全局方法调用前操作
func (ca *Controller) BeforeAction(ctx *gin.Context, action string) bool {
	return true
}

// Success 可变参数args 的格式顺序约定为
//
//	msg string
//	data map[string]any
//	code int
func (ca *Controller) Success(ctx *gin.Context, args ...any) {
	JsonSuccess(ctx, args...)
}

// Error 可变参数args 的格式顺序约定为
//
//	msg string
//	data map[string]any
//	code int
func (ca *Controller) Error(ctx *gin.Context, args ...any) {
	JsonFail(ctx, args...)
}

// Jump 跳转到错误页
func (ca *Controller) Jump(ctx *gin.Context, err error) {
	ErrorHtml(ctx, err)
}

// Render 解析并展示文件 args为额外的
func (ca *Controller) Render(ctx *gin.Context, html string, data map[string]any, args ...string) {
	htmlPath := ca.Id + "/" + html
	boot := strings.Trim(core.G_CONFIG.Route.Admin, "/")
	group := ""
	if boot != "" {
		boot = "/" + boot
		group = boot
	}

	if ca.Group != "" {
		group = group + "/" + ca.Group
		htmlPath = ca.Group + "/" + htmlPath
	}
	if data == nil {
		data = gin.H{}
	}

	//添加路径参数
	data["_root_"] = boot
	data["_group_"] = group
	data["_controller_"] = group + "/" + ca.Id

	//添加部分自定义方法
	data["func"] = TempFunc()

	//解析html
	htmlList := []string{htmlPath}
	for _, v := range args {
		if v != "" {
			htmlList = append(htmlList, v)
		}
	}
	temp := ParseHtml(htmlList...)
	_ = temp.ExecuteTemplate(ctx.Writer, htmlPath, data)
}

// ParseUrl 快捷生成控制器访问的url，用户生成路由
func (ca *Controller) ParseUrl(action string, prefix bool) string {
	path := ca.Id + "/" + action
	if ca.Group != "" {
		path = ca.Group + "/" + path
	}
	if prefix {
		path = AdminCreateUrl(path)
	} else {
		path = "/" + strings.Trim(path, "/")
	}
	return path
}
