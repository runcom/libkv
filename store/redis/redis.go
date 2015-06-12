package redis

import (
	"github.com/docker/libkv/store"
	redis "gopkg.in/redis.v3"
)

type Redis struct {
	config redis.Config
	client redis.Client
}

// InitializeConsul creates a new Consul client given
// a list of endpoints and optional tls config
func InitializeRedis(endpoints []string, options *store.Config) (store.Store, error) {
}
