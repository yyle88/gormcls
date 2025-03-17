package gormrepo

import (
	"gorm.io/gorm"
)

type AbstractRepo[MOD any, CLS any] struct {
	mod *MOD
	cls CLS
}

func NewAbstractRepo[MOD any, CLS any](_ *MOD, cls CLS) *AbstractRepo[MOD, CLS] {
	return &AbstractRepo[MOD, CLS]{
		mod: nil, // 这里就是设置个空值避免共享对象
		cls: cls,
	}
}

func (repo *AbstractRepo[MOD, CLS]) NewRepo(db *gorm.DB) *Repo[MOD, CLS] {
	return NewRepo(db, repo.mod, repo.cls)
}
