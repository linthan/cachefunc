package cachefunc

import "time"

//Option 设置Instance
type Option func(*Block)

//Size 设置cacheSize
func Size(cacheSize int) Option {
	return func(ins *Block) {
		ins.cacheSize = cacheSize
	}
}

//Expire 设置cacheExpire
func Expire(cacheExpire time.Duration) Option {
	return func(ins *Block) {
		ins.cacheExpire = cacheExpire
	}
}

//ShardCount 设置shardCount
func ShardCount(shardCount int) Option {
	return func(ins *Block) {
		ins.shardCount = shardCount
	}
}
