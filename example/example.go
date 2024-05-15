package example

import (
	"time"

	"github.com/yyle88/gormcnm"
)

type Example struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"unique;"`
	Age       int       `gorm:"type:int32;index:idx_example_age;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (*Example) TableName() string {
	return "example"
}

// auto generate code (use github.com/yyle88/gormcngen):
// auto generate code (use github.com/yyle88/gormcngen):
// auto generate code (use github.com/yyle88/gormcngen):

func (*Example) Columns() *ExampleColumns {
	return &ExampleColumns{
		ID:        "id",
		Name:      "name",
		Age:       "age",
		CreatedAt: "created_at",
		UpdatedAt: "updated_at",
	}
}

type ExampleColumns struct {
	gormcnm.ColumnBaseFuncClass //继承操作函数，让查询更便捷
	// 模型各个列名和类型:
	ID        gormcnm.ColumnName[int32]
	Name      gormcnm.ColumnName[string]
	Age       gormcnm.ColumnName[int]
	CreatedAt gormcnm.ColumnName[time.Time]
	UpdatedAt gormcnm.ColumnName[time.Time]
}
