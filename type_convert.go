package cachefunc

//IntToInterface  数据转换
func IntToInterface(ids []int) (ret []interface{}) {
	filter := make(map[interface{}]int, len(ids))
	ret = make([]interface{}, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)

	}
	return
}

//StringsToInterface  数据转换
func StringsToInterface(ids []string) (ret []interface{}) {
	filter := make(map[interface{}]int, len(ids))
	ret = make([]interface{}, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)

	}
	return
}

//StringsToInt64  数据转换
func StringsToInt64(ids []int64) (ret []interface{}) {
	filter := make(map[interface{}]int, len(ids))
	ret = make([]interface{}, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)
	}
	return
}

//StringsToUint32  数据转换
func StringsToUint32(ids []uint32) (ret []interface{}) {
	filter := make(map[interface{}]int, len(ids))
	ret = make([]interface{}, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)
	}
	return
}

//StringsToUint64  数据转换
func StringsToUint64(ids []uint64) (ret []interface{}) {
	filter := make(map[interface{}]int, len(ids))
	ret = make([]interface{}, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		if _, ok := filter[id]; ok {
			continue
		} else {
			filter[id] = 0
		}
		ret = append(ret, id)
	}
	return
}
