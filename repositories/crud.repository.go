package repositories

import (
	"context"
	"duolacloud.com/duolacloud/crud-core/types"
)

type CrudRepository[DTO any, CreateDTO any, UpdateDTO any] interface {
	Create(c context.Context, createDTO *CreateDTO) (*DTO, error)
	Delete(c context.Context, id types.ID) error
	Update(c context.Context, id types.ID, updateDTO *UpdateDTO) (*DTO, error)
	Get(c context.Context, id types.ID) (*DTO, error)
	Query(c context.Context, query *types.PageQuery[DTO]) ([]*DTO, error)
}
