package gormrepo_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/must"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		must.Done(rese.P1(db.DB()).Close())
	}()

	done.Done(db.AutoMigrate(&Account{}))

	must.Done(db.Save(&Account{
		Model:    gorm.Model{},
		Username: "demo-1-username",
		Password: "demo-1-password",
		Nickname: "demo-1-nickname",
	}).Error)
	must.Done(db.Save(&Account{
		Model:    gorm.Model{},
		Username: "demo-2-username",
		Password: "demo-2-password",
		Nickname: "demo-2-nickname",
	}).Error)

	caseDB = db
	m.Run()
}

func TestRepo_First(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormrepo.Umc(db, &Account{}))
	require.True(t, repo.OK())

	{
		var account Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		}, &account).Error)
		require.Equal(t, "demo-1-nickname", account.Nickname)
	}

	{
		var account Account
		require.NoError(t, repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		}, &account).Error)
		require.Equal(t, "demo-2-nickname", account.Nickname)
	}
}

func TestRepo_FirstX(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormrepo.Umc(db, &Account{}))
	require.True(t, repo.OK())

	{
		res, err := repo.FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	{
		res, err := repo.FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-2-username"))
		})
		require.NoError(t, err)
		require.Equal(t, "demo-2-nickname", res.Nickname)
	}
}

func TestRepo_FirstE(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormrepo.Umc(db, &Account{}))
	require.True(t, repo.OK())

	{
		res, erb := repo.FirstE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.Nil(t, erb)
		require.Equal(t, "demo-1-nickname", res.Nickname)
	}

	{
		res, erb := repo.FirstE(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-x-username"))
		})
		require.NotNil(t, erb)
		require.ErrorIs(t, erb.ErrCause, gorm.ErrRecordNotFound)
		require.True(t, erb.NotExist)
		require.Nil(t, res)
	}
}

func TestRepo_Exist(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormrepo.Umc(db, &Account{}))
	require.True(t, repo.OK())

	{
		exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-1-username"))
		})
		require.NoError(t, err)
		require.True(t, exist)
	}

	{
		exist, err := repo.Exist(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Where(cls.Username.Eq("demo-x-username"))
		})
		require.NoError(t, err)
		require.False(t, exist)
	}
}

func TestRepo_Find(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormrepo.Use(db, &Account{}))
	require.True(t, repo.OK())

	var accounts []*Account
	require.NoError(t, repo.Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Like("demo-%-username"))
	}, &accounts).Error)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

func TestRepo_FindX(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormrepo.Use(db, &Account{}))
	require.True(t, repo.OK())

	accounts, err := repo.FindX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Like("demo-%-username"))
	})
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

func TestRepo_FindN(t *testing.T) {
	db := caseDB

	repo := gormrepo.NewRepo(gormrepo.Use(db, &Account{}))
	require.True(t, repo.OK())

	accounts, err := repo.FindN(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	}, 2)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

func TestRepo_Update(t *testing.T) {
	db := caseDB

	username := uuid.New().String()

	require.NoError(t, db.Save(&Account{
		Model:    gorm.Model{},
		Username: username,
		Password: uuid.New().String(),
		Nickname: uuid.New().String(),
	}).Error)

	repo := gormrepo.NewRepo(gormrepo.Use(db, &Account{}))
	require.True(t, repo.OK())

	newNickname := uuid.New().String()
	err := repo.Update(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) (string, interface{}) {
		return cls.Nickname.Kv(newNickname)
	})
	require.NoError(t, err)

	res, err := repo.FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, newNickname, res.Nickname)
}

func TestRepo_Updates(t *testing.T) {
	db := caseDB

	username := uuid.New().String()

	require.NoError(t, db.Save(&Account{
		Model:    gorm.Model{},
		Username: username,
		Password: uuid.New().String(),
		Nickname: uuid.New().String(),
	}).Error)

	repo := gormrepo.NewRepo(gormrepo.Use(db, &Account{}))
	require.True(t, repo.OK())

	newNickname := uuid.New().String()
	newPassword := uuid.New().String()
	err := repo.Updates(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	}, func(cls *AccountColumns) map[string]interface{} {
		return cls.
			Kw(cls.Nickname.Kv(newNickname)).
			Kw(cls.Password.Kv(newPassword)).
			AsMap()
	})
	require.NoError(t, err)

	res, err := repo.FirstX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq(username))
	})
	require.NoError(t, err)
	require.Equal(t, newNickname, res.Nickname)
	require.Equal(t, newPassword, res.Password)
}
