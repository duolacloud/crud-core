package services

import (
	"context"

	"github.com/duolacloud/crud-core/repositories"
	"github.com/duolacloud/crud-core/types"
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

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Create(c context.Context, createDTO *CreateDTO, opts ...types.CreateOption) (*DTO, error) {
	return s.repo.Create(c, createDTO, opts...)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID) error {
	return s.repo.Delete(c, id)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	return s.repo.Update(c, id, updateDTO, opts...)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Get(c context.Context, id types.ID) (*DTO, error) {
	return s.repo.Get(c, id)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery) ([]*DTO, error) {
	return s.repo.Query(c, query)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Count(c context.Context, query *types.PageQuery) (int64, error) {
	return s.repo.Count(c, query)
}

func (s *crudService[DTO, CreateDTO, UpdateDTO]) Aggregate(
	c context.Context,
	filter map[string]interface{},
	aggregateQuery *types.AggregateQuery,
) ([]*types.AggregateResponse, error) {
	return s.repo.Aggregate(c, filter, aggregateQuery)
}