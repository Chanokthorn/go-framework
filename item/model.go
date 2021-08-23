package item

type RelationalItem struct {
	ID        *int
	UUID      *string
	Name      *string
	Locations []Location
}

func NewRelationalItem(item *Item) *RelationalItem {
	return &RelationalItem{
		UUID: item.UUID,
		Name: item.Name,
	}
}

func (r *RelationalItem) Domain() {}

type Item struct {
	UUID      *string
	Name      *string
	Locations []Location
}

func (i *Item) Domain() {}

type Location struct {
	Country    *string
	PostalCode *string
}

func (l *Location) Domain() {}
