package model

import (
	"reflect-test/v2/internal/duck"
	stdMysql "reflect-test/v2/internal/std/mysql"
)

type Duck struct {
	stdMysql.DBRootCommon
	DuckID   *int    `db:"DuckID"`
	DuckUUID *string `db:"DuckUUID"`
	Name     *string `db:"Name"`
	Color    *string `db:"Color"`
	Eggs     []Egg
}

func (d *Duck) GetConfig() stdMysql.RootModelConfig {
	return stdMysql.RootModelConfig{
		TableName: "duck",
		IDField:   "DuckID",
		UUIDField: "DuckUUID",
	}
}

func NewDuck(dReq duck.Duck) *Duck {
	var es []Egg
	if dReq.Eggs != nil {
		for _, e := range dReq.Eggs {
			es = append(es, *NewEgg(e))
		}
	}

	return &Duck{
		DuckID:   dReq.DuckID,
		DuckUUID: dReq.DuckUUID,
		Name:     dReq.Name,
		Color:    dReq.Color,
		Eggs:     es,
	}
}

func (d *Duck) ToModel() duck.Duck {
	var es []duck.Egg
	if d.Eggs != nil {
		for _, e := range d.Eggs {
			es = append(es, e.ToModel())
		}
	}

	return duck.Duck{
		DuckID:   d.DuckID,
		DuckUUID: d.DuckUUID,
		Name:     d.Name,
		Color:    d.Color,
		Eggs:     es,
	}
}

type Egg struct {
	stdMysql.DBAggregateCommon
	EggID  *int    `db:"EggID"`
	DuckID *int    `db:"DuckID"`
	Name   *string `db:"Name"`
	Age    *int    `db:"Age"`
}

func (e *Egg) GetConfig() stdMysql.AggregateModelConfig {
	return stdMysql.AggregateModelConfig{
		TableName:   "egg",
		IDField:     "EggID",
		RootIDField: "DuckID",
	}
}

func NewEgg(eReq duck.Egg) *Egg {
	return &Egg{
		Name: eReq.Name,
		Age:  eReq.Age,
	}
}

func (e *Egg) ToModel() duck.Egg {
	return duck.Egg{
		Name: e.Name,
		Age:  e.Age,
	}
}
