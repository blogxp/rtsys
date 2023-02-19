package test_func

import (
	"rtsys/utils/core"
	"testing"
)

// 测试雪花算法 性能
func BenchmarkSnowFlacke(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		//5秒内20次
		_ = core.G_GENID.New().String()
	}
}
