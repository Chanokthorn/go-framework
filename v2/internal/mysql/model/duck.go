package model

import (
	"reflect-test/v2/internal/duck"
	stdMysql "reflect-test/v2/internal/std/mysql"
)

type Duck struct {
	stdMysql.DBModelCommon
	ID    *int    `db:"DuckID"`
	UUID  *string `db:"DuckUUID"`
	Name  *string `db:"Name"`
	Color *string `db:"Color"`
}

func NewDuck(dReq duck.Duck) *Duck {
	return &Duck{
		Name:  dReq.Name,
		Color: dReq.Color,
	}
}

func (d *Duck) GetConfig() stdMysql.StdConfig {
	return stdMysql.StdConfig{
		TableName:         "duck",
		IDField:           "DuckID",
		UUIDField:         "DuckUUID",
		ParentIDField:     "",
		RecursiveOnGetAll: false,
	}
}

func (d *Duck) ToModel() duck.Duck {
	return duck.Duck{
		Name:  d.Name,
		Color: d.Color,
	}
}
