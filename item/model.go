package item

type RelationalItem struct {
	ID   *int    `db:"id" fake:"{int8}"`
	UUID *string `db:"string" fake:"{uuid}"`
	Name *string `db:"name" fake:"{name}"`
}

func NewRelationalItem(item *Item) *RelationalItem {
	return &RelationalItem{
		UUID: item.UUID,
		Name: item.Name,
	}
}

func (r *RelationalItem) Equal(model interface{}) bool {
	other := model.(RelationalItem)
	return *r.UUID == *other.UUID && *r.Name == *other.Name
}

type Item struct {
	UUID *string `db:"string"`
	Name *string `db:"name"`
}

func (i *Item) Equal(model interface{}) bool {
	other := model.(Item)
	return *i.UUID == *other.UUID && *i.Name == *other.Name
}
