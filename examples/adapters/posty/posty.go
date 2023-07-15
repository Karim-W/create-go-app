// Description: Database context initialization
// posty for postgresql
package posty

import (
	"github.com/karim-w/stdlib/sqldb"
)

// MustInit initializes a database connection pool
// and panics if it fails to do so or if given an
// invalid driver.
func MustInit(
	driver string,
	dsn string,
	conns int,
) sqldb.DB {
	dbCtx := sqldb.NewWithOptions(
		driver,
		dsn,
		&sqldb.Options{
			MaxIdleConns: conns,
			MaxOpenConns: conns,
		},
	)
	return dbCtx
}
