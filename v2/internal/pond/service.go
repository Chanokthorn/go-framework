package pond

import "context"

type Service interface {
	GetAll(ctx context.Context) ([]Pond, error)
	Get(ctx context.Context, uuid string) (Pond, error)
	Create(ctx context.Context, pondReq Pond) (id int, err error)
}
