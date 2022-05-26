package services

import (
	"context"

	"github.com/duolacloud/crud-core/types"
)

type CrudService [DTO any, CreateDTO any, UpdateDTO any] interface {
	Create(c context.Context, dto *CreateDTO, opts ...types.CreateOption) (*DTO, error)
	Delete(c context.Context, id types.ID) error
	Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error)
	Get(c context.Context, id types.ID) (*DTO, error)
	Query(c context.Context, query *types.PageQuery) ([]*DTO, error)
	QueryOne(c context.Context, filter map[string]any) (*DTO, error)
	Count(c context.Context, query *types.PageQuery) (int64, error)
	Aggregate(
		c context.Context,
		filter map[string]any,
		aggregateQuery *types.AggregateQuery,
	) ([]*types.AggregateResponse, error)
}
