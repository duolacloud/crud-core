package services

import (
	"context"
	"testing"
	"duolacloud.com/duolacloud/crud-core/types"
)

type UserDTO struct {
	ID string
}

type UserEntity struct {
	ID string
}

func TestNoopService(t *testing.T) {
	s := NewNoopService[UserDTO, UserDTO, UserDTO]()
	u, err := s.Create(context.TODO(), &UserDTO{
		ID: "1",
	})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("haha: %v", u)

	us, err := s.Query(context.TODO(), &types.PageQuery[UserDTO]{

	})

	if err != nil {
		t.Fatal(err)
	}
	t.Logf("haha: %v", us)
}
