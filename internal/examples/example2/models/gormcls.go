package models

import (
	"sync"

	"github.com/yyle88/gormcls"
	"github.com/yyle88/mutexmap"
)

// 这个全局变量的最佳实践，就是和 models 数据放在一起，而里面的数量就是填写数据表的（随意的预估）个数，当然也可以填0
var cache1 = mutexmap.NewMap[string, interface{}](2)

// Usc 就是用的工具包中的 Usc 逻辑，由于go目前不允许类的成员函数名为泛型，这里只能是定义个普通的函数，函数内部用全局变量缓存信息
func Usc[MOD gormcls.ColumnsTableNameIFace[CLS], CLS any](a MOD) (MOD, CLS) {
	return gormcls.Usc(a, cache1)
}

// 这是第二套方案的缓存变量
var cache2 = &sync.Map{}

// Uss 就是用的工具包中的 Uss 逻辑
func Uss[MOD gormcls.ColumnsTableNameIFace[CLS], CLS any](a MOD) (MOD, CLS) {
	return gormcls.Uss(a, cache2)
}
