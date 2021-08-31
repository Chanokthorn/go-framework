package duck

import "reflect-test/v2/internal/pond"

type Duck struct {
	DuckID   *int    `fake:"skip"`
	DuckUUID *string `fake:"skip"`
	Name     *string `fake:"{name}"`
	Color    *string `fake:"{color}"`
	Eggs     []Egg
	Ponds    []pond.Pond `fake:"skip"`
}

type Egg struct {
	Name *string `fake:"{name}"`
	Age  *int    `fake:"{number:1,50}"`
}
