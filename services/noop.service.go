package services

import (
	"context"
	"github.com/duolacloud/crud-core/types"
)

type noopService [DTO any, CreateDTO any, UpdateDTO any] struct {
}

func NewNoopService[DTO any, CreateDTO any, UpdateDTO any]() CrudService[DTO, CreateDTO, UpdateDTO] {
	return &noopService[DTO, CreateDTO, UpdateDTO]{}
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Create(c context.Context, createDTO *CreateDTO, opts ...types.CreateOption) (*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) CreateMany(c context.Context, items []*CreateDTO, opts ...types.CreateOption) ([]*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID) error {
	return nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Get(c context.Context, id types.ID) (*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery) ([]*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) QueryOne(c context.Context, filter map[string]any) (*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Count(c context.Context, query *types.PageQuery) (int64, error) {
	return 0, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Aggregate(
	c context.Context,
	filter map[string]any,
	aggregateQuery *types.AggregateQuery,
) ([]*types.AggregateResponse, error) {
	return nil, nil
}