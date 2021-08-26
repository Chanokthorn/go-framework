package mysql

import (
	"fmt"
	"reflect-test/v1/internal/mysql/model"
	"reflect-test/v1/internal/std/mysql"
)

type itemRepository struct {
	mysql.DomainRepository
	db *DB
}

func NewItemRepository(db *DB) (*itemRepository, error) {
	stdRepository, err := newStandardRepository(model.DBItem{}, db)
	if err != nil {
		return nil, fmt.Errorf(`unable to generate std repository: %v`, err)
	}

	return &itemRepository{stdRepository, db}, nil
}
