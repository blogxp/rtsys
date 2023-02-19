// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package lib

import (
	"html/template"
	"rtsys/utils/tool"
	"strings"

	"github.com/gin-gonic/gin"
)

// ParseHtml 解析约定的html
func ParseHtml(files ...string) *template.Template {
	//上传需要的页面
	files = append(files, UploadHtml()...)

	for i := 0; i < len(files); i++ {
		files[i] = "template/admin/" + strings.Trim(files[i], "/") + ".html"
	}
	//下面为每个全局加载的页面
	files = append(files, "template/admin/global/layout/include.html")
	files = append(files, "template/admin/global/layout/default.html")
	files = append(files, "template/admin/global/layout/form.html")
	files = append(files, "template/admin/global/layout/full.html")

	return template.Must(template.ParseFiles(files...))
}

// ErrorHtml 错误页面渲染
func ErrorHtml(ctx *gin.Context, err error) {
	temp := ParseHtml("global/common/error")
	_ = temp.ExecuteTemplate(ctx.Writer, "common/error", gin.H{"error": err.Error()})
}

// UploadHtml 上传插件的模板，在render最后参数调用
func UploadHtml() []string {
	return []string{
		"global/upload/file",
		"global/upload/img",
		"global/upload/video",
	}
}

// TempFunc 定义的模中的方法
func TempFunc() map[string]any {
	return map[string]any{
		"inArray": tool.InArray,
		"inMap":   tool.InMap,
		"split": func(s string, sep string) []string {
			return strings.Split(s, sep)
		},
		"join": func(elems []string, sep string) string {
			return strings.Join(elems, sep)
		},
		"fileDomain": func(path string) string {
			return tool.UrlDomain(path, true)
		},
		"makeArray": func(args ...any) []any {
			arr := make([]any, len(args))
			for i := 0; i < len(args); i++ {
				arr[i] = args[i]
			}
			return arr
		},
		//创建集合 参数为字符串 key:value  ...
		"makeMap": func(args ...string) map[string]any {
			maps := map[string]any{}
			for i := 0; i < len(args); i++ {
				item := strings.SplitN(args[i], ":", 2)
				if len(item) == 2 && item[0] != "" {
					maps[item[0]] = item[1]
				}
			}
			return maps
		},
		//从一个map中获取某个键的值，不存在设为默认值
		"mapValue": func(key string, maps map[string]any, defaults string) any {
			if v, exit := maps[key]; exit {
				return v
			}
			return defaults
		},
		//往map[string]any 中设置一个值
		"dataSet": func(key string, val any, data map[string]any) map[string]any {
			data[key] = val
			return data
		},
	}
}
