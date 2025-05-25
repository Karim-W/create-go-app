package factory

import (
	"context"

	"{{.moduleName}}/pkg/infrastructure/tracing"

	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/karim-w/stdlib/httpclient"
	"github.com/karim-w/stdlib/sqldb"
	"go.uber.org/zap"
)

type Service interface {
	context.Context
	Logger() *zap.Logger
	Context() context.Context

	// STORE
	Get(key string) (any, bool)
	Set(key string, value any)

	Span() (string, string)
	Child() Service

	// HTTP
	HttpClient(url string) httpclient.HTTPRequest

	// INFO
	TraceParent() string
	Caller() string
	Trx() tracing.Tracer
	TraceInfo() (ver, tid, pid, rid, flg string)
	SetCaller(caller string)

	// REDIS
	RDB() *redis.Client
	DLock(key string) (*redsync.Mutex, error)

	// SQL
	PSQL() sqldb.DB
	BeginPsqlTx() (*sqldb.Tx, bool, error)
	CommitPsqlTx() error
}
