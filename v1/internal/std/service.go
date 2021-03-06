package std

type Service interface {
	GetAll() ([]DomainModel, error)
	GetByID(id int) (DomainModel, error)
	Search(domain DomainModel) ([]DomainModel, error)
	GetByUUID(uuid string) (DomainModel, error)
	Create(domain DomainModel) (int, error)
	Update(domain DomainModel) error
	Delete(uuid string) error
}
