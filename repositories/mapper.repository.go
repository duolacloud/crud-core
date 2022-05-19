package repositories

import (
	"context"

	"duolacloud.com/duolacloud/crud-core/mappers"
	"duolacloud.com/duolacloud/crud-core/types"
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
	createEntity := r.mapper.ConvertToCreateEntity(c, createDTO)

	entity, err := r.repo.Create(c, createEntity)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTO(c, entity), nil
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Delete(c context.Context, id types.ID) error {
	return r.repo.Delete(c, id)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO, opts ...types.UpdateOption) (*DTO, error) {
	entity := r.mapper.ConvertToUpdateEntity(c, updateDTO)

	updatedEntity, err := r.repo.Update(c, id, entity)
	if err != nil {
		return nil, err
	}
	return r.mapper.ConvertToDTO(c, updatedEntity), nil
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Get(c context.Context, id types.ID) (*DTO, error) {
	entity, err := r.repo.Get(c, id)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTO(c, entity), nil
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Query(c context.Context, query *types.PageQuery[DTO]) ([]*DTO, error) {
	entityQuery := r.mapper.ConvertQuery(c, query)

	entities, err := r.repo.Query(c, entityQuery)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTOs(c, entities), nil
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Count(c context.Context, query *types.PageQuery[DTO]) (int64, error) {
	entityQuery := r.mapper.ConvertQuery(c, query)

	return r.repo.Count(c, entityQuery)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Aggregate(
	c context.Context,
	filter map[string]interface{},
	aggregateQuery *types.AggregateQuery,
) ([]*types.AggregateResponse, error) {
	return r.repo.Aggregate(c, filter, aggregateQuery)
}

func (r *MapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) CursorQuery(c context.Context, query *types.CursorQuery) ([]*DTO, *types.CursorExtra, error) {
	return r.repo.CursorQuery(c, query)
}