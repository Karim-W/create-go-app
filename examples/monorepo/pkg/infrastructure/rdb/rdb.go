package rdb

// Redis client initialization and access.

import (
	"github.com/go-redis/redis/v8"
	"github.com/karim-w/coreruntime_extensions/common/options"
	red "github.com/karim-w/coreruntime_extensions/connectors/rdb"
)

// Init initializes a Redis client with the provided URI.
func Init(
	uri string,
) (*redis.Client, error) {
	return red.Initialize(uri, options.Redis{
		PanicablePings: true,
	})
}
