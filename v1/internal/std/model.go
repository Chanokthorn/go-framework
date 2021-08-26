package std

type DBModel interface {
	ToDomain() DomainModel
}

type DBRootModel interface {
	DBModel
	Set(domain DomainModel)
}

type DBAggregateModel interface {
	DBModel
	Set(domain DomainModel)
}

type DomainModel interface {
	Domain()
}
