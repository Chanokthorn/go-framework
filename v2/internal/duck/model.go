package duck

type Duck struct {
	Name  *string `fake:"{name}"`
	Color *string `fake:"{color}"`
}
