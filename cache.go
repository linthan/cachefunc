package cachefunc

import (
	"context"
	"time"

	lru "github.com/hnlq715/golang-lru"
)

var registryBlocks = map[string]int{}

//Block 实现缓存对象
type Block struct {
	name        string
	cacheSize   int           //设置缓存的大小
	cacheExpire time.Duration //设置缓存的过期时间
	caches      ConcurrentBlock
	shardCount  int
}

//New 初始化
func New(name string, opts ...Option) (ret *Block, err error) {
	//初始化
	if _, ok := registryBlocks[name]; !ok {
		registryBlocks[name] = 0
	} else {
		//确保名称唯一,容易上报指标
		panic("already registry the block cache")
	}
	var instance = &Block{
		name:        name,
		cacheSize:   4000,
		cacheExpire: 5 * time.Minute,
		shardCount:  32,
	}

	for _, opt := range opts {
		opt(instance)
	}
	//初始化cache
	caches := make([]*lru.ARCCache, instance.shardCount)
	for i := 0; i < instance.shardCount; i++ {
		caches[i], err = lru.NewARCWithExpire(instance.cacheSize, instance.cacheExpire)
		if err != nil {
			//初始化必须成功
			panic(err)
		}
	}
	instance.caches = caches
	ret = instance
	return
}

//WrapperReq 请求
type WrapperReq interface {
	GetID() string
}

//BatchWrapperReq 批量获取
type BatchWrapperReq interface {
	GetIDs() []string
}

//Wrapper 包装
func (ins *Block) Wrapper(in func(context.Context, WrapperReq) (interface{}, error)) func(context.Context, WrapperReq) (interface{}, bool, error) {
	if len(ins.caches) == 0 {
		panic("cache cannot be nil")
	}
	return func(ctx context.Context, req WrapperReq) (ret interface{}, ok bool, err error) {
		ret, ok = ins.caches.Get(req.GetID())
		if !ok {
			ret, err = in(ctx, req)
			if err != nil {
				return
			}
			ins.caches.Add(req.GetID(), ret)
			return
		}
		return
	}
}
