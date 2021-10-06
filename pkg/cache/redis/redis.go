package redis

import (
	"context"
	"time"

	cacheroot "github.com/deniskelin/billing-gokit/pkg/cache"
	"github.com/deniskelin/billing-gokit/pkg/helper"
	"github.com/go-redis/redis/v8"
)

const CacheName = "redis"

type cache struct {
	rdb *redis.Client
}

var ctx = context.Background()

// NewCache
// There are two connection types: by tcp socket and by unix socket.
// Tcp connection:
// 		redis://<user>:<password>@<host>:<port>/<db_number>
// Unix connection:
//		unix://<user>:<password>@</path/to/redis.sock>?db=<db_number>
func NewCache(connStr string) (*cache, error) {
	opts, err := redis.ParseURL(connStr)
	if err != nil {
		return nil, err
	}
	return &cache{rdb: redis.NewClient(opts)}, nil
}

func (c *cache) Get(key []byte) ([]byte, error) {
	return c.GetByStringKey(helper.B2S(key))
}

func (c *cache) GetByStringKey(key string) ([]byte, error) {
	val, err := c.rdb.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return nil, cacheroot.ErrCacheKeyNotFound
		} else {
			return nil, err
		}
	}
	return val, nil
}

func (c *cache) Set(key []byte, value []byte, ttl time.Duration) error {
	return c.SetByStringKey(helper.B2S(key), value, ttl)
}

func (c *cache) SetByStringKey(key string, value []byte, ttl time.Duration) error {
	return c.rdb.Set(ctx, key, value, ttl).Err()
}

func (c *cache) Delete(key []byte) error {
	return c.DeleteByStringKey(helper.B2S(key))
}

func (c *cache) DeleteByStringKey(key string) error {
	return c.rdb.Del(ctx, key).Err()
}

func (c *cache) Flush() error {
	return c.rdb.FlushDB(ctx).Err()
}

func (c *cache) Close() error {
	return c.rdb.Close()
}
