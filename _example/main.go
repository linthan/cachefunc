package main

import (
	"context"
	"fmt"
	"time"

	"github.com/linthan/cachefunc"
)

//GetDataFromServerReq 请求
type GetDataFromServerReq struct {
	ID string `json:"id"`
}

//GetID 获取请求的ID用来存储ID
func (req *GetDataFromServerReq) GetID() string {
	return req.ID
}

//GetDataFromServer 获取数据
func GetDataFromServer(ctx context.Context, req cachefunc.WrapperReq) (ret interface{}, err error) {
	innerReq, ok := req.(*GetDataFromServerReq)
	if !ok {
		return
	}
	ret = innerReq.ID
	return
}

//BatchGetDataFromServer 获取数据
func BatchGetDataFromServer(ctx context.Context, keys []interface{}) (ret map[interface{}]interface{}, err error) {
	ret = make(map[interface{}]interface{}, len(keys))
	for _, key := range keys {
		ret[key] = key
	}
	return
}

func main() {

	cblock, _ := cachefunc.New("test", cachefunc.ShardCount(10), cachefunc.Expire(time.Second*10))
	// cacheFunc := cblock.Wrapper(GetDataFromServer)
	batchCacheFunc := cblock.BatchWrapper(BatchGetDataFromServer)
	for {
		time.Sleep(time.Second)
		items := []int{1, 2, 2, 3, 4, 5, 6}
		data, err := batchCacheFunc(context.Background(), cachefunc.IntToInterface(items))
		fmt.Println("-------", data, err)
	}
}
