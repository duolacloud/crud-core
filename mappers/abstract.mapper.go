package mappers

import (
	"context"
	"duolacloud.com/duolacloud/crud-core/types"
)

type AbstractMapper [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] struct {
	fnConvertToDTO func(context.Context, *Entity) *DTO
	fnConvertToEntity func(context.Context, *DTO) *Entity
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToDTOs(c context.Context, entities []*Entity) []*DTO {
	dtos := make([]*DTO, len(entities))
	for i, entity := range entities {
		dtos[i] = m.fnConvertToDTO(c, entity)
	}
	return dtos
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToEntities(c context.Context, dtos []*DTO) []*Entity {
	entities := make([]*Entity, len(dtos))
	for i, dto := range dtos {
		entities[i] = m.fnConvertToEntity(c, dto)
	}
	return entities
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertQuery(c context.Context, query *types.PageQuery) *types.PageQuery {
	return nil
}
