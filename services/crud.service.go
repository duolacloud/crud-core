package services

import (
	"context"

	"duolacloud.com/duolacloud/crud-core/types"
)

type CrudService [DTO any, CreateDTO any, UpdateDTO any] interface {
	Create(c context.Context, dto *CreateDTO) (*DTO, error)
	Delete(c context.Context, id types.ID) error
	Update(c context.Context, id types.ID, updateDTO *UpdateDTO) (*DTO, error)
	Get(c context.Context, id types.ID) (*DTO, error)
	Query(c context.Context, query *types.PageQuery[DTO]) ([]*DTO, error)
}
