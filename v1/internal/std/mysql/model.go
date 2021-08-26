package mysql

import "time"

type DBModelCommon struct {
	CreatedBy   *string    `db:"CreatedBy"`
	CreatedDate *time.Time `db:"CreatedDate"`
	UpdatedBy   *string    `db:"UpdatedBy"`
	UpdatedDate *time.Time `db:"UpdatedDate"`
	IsDeleted   *bool      `db:"IsDeleted"`
}
