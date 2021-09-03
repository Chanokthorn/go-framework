package duck

import (
	"context"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetByID(ctx context.Context, id int) (Duck, error) {
	return s.repository.GetByID(ctx, id)
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
	return s.repository.Create(ctx, dReq)
}

func (s *service) Update(ctx context.Context, dReq Duck) error {
	return s.repository.Update(ctx, dReq)
}

func (s *service) Delete(ctx context.Context, uuid string) error {
	return s.repository.Delete(ctx, uuid)
}
