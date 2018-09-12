package migrations

import (
	"github.com/rubenv/sql-migrate"
)

// Default - does the thing
func Default() *migrate.MemoryMigrationSource {

	migrationsList := migrate.MemoryMigrationSource {
		Migrations: []*migrate.Migration{
			CreateInitial(),
		},
	}

	return &migrationsList
}