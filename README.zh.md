# gormrepo - 隔离临时变量的作用域，简化 GORM 操作

`gormrepo` 在使用 `GORM` 时，**隔离临时变量的作用域**，简化数据库操作，使代码更加简洁。

`gormrepo` 跟 [gormcnm](https://github.com/yyle88/gormcnm) 和 [gormcngen](https://github.com/yyle88/gormcngen) 配合使用，能简化 GORM 开发并优化临时变量作用域的管理。

---

## 英文文档

[ENGLISH README](README.md)

---

## 安装

```bash
go get github.com/yyle88/gormrepo
```

---

## 快速开始

### 示例代码

#### 查询数据

```go
var example Example
if cls := gormcls.Cls(&Example{}); cls.OK() {
	err := db.Table(example.TableName()).Where(cls.Name.Eq("test")).First(&example).Error
    must.Done(err)
    fmt.Println("Fetched Name:", example.Name)
}
```

#### 更新数据

```go
if one, cls := gormcls.Use(&Example{}); cls.OK() {
    err := db.Model(one).Where(cls.Name.Eq("test")).Update(cls.Age.Kv(30)).Error
    must.Done(err)
    fmt.Println("Age updated to:", 30)
}
```

#### 获取最大值

```go
var maxAge int
if one, cls := gormcls.Use(&Example{}); cls.OK() {
	err := db.Model(one).Select(cls.Age.COALESCE().MaxStmt("max_age")).First(&maxAge).Error
	must.Done(err)
    fmt.Println("Max Age:", maxAge)
}
```

---

## API 概览

| 函数    | 参数    | 返回                | 描述                                                    | 
|-------|-------|-------------------|-------------------------------------------------------|
| `Cls` | `MOD` | `CLS`             | 返回列信息（`cls`），适用于仅需要列数据的场景。                            |
| `Use` | `MOD` | `MOD, CLS`        | 返回模型（`mod`）、关联的列（`cls`），适用于需要同时获取模型和列数据的查询或操作。        |
| `Umc` | `MOD` | `MOD, CLS`        | 返回模型（`mod`）、关联的列（`cls`），功能与 `Use` 函数相同。               |
| `Usc` | `MOD` | `[]MOD, CLS`      | 返回多个模型（`MOD`）、关联的列（`cls`），适用于返回多个模型的查询（如 `Find` 查询）。  |
| `Msc` | `MOD` | `MOD, []MOD, CLS` | 返回模型（`mod`）、模型切片（`[]MOD`）、关联的列（`cls`），适用于需要模型和列数据的查询。 |
| `One` | `MOD` | `MOD`             | 返回模型（`mod`），通过编译时检查确保类型安全。                            |
| `Ums` | `MOD` | `[]MOD`           | 返回模型（`MOD`）切片，适用于需要模型切片的查询（例如 `Find` 查询）。             |
| `Uss` | -     | `[]MOD`           | 返回一个空的模型（`MOD`）切片，通常用于初始化或为未来填充对象做准备，无需关联列（`cls`）。    |
| `Usn` | `int` | `[]MOD`           | 返回一个具有指定初始容量的模型（`MOD`）切片，优化内存分配以适应预期的对象数量。            |

---

## 许可

`gormrepo` 是一个开源项目，发布于 MIT 许可证下。有关更多信息，请参阅 [LICENSE](LICENSE) 文件。

## 贡献与支持

欢迎通过提交 pull request 或报告问题来贡献此项目。

如果你觉得这个包对你有帮助，请在 GitHub 上给个 ⭐，感谢支持！！！

**感谢你的支持！**

**祝编程愉快！** 🎉

Give me stars. Thank you!!!
