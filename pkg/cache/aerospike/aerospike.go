package aerospike

import (
	"net/url"
	"strconv"
	"time"

	"github.com/aerospike/aerospike-client-go"
	"github.com/deniskelin/billing-gokit/pkg/helper"
	"github.com/pkg/errors"
)

const CacheName = "aerospike"

type cache struct {
	aero      *aerospike.Client
	namespace string
	setName   string
}

// NewCache
// There are two connection types: by tcp socket
// Tcp connection:
// 		aerospike://<host>:<port>/?namespace=<namespace>&set=<set_name>
func NewCache(connStr string) (*cache, error) {
	uri, err := url.ParseRequestURI(connStr)
	if err != nil {
		return nil, err
	}
	if uri.Scheme != CacheName {
		return nil, errors.New("schema not equals aerospike")
	}
	port, err := strconv.Atoi(uri.Port())
	if err != nil {
		return nil, err
	}

	namespace := uri.Query().Get("namespace")
	if namespace == "" {
		return nil, errors.New("empty namespace not allowed")
	}
	setName := uri.Query().Get("set")
	if setName == "" {
		return nil, errors.New("empty set not allowed")
	}

	client, err := aerospike.NewClient(uri.Host, port)
	if err != nil {
		return nil, err
	}
	return &cache{aero: client, namespace: namespace, setName: setName}, nil
}

func (c *cache) Get(key []byte) ([]byte, error) {
	return c.GetByStringKey(helper.B2S(key))
}

func (c *cache) GetByStringKey(key string) (val []byte, err error) {
	ckey, err := aerospike.NewKey(c.namespace, c.setName, key)
	if err != nil {
		return nil, err
	}

	err = c.aero.GetObject(nil, ckey, &val)
	if err != nil {
		return nil, err
	}
	return
}

func (c *cache) Set(key []byte, value []byte, ttl time.Duration) error {
	return c.SetByStringKey(helper.B2S(key), value, ttl)
}

func (c *cache) SetByStringKey(key string, value []byte, ttl time.Duration) error {
	ckey, err := aerospike.NewKey(c.namespace, c.setName, key)
	if err != nil {
		return err
	}
	policy := aerospike.NewWritePolicy(0, uint32(ttl.Seconds()))
	return c.aero.PutObject(policy, ckey, value)
}

func (c *cache) Delete(key []byte) error {
	return c.DeleteByStringKey(helper.B2S(key))
}

func (c *cache) DeleteByStringKey(key string) error {
	ckey, err := aerospike.NewKey(c.namespace, c.setName, key)
	if err != nil {
		return err
	}
	_, err = c.aero.Delete(nil, ckey)
	return err
}

func (c *cache) Flush() error {
	return c.aero.Truncate(nil, c.namespace, c.setName, nil)
}

func (c *cache) Close() error {
	c.aero.Close()
	return nil
}
