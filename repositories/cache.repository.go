package repositories

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	core_cache "github.com/duolacloud/crud-core/cache"
	"github.com/duolacloud/crud-core/types"
)

type CacheRepositoryOptions struct {
	Expiration time.Duration
}

type CacheRepositoryOption func(*CacheRepositoryOptions)

func WithExpiration(expiration time.Duration) CacheRepositoryOption {
	return func(cro *CacheRepositoryOptions) {
		cro.Expiration = expiration
	}
}

type CacheRepository[DTO any, CreateDTO any, UpdateDTO any] struct {
	CrudRepository[DTO, CreateDTO, UpdateDTO]
	cache     core_cache.Cache
	mutex     sync.Mutex
	options   *CacheRepositoryOptions
	modelName string
}

func NewCacheRepository[DTO any, CreateDTO any, UpdateDTO any](
	repository CrudRepository[DTO, CreateDTO, UpdateDTO],
	cache core_cache.Cache,
	opts ...CacheRepositoryOption,
) CrudRepository[DTO, CreateDTO, UpdateDTO] {
	options := &CacheRepositoryOptions{}
	for _, opt := range opts {
		opt(options)
	}

	var dto DTO
	modelName := strings.ToLower(reflect.TypeOf(dto).Name())

	repo := &CacheRepository[DTO, CreateDTO, UpdateDTO]{
		CrudRepository: repository,
		cache:          cache,
		options:        options,
		modelName:      modelName,
	}

	return repo
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) Create(c context.Context, createDTO *CreateDTO, opts ...types.CreateOption) (*DTO, error) {
	return r.CrudRepository.Create(c, createDTO, opts...)
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) CreateMany(c context.Context, items []*CreateDTO, opts ...types.CreateManyOption) ([]*DTO, error) {
	return r.CrudRepository.CreateMany(c, items, opts...)
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID, opts ...types.DeleteOption) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if err := r.CrudRepository.Delete(c, id, opts...); err != nil {
		return err
	}

	if err := r.cache.Delete(c, r.formatCacheKey(id)); err != nil {
		return err
	}
	return nil
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	if err := r.cache.Delete(c, r.formatCacheKey(id)); err != nil {
		return nil, err
	}
	return r.CrudRepository.Update(c, id, updateDTO, opts...)
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) Get(c context.Context, id types.ID, opts ...types.GetOption) (*DTO, error) {
	// 查缓存用双重检查锁
	cacheKey := r.formatCacheKey(id)
	dto := new(DTO)

	err := r.cache.Get(c, cacheKey, dto)
	if err != nil && !errors.Is(err, types.ErrNotFound) {
		// 缓存查询错误
		return nil, err
	}
	if err == nil {
		// 命中缓存，直接返回
		return dto, nil
	}

	r.mutex.Lock()
	defer r.mutex.Unlock()

	err = r.cache.Get(c, cacheKey, dto)
	if err != nil && !errors.Is(err, types.ErrNotFound) {
		// 缓存查询错误
		return nil, err
	}
	if err == nil {
		// 命中缓存，直接返回
		return dto, nil
	}
	// 未命中缓存
	dto, err = r.CrudRepository.Get(c, id, opts...)
	if err != nil {
		return nil, err
	}
	if dto != nil {
		opts := make([]core_cache.SetOption, 0)
		if r.options.Expiration.Seconds() > 0 {
			opts = append(opts, core_cache.WithExpiration(r.options.Expiration))
		}
		_ = r.cache.Set(c, cacheKey, dto, opts...)
	}
	return dto, nil
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery) ([]*DTO, error) {
	return r.CrudRepository.Query(c, query)
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) Count(c context.Context, query *types.PageQuery) (int64, error) {
	return r.CrudRepository.Count(c, query)
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) QueryOne(c context.Context, filter map[string]any) (*DTO, error) {
	return r.CrudRepository.QueryOne(c, filter)
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) Aggregate(
	c context.Context,
	filter map[string]any,
	aggregateQuery *types.AggregateQuery,
) ([]*types.AggregateResponse, error) {
	return r.CrudRepository.Aggregate(c, filter, aggregateQuery)
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) CursorQuery(c context.Context, query *types.CursorQuery) ([]*DTO, *types.CursorExtra, error) {
	return r.CrudRepository.CursorQuery(c, query)
}

func (r *CacheRepository[DTO, CreateDTO, UpdateDTO]) formatCacheKey(key any) string {
	return types.FormatID(fmt.Sprintf("%s.%v", r.modelName, key))
}
