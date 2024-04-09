package migrations

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// Migrate runs migrations
//
// dsn: the database connection string
//
// migration_source: the path to the migrations
//
// returns: error if any
func Run(
	dsn string,
	migration_source string,
) error {
	log.Println("running migrations from", migration_source, "to", dsn)

	m, err := migrate.New(
		migration_source,
		dsn,
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
