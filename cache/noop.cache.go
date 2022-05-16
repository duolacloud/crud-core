package cache

import "context"

type noopCache struct {
}

func NewNoopCache() Cache {
	return &noopCache{}
}

func (n *noopCache) Get(c context.Context, key string, opts ...GetOption) (interface{}, error) {
	return nil, nil
}

func (n *noopCache) Set(c context.Context, key string, value interface{}, opts ...SetOption) error {
	return nil
}
