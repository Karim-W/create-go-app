package rsync

import (
	"log/slog"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
)

func Init(
	redisUri string,
) *redsync.Redsync {
	// Parse the Redis URI.
	opts, err := redis.ParseURL(redisUri)
	if err != nil {
		slog.Error("Failed to parse redis uri", slog.String("error", err.Error()))
		panic(err)
	}

	// Create a Redis client.
	client := redis.NewClient(opts)

	// Create a pool of Redis connections.
	pool := goredis.NewPool(client)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	return redsync.New(pool)
}
