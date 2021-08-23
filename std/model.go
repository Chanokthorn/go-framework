package std

type DomainModel interface {
	Domain()
}

type DBModel interface {
	ToDomain() DomainModel
}

type DBRootModel interface {
	DBModel
	Set(domain DomainModel)
}

type DBAggregateModel interface {
	DBModel
	Set(rootID int, domain DomainModel)
}
