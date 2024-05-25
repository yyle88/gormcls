package gormcls

// Use 这个函数起到隔离【作用域】的功能，避免临时变量在函数中的作用域过大，避免重名变量混淆
func Use[MOD columnsInterface[CLS], CLS any](a MOD) (MOD, CLS) {
	return a, a.Columns()
}

// 配合 https://github.com/yyle88/gormcngen 使用，因为里面的默认函数就是 Columns
type columnsInterface[CLS any] interface {
	Columns() CLS
}
