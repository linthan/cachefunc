package cachefunc

import "context"

//BatchWrapper 批量获取包装
func (ins *Block) BatchWrapper(in BatchCacheFuncWithInterface) BatchCacheFuncWithInterface {
	if len(ins.caches) == 0 {
		panic("cache cannot be nil")
	}
	return func(ctx context.Context, keys []interface{}) (ret map[interface{}]interface{}, err error) {
		ret = make(map[interface{}]interface{})
		if len(keys) == 0 {
			return
		}
		notInKeys := make([]interface{}, 0, len(keys))
		for _, key := range keys {
			item, ok := ins.caches.Get(key)
			if !ok {
				notInKeys = append(notInKeys, key)
				continue
			}
			//判空,防止使用方空指针
			if item != nil {
				ret[key] = item
			}

		}
		if len(notInKeys) == 0 {
			return
		}

		//获取为命中数据
		items, err := in(ctx, notInKeys)
		if err != nil {
			return
		}
		for _, key := range notInKeys {
			item, ok := items[key]
			if !ok {
				continue
			}
			if item != nil {
				ret[key] = item
			}
			//这个的item 可为nil,也可以不为nil,主要看是否需要防止缓存穿透
			ins.caches.Add(key, item)
		}
		return
	}
}
