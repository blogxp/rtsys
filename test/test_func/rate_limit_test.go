// +----------------------------------------------------------------------
// | B5GoCMF V1.0 [快捷通用基础管理开发平台]
// +----------------------------------------------------------------------
// | Author: 冰舞 <357145480@qq.com>
// +----------------------------------------------------------------------

////////   简单实现的 固定窗口和浮动窗口限流 方法  令牌桶和漏桶需要定时不写了    ///////////

package test_func

import (
	"fmt"
	"rtsys/utils/tool"
	"testing"
	"time"
)

// 测试固定窗口方法
func BenchmarkNewFixRateLimit(b *testing.B) {
	b.ResetTimer()
	start := time.Now()
	for i := 0; i < b.N; i++ {
		//10秒内20次
		tool.NewFixRateLimit(map[string]int{"limit": 20, "exp": 10}).CheckLimit("bench_fix_rate_limit")
	}
	b.Log(b.N, time.Since(start))
}

// 测试浮动窗口方法
func TestNewSlideRateLimit(t *testing.T) {
	s := tool.NewSlideRateLimit(map[string]int{"limit": 10, "exp": 1, "len": 10}).CheckLimit("bench_slide_rate_limit")
	fmt.Println(s)
}
func BenchmarkNewSlideRateLimit(b *testing.B) {
	j := 1
	b.ResetTimer()
	start := time.Now()
	for i := 0; i < b.N; i++ {
		//1秒内20次, 5秒 100次
		s := tool.NewSlideRateLimit(map[string]int{"limit": 10, "exp": 1, "len": 10}).CheckLimit("bench_slide_rate_limit")
		if s {
			j++
		}
	}
	b.Log(b.N, time.Since(start), j)
}
