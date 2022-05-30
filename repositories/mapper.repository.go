package repositories

import (
	"context"

	"github.com/duolacloud/crud-core/mappers"
	"github.com/duolacloud/crud-core/types"
)

type MapperRepository [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] struct {
	repo   CrudRepository[Entity, CreateEntity, UpdateEntity]
	mapper mappers.Mapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]
}

func NewMapperRepository[DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any](
	repo CrudRepository[Entity, CreateEntity, UpdateEntity],
	mapper mappers.Mapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity],
) *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity] {
	return &MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]{
		repo:   repo,
		mapper: mapper,
	}
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Create(c context.Context, createDTO *CreateDTO, opts ...types.CreateOption) (*DTO, error) {
	createEntity, err := r.mapper.ConvertToCreateEntity(c, createDTO)
	if err != nil {
		return nil, err
	}

	entity, err := r.repo.Create(c, createEntity)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTO(c, entity)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) CreateMany(c context.Context, items []*CreateDTO, opts ...types.CreateOption) ([]*DTO, error) {
	converted, err := r.mapper.ConvertToCreateEntities(c, items)
	if err != nil {
		return nil, err
	}

	entities, err := r.repo.CreateMany(c, converted)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTOs(c, entities)
}


func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Delete(c context.Context, id types.ID) error {
	return r.repo.Delete(c, id)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	entity, err := r.mapper.ConvertToUpdateEntity(c, updateDTO)
	if err != nil {
		return nil, err
	}

	updatedEntity, err := r.repo.Update(c, id, entity)
	if err != nil {
		return nil, err
	}
	return r.mapper.ConvertToDTO(c, updatedEntity)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Get(c context.Context, id types.ID) (*DTO, error) {
	entity, err := r.repo.Get(c, id)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTO(c, entity)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Query(c context.Context, query *types.PageQuery) ([]*DTO, error) {
	entityQuery, err := r.mapper.ConvertQuery(c, query)
	if err != nil {
		return nil, err
	}

	entities, err := r.repo.Query(c, entityQuery)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTOs(c, entities)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Count(c context.Context, query *types.PageQuery) (int64, error) {
	entityQuery, err := r.mapper.ConvertQuery(c, query)
	if err != nil {
		return 0, err
	}

	return r.repo.Count(c, entityQuery)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) QueryOne(c context.Context, filter map[string]any) (*DTO, error) {
	entityQuery, err := r.mapper.ConvertQuery(c, &types.PageQuery{ Filter: filter })
	if err != nil {
		return nil, err
	}

	entity, err := r.repo.QueryOne(c, entityQuery.Filter)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTO(c, entity)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Aggregate(
	c context.Context,
	filter map[string]any,
	aggregateQuery *types.AggregateQuery,
) ([]*types.AggregateResponse, error) {
	return r.repo.Aggregate(c, filter, aggregateQuery)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) CursorQuery(c context.Context, query *types.CursorQuery) ([]*DTO, *types.CursorExtra, error) {
	entities, extra, err := r.repo.CursorQuery(c, query)
	if err != nil {
		return nil, nil, err
	}

	dtos, err := r.mapper.ConvertToDTOs(c, entities)
	if err != nil {
		return nil, nil, err
	}

	return dtos, extra, nil
}