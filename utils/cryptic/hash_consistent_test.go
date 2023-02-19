// group cache 中的一致性hash算法

package cryptic

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T)  {
	hash := NewConsistentHash(3, nil)
	hash.Add("192.168.1.1","192.168.1.2","192.168.1.3")


	fmt.Println(hash.Get("config_cache_list"))
	fmt.Println(hash.Get("struct_cache_list"))
	fmt.Println(hash)
}

func TestConsistency(t *testing.T) {
	hash1 := NewConsistentHash(1, nil)
	hash2 := NewConsistentHash(1, nil)

	hash1.Add("Bill", "Bob", "Bonny")
	hash2.Add("Bob", "Bonny", "Bill")

	if hash1.Get("Ben") != hash2.Get("Ben") {
		t.Errorf("Fetching 'Ben' from both hashes should be the same")
	}

	hash2.Add("Becky", "Ben", "Bobby")

	if hash1.Get("Ben") != hash2.Get("Ben") ||
		hash1.Get("Bob") != hash2.Get("Bob") ||
		hash1.Get("Bonny") != hash2.Get("Bonny") {
		t.Errorf("Direct matches should always return the same entry")
	}

}

func BenchmarkGet8(b *testing.B)   { benchmarkGet(b, 8) }
func BenchmarkGet32(b *testing.B)  { benchmarkGet(b, 32) }
func BenchmarkGet128(b *testing.B) { benchmarkGet(b, 128) }
func BenchmarkGet512(b *testing.B) { benchmarkGet(b, 512) }

func benchmarkGet(b *testing.B, shards int) {

	hash := NewConsistentHash(50, nil)

	var buckets []string
	for i := 0; i < shards; i++ {
		buckets = append(buckets, fmt.Sprintf("shard-%d", i))
	}

	hash.Add(buckets...)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		hash.Get(buckets[i&(shards-1)])
	}
}