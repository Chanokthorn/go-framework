package duck

type Duck struct {
	ID             *int    `fake:"skip"`
	UUID           *string `fake:"skip"`
	Name           *string `fake:"{name}"`
	Color          *string `fake:"{color}"`
	Eggs           []Egg
	FavouriteMenus []FavouriteMenu
}

type Egg struct {
	Name *string `fake:"{name}"`
	Age  *int    `fake:"{number:1,50}"`
}

type FavouriteMenu struct {
	UUID *string `fake:"skip"`
	Name *string `fake:"{vegetable}"`
}
