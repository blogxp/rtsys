// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package lib

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义特定的code值
const (
	RSuccess  int = 0
	RFail     int = 300
	RNoLogin int = 305
)

// JsonSuccess 成功返回json
func JsonSuccess(ctx *gin.Context, args ...any) {
	ctx.JSON(http.StatusOK, JsonArgsParse(true, args...))
}

// JsonFail 失败返回json
func JsonFail(ctx *gin.Context, args ...any) {
	ctx.JSON(http.StatusOK, JsonArgsParse(false, args...))
}

// JsonArgsParse 解析生成JSON返回的格式
// 可变参数args 的格式顺序约定为
//
//	0 msg string
//	1 data map[string]any
//	2 code int
//	3 url string
func JsonArgsParse(success bool, args ...any) map[string]any {
	var msg, data, code,url any
	if len(args) > 0 && args[0] != nil {
		msg = args[0]
	} else {
		if success {
			msg = "操作成功"
		} else {
			msg = "操作失败"
		}
	}
	if len(args) > 1 && args[1] != nil {
		data = args[1]
	} else {
		data = make(map[string]any)
	}
	if len(args) > 2 && args[2] != nil {
		code = args[2]
	} else {
		if success {
			code = RSuccess
		} else {
			code = RFail
		}
	}
	if len(args)>3 && args[3] !=nil {
		url = args[3]
	}else{
		url = ""
	}
	return map[string]any{"code": code, "msg": msg, "data": data, "success": success,"url":url}
}
