// Description: Database context initialization
// posty for postgresql
package posty

import (
	"database/sql"

	"github.com/karim-w/coreruntime_extensions/common/options"
	"github.com/karim-w/coreruntime_extensions/connectors/psql"
	"github.com/karim-w/gopts"
)

// Init initializes a database connection pool
// with the provided Data Source Name (DSN) and connection options.
// It returns a pointer to sql.DB and an error if any occurs during initialization.
func Init(
	dsn string,
	conns int,
) (*sql.DB, error) {
	return psql.MustInit(
		dsn,
		options.SQL_Database{
			MaxIdleConns:   gopts.Some(conns),
			MaxOpenConns:   gopts.Some(conns),
			PanicablePings: true,
		},
	)
}
