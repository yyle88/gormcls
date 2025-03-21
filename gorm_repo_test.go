package gormrepo_test

import (
	"sort"
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

func TestGormRepo_First(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))
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

func TestGormRepo_FirstX(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))
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

func TestGormRepo_FirstE(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))
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

func TestGormRepo_Where(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))
	require.True(t, repo.OK())

	{
		var nicknames []string
		db := repo.Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).
				Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"})).
				Distinct(cls.Nickname.Name())
		})
		require.NoError(t, db.Find(&nicknames).Error)
		t.Log(nicknames)
		require.Len(t, nicknames, 2)
		sort.Strings(nicknames)
		require.Equal(t, []string{"demo-1-nickname", "demo-2-nickname"}, nicknames)
	}
	{
		var nickname string
		db := repo.Where(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
			return db.Model(&Account{}).
				Select("MIN(" + cls.Nickname.Name() + ")").
				Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
		})
		require.NoError(t, db.First(&nickname).Error)
		t.Log(nickname)
		require.Equal(t, "demo-1-nickname", nickname)
	}
}

func TestGormRepo_Exist(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Umc(caseDB, &Account{}))
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

func TestGormRepo_Find(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))
	require.True(t, repo.OK())

	var accounts []*Account
	require.NoError(t, repo.Find(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Like("demo-%-username"))
	}, &accounts).Error)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

func TestGormRepo_FindX(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))
	require.True(t, repo.OK())

	accounts, err := repo.FindX(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Like("demo-%-username"))
	})
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

func TestGormRepo_FindN(t *testing.T) {
	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))
	require.True(t, repo.OK())

	accounts, err := repo.FindN(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.In([]string{"demo-1-username", "demo-2-username"}))
	}, 2)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	t.Log(neatjsons.S(accounts))
}

func TestGormRepo_Update(t *testing.T) {
	username := uuid.New().String()

	require.NoError(t, caseDB.Save(&Account{
		Model:    gorm.Model{},
		Username: username,
		Password: uuid.New().String(),
		Nickname: uuid.New().String(),
	}).Error)

	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))
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

func TestGormRepo_Updates(t *testing.T) {
	username := uuid.New().String()

	require.NoError(t, caseDB.Save(&Account{
		Model:    gorm.Model{},
		Username: username,
		Password: uuid.New().String(),
		Nickname: uuid.New().String(),
	}).Error)

	repo := gormrepo.NewGormRepo(gormrepo.Use(caseDB, &Account{}))
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
