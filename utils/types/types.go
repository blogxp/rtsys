// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package types

import (
	"html/template"
)

// KeyVal 定义的通用的 key=>val的结构体，主要用于有序的配置输出
type KeyVal struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type SimpleId struct {
	Id  string `json:"id"`
	Ids string `json:"ids"`
}

type AutoKey struct {
	Where map[string]string `json:"where"`
}

type HtmlShow struct {
	Html template.HTML
}