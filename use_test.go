package gormcls

import (
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

	var account models.Account
	if cls := account.Columns(); cls.OK() {
		//if err := db.Where(***).First(&account).Error; err != nil {***}
	}

	var example models.Example
	if cls := example.Columns(); cls.OK() {
		//if err := db.Where(***).First(&example).Error; err != nil {***}
	}

	//这个时候你就很清楚自己用的是哪些变量，要不然代码复杂时就容易混淆
	t.Log(account.TableName(), example.TableName())
}
