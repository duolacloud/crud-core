package mappers

import (
	"context"
	"github.com/duolacloud/crud-core/types"
	"github.com/mitchellh/mapstructure"
)

type MapStructureMapper [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] struct {
	AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]
}

func NewMapStructureMapper[DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any]() *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity] {
	m := &MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]{	
	}

	m.fnConvertToDTO = m.ConvertToDTO
	m.fnConvertToEntity = m.ConvertToEntity 

	return m
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToDTO(c context.Context, entity *Entity) *DTO {
	var dto *DTO
	_ = mapstructure.Decode(entity, &dto)
	return dto
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToEntity(c context.Context, dto *DTO) *Entity {
	var entity *Entity
	_ = mapstructure.Decode(dto, &entity)
	return entity
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToCreateEntity(c context.Context, createDTO *CreateDTO) *CreateEntity {
	var createEntity *CreateEntity
	_ = mapstructure.Decode(createDTO, &createEntity)
	return createEntity
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToUpdateEntity(c context.Context, updateDTO *UpdateDTO) *UpdateEntity {
	var updateEntity *UpdateEntity
	_ = mapstructure.Decode(updateDTO, &updateEntity)
	return updateEntity
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertQuery(c context.Context, query *types.PageQuery) *types.PageQuery {
	return query
}