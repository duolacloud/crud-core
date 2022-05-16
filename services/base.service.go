package services

import (
	"context"

	"duolacloud.com/duolacloud/crud-core/repositories"
	"duolacloud.com/duolacloud/crud-core/types"
)

type crudService [DTO any, CreateDTO any, UpdateDTO any] struct {
	repo repositories.CrudRepository[DTO, CreateDTO, UpdateDTO]
}

func NewCrudService[DTO any, CreateDTO any, UpdateDTO any](
	repo repositories.CrudRepository[DTO, CreateDTO, UpdateDTO],
) CrudService[DTO, CreateDTO, UpdateDTO] {
	return &crudService[DTO, CreateDTO, UpdateDTO]{
		repo: repo,
	}
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Create(c context.Context, createDTO *CreateDTO) (*DTO, error) {
	return s.repo.Create(c, createDTO)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID) error {
	return s.repo.Delete(c, id)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO) (*DTO, error) {
	return s.repo.Update(c, id, updateDTO)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Get(c context.Context, id types.ID) (*DTO, error) {
	return s.repo.Get(c, id)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery[DTO]) ([]*DTO, error) {
	return s.repo.Query(c, query)
}