# gormcls
跟前面的gormcnm配合使用，最终效果是很棒的

以下是个简单的demo，说明，就是当你使用gorm定义了些models，在操作时不可避免的要写
```
db.Table("example").Where("name=?", "bbb")).Update("age", 18).Error
```
但是很明显的，假如你的models很多，而且字段也很多的时候，将会造成混淆。

虽然在最开始就定义好所有的 models 的做法是可行的，但中间增删字段，或者修改字段名称或类型，这些操作在软件重构的时候都是很常见的

只要有一处漏改就会导致运行时BUG

而假如我们不实用硬编码，而是使用某种结构，把字段和类型存起来，当修改models以后自动修改结构，就能避免这个问题
```
var res Example
var cls = res.Columns()

// UPDATE `example` SET `age`=18 WHERE name="bbb"
db.Table(one.TableName()).Where(cls.Name.Eq("bbb")).Update(cls.Age.Kv(18)).Error

// SELECT * FROM `example` WHERE name="bbb" ORDER BY `example`.`id` LIMIT 1
db.Table(one.TableName()).Where(cls.Name.Eq("bbb")).First(&res).Error

t.Log("name:", res.Name, "age:", res.Age)
require.Equal(t, 18, res.Age)

// UPDATE `example` SET `age`=age + 2 WHERE name="bbb"
db.Table(one.TableName()).Where(cls.Name.Eq("bbb")).Update(cls.Age.KeAdd(2)).Error
```
以上代码中，res就是一个实体对象，而cls则是它的影子，这就叫做“如影随形”，能够让你在使用gorm时，避免使用任何的硬编码

详情请看完整的example:
[Example 文件](/example/example.go)
只是使用文件：
[Example 使用](/example/example_test.go)
