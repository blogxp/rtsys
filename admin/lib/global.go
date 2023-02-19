// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package lib

import (
	"rtsys/utils/core"
	"strings"

	"github.com/gin-gonic/gin"
)

// AdminGlobalSetting 全局加载设置
func AdminGlobalSetting(engine *gin.Engine) {

}

// AdminCreateUrl 生成后端的包含前缀的url
func AdminCreateUrl(url string) string {
	if url == "" {
		return url
	}
	if strings.Index(url, "http") == 0 {
		return url
	}
	root := strings.Trim(core.G_CONFIG.Route.Admin, "/")
	if root != "" {
		root = "/" + root
	}
	return root + "/" + strings.Trim(url, "/")
}

// AdminPathParse 解析后端url的模块名，控制器名以及方法名
func AdminPathParse(path string) map[string]string {
	prefix := core.G_CONFIG.Route.Admin
	prefixLen := len(prefix)
	if prefixLen > 0 {
		path = strings.Trim(path[prefixLen:], "/")
	}
	pathArr := strings.Split(path, "/")
	if len(pathArr) != 2 && len(pathArr) != 3 {
		return nil
	}
	var group, controller, action string
	if len(pathArr) == 2 {
		group = ""
		controller = strings.ToLower(pathArr[0])
		action = strings.ToLower(pathArr[1])
	} else {
		group = strings.ToLower(pathArr[0])
		controller = strings.ToLower(pathArr[1])
		action = strings.ToLower(pathArr[2])
	}
	return map[string]string{"group": group, "controller": controller, "action": action}
}
