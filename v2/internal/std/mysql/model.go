package std_mysql

import (
	"reflect-test/v2/internal/std"
	"time"
)

type DBModelCommon struct {
	DBModel
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}

type StdConfig struct {
	TableName         string
	IDField           string
	UUIDField         string
	ParentIDField     string
	RecursiveOnGetAll bool
}

type DBModel interface {
	GetConfig() StdConfig
}

type DBRootModel interface {
	DBModel
	Set(domain std.DomainModel)
}

type DBAggregateModel interface {
	DBModel
	Set(domain std.DomainModel)
}
