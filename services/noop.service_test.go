package services

import (
	"context"
	"testing"
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
}
