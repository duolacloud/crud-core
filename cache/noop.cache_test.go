package cache

import (
	"context"
	"testing"
)

func TestNoopCache(t *testing.T) {
	c := NewNoopCache()
	c.Set(context.TODO(), "k", "v")
	v := new(string)
	c.Get(context.TODO(), "k", v)
	c.Delete(context.TODO(), "k")
}
