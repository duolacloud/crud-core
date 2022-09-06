package cache

import (
	"time"
)

type CacheOptions struct{}
type CacheOption func(*CacheOptions)

type GetOptions struct{}
type GetOption func(*GetOptions)

type SetOptions struct {
	Exipration time.Duration
}
type SetOption func(*SetOptions)

func WithExpiration(exipration time.Duration) SetOption {
	return func(o *SetOptions) {
		o.Exipration = exipration
	}
}

type DeleteOptions struct{}
type DeleteOption func(*DeleteOptions)
