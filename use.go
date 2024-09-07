package gormcls

import (
	"sync"

	"github.com/yyle88/erero"
	"github.com/yyle88/mutexmap"
)

// Use 这个函数起到隔离【作用域】的功能，避免临时变量在函数中的作用域过大，避免重名变量混淆
func Use[MOD ColumnsIFace[CLS], CLS any](a MOD) (MOD, CLS) {
	return a, a.Columns()
}

// ColumnsIFace 配合 https://github.com/yyle88/gormcngen 使用，因为里面的默认函数就是 Columns
type ColumnsIFace[CLS any] interface {
	Columns() CLS
}

// Usa 当你需要 Find 而不是 First 就很有用
func Usa[MOD ColumnsIFace[CLS], CLS any](a MOD) ([]MOD, CLS) {
	return []MOD{}, a.Columns()
}

// Uas 当你需要 Find 而且需要 Model 时有用
func Uas[MOD ColumnsIFace[CLS], CLS any](a MOD) (MOD, []MOD, CLS) {
	return a, []MOD{}, a.Columns()
}

// One 这个函数也是有神奇的功能，比如gorm的Create或者Save函数只接受指针类型，这个函数能在编译阶段就判定传的是不是指针类型，以便于后面调用Create或者Save函数
func One[MOD ColumnsIFace[CLS], CLS any](a MOD) MOD {
	return a //把数据原封不动的返回来，因为按照 gormcngen 的默认规则，只给类型生成 func (*X) Columns() XColumns {} 这样的成员函数
}

// Usc 在实现 Use 隔离【作用域】的同时，增加缓存效果，即避免总是调用 Columns 函数生成 gormcnm 的对象，而是把这个对象缓存起来，能提高些性能(当然不介意这部分性能的也可以不用它)
func Usc[MOD ColumnsTableNameIFace[CLS], CLS any](a MOD, cache *mutexmap.Map[string, interface{}]) (MOD, CLS) {
	vax, _ := cache.GetOrzSet(a.TableName(), func() interface{} {
		return a.Columns()
	})
	cls, ok := vax.(CLS)
	if !ok {
		panic(erero.Errorf("wrong TABLE_NAME=%s", a.TableName()))
	}
	return a, cls
}

// ColumnsTableNameIFace 配合 gorm 的基础使用，因为里面用的就是 TableName() 函数作为自定义表名的函数
type ColumnsTableNameIFace[CLS any] interface {
	Columns() CLS
	TableName() string
}

// Uss 在实现 Use 隔离【作用域】的同时，增加缓存效果，即避免总是调用 Columns 函数生成 gormcnm 的对象，而是把这个对象缓存起来，能提高些性能(当然不介意这部分性能的也可以不用它)
func Uss[MOD ColumnsTableNameIFace[CLS], CLS any](a MOD, cache *sync.Map) (MOD, CLS) {
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
