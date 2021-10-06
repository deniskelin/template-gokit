package connector

import (
	"github.com/deniskelin/billing-gokit/pkg/cache"
	"github.com/deniskelin/billing-gokit/pkg/cache/aerospike"
	"github.com/deniskelin/billing-gokit/pkg/cache/dummy"
	"github.com/deniskelin/billing-gokit/pkg/cache/memory"
	"github.com/deniskelin/billing-gokit/pkg/cache/redis"
)

func NewCache(cacheType, connString string) (cache.ICache, error) {
	switch cacheType {
	case dummy.CacheName:
		return dummy.NewCache()
	case memory.CacheName:
		return memory.NewCache()
	case redis.CacheName:
		return redis.NewCache(connString)
	case aerospike.CacheName:
		return aerospike.NewCache(connString)
	default:
		return nil, cache.ErrUnsupportedCacheType
	}
}
