package rdb

// Redis client initialization and access.

import (
	"github.com/go-redis/redis/v8"
)

// InitRedisOrDie initializes a redis client and panics if it fails.
func InitRedisOrDie(
	uri string,
) *redis.Client {
	rop, err := redis.ParseURL(uri)
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(
		rop,
	)

	err = client.Ping(client.Context()).Err()
	if err != nil {
		panic(err)
	}

	return client
}
