package duck

import (
	"context"
	"fmt"
	"reflect-test/v2/internal/pond"
)

type service struct {
	repository  Repository
	pondService pond.Service
}

func NewService(repository Repository, pondService pond.Service) *service {
	return &service{repository: repository, pondService: pondService}
}

func (s *service) GetByID(ctx context.Context, id int) (Duck, error) {
	d, err := s.repository.GetByID(ctx, id)
	if err != nil {
		return Duck{}, fmt.Errorf(`unable to get by id: %v`, err)
	}

	err = s.pondService.FillStructsByID(ctx, &d.Ponds)
	if err != nil {
		return Duck{}, fmt.Errorf(`unable to fill pond structs by id: %v`, err)
	}

	return d, nil
}

func (s *service) GetAll(ctx context.Context) ([]Duck, error) {
	return s.repository.GetAll(ctx)
}

func (s *service) Search(ctx context.Context, dReq Duck) ([]Duck, error) {
	return s.repository.Search(ctx, dReq)
}

func (s *service) GetByUUID(ctx context.Context, uuid string) (Duck, error) {
	return s.repository.Get(ctx, uuid)
}

func (s *service) Create(ctx context.Context, dReq Duck) (int, error) {
	err := s.pondService.FillStructsByUUID(ctx, &dReq.Ponds)
	if err != nil {
		return 0, fmt.Errorf(`unable to fill pond structs by uuid: %v`, err)
	}

	return s.repository.Create(ctx, dReq)
}

func (s *service) Update(ctx context.Context, dReq Duck) error {
	err := s.pondService.FillStructsByUUID(ctx, &dReq.Ponds)
	if err != nil {
		return fmt.Errorf(`unable to fill pond structs by uuid: %v`, err)
	}

	return s.repository.Update(ctx, dReq)
}

func (s *service) Delete(ctx context.Context, uuid string) error {
	return s.repository.Delete(ctx, uuid)
}
