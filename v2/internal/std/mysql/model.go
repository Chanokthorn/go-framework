package std_mysql

import (
	"time"
)

type dbModel interface {
	dbModelFunc()
}

// RootModel represent root model that the repository stores, the model must implement this interface
// by embedding RootCommon and has GetConfig() that returns RootConfig
// for example:
//	type Duck struct {
//		std_mysql.RootCommon
//		DuckID   *int    `db:"DuckID"`
//		DuckUUID *string `db:"DuckUUID"`
//		Name     *string `db:"Name"`
//		Color    *string `db:"Color"`
//		Eggs     []Egg
//	}
//
//	func (d *Duck) GetConfig() stdMysql.RootConfig {
//		return std_mysql.RootConfig{
//			TableName: "duck",
//			IDField:   "DuckID",
//			UUIDField: "DuckUUID",
//		}
//	}
type RootModel interface {
	dbModel
	GetConfig() RootConfig
}

// AggregateModel represent aggregate model that the root model contains, the model must implement this interface
// by embedding AggregateCommon and has GetConfig() that returns AggregateConfig
// for example:
//	type Egg struct {
//		std_mysql.AggregateCommon
//		EggID  *int    `db:"EggID"`
//		DuckID *int    `db:"DuckID"`
//		Name   *string `db:"Name"`
//		Age    *int    `db:"Age"`
//	}
//
//	func (e *Egg) GetConfig() stdMysql.AggregateConfig {
//		return std_mysql.AggregateConfig{
//			TableName:   "egg",
//			IDField:     "EggID",
//			RootIDField: "DuckID",
//		}
//	}
type AggregateModel interface {
	dbModel
	GetConfig() AggregateConfig
}

// RootConfig contains property of root model IDField, UUIDField, RootIDField must be the same name for both DB column
// name and struct field name. RootIDField is optional in case the Root model contains ID of another Root
type RootConfig struct {
	TableName   string
	IDField     string
	UUIDField   string
	RootIDField string
}

// AggregateConfig contains property of aggregate model IDField, RootIDField must be the same name for both DB column
// name and struct field name. All fields are required.
type AggregateConfig struct {
	TableName   string
	IDField     string
	RootIDField string
}

type dbCommon interface {
	dbCommonFunc()
}

// RootCommon contains basic MySQL fields and implements RootModel. Can be embedded by struct to represent Root model.
type RootCommon struct {
	dbCommon
	RootModel
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}

// AggregateCommon contains basic MySQL fields and implements AggregateCommon. Can be embedded by struct to
//	represent Aggregate model.
type AggregateCommon struct {
	dbCommon
	AggregateModel
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}
