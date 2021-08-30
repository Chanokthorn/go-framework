package std_mysql

import (
	"time"
)

type DBRootCommon struct {
	DBRootModel
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}

type DBAggregateCommon struct {
	DBAggregateModel
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}

type DBUpdatableAggregateCommon struct {
	DBUpdatableAggregateModel
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}

type RootModelConfig struct {
	TableName string
	IDField   string
	UUIDField string
}

type AggregateModelConfig struct {
	TableName         string
	IDField           string
	RootIDField       string
	RecursiveOnGetAll bool
}

type UpdatableAggregateModelConfig struct {
	TableName   string
	IDField     string
	UUIDField   string
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

type DBUpdatableAggregateModel interface {
	DBModel
	GetConfig() UpdatableAggregateModelConfig
}
