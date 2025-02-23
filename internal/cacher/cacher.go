package cacher

import (
	"context"
	"sync"

	"github.com/go-gorm/caches/v4"
)

type Cacher struct {
	store *sync.Map
}

func (c *Cacher) init() {
	if c.store == nil {
		c.store = &sync.Map{}
	}
}

func (c *Cacher) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	c.init()
	val, ok := c.store.Load(key)
	if !ok {
		return nil, nil
	}

	if err := q.Unmarshal(val.([]byte)); err != nil {
		return nil, err
	}

	return q, nil
}

func (c *Cacher) Store(ctx context.Context, key string, val *caches.Query[any]) error {
	c.init()
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	c.store.Store(key, res)
	return nil
}

func (c *Cacher) Invalidate(ctx context.Context) error {
	c.store = &sync.Map{}
	return nil
}
