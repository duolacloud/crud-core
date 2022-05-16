package services

import (
	"context"
	"duolacloud.com/duolacloud/crud-core/types"
)

type noopService [DTO any, CreateDTO any, UpdateDTO any] struct {
}

func NewNoOpService[DTO any, CreateDTO any, UpdateDTO any]() CrudService[DTO, CreateDTO, UpdateDTO] {
	return &noopService[DTO, CreateDTO, UpdateDTO]{}
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Create(c context.Context, createDTO *CreateDTO) (*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Delete(c context.Context, id types.ID) error {
	return nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Update(c context.Context, id types.ID, updateDTO *UpdateDTO) (*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Get(c context.Context, id types.ID) (*DTO, error) {
	return nil, nil
}

func (s *noopService[DTO, CreateDTO, UpdateDTO]) Query(c context.Context, query *types.PageQuery[DTO]) ([]*DTO, error) {
	return nil, nil
}
