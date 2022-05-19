package services

import (
	"context"
	"testing"
	"duolacloud.com/duolacloud/crud-core/types"
)

type UserDTO2 struct {
	ID string
}

func TestNoopService(t *testing.T) {
	s := NewNoopService[UserDTO2, UserDTO2, UserDTO2]()
	u, err := s.Create(context.TODO(), &UserDTO2{
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
