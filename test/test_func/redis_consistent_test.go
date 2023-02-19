package test_func

import (
	"fmt"
	"rtsys/utils/core"
	"testing"
	"time"
)

func TestRedisHash(t *testing.T) {
	key := "config_sys_name"
	val := "b5gocmf"
	core.G_Redis.HashConn(key).Set(key, "b5gocmf", 0)

	result, _ := core.G_Redis.HashConn(key).Get(key).Result()
	fmt.Println(result)
	if result != val {
		t.Error("值不相等")
	}
}

// 测试通过一致性hash获取
func BenchmarkRedisHash(b *testing.B) {
	b.ResetTimer()
	start := time.Now()
	for i := 0; i < b.N; i++ {
		core.G_Redis.HashConn("config_sys_name").Get("config_sys_name").Result()
	}
	b.Log(b.N, time.Since(start))
}
