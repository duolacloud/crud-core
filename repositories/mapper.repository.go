package repositories

import (
	"context"

	"duolacloud.com/duolacloud/crud-core/mappers"
	"duolacloud.com/duolacloud/crud-core/types"
)

type mapperRepository [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] struct {
	repo   CrudRepository[Entity, CreateEntity, UpdateEntity]
	mapper mappers.Mapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]
}

func NewMapperRepository[DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any](
	repo CrudRepository[Entity, CreateEntity, UpdateEntity],
	mapper mappers.Mapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity],
) CrudRepository[DTO, CreateDTO, UpdateDTO] {
	return &mapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]{
		repo:   repo,
		mapper: mapper,
	}
}

func (r *mapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Create(c context.Context, createDTO *CreateDTO) (*DTO, error) {
	createEntity := r.mapper.ConvertToCreateEntity(c, createDTO)

	entity, err := r.repo.Create(c, createEntity)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTO(c, entity), nil
}

func (r *mapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Delete(c context.Context, id types.ID) error {
	return r.repo.Delete(c, id)
}

func (r *mapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO) (*DTO, error) {
	entity := r.mapper.ConvertToUpdateEntity(c, updateDTO)

	updatedEntity, err := r.repo.Update(c, id, entity)
	if err != nil {
		return nil, err
	}
	return r.mapper.ConvertToDTO(c, updatedEntity), nil
}

func (r *mapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Get(c context.Context, id types.ID) (*DTO, error) {
	entity, err := r.repo.Get(c, id)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTO(c, entity), nil
}

func (r *mapperRepository[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) Query(c context.Context, query *types.PageQuery[DTO]) ([]*DTO, error) {
	entityQuery := r.mapper.ConvertQuery(c, query)

	entities, err := r.repo.Query(c, entityQuery)
	if err != nil {
		return nil, err
	}

	return r.mapper.ConvertToDTOs(c, entities), nil
}