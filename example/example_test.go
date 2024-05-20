package example

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcls"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var caseDB *gorm.DB

func TestMain(m *testing.M) {
	fmt.Println("run_test_main")
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()

	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	done.Done(db.AutoMigrate(&Example{}))

	caseDB = db //a global variable caseDB
	m.Run()
}

func TestExample(t *testing.T) {
	var db = caseDB

	example1 := &Example{
		ID:        0,
		Name:      "aaa",
		Age:       1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	example2 := &Example{
		ID:        0,
		Name:      "bbb",
		Age:       2,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	require.NoError(t, db.Create(example1).Error)
	require.NoError(t, db.Create(example2).Error)

	var resA Example
	if one, cls := gormcls.Use(&Example{}); cls.OK() {
		require.NoError(t, db.Table(one.TableName()).Where(cls.Name.Eq("aaa")).First(&resA).Error)
	}
	t.Log("select res.name:", resA.Name)
	require.Equal(t, "aaa", resA.Name)

	var maxAge int
	if one, cls := gormcls.Use(&Example{}); cls.OK() {
		require.NoError(t, db.Model(one).Where(cls.Age.Gt(0)).Select(cls.Age.CoalesceMaxStmt("age_alias")).First(&maxAge).Error)
	}
	t.Log("max_age:", maxAge)
	require.Equal(t, 2, maxAge)

	if one, cls := gormcls.Use(&Example{}); cls.OK() {
		var resB Example

		require.NoError(t, db.Table(one.TableName()).Where(cls.Name.Eq("bbb")).Update(cls.Age.Kv(18)).Error)

		require.NoError(t, db.Table(one.TableName()).Where(cls.Name.Eq("bbb")).First(&resB).Error)

		t.Log("new_age:", resB.Age)
		require.Equal(t, 18, resB.Age)
	}

	var resB Example
	if one, cls := gormcls.Use(&Example{}); cls.OK() {
		require.NoError(t, db.Table(one.TableName()).Where(cls.Name.Eq("bbb")).Update(cls.Age.KeAdd(2)).Error)

		require.NoError(t, db.Table(one.TableName()).Where(cls.Name.Eq("bbb")).First(&resB).Error)
	}
	t.Log("new_age:", resB.Age)
	require.Equal(t, 20, resB.Age)
}
