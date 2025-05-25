package wires

import (
	"{{.moduleName}}/pkg/adapters/id"
	"{{.moduleName}}/pkg/infrastructure/tracing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/karim-w/stdlib/sqldb"
)

type (
	AdapterOption int64
)

const (
	IDIER AdapterOption = iota
	PSQL
	REDIS
)

type AdapterFlags struct {
	Idier bool
	PSQL  bool
	Redis bool
	Rsync bool
	Extra []AdapterOption
}

type AdapterResults struct {
	Idier id.Id
	PSQL  []sqldb.DB
	Redis []*redis.Client
	Rsync *redsync.Redsync
}

type InfraOptions struct {
	MigrationsFlag bool
	Tracer         tracing.Tracer
}

type InfraResults struct {
	Trx   tracing.Tracer
	Redis []*redis.Client
}

type Config struct {
	ServiceName string
	Adapters    AdapterFlags
	Infra       InfraOptions
}

type Result struct {
	Adapters AdapterResults
	Infra    InfraResults
}
