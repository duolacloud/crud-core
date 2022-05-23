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

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToDTO(c context.Context, entity *Entity) (*DTO, error) {
	var dto *DTO
	err := mapstructure.Decode(entity, &dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToEntity(c context.Context, dto *DTO) (*Entity, error) {
	var entity *Entity
	err := mapstructure.Decode(dto, &entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToCreateEntity(c context.Context, createDTO *CreateDTO) (*CreateEntity, error) {
	var createEntity *CreateEntity
	err := mapstructure.Decode(createDTO, &createEntity)
	if err != nil {
		return nil, err
	}

	return createEntity, nil
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToUpdateEntity(c context.Context, updateDTO *UpdateDTO) (*UpdateEntity, error) {
	var updateEntity *UpdateEntity
	err := mapstructure.Decode(updateDTO, &updateEntity)
	if err != nil {
		return nil, err
	}

	return updateEntity, nil
}

func (m *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertQuery(c context.Context, query *types.PageQuery) (*types.PageQuery, error) {
	return query, nil
}