package cache

import (
	"context"
	"errors"
)

var (
	ErrNotExsit = errors.New("cache not exist.")
)

type Cache interface {
	Get(c context.Context, key string, value any, opts ...GetOption) error
	Set(c context.Context, key string, value any, opts ...SetOption) error
	Delete(c context.Context, key string, opts ...DeleteOption) error
}
