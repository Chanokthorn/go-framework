package duck

import (
	"context"
)

type Repository interface {
	GetByIDs(ctx context.Context, ids []int) ([]Duck, error)
	GetByUUIDs(ctx context.Context, uuids []string) ([]Duck, error)
	Get(ctx context.Context, uuid string) (Duck, error)
	GetByID(ctx context.Context, id int) (Duck, error)
	GetAll(ctx context.Context) ([]Duck, error)
	Search(ctx context.Context, dReq Duck) ([]Duck, error)
	Create(ctx context.Context, dReq Duck) (id int, err error)
	Update(ctx context.Context, dReq Duck) error
	Delete(ctx context.Context, uuid string) error
}
