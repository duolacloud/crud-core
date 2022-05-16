package cache

import (
	"context"
)

type Cache interface {
	Get(c context.Context, key string, opts ...GetOption) (interface{}, error)
	Set(c context.Context, key string, value interface{}, opts ...SetOption) error
}
