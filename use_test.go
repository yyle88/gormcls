package gormcls

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/gormcls/internal/examples/example1/models"
)

func TestUse(t *testing.T) {
	//假如需要使用 cls 这个变量
	if account, cls := Use(&models.Account{}); cls.OK() {
		t.Log(account.TableName())
	}

	{ //跟这个是等效的，把变量 account 和 cls 限制在较小的作用域内
		var account models.Account
		var cls = account.Columns()
		require.True(t, cls.OK())
	}

	//假如需要再次使用 cls 这个变量
	if example, cls := Use(&models.Example{}); cls.OK() {
		t.Log(example.TableName())
	}

	{ //跟这个是等效的，把变量 example 和 cls 限制在较小的作用域内
		var example models.Example
		var cls = example.Columns()
		require.True(t, cls.OK())
	}

	//当需要暴露这个变量到下文时
	var account models.Account
	if cls := account.Columns(); cls.OK() {
		//if err := db.Where(***).First(&account).Error; err != nil {***}
		t.Log("-")
	}

	//当需要暴露这个变量到下文时
	var example models.Example
	if cls := example.Columns(); cls.OK() {
		//if err := db.Where(***).First(&example).Error; err != nil {***}
		t.Log("-")
	}

	//这个时候你就很清楚自己用的是哪些变量，要不然代码复杂时就容易混淆
	t.Log(account.TableName(), example.TableName())
}

func TestOne(t *testing.T) {
	{
		var account = makeAccount()
		// 由于这里返回的不是指针，在存储的时候就需要取地址，这样运行时才不会报错
		// if err := db.Create(&account).Error; err != nil {***}

		// 2000 THOUSAND YEARS LATER.
		// 假设过了很长时间。
		// 假如哪天重构代码时，我又把前面的函数改为了返回指针类型，这时候代码在静态检查时感知不到错误，但实际传的是二级指针类型，这就完蛋啦

		// 因此我发明了这个函数，能够确保你传给 db.Create 的就是指针类型的数据
		// 这样当重构函数时，把返回值改为二级指针也没事，在静态检查时就会报错，也就能让人更放心大胆的写逻辑
		one := One(&account)
		// 只有 one 是指针的时候才能调用 Save 函数
		// if err := db.Save(one).Error; err != nil {***}
		// 这样就没问题啦
		require.Equal(t, reflect.Ptr, reflect.TypeOf(one).Kind())
	}
	{
		example := models.Example{}
		// one := One(example) // 静态检查不过，编译不过，能有效避免问题
		one := One(&example)
		// if err := db.Create(one).Error; err != nil {***}
		// 这样就没问题啦
		require.Equal(t, reflect.Ptr, reflect.TypeOf(one).Kind())
	}
}

func makeAccount() models.Account {
	return models.Account{}
}
