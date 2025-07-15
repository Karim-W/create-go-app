package migrations

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	cas "github.com/golang-migrate/migrate/v4/database/cassandra"
	"github.com/karim-w/coreruntime_extensions/common/options"
	"github.com/karim-w/coreruntime_extensions/connectors/cassandra"
)

// Migrate runs migrations
//
// dsn: the database connection string
//
// migration_source: the path to the migrations
//
// returns: error if any
func Cassandra(
	dsn string,
	migration_source string,
) error {
	log.Println("running cassandra migrations from", migration_source)

	cdb, err := cassandra.Initialize(dsn, options.Cassandra{})
	if err != nil {
		log.Println("failed to initialize cassandra: ", err)
		return err
	}

	cc, err := cas.WithInstance(cdb, &cas.Config{})
	if err != nil {
		log.Println("failed to create cassandra instance: ", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migration_source,
		dsn,
		cc,
	)
	if err != nil {
		log.Println("failed to create migration: ", err)
		return err
	}

	err = m.Up()
	if err == nil {
		log.Println("migrated successfully")
		return nil
	}

	if err == migrate.ErrNoChange {
		log.Println("zero difference in migrations")
		return nil
	}

	log.Println("failed to migrate", err)
	return err
}
