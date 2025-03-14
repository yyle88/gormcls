package gormrepo_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcls/gormrepo"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/must"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/rese"
	"github.com/yyle88/runpath"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Account struct {
	gorm.Model
	Username string `gorm:"unique;"`
	Password string `gorm:"size:255;"`
	Nickname string `gorm:"column:nickname;"`
}

func (*Account) TableName() string {
	return "accounts"
}

func (a *Account) Columns() *AccountColumns {
	return &AccountColumns{
		ID:        gormcnm.Cnm(a.ID, "id"),
		CreatedAt: gormcnm.Cnm(a.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(a.UpdatedAt, "updated_at"),
		DeletedAt: gormcnm.Cnm(a.DeletedAt, "deleted_at"),
		Username:  gormcnm.Cnm(a.Username, "username"),
		Password:  gormcnm.Cnm(a.Password, "password"),
		Nickname:  gormcnm.Cnm(a.Nickname, "nickname"),
	}
}

type AccountColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[uint]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
	DeletedAt gormcnm.ColumnName[gorm.DeletedAt]
	Username  gormcnm.ColumnName[string]
	Password  gormcnm.ColumnName[string]
	Nickname  gormcnm.ColumnName[string]
}

// Tests the generation of columns for models.
// 测试模型列的生成。
func TestGenerateColumns(t *testing.T) {
	absPath := runpath.Path() // Retrieve the absolute path of the source file based on the current test file's location
	// 获取当前测试文件位置基础上的源文件绝对路径
	t.Log(absPath)

	// Verify the existence of the target file. The file should be created manually to ensure it can be located by the code.
	// 检查目标文件是否存在。文件应手动创建，确保代码能够找到它。
	require.True(t, osmustexist.IsFile(absPath))

	// List the models for which columns will be generated. Both pointer and non-pointer types are supported.
	// 设置需要生成列的模型，这里支持指针类型和非指针类型。
	objects := []any{&Account{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable struct names (e.g., ExampleColumns) // 生成可导出的结构体名称（例如 ExampleColumns）
		WithColumnsMethodRecvName("a").
		WithColumnsCheckFieldType(true)

	// Configure code generation settings
	// 配置代码生成设置
	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen() // Generate and write the code to the target location (e.g., "gormcnm.gen.go") // 生成并将代码写入目标位置（例如 "gormcnm.gen.go"）
}

func TestRepo_First(t *testing.T) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		must.Done(rese.P1(db.DB()).Close())
	}()

	done.Done(db.AutoMigrate(&Account{}))

	require.NoError(t, db.Save(&Account{
		Model:    gorm.Model{},
		Username: "demo-1-username",
		Password: "demo-1-password",
		Nickname: "demo-1-nickname",
	}).Error)
	require.NoError(t, db.Save(&Account{
		Model:    gorm.Model{},
		Username: "demo-2-username",
		Password: "demo-2-password",
		Nickname: "demo-2-nickname",
	}).Error)

	repo := gormrepo.NewRepo(gormrepo.Dmc(db, &Account{}))
	require.True(t, repo.OK())

	one, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-1-username"))
	})
	require.NoError(t, err)
	require.Equal(t, "demo-1-nickname", one.Nickname)

	two, err := repo.First(func(db *gorm.DB, cls *AccountColumns) *gorm.DB {
		return db.Where(cls.Username.Eq("demo-2-username"))
	})
	require.NoError(t, err)
	require.Equal(t, "demo-2-nickname", two.Nickname)
}
