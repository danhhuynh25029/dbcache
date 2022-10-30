package dbcache

import (
	"context"
	"dbcache/cache"
)

type DBCache interface {
	Set(context.Context, string, string) error
	Get(context.Context, string) (string, error)
}
type dbCache struct {
	cache *cache.Cache
}

func NewCache(ctx context.Context) DBCache {
	c := cache.NewCache()
	return &dbCache{
		cache: c,
	}
}

func (c dbCache) Set(ctx context.Context, key string, value string) error {
	err := c.cache.Set(key, value)
	if err != nil {
		return err
	}
	return nil
}
func (c dbCache) Get(ctx context.Context, key string) (string, error) {
	res, err := c.cache.Get(key)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
