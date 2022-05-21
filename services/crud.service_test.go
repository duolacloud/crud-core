package services

import (
	"context"
	"testing"
	"github.com/duolacloud/crud-core/types"
	"github.com/duolacloud/crud-core/repositories"
	"github.com/duolacloud/crud-core/mappers"
)

type UserDTO struct {
	ID string
}

type UserEntity struct {
	ID string
}

func TestCrudService(t *testing.T) {
	s := NewCrudService[UserDTO, UserDTO, UserDTO](
		repositories.NewMapperRepository[UserDTO, UserDTO, UserDTO, UserEntity, UserEntity, UserEntity](
			repositories.NewNoopRepository[UserEntity, UserEntity, UserEntity](),
			mappers.NewMapStructureMapper[UserDTO, UserDTO, UserDTO, UserEntity, UserEntity, UserEntity](),
		),
	)
	u, err := s.Create(context.TODO(), &UserDTO{
		ID: "1",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("haha: %v", u)

	us, err := s.Query(context.TODO(), &types.PageQuery{

	})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("haha: %v", us)
}
