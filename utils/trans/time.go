// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

package trans

import (
	"rtsys/utils/core"
	"strings"
	"time"
)

// TimeFormat 时间格式化
func TimeFormat(times time.Time, format string) string {
	if format == "" {
		format = "Y-m-d H:i:s"
	}
	maps := map[string]string{"Y": "2006", "m": "01", "d": "02", "H": "15", "i": "04", "s": "05"}
	for i, v := range maps {
		format = strings.Replace(format, i, v, -1)
	}
	return times.Format(format)
}

// TimeStringFormat 对时间字符串重新格式化
// 主要用于从数据库中取出 datetime后的展示效果
func TimeStringFormat(times string, format string) string {
	parse, err := time.Parse(times, core.G_TIME)
	if err != nil {
		return times
	}
	return parse.Format(format)
}
