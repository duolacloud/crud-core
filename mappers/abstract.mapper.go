package mappers

import (
	"context"
	"duolacloud.com/duolacloud/crud-core/types"
)

type AbstractMapper [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] struct {
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToDTO(c context.Context, entity *Entity) *DTO {
	return nil
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToDTOs(c context.Context, entities []*Entity) []*DTO {
	dtos := make([]*DTO, len(entities))
	for i, entity := range entities {
		dtos[i] = m.ConvertToDTO(c, entity)
	}
	return dtos
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToEntity(c context.Context, dto *DTO) *Entity {
	return nil
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToCreateEntity(c context.Context, createDTO *CreateDTO) *CreateEntity {
	return nil
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToUpdateEntity(c context.Context, updateDTO *UpdateDTO) *UpdateEntity {
	return nil
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertToEntities(c context.Context, dtos []*DTO) []*Entity {
	entities := make([]*Entity, len(dtos))
	for i, dto := range dtos {
		entities[i] = m.ConvertToEntity(c, dto)
	}
	return entities
}

func (m *AbstractMapper[DTO, CreateDTO, UpdateDTO, Entity, CreateEntity, UpdateEntity]) ConvertQuery(c context.Context, query *types.PageQuery) *types.PageQuery {
	return nil
}
