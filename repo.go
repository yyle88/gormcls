package gormrepo

import (
	"gorm.io/gorm"
)

type Repo[MOD any, CLS any] struct {
	mod *MOD
	cls CLS
}

func NewRepo[MOD any, CLS any](_ *MOD, cls CLS) *Repo[MOD, CLS] {
	return &Repo[MOD, CLS]{
		mod: nil, // 这里就是设置个空值避免共享对象
		cls: cls,
	}
}

func (repo *Repo[MOD, CLS]) Gorm(db *gorm.DB) *GormRepo[MOD, CLS] {
	return NewGormRepo(db, repo.mod, repo.cls)
}

func (repo *Repo[MOD, CLS]) DBMo(db *gorm.DB) *GormRepo[MOD, CLS] {
	return NewGormRepo(db.Model((*MOD)(nil)), repo.mod, repo.cls)
}
