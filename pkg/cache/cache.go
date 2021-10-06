package cache

import (
	"time"

	"github.com/pkg/errors"
)

type ICache interface {
	Get(key []byte) (value []byte, err error)
	GetByStringKey(key string) (value []byte, err error)
	Set(key []byte, value []byte, ttl time.Duration) (err error)
	SetByStringKey(key string, value []byte, ttl time.Duration) (err error)
	Delete(key []byte) (err error)
	DeleteByStringKey(key string) (err error)
	Flush() error
	Close() error
}

var (
	ErrUnsupportedCacheType = errors.New("unsupported cache type")
	ErrCacheKeyNotFound     = errors.New("key not found")
	// ErrCacheServerNotWorks = errors.New("cache server not works")
	// ErrCacheServerError    = errors.New("cache server error")
)

const (
	NoExpiration time.Duration = -1
)
