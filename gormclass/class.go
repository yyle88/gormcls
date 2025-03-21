package gormclass

// GormClass is used for models that implement the Columns method to return associated columns (cls).
// GormClass 用于实现 Columns 方法以返回关联列（cls）的模型。
type GormClass[CLS any] interface {
	Columns() CLS
}

// ClassType 配合 gorm 的基础使用，使用 TableName() 作为自定义表名函数
// ClassType is used in combination with gorm for basic usage, where TableName() is used as a custom table name function.
type ClassType[CLS any] interface {
	TableName() string
	Columns() CLS
}
