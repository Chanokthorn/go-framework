package model

import (
	"reflect-test/v2/internal/duck"
	stdMysql "reflect-test/v2/internal/std/mysql"
)

type Duck struct {
	stdMysql.DBRootCommon
	ID    *int    `db:"DuckID"`
	UUID  *string `db:"DuckUUID"`
	Name  *string `db:"Name"`
	Color *string `db:"Color"`
	Eggs  []Egg   `fakesize:"3"`
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
		ID:    dReq.ID,
		UUID:  dReq.UUID,
		Name:  dReq.Name,
		Color: dReq.Color,
		Eggs:  es,
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
		ID:    d.ID,
		UUID:  d.UUID,
		Name:  d.Name,
		Color: d.Color,
		Eggs:  es,
	}
}

type Egg struct {
	stdMysql.DBAggregateCommon
	ID     *int    `db:"EggID"`
	RootID *int    `db:"DuckID"`
	Name   *string `db:"Name"`
	Age    *int    `db:"Age"`
}

func (e *Egg) GetConfig() stdMysql.AggregateModelConfig {
	return stdMysql.AggregateModelConfig{
		TableName:         "egg",
		IDField:           "EggID",
		UUIDField:         "EggUUID",
		RootIDField:       "DuckID",
		RecursiveOnGetAll: false,
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
