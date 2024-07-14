package migrations

import (
	"github.com/karim-w/cafe"
)

const (
	CONFIG_KEY    = "MIGRATIONS"
	POSTGRES_KEY  = "postgres"
	CASSANDRA_KEY = "cassandra"
)

// Cafe is the config schema for the rdb package.
// It is used by the config package to initialize the config.
var Cafe = cafe.SubSchema(
	CONFIG_KEY,
	cafe.Schema{
		"DB_DRIVER": cafe.String("DB_DRIVER").Default("postgres"),
		"DB_DSN":    cafe.String("DB_DSN").Require(),
		"DB_MIGRATION_SOURCE": cafe.String("DB_MIGRATION_SOURCE").
			Default("file://internal/sql_migrations"),
	},
)

func Initialize(c *cafe.Cafe) error {
	config, err := c.GetSubSchema(CONFIG_KEY)
	if err != nil {
		panic(err)
	}

	dsn, err := config.GetString("DB_DSN")
	if err != nil {
		panic(err)
	}

	driver, err := config.GetString("DB_DRIVER")
	if err != nil {
		panic(err)
	}

	migrationSource, err := config.GetString("DB_MIGRATION_SOURCE")
	if err != nil {
		panic(err)
	}

	switch driver {
	case "postgres":
		return Run(dsn, migrationSource)
	case "cassandra":
		return Cassandra(dsn, migrationSource)
	default:
		panic("unsupported driver")
	}
}
