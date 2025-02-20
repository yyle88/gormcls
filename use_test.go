package gormcls_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormcls"
	"github.com/yyle88/gormcngen"
	"github.com/yyle88/gormcnm"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
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

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"unique;"`
	Age       int       `gorm:"type:int32;index:idx_example_age;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (*Example) TableName() string {
	return "examples"
}

func (a *Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        gormcnm.Cnm(a.ID, "id"),
		Name:      gormcnm.Cnm(a.Name, "name"),
		Age:       gormcnm.Cnm(a.Age, "age"),
		CreatedAt: gormcnm.Cnm(a.CreatedAt, "created_at"),
		UpdatedAt: gormcnm.Cnm(a.UpdatedAt, "updated_at"),
	}
}

type ExampleColumns struct {
	// Embedding operation functions make it easy to use // 继承操作函数便于使用
	gormcnm.ColumnOperationClass
	// The column names and types of the model's columns // 模型各列的列名和类型
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	Age       gormcnm.ColumnName[int]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
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
	objects := []any{&Account{}, &Example{}}

	options := gormcngen.NewOptions().
		WithColumnClassExportable(true). // Generate exportable struct names (e.g., ExampleColumns) // 生成可导出的结构体名称（例如 ExampleColumns）
		WithColumnsMethodRecvName("a").
		WithColumnsCheckFieldType(true)

	// Configure code generation settings
	// 配置代码生成设置
	cfg := gormcngen.NewConfigs(objects, options, absPath)
	cfg.Gen() // Generate and write the code to the target location (e.g., "gormcnm.gen.go") // 生成并将代码写入目标位置（例如 "gormcnm.gen.go"）
}

// Demonstrates the usage of gormcls with the Account struct.
// 演示如何使用 gormcls 处理 Account 结构体。
func TestUseAccount(t *testing.T) {
	// Example: Using gormcls with Account model
	// 示例：使用 gormcls 处理 Account 模型
	if account, cls := gormcls.Use(&Account{}); cls.OK() {
		t.Logf("TableName: %s", account.TableName())
		t.Logf("Columns: %s", neatjsons.S(cls))
	}

	// Alternative: Limit the scope of the account and cls variables for better control
	// 其它方式：限制 account 和 cls 变量的作用范围以更好地控制
	{
		var account Account
		var cls = account.Columns()
		require.True(t, cls.OK(), "Expected Columns to be OK")
	}
}

// Demonstrates the usage of gormcls with the Example struct.
// 演示如何使用 gormcls 处理 Example 结构体。
func TestUseExample(t *testing.T) {
	// Example: Using gormcls with Example model
	// 示例：使用 gormcls 处理 Example 模型
	if example, cls := gormcls.Use(&Example{}); cls.OK() {
		t.Logf("TableName: %s", example.TableName())
		t.Logf("Columns: %s", neatjsons.S(cls))
	}

	// Alternative: Limit the scope of the example and cls variables for better control
	// 其它方式：限制 example 和 cls 变量的作用范围以更好地控制
	{
		var example Example
		var cls = example.Columns()
		require.True(t, cls.OK(), "Expected Columns to be OK")
	}
}

// Demonstrates the usage of both Account and Example models in the same test case.
// 演示如何在同一测试中使用 Account 和 Example 模型。
func TestAccountAndExample(t *testing.T) {
	var account Account
	if cls := account.Columns(); cls.OK() {
		t.Log("Account columns are valid")
	}

	var example Example
	if cls := example.Columns(); cls.OK() {
		t.Log("Example columns are valid")
	}

	t.Logf("Account TableName: %s", account.TableName())
	t.Logf("Example TableName: %s", example.TableName())
}

// Demonstrates how to retrieve column information for the Account model.
// 演示如何访问 Account 模型的列信息。
func TestColumnsWithAccount(t *testing.T) {
	cls := gormcls.Cls(&Account{})
	require.True(t, cls.OK(), "Expected cls to be OK for Account")
	t.Logf("Account Columns: %s", neatjsons.S(cls))
}

// Demonstrates how gormcls ensures the model is treated as a pointer type.
// 演示 gormcls 如何确保模型被当作指针类型处理。
func TestEnsurePointerInputType(t *testing.T) {

	{
		var account Account
		one := gormcls.One(&account)
		require.Equal(t, reflect.Ptr, reflect.TypeOf(one).Kind(), "Expected pointer type")
	}

	{
		example := Example{}
		one := gormcls.One(&example)
		require.Equal(t, reflect.Ptr, reflect.TypeOf(one).Kind(), "Expected pointer type")
	}
}

// Demonstrates the usage of the Ums function with the Example model.
// 演示如何使用 Ums 函数处理 Example 模型。
func TestUmsWithExample(t *testing.T) {
	examples := gormcls.Ums(&Example{})
	t.Logf("Ums result: %s", neatjsons.S(examples))
}

// Demonstrates the usage of the Uss function with the Example model.
// 演示如何使用 Uss 函数处理 Example 模型。
func TestUssWithExample(t *testing.T) {
	examples := gormcls.Uss[*Example]()
	t.Logf("Uss result: %s", neatjsons.S(examples))
	t.Logf("result cap: %d", cap(examples))
}

// Demonstrates the usage of the Usn function with the Example model.
// 演示如何使用 Usn 函数处理 Example 模型。
func TestUsnWithExample(t *testing.T) {
	examples := gormcls.Usn[*Example](100)
	t.Logf("Usn result: %s", neatjsons.S(examples))
	t.Logf("result cap: %d", cap(examples))
}

// Demonstrates the usage of the Usc function with the Example model.
// 演示如何使用 Usc 函数处理 Example 模型。
func TestUscWithExample(t *testing.T) {
	examples, cls := gormcls.Usc(&Example{})
	require.True(t, cls.OK(), "Expected cls to be OK for Example")
	t.Logf("Usc result: %s", neatjsons.S(examples))
}

// Demonstrates the usage of the Msc function with the Example model.
// 演示如何使用 Msc 函数处理 Example 模型。
func TestMscWithExample(t *testing.T) {
	one, examples, cls := gormcls.Msc(&Example{})
	require.True(t, cls.OK(), "Expected cls to be OK for Example")
	t.Logf("Msc TableName: %s", one.TableName())
	t.Logf("Msc examples: %s", neatjsons.S(examples))
}

func TestExample(t *testing.T) {
	db := done.VCE(gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})).Nice()
	defer func() {
		done.Done(done.VCE(db.DB()).Nice().Close())
	}()

	done.Done(db.AutoMigrate(&Example{}))

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
	if cls := gormcls.Cls(&Example{}); cls.OK() {
		require.NoError(t, db.Table(resA.TableName()).Where(cls.Name.Eq("aaa")).First(&resA).Error)
		require.Equal(t, "aaa", resA.Name)
	}
	t.Log("select res.name:", resA.Name)

	var maxAge int
	if one, cls := gormcls.Use(&Example{}); cls.OK() {
		require.NoError(t, db.Model(one).Where(cls.Age.Gt(0)).Select(cls.Age.COALESCE().MaxStmt("age_alias")).First(&maxAge).Error)
		require.Equal(t, 2, maxAge)
	}
	t.Log("max_age:", maxAge)

	if one, cls := gormcls.Use(&Example{}); cls.OK() {
		require.NoError(t, db.Model(one).Where(cls.Name.Eq("bbb")).Update(cls.Age.Kv(18)).Error)
		require.Equal(t, 18, one.Age)
	}

	var resB Example
	if cls := resB.Columns(); cls.OK() {
		require.NoError(t, db.Table(resB.TableName()).Where(cls.Name.Eq("bbb")).Update(cls.Age.KeAdd(2)).Error)

		require.NoError(t, db.Table(resB.TableName()).Where(cls.Name.Eq("bbb")).First(&resB).Error)
		require.Equal(t, 20, resB.Age)
	}
	t.Log(resB)
}
