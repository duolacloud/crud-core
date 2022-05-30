package services

import (
	"context"
	"github.com/duolacloud/crud-core/types"
	"github.com/duolacloud/crud-core/cache"
)

type cacheService [DTO any, CreateDTO any, UpdateDTO any] struct {
	service CrudService[DTO, CreateDTO, UpdateDTO]
	cache cache.Cache
}

func NewCacheService[DTO any, CreateDTO any, UpdateDTO any](
	service CrudService[DTO, CreateDTO, UpdateDTO],
	cache cache.Cache,
) CrudService[DTO, CreateDTO, UpdateDTO] {
	return &cacheService[DTO, CreateDTO, UpdateDTO]{
		service: service,
		cache: cache,
	}
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) Create(c context.Context, createDTO *CreateDTO, opts ...types.CreateOption) (*DTO, error) {
	return s.service.Create(c, createDTO, opts...)
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) CreateMany(c context.Context, items []*CreateDTO, opts ...types.CreateManyOption) ([]*DTO, error) {
	return s.service.CreateMany(c, items, opts...)
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID) error {
	return s.service.Delete(c, id)
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	return s.service.Update(c, id, updateDTO, opts...)
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) Get(c context.Context, id types.ID) (*DTO, error) {
	return s.service.Get(c, id)
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery) ([]*DTO, error) {
	return s.service.Query(c, query)
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) Count(c context.Context, query *types.PageQuery) (int64, error) {
	return s.service.Count(c, query)
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) QueryOne(c context.Context, filter map[string]any) (*DTO, error) {
	return s.service.QueryOne(c, filter)
}

func (s *cacheService[DTO, CreateDTO, UpdateDTO]) Aggregate(
	c context.Context,
	filter map[string]any,
	aggregateQuery *types.AggregateQuery,
) ([]*types.AggregateResponse, error) {
	return s.service.Aggregate(c, filter, aggregateQuery)
}