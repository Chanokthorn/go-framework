package model

import (
	"reflect-test/item"
	"reflect-test/std"
)

type DBItem struct {
	ID        *int    `db:"id" fake:"{int8}"`
	UUID      *string `db:"uuid" fake:"{uuid}"`
	Name      *string `db:"name" fake:"{name}"`
	CreatedBy *string `db:"createdBy" fake:"{name}"`
}

func (d *DBItem) Set(domain std.DomainModel) {
	n, _ := domain.(*item.RelationalItem)
	d.ID = n.ID
	d.UUID = n.UUID
	d.Name = n.Name
}

func (d DBItem) ToModel() std.DomainModel {
	return &item.RelationalItem{
		ID:   d.ID,
		UUID: d.UUID,
		Name: d.Name,
	}
}
