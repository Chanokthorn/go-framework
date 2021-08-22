package item

import "reflect-test/std"

type Service interface {
	std.Service
}

type service struct {
	repository RelationalRepository
}

func NewService(repository RelationalRepository) std.Service {
	return &service{repository: repository}
}

func (s *service) GetAll() ([]std.DomainModel, error) {
	return s.repository.GetAll()
}

func (s *service) GetByID(id int) (std.DomainModel, error) {
	return s.repository.GetByID(id)
}

func (s *service) GetByUUID(uuid string) (std.DomainModel, error) {
	panic("implement me")
}

func (s *service) Create(domain std.DomainModel) (int, error) {
	return s.repository.Insert(domain)
}

func (s *service) Update(domain std.DomainModel) error {
	item := domain.(*Item)
	rel := NewRelationalItem(item)
	return s.repository.Update(rel)
}

func (s *service) Delete(uuid string) error {
	return s.repository.Delete(uuid)
}
