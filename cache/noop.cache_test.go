package cache

import (
	"context"
	"testing"
)

func TestNoopCache(t *testing.T) {
	c := NewNoopCache()
	c.Set(context.TODO(), "k", "v")

	c.Get(context.TODO(), "k")
}
