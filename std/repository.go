package std

type Repository interface {
	GetByID(id int) (DomainModel, error)
	GetAll() ([]DomainModel, error)
	Insert(domain DomainModel) (id int, err error)
	Update(domain DomainModel) error
	Delete(id int) error
}
