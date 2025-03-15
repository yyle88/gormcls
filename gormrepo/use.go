package gormrepo

import (
	"github.com/yyle88/gormcls"
	"github.com/yyle88/gormcls/internal/classtype"
	"gorm.io/gorm"
)

// Use returns the database(db) model (mod) and the associated columns (cls).
// Use 返回 数据库(db) 模型（mod）和关联的列（cls）。
func Use[MOD classtype.ModelType[CLS], CLS any](db *gorm.DB, one MOD) (*gorm.DB, MOD, CLS) {
	one, cls := gormcls.Use(one)
	return db, one, cls
}

// Umc returns the database(db) model (mod) and the associated columns (cls), functioning identically to the Use function.
// Umc 返回 数据库(db) 模型（mod）和关联的列（cls），功能与 Use 函数相同。
func Umc[MOD classtype.ModelType[CLS], CLS any](db *gorm.DB, one MOD) (*gorm.DB, MOD, CLS) {
	one, cls := gormcls.Umc(one)
	return db, one, cls
}
