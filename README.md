[![GitHub Workflow Status (branch)](https://img.shields.io/github/actions/workflow/status/yyle88/gormrepo/release.yml?branch=main&label=BUILD)](https://github.com/yyle88/gormrepo/actions/workflows/release.yml?query=branch%3Amain)
[![GoDoc](https://pkg.go.dev/badge/github.com/yyle88/gormrepo)](https://pkg.go.dev/github.com/yyle88/gormrepo)
[![Coverage Status](https://img.shields.io/coveralls/github/yyle88/gormrepo/master.svg)](https://coveralls.io/github/yyle88/gormrepo?branch=main)
![Supported Go Versions](https://img.shields.io/badge/Go-1.22%2C%201.23-lightgrey.svg)
[![GitHub Release](https://img.shields.io/github/release/yyle88/gormcls.svg)](https://github.com/yyle88/gormrepo/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/yyle88/gormrepo)](https://goreportcard.com/report/github.com/yyle88/gormrepo)

# gormrepo - Isolate the scope of temporary variables when using GORM, making the code more concise

`gormrepo` isolates the **scope of temporary variables** when working with `GORM`, simplifying database operations and making the code more concise.

`gormrepo` works in conjunction with [gormcnm](https://github.com/yyle88/gormcnm) and [gormcngen](https://github.com/yyle88/gormcngen), simplifying GORM development and optimizing the management of temporary variable scopes.

---

## CHINESE README

[中文说明](README.zh.md)

---

## Installation

```bash
go get github.com/yyle88/gormrepo
```

---

## Quick Start

### Example Code

#### Query Data

```go
var example Example
if cls := gormcls.Cls(&Example{}); cls.OK() {
	err := db.Table(example.TableName()).Where(cls.Name.Eq("test")).First(&example).Error
    must.Done(err)
    fmt.Println("Fetched Name:", example.Name)
}
```

#### Update Data

```go
if one, cls := gormcls.Use(&Example{}); cls.OK() {
    err := db.Model(one).Where(cls.Name.Eq("test")).Update(cls.Age.Kv(30)).Error
    must.Done(err)
    fmt.Println("Age updated to:", 30)
}
```

#### Get Maximum Value

```go
var maxAge int
if one, cls := gormcls.Use(&Example{}); cls.OK() {
	err := db.Model(one).Select(cls.Age.COALESCE().MaxStmt("max_age")).First(&maxAge).Error
	must.Done(err)
    fmt.Println("Max Age:", maxAge)
}
```

---

## API Overview

| Function | Param | Return            | Description                                                                                                                                        | 
|----------|-------|-------------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| `Cls`    | `MOD` | `CLS`             | Returns the column information (`cls`), useful when only column data is needed.                                                                    |
| `Use`    | `MOD` | `MOD, CLS`        | Returns the model (`mod`) and its associated columns (`cls`), ideal for queries or operations that need both.                                      |
| `Umc`    | `MOD` | `MOD, CLS`        | Returns the model (`mod`) and its associated columns (`cls`), functioning identically to the `Use` function.                                       |
| `Usc`    | `MOD` | `[]MOD, CLS`      | Returns a slice of models (`MOD`) and the associated columns (`cls`), suitable for queries returning multiple models (e.g., `Find` queries).       |
| `Msc`    | `MOD` | `MOD, []MOD, CLS` | Returns the model (`mod`), the model slice (`[]MOD`), and the associated columns (`cls`), useful for queries requiring both model and column data. |
| `One`    | `MOD` | `MOD`             | Returns the model (`mod`), ensuring type safety by checking whether the argument is a pointer type at compile-time.                                |
| `Ums`    | `MOD` | `[]MOD`           | Returns a slice of models (`MOD`), useful for queries that expect a slice of models (e.g., `Find` queries).                                        |
| `Uss`    | -     | `[]MOD`           | Returns an empty slice of models (`MOD`), typically used for initialization or preparing for future object population without needing the columns. |
| `Usn`    | `int` | `[]MOD`           | Returns a slice of models (`MOD`) with a specified initial capacity, optimizing memory allocation based on the expected number of objects (`MOD`). |

---

## License

`gormrepo` is open-source and released under the MIT License. See the [LICENSE](LICENSE) file for more information.

---

## Support

Welcome to contribute to this project by submitting pull requests or reporting issues.

If you find this package helpful, give it a star on GitHub!

**Thank you for your support!**

**Happy Coding with `gormrepo`!** 🎉

Give me stars. Thank you!!!

## Starring

[![starring](https://starchart.cc/yyle88/gormcls.svg?variant=adaptive)](https://starchart.cc/yyle88/gormrepo)
