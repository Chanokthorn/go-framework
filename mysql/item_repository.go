package mysql

import (
	"fmt"
	"reflect-test/mysql/model"
	"reflect-test/std"
)

type itemRepository struct {
	std.Repository
	db *DB
}

func NewItemRepository(db *DB) (*itemRepository, error) {
	stdRepository, err := newStandardRepository(model.DBItem{}, db)
	if err != nil {
		return nil, fmt.Errorf(`unable to generate std repository: %v`, err)
	}

	return &itemRepository{stdRepository, db}, nil
}
