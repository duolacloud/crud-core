package cache

import (
	"context"
)

type Cache interface {
	Get(c context.Context, key string, opts ...GetOption) (any, error)
	Set(c context.Context, key string, value any, opts ...SetOption) error
}
