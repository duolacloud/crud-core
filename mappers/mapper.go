package mappers

import (
	"context"
	"github.com/duolacloud/crud-core/types"
)

type Mapper [DTO any, CreateDTO any, UpdateDTO any, Entity any, CreateEntity any, UpdateEntity any] interface {
	ConvertToDTO(c context.Context, entity *Entity) (*DTO, error)
	ConvertToDTOs(c context.Context, entities []*Entity) ([]*DTO, error)
	ConvertToEntity(c context.Context, dto *DTO) (*Entity, error)
	ConvertToCreateEntity(c context.Context, createDTO *CreateDTO) (*CreateEntity, error)
	ConvertToCreateEntities(c context.Context, createDTO []*CreateDTO) ([]*CreateEntity, error)
	ConvertToUpdateEntity(c context.Context, updateDTO *UpdateDTO) (*UpdateEntity, error)
	ConvertToEntities(c context.Context, dto []*DTO) ([]*Entity, error)
	ConvertQuery(c context.Context, query *types.PageQuery) (*types.PageQuery, error)
}
