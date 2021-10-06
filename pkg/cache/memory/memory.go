package memory

import (
	"time"

	cacheroot "github.com/deniskelin/billing-gokit/pkg/cache"
	"github.com/deniskelin/billing-gokit/pkg/helper"
	pcache "github.com/patrickmn/go-cache"
)

const CacheName = "memory"

type cache struct {
	cacher *pcache.Cache
}

func NewCache() (*cache, error) {
	return &cache{
		cacher: pcache.New(5*time.Minute, 10*time.Minute),
	}, nil
}

func (c *cache) Get(key []byte) ([]byte, error) {
	return c.GetByStringKey(helper.B2S(key))
}

func (c *cache) GetByStringKey(key string) ([]byte, error) {
	val, found := c.cacher.Get(key)
	if !found {
		return nil, cacheroot.ErrCacheKeyNotFound
	}
	return val.([]byte), nil
}

func (c *cache) Set(key []byte, val []byte, ttl time.Duration) error {
	return c.SetByStringKey(helper.B2S(key), val, ttl)
}

func (c *cache) SetByStringKey(key string, val []byte, ttl time.Duration) error {
	c.cacher.Set(key, val, ttl)
	return nil
}

func (c *cache) Delete(key []byte) error {
	return c.DeleteByStringKey(helper.B2S(key))
}

func (c *cache) DeleteByStringKey(key string) error {
	c.cacher.Delete(key)
	return nil
}

func (c *cache) Flush() error {
	c.cacher.Flush()
	return nil
}

func (c *cache) Close() error {
	return nil
}
