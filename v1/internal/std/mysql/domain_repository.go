package mysql

import (
	"reflect-test/v1/internal/std"
)

type DomainRepository interface {
	GetByID(id int) (std.DomainModel, error)
	GetAll() ([]std.DomainModel, error)
	Search(domain std.DomainModel) ([]std.DomainModel, error)
	Insert(domain std.DomainModel) (id int, err error)
	Update(domain std.DomainModel) error
	Delete(uuid string) error
}
