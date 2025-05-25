package wires

import (
	"log"

	"{{.moduleName}}/pkg/adapters/id"
	"{{.moduleName}}/pkg/adapters/posty"
	"{{.moduleName}}/pkg/adapters/rdb"
	"{{.moduleName}}/pkg/adapters/rsync"
	"{{.moduleName}}/pkg/infrastructure/logger"
	"{{.moduleName}}/pkg/infrastructure/migrations"
	"{{.moduleName}}/pkg/infrastructure/redistracer"
	"{{.moduleName}}/pkg/infrastructure/tracing"
	"{{.moduleName}}/services/factory"

	"github.com/go-redis/redis/v8"
	"github.com/karim-w/cafe"
	"github.com/karim-w/stdlib/sqldb"
)

func Run(cfg *cafe.Cafe, opt Config) (res Result, err error) {
	deps := factory.Dependencies{}

	// Adapters
	if opt.Adapters.Idier {
		res.Adapters.Idier = id.New()
		deps.Id = res.Adapters.Idier
	}

	if opt.Adapters.PSQL {
		res.Adapters.PSQL = make([]sqldb.DB, 1)
		res.Adapters.PSQL[0] = posty.Initialize(cfg)
		deps.PSQL = res.Adapters.PSQL[0]
	}

	if opt.Adapters.Redis {
		res.Adapters.Redis = make([]*redis.Client, 1)
		res.Adapters.Redis[0] = rdb.MustInit(cfg)
		deps.Redis = res.Adapters.Redis[0]
	}

	if opt.Adapters.Rsync {
		res.Adapters.Rsync = rsync.MustInit(cfg)
		deps.RSync = res.Adapters.Rsync
	}

	for _, v := range opt.Adapters.Extra {
		switch v {
		case PSQL:
			if len(res.Adapters.PSQL) == 0 {
				res.Adapters.PSQL = make([]sqldb.DB, 0, 1)
			}

			res.Adapters.PSQL = append(res.Adapters.PSQL, posty.Initialize(cfg))

		case REDIS:
			if len(res.Adapters.Redis) == 0 {
				res.Adapters.Redis = make([]*redis.Client, 0, 1)
			}

			res.Adapters.Redis = append(res.Adapters.Redis, rdb.MustInit(cfg))

		default:
			log.Printf("Value %d is not supported as an adapter extra\n", v)
		}
	}

	// Infra

	/// Logger
	logger.InitOrDie()

	/// Tracer
	res.Infra.Trx = tracing.InitOrDie(opt.Infra.Tracer)
	deps.Trx = res.Infra.Trx

	res.Infra.Redis = make([]*redis.Client, 0, len(res.Adapters.Redis))

	for _, r := range res.Adapters.Redis {
		res.Infra.Redis = append(res.Infra.Redis, redistracer.WrapWithTracing(
			r,
			opt.ServiceName,
			res.Infra.Trx,
		))
	}

	if opt.Infra.MigrationsFlag {
		err = migrations.Initialize(cfg)
		if err != nil {
			return
		}
	}

	// SERVICES
	factory.Init(deps)

	return
}
