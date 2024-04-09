package factory

import (
	"{{.moduleName}}/pkg/infrastructure/tracing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/karim-w/stdlib/sqldb"
)

type Dependencies struct {
	Trx   tracing.Tracer
	PSQL  sqldb.DB
	Redis *redis.Client
	RSync *redsync.Redsync
}

var deps *Dependencies

// not thread safe
func SetUpDependencies(trx tracing.Tracer) {
	if deps != nil {
		return
	}

	deps = &Dependencies{Trx: trx}
}

// Init initializes the dependencies manually to avoid breaking
// SetDependencies in the future
// Allows u to set the dependencies based to what is available
// NOT THREAD SAFE, i assume you know what you are doing
func Init(d Dependencies) {
	if deps != nil {
		return
	}

	deps = &d
}
