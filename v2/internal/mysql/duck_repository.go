package mysql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"reflect-test/v2/internal/duck"
	"reflect-test/v2/internal/mysql/model"
	stdMysql "reflect-test/v2/internal/std/mysql"
)

type duckRepository struct {
	stdMysql.Repository
}

func NewDuckRepository(db *sqlx.DB) (duck.Repository, error) {
	stdRepository, err := stdMysql.NewRepository(model.Duck{}, db)
	if err != nil {
		return nil, fmt.Errorf(`unable to instantiate std repository: %v`, err)
	}
	return &duckRepository{Repository: stdRepository}, nil
}

func (dr *duckRepository) GetByIDs(ctx context.Context, ids []int) ([]duck.Duck, error) {
	ds := []model.Duck{}

	err := dr.Repository.GetByIDs(ctx, &ds, ids)
	if err != nil {
		return nil, fmt.Errorf(`unable to get by ids %v`, err)
	}

	dRes := []duck.Duck{}

	for _, d := range ds {
		dRes = append(dRes, d.ToModel())
	}

	return dRes, nil
}

func (dr *duckRepository) GetByUUIDs(ctx context.Context, uuids []string) ([]duck.Duck, error) {
	ds := []model.Duck{}

	err := dr.Repository.GetByUUIDs(ctx, &ds, uuids)
	if err != nil {
		return nil, fmt.Errorf(`unable to get by uuids %v`, err)
	}

	dRes := []duck.Duck{}

	for _, d := range ds {
		dRes = append(dRes, d.ToModel())
	}

	return dRes, nil
}

func (dr *duckRepository) Get(ctx context.Context, uuid string) (duck.Duck, error) {
	var d model.Duck

	err := dr.Repository.GetByUUID(ctx, &d, uuid)
	if err != nil {
		return duck.Duck{}, fmt.Errorf(`unable to get by uuid: %v`, err)
	}

	return d.ToModel(), nil
}

func (dr *duckRepository) GetByID(ctx context.Context, id int) (duck.Duck, error) {
	var d model.Duck

	err := dr.Repository.GetByID(ctx, &d, id)
	if err != nil {
		return duck.Duck{}, fmt.Errorf(`unable to get by id: %v`, err)
	}

	return d.ToModel(), nil
}

func (dr *duckRepository) GetAll(ctx context.Context) ([]duck.Duck, error) {
	var ds []model.Duck

	err := dr.Repository.GetAll(ctx, &ds)
	if err != nil {
		return nil, fmt.Errorf(`unable to get all: %v`, err)
	}

	dRes := []duck.Duck{}

	for _, d := range ds {
		dRes = append(dRes, d.ToModel())
	}

	return dRes, nil
}

func (dr *duckRepository) Search(ctx context.Context, dReq duck.Duck) ([]duck.Duck, error) {
	var ds []model.Duck

	err := dr.Repository.Search(ctx, &ds, model.NewDuck(dReq))
	if err != nil {
		return nil, fmt.Errorf(`unable to search: %v`, err)
	}

	dRes := []duck.Duck{}

	for _, d := range ds {
		dRes = append(dRes, d.ToModel())
	}

	return dRes, nil
}

func (dr *duckRepository) Create(ctx context.Context, dReq duck.Duck) (id int, err error) {
	id, err = dr.Repository.Insert(ctx, model.NewDuck(dReq))
	if err != nil {
		return 0, fmt.Errorf(`unable to insert: %v`, err)
	}

	return id, nil
}

func (dr *duckRepository) Update(ctx context.Context, dReq duck.Duck) error {
	err := dr.Repository.Update(ctx, model.NewDuck(dReq))
	if err != nil {
		return fmt.Errorf(`unable to update: %v`, err)
	}

	return nil
}
