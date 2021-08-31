package model

import (
	"reflect-test/v2/internal/duck"
	"reflect-test/v2/internal/pond"
	stdMysql "reflect-test/v2/internal/std/mysql"
)

type Duck struct {
	stdMysql.DBRootCommon
	DuckID    *int    `db:"DuckID"`
	DuckUUID  *string `db:"DuckUUID"`
	Name      *string `db:"Name"`
	Color     *string `db:"Color"`
	Eggs      []Egg
	DuckPonds []DuckPond
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

	var dps []DuckPond

	if dReq.Ponds != nil {
		for _, p := range dReq.Ponds {
			dps = append(dps, NewDuckPond(*p.PondID))
		}
	}

	return &Duck{
		DuckID:    dReq.DuckID,
		DuckUUID:  dReq.DuckUUID,
		Name:      dReq.Name,
		Color:     dReq.Color,
		Eggs:      es,
		DuckPonds: dps,
	}
}

func (d *Duck) ToModel() duck.Duck {
	var es []duck.Egg

	if d.Eggs != nil {
		for _, e := range d.Eggs {
			es = append(es, e.ToModel())
		}
	}

	var ps []pond.Pond

	if d.DuckPonds != nil {
		for _, dp := range d.DuckPonds {
			ps = append(ps, dp.ToModel())
		}
	}

	return duck.Duck{
		DuckID:   d.DuckID,
		DuckUUID: d.DuckUUID,
		Name:     d.Name,
		Color:    d.Color,
		Eggs:     es,
		Ponds:    ps,
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

type DuckPond struct {
	stdMysql.DBAggregateCommon
	DuckPondID *int `db:"DuckPondID"`
	DuckID     *int `db:"DuckID"`
	PondID     *int `db:"PondID"`
}

func (dp *DuckPond) GetConfig() stdMysql.AggregateModelConfig {
	return stdMysql.AggregateModelConfig{
		TableName:   "duck_pond",
		IDField:     "DuckPondID",
		RootIDField: "DuckID",
	}
}

func NewDuckPond(pondID int) DuckPond {
	return DuckPond{PondID: &pondID}
}

func (dp *DuckPond) ToModel() pond.Pond {
	return pond.Pond{PondID: dp.PondID}
}
