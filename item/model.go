package item

import "reflect-test/std"

type RelationalItem struct {
	std.DomainModel
	ID        *int       `fake:"skip"`
	UUID      *string    `fake:"{uuid}"`
	Name      *string    `fake:"{name}"`
	Locations []Location `fakesize:"3"`
}

func NewRelationalItem(item *Item) *RelationalItem {
	return &RelationalItem{
		UUID: item.UUID,
		Name: item.Name,
	}
}

type Item struct {
	std.DomainModel
	UUID      *string
	Name      *string
	Locations []Location
}

type Location struct {
	std.DomainModel
	Country    *string `fake:"{country}"`
	PostalCode *string `fake:"{randomstring:[13212,13142,14283,12058,19640]}"`
}
