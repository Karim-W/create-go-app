package posty

import (
	"github.com/karim-w/cafe"
	"github.com/karim-w/stdlib/sqldb"
)

const (
	CONFIG_KEY = "PSQL"
	DRIVER_KEY = "postgres"
)

// Cafe is the config schema for the rdb package.
// It is used by the config package to initialize the config.
var Cafe = cafe.SubSchema(
	CONFIG_KEY,
	cafe.Schema{
		"DB_DSN":       cafe.String("DB_DSN").Require(),
		"DB_MAX_CONNS": cafe.Int("DB_MAX_OPEN_CONNS").Default(30),
	},
)

func Initialize(c *cafe.Cafe) sqldb.DB {
	config, err := c.GetSubSchema(CONFIG_KEY)
	if err != nil {
		panic(err)
	}

	dsn, err := config.GetString("DB_DSN")
	if err != nil {
		panic(err)
	}

	maxConns, err := config.GetInt("DB_MAX_CONNS")
	if err != nil {
		panic(err)
	}

	return MustInit(DRIVER_KEY, dsn, maxConns)
}
