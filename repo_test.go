package gormrepo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

func TestRepo_NewGormRepo(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		res, err := repo.NewGormRepo(db).FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	require.NoError(t, db.Transaction(func(db *gorm.DB) error {
		res, err := repo.NewGormRepo(db).FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
		return nil
	}))
}

func TestRepo_GormRepo(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		res, err := repo.GormRepo(db).FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	require.NoError(t, db.Transaction(func(db *gorm.DB) error {
		res, err := repo.GormRepo(db).FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
		return nil
	}))
}

func TestRepo_Gorm(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		res, err := repo.Gorm(db).FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	require.NoError(t, db.Transaction(func(db *gorm.DB) error {
		res, err := repo.Gorm(db).FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
		return nil
	}))
}
