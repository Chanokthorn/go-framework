package std

import "time"

type DomainModel interface {
	Domain()
}

type DBModel interface {
	ToDomain() DomainModel
}

type DBModelCommon struct {
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}

type DBRootModel interface {
	DBModel
	Set(domain DomainModel)
}

type DBAggregateModel interface {
	DBModel
	Set(domain DomainModel)
}
