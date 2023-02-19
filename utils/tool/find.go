// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package tool

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"reflect"
	"strings"
)

// InArray 判断某个值是否在数组中
func InArray(value any, arr interface{}) bool {
	t := reflect.TypeOf(arr)
	if t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		v := reflect.ValueOf(arr)
		for i := 0; i < v.Len(); i++ {
			if v.Index(i).Interface() == value {
				return true
			}
		}
	}
	return false
}

// InMap 判断某个键是否存在map种
func InMap(value any,target any) bool {
	t := reflect.TypeOf(target)
	fmt.Println(t.Kind())
	if t.Kind() == reflect.Map {
		v := reflect.ValueOf(target)
		keys := v.MapKeys()
		for i := 0; i < len(keys); i++ {
			if keys[i].Interface() == value {
				return true
			}
		}
	}
	return false
}

func IsAjax(ctx *gin.Context) bool {
	return ctx.GetHeader("X-Requested-With") == "XMLHttpRequest"
}

func IsRender(ctx *gin.Context) bool  {
	return !IsAjax(ctx) && strings.ToLower(ctx.Request.Method) == "get"
}
