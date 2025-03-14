package gormrepo

import (
	"github.com/yyle88/gormcls/internal/classtype"
	"github.com/yyle88/must"
	"gorm.io/gorm"
)

// Dmc returns the database(db) model (mod) and the associated columns (cls).
// Dmc 返回 数据库(db) 模型（mod）和关联的列（cls）。
func Dmc[MOD classtype.ModelType[CLS], CLS any](db *gorm.DB, one MOD) (*gorm.DB, MOD, CLS) {
	return db, one, one.Columns()
}

type Repo[MOD any, CLS any] struct {
	db  *gorm.DB
	cls CLS
}

func NewRepo[MOD any, CLS any](db *gorm.DB, one *MOD, cls CLS) *Repo[MOD, CLS] {
	must.Nice(one) //只存类型信息，而不存实体避免共享内存
	return &Repo[MOD, CLS]{
		db:  db,
		cls: cls,
	}
}

func (repo *Repo[MOD, CLS]) OK() bool {
	return true
}

func (repo *Repo[MOD, CLS]) First(where func(db *gorm.DB, cls CLS) *gorm.DB) (*MOD, error) {
	var result = new(MOD)
	db := where(repo.db, repo.cls)
	if err := db.First(result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *Repo[MOD, CLS]) Exist(where func(db *gorm.DB, cls CLS) *gorm.DB) (bool, error) {
	var exists bool
	db := where(repo.db, repo.cls)
	if err := db.Model(new(MOD)).Select("1").Limit(1).Find(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}
