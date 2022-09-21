package cache

import (
	"context"
)

type Cache interface {
	Get(c context.Context, key string, value any, opts ...GetOption) error
	Set(c context.Context, key string, value any, opts ...SetOption) error
	Delete(c context.Context, key string, opts ...DeleteOption) error
}
