package cache

import "context"

type noopCache struct {
}

func NewNoopCache() Cache {
	return &noopCache{}
}

func (n *noopCache) Get(c context.Context, key string, value any, opts ...GetOption) error {
	return nil
}

func (n *noopCache) Set(c context.Context, key string, value any, opts ...SetOption) error {
	return nil
}

func (n *noopCache) Delete(c context.Context, key string, opts ...DeleteOption) error {
	return nil
}
