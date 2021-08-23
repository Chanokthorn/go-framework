package model

import (
	"reflect-test/item"
	"reflect-test/std"
)

type DBItem struct {
	Config      struct{} `std:"tableName:item,idField:id,uuidField:uuid"`
	ID          *int     `db:"id"`
	UUID        *string  `db:"uuid"`
	Name        *string  `db:"name"`
	Locations   []DBLocation
	CreatedBy   *string `db:"CreatedBy"`
	UpdatedDate *string `db:"UpdatedDate"`
	IsDeleted   *bool   `db:"IsDeleted"`
}

func (d DBItem) ToDomain() std.DomainModel {
	return &item.RelationalItem{
		ID:   d.ID,
		UUID: d.UUID,
		Name: d.Name,
	}
}

func (d *DBItem) Set(domain std.DomainModel) {
	n, _ := domain.(*item.RelationalItem)
	d.ID = n.ID
	d.UUID = n.UUID
	d.Name = n.Name
}

type DBLocation struct {
	Config      struct{} `std:"tableName:locations,idField:id,parentIDField:itemID"`
	ID          *int     `db:"id"`
	ItemID      *int     `db:"itemID"`
	Country     *string  `db:"country"`
	PostalCode  *string  `db:"postalCode"`
	CreatedBy   *string  `db:"CreatedBy"`
	UpdatedDate *string  `db:"UpdatedDate"`
	IsDeleted   *bool    `db:"IsDeleted"`
}

func (d *DBLocation) ToDomain() std.DomainModel {
	return &item.Location{
		Country:    d.Country,
		PostalCode: d.PostalCode,
	}
}

func (d *DBLocation) Set(rootID int, domain std.DomainModel) {
	n, _ := domain.(*item.Location)
	d.Country = n.Country
	d.PostalCode = n.PostalCode
}
