package mysql

import (
	"reflect-test/mysql/model"
	"reflect-test/std"
)

type itemRepository struct {
	std.Repository
	db *DB
}

func NewItemRepository(db *DB) *itemRepository {
	stdRepository := newStandardRepository("item", model.DBItem{}, "id", db)
	return &itemRepository{stdRepository, db}
}
