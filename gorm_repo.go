package gormrepo

import (
	"gorm.io/gorm"
)

type GormRepo[MOD any, CLS any] struct {
	db  *gorm.DB
	mod *MOD
	cls CLS
}

func NewGormRepo[MOD any, CLS any](db *gorm.DB, _ *MOD, cls CLS) *GormRepo[MOD, CLS] {
	return &GormRepo[MOD, CLS]{
		db:  db,
		mod: nil, // 这里就是设置个空值避免共享对象
		cls: cls,
	}
}

func (repo *GormRepo[MOD, CLS]) OK() bool {
	return true
}

func (repo *GormRepo[MOD, CLS]) First(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *MOD) *gorm.DB {
	return where(repo.db, repo.cls).First(dest)
}

func (repo *GormRepo[MOD, CLS]) FirstX(where func(db *gorm.DB, cls CLS) *gorm.DB) (*MOD, error) {
	var result = new(MOD)
	if err := repo.First(where, result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

func (repo *GormRepo[MOD, CLS]) FirstE(where func(db *gorm.DB, cls CLS) *gorm.DB) (*MOD, *ErrorOrNotExist) {
	var result = new(MOD)
	if err := repo.First(where, result).Error; err != nil {
		return nil, NewErrorOrNotExist(err)
	}
	return result, nil
}

func (repo *GormRepo[MOD, CLS]) Exist(where func(db *gorm.DB, cls CLS) *gorm.DB) (bool, error) {
	var exists bool
	if err := where(repo.db, repo.cls).Model(new(MOD)).Select("1").Limit(1).Find(&exists).Error; err != nil {
		return false, err
	}
	return exists, nil
}

func (repo *GormRepo[MOD, CLS]) Find(where func(db *gorm.DB, cls CLS) *gorm.DB, dest *[]*MOD) *gorm.DB {
	return where(repo.db, repo.cls).Find(dest)
}

func (repo *GormRepo[MOD, CLS]) FindX(where func(db *gorm.DB, cls CLS) *gorm.DB) ([]*MOD, error) {
	var results []*MOD
	if err := repo.Find(where, &results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *GormRepo[MOD, CLS]) FindN(where func(db *gorm.DB, cls CLS) *gorm.DB, size int) ([]*MOD, error) {
	var results = make([]*MOD, 0, size)
	if err := where(repo.db, repo.cls).Limit(size).Find(&results).Error; err != nil {
		return nil, err
	}
	return results, nil
}

func (repo *GormRepo[MOD, CLS]) Update(where func(db *gorm.DB, cls CLS) *gorm.DB, valueFunc func(cls CLS) (string, interface{})) error {
	column, value := valueFunc(repo.cls)
	if err := where(repo.db, repo.cls).Model(new(MOD)).Update(column, value).Error; err != nil {
		return err
	}
	return nil
}

func (repo *GormRepo[MOD, CLS]) Updates(where func(db *gorm.DB, cls CLS) *gorm.DB, valuesFunc func(cls CLS) map[string]interface{}) error {
	mp := valuesFunc(repo.cls)
	if err := where(repo.db, repo.cls).Model(new(MOD)).Updates(mp).Error; err != nil {
		return err
	}
	return nil
}
