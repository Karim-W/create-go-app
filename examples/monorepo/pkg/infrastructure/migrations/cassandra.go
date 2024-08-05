package migrations

import (
	"log"

	"{{.moduleName}}/pkg/adapters/cassie"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/cassandra"
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

	user, pass, host, keyspace, port, tls, err := cassie.ParseUri(dsn)
	if err != nil {
		log.Println("failed to parse dsn: ", err)
		return err
	}

	cdb, err := cassie.Connect(host, port, keyspace, user, pass, tls)
	if err != nil {
		log.Println("failed to connect to cassandra: ", err)
		return err
	}

	cassie, err := cassandra.WithInstance(cdb, &cassandra.Config{
		KeyspaceName: keyspace,
	})
	if err != nil {
		log.Println("failed to create cassandra instance: ", err)
		return nil
	}

	m, err := migrate.NewWithDatabaseInstance(
		migration_source,
		dsn,
		cassie,
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
