package migrations_test

import (
	"ams/pkg/infrastructure/migrations"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	cassUrl := os.Getenv("CASSANDRA_URI")
	if cassUrl == "" {
		t.Skip("CASSANDRA_URI is not set")
	}

	path := os.Getenv("DB_MIGRATIONS_PATH")
	if path == "" {
		t.Skip("DB_MIGRATIONS_PATH is not set")
	}

	err := migrations.Cassandra(cassUrl, path)
	if err != nil {
		panic(err)
		t.Error(err)
	}
}
