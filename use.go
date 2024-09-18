package gormcls

// Use 这个函数起到隔离【作用域】的功能，避免临时变量在函数中的作用域过大，避免重名变量混淆
func Use[MOD ColumnsIFace[CLS], CLS any](one MOD) (MOD, CLS) {
	return one, one.Columns()
}

// Cls 当你完全不需要mod而只需要cls的时候 就很有用
func Cls[MOD ColumnsIFace[CLS], CLS any](one MOD) CLS {
	return one.Columns()
}

// ColumnsIFace 配合 https://github.com/yyle88/gormcngen 使用，因为里面的默认函数就是 Columns
type ColumnsIFace[CLS any] interface {
	Columns() CLS
}

// Ucs 就是返回模型的数组 以便于使用 Find 查询. s means slice
func Ucs[MOD ColumnsIFace[CLS], CLS any](one MOD) ([]MOD, CLS) {
	return []MOD{}, one.Columns()
}

// Usc 当你需要 Find 而且需要 Model 时有用
func Usc[MOD ColumnsIFace[CLS], CLS any](one MOD) (MOD, []MOD, CLS) {
	return one, []MOD{}, one.Columns()
}

// One 这个函数也是有神奇的功能，比如gorm的Create或者Save函数只接受指针类型，这个函数能在编译阶段就判定传的是不是指针类型，以便于后面调用Create或者Save函数
func One[MOD ColumnsIFace[CLS], CLS any](one MOD) MOD {
	return one //把数据原封不动的返回来，因为按照 gormcngen 的默认规则，只给类型生成 func (*X) Columns() XColumns {} 这样的成员函数
}

// Ums 得到对象数组. 当你需要 Find 而不是 First 就很有用，s means slice
func Ums[MOD ColumnsIFace[CLS], CLS any](MOD) []MOD {
	return []MOD{}
}

// Uss 得到对象数组
func Uss[MOD ColumnsIFace[CLS], CLS any]() []MOD {
	return []MOD{}
}

// Usn 得到对象数组
func Usn[MOD ColumnsIFace[CLS], CLS any](cap int) []MOD {
	return make([]MOD, 0, cap)
}
