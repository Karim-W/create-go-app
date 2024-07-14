package cassie

import (
	"github.com/gocql/gocql"
	"github.com/karim-w/cafe"
)

const (
	CONFIG_KEY = "CDB"
	DRIVER_KEY = "cassandra"
)

// Cafe is the config schema for the cassie package.
// It is used by the config package to initialize the config.
var Cafe = cafe.SubSchema(
	CONFIG_KEY,
	cafe.Schema{
		"DB_DSN": cafe.String("DB_DSN").Require(),
	},
)

func Initialize(c *cafe.Cafe) (*gocql.Session, error) {
	config, err := c.GetSubSchema(CONFIG_KEY)
	if err != nil {
		panic(err)
	}

	dsn, err := config.GetString("DB_DSN")
	if err != nil {
		panic(err)
	}

	username, password, hosts, keyspace, port, ssl, err := ParseUri(dsn)
	if err != nil {
		return nil, err
	}

	return Connect(hosts, port, keyspace, username, password, ssl)
}
