package mysql

import (
	"context"
	"fmt"
	"reflect-test/v2/internal/pond"
	stdMysql "reflect-test/v2/internal/std/mysql"
)

type pondService struct {
	stdMysql.Repository
}

func NewPondService(db *stdMysql.DB) (pond.Service, error) {
	stdRepository, err := stdMysql.NewRepository(pond.Pond{}, db)
	if err != nil {
		return nil, fmt.Errorf(`unable to instantiate std repository: %v`, err)
	}
	return &pondService{stdRepository}, nil
}

func (ps *pondService) GetAll(ctx context.Context) ([]pond.Pond, error) {
	ponds := []pond.Pond{}

	err := ps.Repository.GetAll(ctx, &ponds)
	if err != nil {
		return nil, fmt.Errorf(`unable to get all ponds: %v`, err)
	}

	return ponds, nil
}

func (ps *pondService) Get(ctx context.Context, uuid string) (pond.Pond, error) {
	var p pond.Pond

	err := ps.Repository.GetByUUID(ctx, &p, uuid)
	if err != nil {
		return pond.Pond{}, fmt.Errorf(`unable to get pond by uuid: %v`, err)
	}

	return p, nil
}

func (ps *pondService) Create(ctx context.Context, pondReq pond.Pond) (int, error) {
	id, err := ps.Repository.Insert(ctx, &pondReq)
	if err != nil {
		return 0, fmt.Errorf(`unable to insert: %v`, err)
	}

	return id, nil
}
