// group cache 中的一致性hash算法

package cryptic

import (
	"hash/crc32"
	"sort"
	"strconv"
)


type ConsistentHash func(data []byte) uint32

type ConsistentHashMap struct {
	hash     ConsistentHash //对传入的值进行hash算法的方法
	replicas int  //虚拟节点数
	keys     []int // 所有节点的hash ,为了快速找到所属hash,长度为节点的数量 * 虚拟节点数
	hashMap  map[int]string //int键为hash值  string值为节点标识  长度和keys相等
}

func NewConsistentHash(replicas int, fn ConsistentHash) *ConsistentHashMap {
	m := &ConsistentHashMap{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

// Add adds some keys to the hash. 添加一个或多个节点标识
func (m *ConsistentHashMap) Add(nodes ...string) {
	for _, node := range nodes {
		for i := 0; i < m.replicas; i++ {
			hash := int(m.hash([]byte(strconv.Itoa(i) + node)))
			m.keys = append(m.keys, hash)
			m.hashMap[hash] = node
		}
	}
	sort.Ints(m.keys) //递增排序，生成一个有序的hash环（索引为环的长度 = 0）
}

// Get gets the closest item in the hash to the provided key.  传入一个字符串 返回被分配的节点标识
func (m *ConsistentHashMap) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}

	hash := int(m.hash([]byte(key)))

	// Binary search for appropriate replica.
	// Search(n int, f func(int) bool) int采用二分法搜索找到[0, n)区间内最小的满足f(i)==true的值i，如果没有该值，函数会返回n
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })

	// Means we have cycled back to the first replica.
	// 没有找到，则默认存为第一个.
	if idx == len(m.keys) {
		idx = 0
	}

	return m.hashMap[m.keys[idx]]
}

// IsEmpty returns true if there are no items available.
func (m *ConsistentHashMap) IsEmpty() bool {
	return len(m.keys) == 0
}
