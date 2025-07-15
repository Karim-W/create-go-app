package wires

import (
	"database/sql"
	"fmt"
	"log"

	"{{.moduleName}}/pkg/infrastructure/logger"
	"{{.moduleName}}/pkg/infrastructure/migrations"
	"{{.moduleName}}/pkg/infrastructure/posty"
	"{{.moduleName}}/pkg/infrastructure/rdb"
	"{{.moduleName}}/pkg/infrastructure/rsync"
	"{{.moduleName}}/pkg/infrastructure/tracer"

	"github.com/go-redis/redis/v8"
	"github.com/karim-w/cafe"
	"github.com/karim-w/coreruntime"
	"github.com/karim-w/coreruntime/runtimekeys"
)

func Run(cfg *cafe.Cafe, opt Config) (err error) {
	deps := coreruntime.Dependencies{}
	deps.SQLDatabases = make(map[string]*sql.DB)
	deps.RedisClients = make(map[string]*redis.Client)

	// Infrastructure

	if opt.Infra.PSQL {
		deps.SQLDatabases[runtimekeys.DEFAULT_SQL_KEY], err = posty.Initialize(cfg)
		if err != nil {
			return
		}
	}

	if opt.Infra.Redis {
		deps.RedisClients[runtimekeys.DEFAULT_REDIS_KEY], err = rdb.Initialize(cfg)
	}

	if opt.Infra.Rsync {
		deps.Redsync, err = rsync.Initialize(cfg)
		if err != nil {
			return
		}

	}

	// TODO:
	for _, v := range opt.Infra.Extra {
		switch v.Option {
		case PSQL:
			_, ok := deps.SQLDatabases[v.Name]
			if ok {
				log.Printf("Database %s already exists in the dependencies\n", v.Name)
				err = fmt.Errorf("Database %s already exists in the dependencies", v.Name)
				return
			}

			deps.SQLDatabases[v.Name], err = posty.Initialize(cfg)

		case REDIS:

			_, ok := deps.RedisClients[v.Name]
			if ok {
				log.Printf("Redis client %s already exists in the dependencies\n", v.Name)
				err = fmt.Errorf("Redis client %s already exists in the dependencies", v.Name)
				return
			}

			deps.RedisClients[v.Name], err = rdb.Initialize(cfg)
			if err != nil {
				log.Printf("Failed to initialize Redis client %s: %v\n", v.Name, err)
				return
			}

		default:
			log.Printf("Value %d is not supported as an adapter extra\n", v)
		}
	}

	/// Logger
	deps.Logger, err = logger.Initialize()
	if err != nil {
		return
	}

	/// Tracer
	deps.Tracer, err = tracer.Initialize()
	if err != nil {
		return
	}

	if opt.Infra.MigrationsFlag {
		err = migrations.Initialize(cfg)
		if err != nil {
			return
		}
	}

	// Setup Runtime Backbone
	coreruntime.Setup(deps)

	return
}
