package repositories

import (
	"context"
	"github.com/duolacloud/crud-core/types"
)

type CrudRepository[DTO any, CreateDTO any, UpdateDTO any] interface {
	Create(c context.Context, createDTO *CreateDTO, opts ...types.CreateOption) (*DTO, error)
	CreateMany(c context.Context, items []*CreateDTO, opts ...types.CreateOption) ([]*DTO, error)
	Delete(c context.Context, id types.ID) error
	Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error)
	Get(c context.Context, id types.ID) (*DTO, error)
	Query(c context.Context, query *types.PageQuery) ([]*DTO, error)
	Count(c context.Context, query *types.PageQuery) (int64, error)
	QueryOne(c context.Context, filter map[string]any) (*DTO, error)
	Aggregate(
		c context.Context,
		filter map[string]any,
		aggregateQuery *types.AggregateQuery,
	) ([]*types.AggregateResponse, error)

	CursorQuery(c context.Context, query *types.CursorQuery) ([]*DTO, *types.CursorExtra, error)
}
