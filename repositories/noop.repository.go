package repositories

import (
	"context"
	"github.com/duolacloud/crud-core/types"
)

type noopRepository[DTO any, CreateDTO any, UpdateDTO any] struct {
}

func NewNoopRepository[DTO any, CreateDTO any, UpdateDTO any]() CrudRepository[DTO, CreateDTO, UpdateDTO] {
	return &noopRepository[DTO, CreateDTO, UpdateDTO]{}
}

func(r *noopRepository[DTO, CreateDTO, UpdateDTO])	Create(c context.Context, createDTO *CreateDTO, opts ...types.CreateOption) (*DTO, error) {
	return nil, nil
}

func (r *noopRepository[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID) error {
	return nil
}

func (r *noopRepository[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	return nil, nil
}

func (r *noopRepository[DTO, CreateDTO, UpdateDTO])	Get(c context.Context, id types.ID) (*DTO, error) {
	return nil, nil
}

func (r *noopRepository[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery) ([]*DTO, error) {
	return nil, nil
}

func (r *noopRepository[DTO, CreateDTO, UpdateDTO]) Count(c context.Context, query *types.PageQuery) (int64, error) {
	return 0, nil
}

func (r *noopRepository[DTO, CreateDTO, UpdateDTO]) Aggregate(
		c context.Context,
		filter map[string]interface{},
		aggregateQuery *types.AggregateQuery,
	) ([]*types.AggregateResponse, error) {
		return nil, nil
}

func (r *noopRepository[DTO, CreateDTO, UpdateDTO]) CursorQuery(c context.Context, query *types.CursorQuery) ([]*DTO, *types.CursorExtra, error) {
	return nil, nil, nil
}