package repositories

import (
	"context"
	"fmt"
	"sync"

	core_cache "github.com/duolacloud/crud-core/cache"
	"github.com/duolacloud/crud-core/types"
)

type cacheRepository[DTO any, CreateDTO any, UpdateDTO any] struct {
	CrudRepository[DTO, CreateDTO, UpdateDTO]
	cache core_cache.Cache
	mutex sync.Mutex
}

func NewCacheRepository[DTO any, CreateDTO any, UpdateDTO any](
	repository CrudRepository[DTO, CreateDTO, UpdateDTO],
	cache core_cache.Cache,
) CrudRepository[DTO, CreateDTO, UpdateDTO] {
	return &cacheRepository[DTO, CreateDTO, UpdateDTO]{
		CrudRepository: repository,
		cache:          cache,
	}
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) Create(c context.Context, createDTO *CreateDTO, opts ...types.CreateOption) (*DTO, error) {
	return r.CrudRepository.Create(c, createDTO, opts...)
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) CreateMany(c context.Context, items []*CreateDTO, opts ...types.CreateManyOption) ([]*DTO, error) {
	return r.CrudRepository.CreateMany(c, items, opts...)
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if err := r.cache.Delete(c, fmt.Sprintf("%v", id)); err != nil {
		return err
	}
	return r.CrudRepository.Delete(c, id)
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if err := r.cache.Delete(c, fmt.Sprintf("%v", id)); err != nil {
		return nil, err
	}
	return r.CrudRepository.Update(c, id, updateDTO, opts...)
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) Get(c context.Context, id types.ID) (*DTO, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	cacheKey := fmt.Sprintf("%v", id)
	dto := new(DTO)
	err := r.cache.Get(c, cacheKey, dto)
	if err != nil && err != core_cache.ErrNotExsit {
		// 缓存查询错误
		return nil, err
	}
	if err == nil {
		// 命中缓存，直接返回
		return dto, nil
	}
	// 未命中缓存
	dto, err = r.CrudRepository.Get(c, id)
	if err != nil {
		return nil, err
	}
	_ = r.cache.Set(c, cacheKey, dto)
	return dto, nil
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery) ([]*DTO, error) {
	return r.CrudRepository.Query(c, query)
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) Count(c context.Context, query *types.PageQuery) (int64, error) {
	return r.CrudRepository.Count(c, query)
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) QueryOne(c context.Context, filter map[string]any) (*DTO, error) {
	return r.CrudRepository.QueryOne(c, filter)
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) Aggregate(
	c context.Context,
	filter map[string]any,
	aggregateQuery *types.AggregateQuery,
) ([]*types.AggregateResponse, error) {
	return r.CrudRepository.Aggregate(c, filter, aggregateQuery)
}

func (r *cacheRepository[DTO, CreateDTO, UpdateDTO]) CursorQuery(c context.Context, query *types.CursorQuery) ([]*DTO, *types.CursorExtra, error) {
	return r.CrudRepository.CursorQuery(c, query)
}
