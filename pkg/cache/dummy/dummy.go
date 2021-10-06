package dummy

import (
	"time"

	cacheroot "github.com/deniskelin/billing-gokit/pkg/cache"
)

const CacheName = "dummy"

type cache struct{}

func NewCache() (*cache, error) {
	return &cache{}, nil
}

func (c *cache) Get(_ []byte) ([]byte, error) {
	return nil, cacheroot.ErrCacheKeyNotFound
}

func (c *cache) GetByStringKey(_ string) ([]byte, error) {
	return nil, cacheroot.ErrCacheKeyNotFound
}

func (c *cache) Set(_ []byte, _ []byte, _ time.Duration) error {
	return nil
}

func (c *cache) SetByStringKey(_ string, _ []byte, _ time.Duration) error {
	return nil
}

func (c *cache) Delete(_ []byte) error {
	return nil
}

func (c *cache) DeleteByStringKey(_ string) error {
	return nil
}

func (c *cache) Flush() error {
	return nil
}

func (c *cache) Close() error {
	return nil
}
