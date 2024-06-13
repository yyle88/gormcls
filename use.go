package gormcls

// Use 这个函数起到隔离【作用域】的功能，避免临时变量在函数中的作用域过大，避免重名变量混淆
func Use[MOD columnsInterface[CLS], CLS any](a MOD) (MOD, CLS) {
	return a, a.Columns()
}

// 配合 https://github.com/yyle88/gormcngen 使用，因为里面的默认函数就是 Columns
type columnsInterface[CLS any] interface {
	Columns() CLS
}

// One 这个函数也是有神奇的功能，比如gorm的Create或者Save函数只接受指针类型，这个函数能在编译阶段就判定传的是不是指针类型，以便于后面调用Create或者Save函数
func One[MOD columnsInterface[CLS], CLS any](a MOD) MOD {
	return a //把数据原封不动的返回来，因为按照 gormcngen 的默认规则，只给类型生成 func (*X) Columns() XColumns {} 这样的成员函数
}
