package mappers

import (
	"context"
	"github.com/duolacloud/crud-core/types"
	"encoding/json"
)

type JSONMapper [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] struct {
	AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]
}

func NewJSONMapper[DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any]() *MapStructureMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity] {
	m := &JSONMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]{	
	}

	m.fnConvertToDTO = m.ConvertToDTO
	m.fnConvertToEntity = m.ConvertToEntity
	m.fnConvertToCreateEntity = m.ConvertToCreateEntity

	return m
}

func (m *JSONMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToDTO(c context.Context, entity *Entity) (*DTO, error) {
	bytes, err := json.Marshal(entity)
	if err != nil {
		return nil, nil
	}

	var dto *DTO
	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (m *JSONMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToEntity(c context.Context, dto *DTO) (*Entity, error) {
	bytes, err := json.Marshal(dto)
	if err != nil {
		return nil, nil
	}

	var entity *Entity
	err = json.Unmarshal(bytes, &entity)
	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (m *JSONMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToCreateEntity(c context.Context, createDTO *CreateDTO) (*CreateEntity, error) {
	bytes, err := json.Marshal(createDTO)
	if err != nil {
		return nil, nil
	}

	var createEntity *CreateEntity
	err = json.Unmarshal(bytes, &createEntity)
	if err != nil {
		return nil, err
	}

	return createEntity, nil
}

func (m *JSONMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToUpdateEntity(c context.Context, updateDTO *UpdateDTO) (*UpdateEntity, error) {
	bytes, err := json.Marshal(updateDTO)
        if err != nil {
                return nil, nil
        }

	var updateEntity *UpdateEntity
	err = json.Unmarshal(bytes, &updateEntity)
	if err != nil {
		return nil, err
	}

	return updateEntity, nil
}

func (m *JSONMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertQuery(c context.Context, query *types.PageQuery) (*types.PageQuery, error) {
	return query, nil
}
