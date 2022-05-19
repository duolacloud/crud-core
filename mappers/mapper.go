package mappers

import (
	"context"
	"duolacloud.com/duolacloud/crud-core/types"
)

type Mapper [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] interface {
	ConvertToDTO(c context.Context, entity *Entity) *DTO
	ConvertToDTOs(c context.Context, entities []*Entity) []*DTO
	ConvertToEntity(c context.Context, dto *DTO) *Entity
	ConvertToCreateEntity(c context.Context, createDTO *CreateDTO) *CreateEntity
	ConvertToUpdateEntity(c context.Context, updateDTO *UpdateDTO) *UpdateEntity
	ConvertToEntities(c context.Context, dto []*DTO) []*Entity
	ConvertQuery(c context.Context, query *types.PageQuery) *types.PageQuery
}
