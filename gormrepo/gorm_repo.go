package gormrepo

import (
	"github.com/yyle88/must"
	"gorm.io/gorm"
)

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

func (repo *Repo[MOD, CLS]) Find(where func(db *gorm.DB, cls CLS) *gorm.DB) ([]*MOD, error) {
	var results []*MOD
	db := where(repo.db, repo.cls)
	if err := db.Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *Repo[MOD, CLS]) FindN(where func(db *gorm.DB, cls CLS) *gorm.DB, size int) ([]*MOD, error) {
	var results = make([]*MOD, 0, size)
	db := where(repo.db, repo.cls)
	if err := db.Limit(size).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *Repo[MOD, CLS]) Update(where func(db *gorm.DB, cls CLS) *gorm.DB, valueFunc func(cls CLS) (string, interface{})) error {
	db := where(repo.db, repo.cls)
	column, value := valueFunc(repo.cls)
	if err := db.Model(new(MOD)).Update(column, value).Error; err != nil {
		return err
	}
	return nil
}

func (repo *Repo[MOD, CLS]) Updates(where func(db *gorm.DB, cls CLS) *gorm.DB, valuesFunc func(cls CLS) map[string]interface{}) error {
	db := where(repo.db, repo.cls)
	mp := valuesFunc(repo.cls)
	if err := db.Model(new(MOD)).Updates(mp).Error; err != nil {
		return err
	}
	return nil
}
