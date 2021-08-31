package pond

import "context"

type Service interface {
	GetAll(ctx context.Context) ([]Pond, error)
	Get(ctx context.Context, uuid string) (Pond, error)
	FillStructsByUUID(ctx context.Context, ponds *[]Pond) error
	FillStructsByID(ctx context.Context, ponds *[]Pond) error
	Create(ctx context.Context, pondReq Pond) (id int, err error)
}
