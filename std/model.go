package std

type DomainModel interface {
	Equal(model interface{}) bool
}

type DBModel interface {
	Set(domain DomainModel)
	ToModel() DomainModel
}
