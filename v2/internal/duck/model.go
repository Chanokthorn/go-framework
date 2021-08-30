package duck

type Duck struct {
	DuckID   *int    `fake:"skip"`
	DuckUUID *string `fake:"skip"`
	Name     *string `fake:"{name}"`
	Color    *string `fake:"{color}"`
	Eggs     []Egg
}

type Egg struct {
	Name *string `fake:"{name}"`
	Age  *int    `fake:"{number:1,50}"`
}
