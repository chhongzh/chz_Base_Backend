package cacher

import (
	"context"
	"sync"

	"github.com/go-gorm/caches/v4"
	"go.uber.org/zap"
)

type Cacher struct {
	store  *sync.Map
	logger *zap.Logger

	totalHitCount   int
	totalStoreCount int
}

func NewCacher(logger *zap.Logger) *Cacher {
	return &Cacher{
		store:  &sync.Map{},
		logger: logger,
	}
}

func (c *Cacher) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
	val, ok := c.store.Load(key)
	if !ok {
		return nil, nil
	}

	if err := q.Unmarshal(val.([]byte)); err != nil {
		return nil, err
	}

	c.totalHitCount++

	return q, nil
}

func (c *Cacher) Store(ctx context.Context, key string, val *caches.Query[any]) error {
	res, err := val.Marshal()
	if err != nil {
		return err
	}

	c.totalStoreCount++

	c.store.Store(key, res)
	return nil
}

func (c *Cacher) Invalidate(ctx context.Context) error {
	c.logger.Info("Cache Invalidated", zap.Int("Cleanup Count", c.len()))

	c.store.Clear()
	return nil
}

func (c *Cacher) len() int {
	length := 0

	c.store.Range(func(key, value any) bool {
		length++
		return true
	})

	return length
}

func (c *Cacher) GetCacheInfo() (totalHit int, totalStore int) {
	return c.totalHitCount, c.totalStoreCount
}
