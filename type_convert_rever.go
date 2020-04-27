package cachefunc

//InterfaceToInt  数据转换
func InterfaceToInt(ids []interface{}) (ret []int) {
	ret = make([]int, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		tmpID, ok := id.(int)
		if !ok {
			continue
		}
		ret = append(ret, tmpID)
	}
	return
}

//InterfaceToStrings  数据转换
func InterfaceToStrings(ids []interface{}) (ret []string) {
	ret = make([]string, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		tmpID, ok := id.(string)
		if !ok {
			continue
		}
		ret = append(ret, tmpID)

	}
	return
}

//InterfaceToInt64  数据转换
func InterfaceToInt64(ids []interface{}) (ret []int64) {
	ret = make([]int64, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		tmpID, ok := id.(int64)
		if !ok {
			continue
		}
		ret = append(ret, tmpID)

	}
	return
}

//InterfaceToUint32  数据转换
func InterfaceToUint32(ids []interface{}) (ret []uint32) {
	ret = make([]uint32, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		tmpID, ok := id.(uint32)
		if !ok {
			continue
		}
		ret = append(ret, tmpID)

	}
	return
}

//InterfaceToUint64  数据转换
func InterfaceToUint64(ids []interface{}) (ret []uint64) {
	ret = make([]uint64, 0, len(ids))
	if len(ids) == 0 {
		return
	}
	for _, id := range ids {
		tmpID, ok := id.(uint64)
		if !ok {
			continue
		}
		ret = append(ret, tmpID)
	}
	return
}
