package std

import "context"

type Service interface {
	GetAll(ctx context.Context) ([]DomainModel, error)
	GetByID(ctx context.Context, id int) (DomainModel, error)
	Search(ctx context.Context, domain DomainModel) ([]DomainModel, error)
	GetByUUID(ctx context.Context, uuid string) (DomainModel, error)
	Create(ctx context.Context, domain DomainModel) (int, error)
	Update(ctx context.Context, domain DomainModel) error
	Delete(ctx context.Context, uuid string) error
}
