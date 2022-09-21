package services

import (
	"context"
	"testing"

	"github.com/duolacloud/crud-core/cache"
)

type UserDTO1 struct {
	ID string
}

func TestCacheService(t *testing.T) {
	cache := cache.NewNoopCache()
	service := NewNoopService[UserDTO1, UserDTO1, UserDTO1]()

	s := NewCacheService(service, cache)
	u, err := s.Create(context.TODO(), &UserDTO1{
		ID: "1",
	})

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("haha: %v", u)
}
