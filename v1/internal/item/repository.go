package item

import (
	"reflect-test/v1/internal/std/mysql"
)

type RelationalRepository interface {
	mysql.DomainRepository
}
