package std_mysql

import (
	"time"
)

type DBCommon interface {
	dbCommon()
}

type DBRootCommon struct {
	DBCommon
	DBRootModel
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}

type DBAggregateCommon struct {
	DBCommon
	DBAggregateModel
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}

type RootModelConfig struct {
	TableName   string
	IDField     string
	UUIDField   string
	RootIDField string
}

type AggregateModelConfig struct {
	TableName   string
	IDField     string
	RootIDField string
}

type DBModel interface {
}

type DBRootModel interface {
	DBModel
	GetConfig() RootModelConfig
}

type DBAggregateModel interface {
	DBModel
	GetConfig() AggregateModelConfig
}
