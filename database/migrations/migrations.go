package migrations

import (
	"github.com/rubenv/sql-migrate"
)

func Default() *migrate.MemoryMigrationSource {

	migrationsList := migrate.MemoryMigrationSoruce {
		Migrations: []*migrate.Migration{
			CreateInitial(),
		},
	}

	return &migrationsList
}