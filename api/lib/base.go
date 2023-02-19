// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package lib

import (
	"github.com/gin-gonic/gin"
)

type BaseApi struct {
	Id           string                                     //控制器名称
	BeforeAction func(ctx *gin.Context, action string) bool //该BaseApi所有路由进行请求前的逻辑判断处理
}

// Dispatch 路由分配器
// 生成具体访问ulr
// 调用控制器构造函数 Construct
func (ca *BaseApi) Dispatch(action string, handler func(*gin.Context)) (string, gin.HandlerFunc) {
	path := "/" + ca.Id + "/" + action
	return path, func(ctx *gin.Context) {
		n := ca.Construct(ctx, action)
		if n {
			handler(ctx)
		}
	}
}

// Construct 构造函数
// 赋值属性及调用前置操作BeforeAction
func (ca *BaseApi) Construct(ctx *gin.Context, action string) bool {
	if ca.BeforeAction != nil {
		return ca.BeforeAction(ctx, action)
	}
	return true
}

// Success / Error 返回json数据  可变参数args 的格式顺序约定为
//
//	msg string
//	data map[string]any
//	code int
func (ca *BaseApi) Success(ctx *gin.Context, args ...any) {
	ApiSuccess(ctx, args...)
}
func (ca *BaseApi) Error(ctx *gin.Context, args ...any) {
	ApiFail(ctx, args...)
}
