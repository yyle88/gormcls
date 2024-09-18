package gormclsm

import (
	"sync"

	"github.com/yyle88/erero"
	"github.com/yyle88/mutexmap"
)

// ColumnsTableNameIFace 配合 gorm 的基础使用，因为里面用的就是 TableName() 函数作为自定义表名的函数
type ColumnsTableNameIFace[CLS any] interface {
	TableName() string
	Columns() CLS
}

// UmcV2 在实现 Use 隔离【作用域】的同时，增加缓存效果，即避免总是调用 Columns 函数生成 gormcnm 的对象，而是把这个对象缓存起来，能提高些性能(当然不介意这部分性能的也可以不用它)
func UmcV2[MOD ColumnsTableNameIFace[CLS], CLS any](a MOD, cache *mutexmap.Map[string, interface{}]) (MOD, CLS) {
	vax, _ := cache.GetOrzSet(a.TableName(), func() interface{} {
		return a.Columns()
	})
	cls, ok := vax.(CLS)
	if !ok {
		panic(erero.Errorf("wrong TABLE_NAME=%s", a.TableName()))
	}
	return a, cls
}

// UmcV3 在实现 Use 隔离【作用域】的同时，增加缓存效果，即避免总是调用 Columns 函数生成 gormcnm 的对象，而是把这个对象缓存起来，能提高些性能(当然不介意这部分性能的也可以不用它)
func UmcV3[MOD ColumnsTableNameIFace[CLS], CLS any](a MOD, cache *sync.Map) (MOD, CLS) {
	value, ok := cache.Load(a.TableName())
	if !ok {
		value = a.Columns()
		cache.Store(a.TableName(), value) //这里在并发时有可能多次存储，但是这并不影响任何功能，因此忽略这个情况
	}
	cls, ok := value.(CLS)
	if !ok {
		panic(erero.Errorf("wrong TABLE_NAME=%s", a.TableName()))
	}
	return a, cls
}
