package cachefunc

import "context"

//BatchCacheFuncWithInterface key为interface{}
type BatchCacheFuncWithInterface func(context.Context, []interface{}) (map[interface{}]interface{}, error)
