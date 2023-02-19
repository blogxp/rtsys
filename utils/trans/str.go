// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package trans

import (
	. "rtsys/utils/types"
	"strings"
)

// StrToKeyValByLine 解析配置字符串成数组KeyVal
// 主要为了有序， 不需要顺序可使用StrToMapByLine解析成集合
// 例 a:1,b:2 =>(sep1=, sep2=:)  =>KeyVal[{Key:a, Value:1} {Key:b, Value:2}]
func StrToKeyValByLine(str string, sep1 string, sep2 string) []KeyVal {
	result := make([]KeyVal, 0)
	if sep1 == "" || sep2 == "" {
		return result
	}
	list := strings.Split(str, sep1)
	if len(list) < 1 {
		return result
	}
	for _, val := range list {
		val = strings.Trim(val, " ")
		if val == "" {
			continue
		}
		item := strings.SplitN(val, sep2, 2)
		if len(item) == 2 {
			result = append(result, KeyVal{Key: item[0], Value: item[1]})
		}
	}
	return result
}

// StrToMapByLine 解析配置字符串成数组集合
// 例 a:1,b:2 =>(sep1=, sep2=:)  =>{a:1,b:2}
func StrToMapByLine(str string, sep1 string, sep2 string) map[string]string {
	var result map[string]string
	if sep1 == "" || sep2 == "" {
		return result
	}
	list := strings.Split(str, sep1)
	if len(list) < 1 {
		return result
	}
	for _, val := range list {
		val = strings.Trim(val, " ")
		if val == "" {
			continue
		}
		item := strings.SplitN(val, sep2, 2)
		if len(item) == 2 {
			result[item[0]] = item[1]
		}
	}
	return result
}
