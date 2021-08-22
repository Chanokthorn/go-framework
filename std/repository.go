package std

type Repository interface {
	Get(id int) (DomainModel, error)
	GetAll() ([]DomainModel, error)
	Insert(domain DomainModel) (id int, err error)
}
