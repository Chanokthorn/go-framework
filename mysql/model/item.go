package model

import (
	"reflect-test/item"
	"reflect-test/std"
)

type DBItem struct {
	std.DBRootModel
	std.DBModelCommon
	Config    struct{} `std:"tableName:item,idField:id,uuidField:uuid"`
	ID        *int     `db:"id"`
	UUID      *string  `db:"uuid"`
	Name      *string  `db:"name"`
	Locations []DBLocation
	//CreatedBy   *string `db:"CreatedBy"`
	//UpdatedDate *string `db:"UpdatedDate"`
	//IsDeleted   *bool   `db:"IsDeleted"`
}

func (d DBItem) ToDomain() std.DomainModel {
	locations := []item.Location{}

	for _, l := range d.Locations {
		location := l.ToDomain().(*item.Location)
		locations = append(locations, *location)
	}

	return &item.RelationalItem{
		ID:        d.ID,
		UUID:      d.UUID,
		Name:      d.Name,
		Locations: locations,
	}
}

func (d *DBItem) Set(domain std.DomainModel) {
	n, _ := domain.(*item.RelationalItem)

	locations := []DBLocation{}

	for _, l := range n.Locations {
		var lDB DBLocation
		lDB.Set(l)

		locations = append(locations, lDB)
	}

	d.ID = n.ID
	d.UUID = n.UUID
	d.Name = n.Name
	d.Locations = locations
}

type DBLocation struct {
	std.DBModelCommon
	Config     struct{} `std:"tableName:item_location,idField:id,parentIDField:itemID"`
	ID         *int     `db:"id"`
	ItemID     *int     `db:"itemID"`
	Country    *string  `db:"country"`
	PostalCode *string  `db:"postalCode"`
}

func (d *DBLocation) ToDomain() std.DomainModel {
	return &item.Location{
		Country:    d.Country,
		PostalCode: d.PostalCode,
	}
}

func (d *DBLocation) Set(domain std.DomainModel) {
	n, _ := domain.(item.Location)
	d.Country = n.Country
	d.PostalCode = n.PostalCode
}
