package tool

import "strings"

//StrToHump 字符串转大驼峰 对 - _进行
func StrToHump(s string) string {
	if s == "" {
		return ""
	}
	strArr := strings.Split(s,"_")
	list := make([]string,0)
	for _,v := range strArr{
		if v == "" {
			continue
		}
		vArr := strings.Split(v,"-")
		for _,v1 := range vArr{
			if v1=="" {
				continue
			}
			list = append(list,strings.ToUpper(v1[:1]) +v1[1:] )
		}
	}
	return strings.Join(list,"")
}
