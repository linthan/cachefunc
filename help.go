package cachefunc

import (
	"fmt"

	lru "github.com/hnlq715/golang-lru"
)

func fnv32(key string) uint32 {
	hash := uint32(2166136261)
	const prime32 = uint32(16777619)
	for i := 0; i < len(key); i++ {
		hash *= prime32
		hash ^= uint32(key[i])
	}
	return hash
}

// GetShard returns shard under given key
func (ins *Block) GetShard(key string) *lru.ARCCache {
	return ins.GetShard(key)
}

//ConcurrentBlock 并发map
type ConcurrentBlock []*lru.ARCCache

// GetShard returns shard under given key
func (m ConcurrentBlock) GetShard(key string) *lru.ARCCache {
	return m[uint(fnv32(key))%uint(len(m))]
}

// Get retrieves an element from map under given key.
func (m ConcurrentBlock) Get(key interface{}) (interface{}, bool) {
	// Get shard
	shard := m.GetShard(fmt.Sprintf("%v", key))
	// Get item from shard.
	val, ok := shard.Get(key)
	return val, ok
}

//Add 添加
func (m ConcurrentBlock) Add(key interface{}, value interface{}) {
	shard := m.GetShard(fmt.Sprintf("%v", key))
	shard.Add(key, value)
	return
}
