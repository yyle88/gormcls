package gormrepo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/gormclass"
	"gorm.io/gorm"
)

func TestRepo_Gorm(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		res, err := repo.Gorm(caseDB).FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		res, err := repo.Gorm(db).FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
		return nil
	}))
}

func TestRepo_MoDB(t *testing.T) {
	repo := gormrepo.NewRepo(gormclass.Use(&Account{}))

	{
		var nickname string
		require.NoError(t, repo.MoDB(caseDB).WhereE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		}, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Select(string(cls.Nickname)).First(&nickname)
		}))
		require.Equal(t, "demo-1-nickname", nickname)
	}

	require.NoError(t, caseDB.Transaction(func(db *gorm.DB) error {
		var nickname string
		require.NoError(t, repo.MoDB(db).WhereE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		}, func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Select(string(cls.Nickname)).First(&nickname)
		}))
		require.Equal(t, "demo-2-nickname", nickname)
		return nil
	}))
}
