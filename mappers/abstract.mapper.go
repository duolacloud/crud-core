package mappers

import (
	"context"
	"github.com/duolacloud/crud-core/types"
)

type AbstractMapper [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] struct {
	fnConvertToDTO func(context.Context, *Entity) (*DTO, error)
	fnConvertToEntity func(context.Context, *DTO) (*Entity, error)
	fnConvertToCreateEntity func(context.Context, *CreateDTO) (*CreateEntity, error)
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToDTOs(c context.Context, entities []*Entity) ([]*DTO, error) {
	var err error
	dtos := make([]*DTO, len(entities))
	for i, entity := range entities {
		dtos[i], err = m.fnConvertToDTO(c, entity)
		if err != nil {
			return nil, err
		}
	}
	return dtos, nil
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToEntities(c context.Context, dtos []*DTO) ([]*Entity, error) {
	var err error
	entities := make([]*Entity, len(dtos))
	for i, dto := range dtos {
		entities[i], err = m.fnConvertToEntity(c, dto)
		if err != nil {
			return nil, err
		}
	}
	return entities, nil
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToCreateEntities(c context.Context, items []*CreateDTO) ([]*CreateEntity, error) {
	var err error
	entities := make([]*CreateEntity, len(items))
	for i, dto := range items {
		entities[i], err = m.fnConvertToCreateEntity(c, dto)
		if err != nil {
			return nil, err
		}
	}
	return entities, nil
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertQuery(c context.Context, query *types.PageQuery) (*types.PageQuery, error) {
	return query, nil
}
