package database

import (
    "database/sql"
    "log"
    "github.com/jmoiron/sqlx"
	"github.com/rubenv/sql-migrate"
	
)

type SQL struct {
    Dbx         *sqlx.DB 
    migrations  *migrate.MemoryMigrationSource
}

func DefaultSQL() *SQL {
    // create db connection
    connectionString := context.Config.DbUser + ":" + context.Config.DbPassword + "@" + context.Config.DbServer + "/" + context.Config.DbName + "?parseTime=true"
    dbHandle, err := sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal ("Database Error: ", err.Error())
    }
    dbx := sqlx.NewDb(dbHandle, "mysql")

    sql := &SQL {
        Dbx:        dbx,
        migrations: migrations.Default(),
    }

    //apply migrations up by Default
    return sql
}

func (database *Database) MigrateSql() error {
    tableName := "cqlw_migrations"
    migrate.SetTable(tableName)
    n, err := migrate.Exec(database.SQL.Dbx.DB, "mysql", database.SQL.migrations, migrate.Up)
    if err != nil {
        log.Printf("MIGRATION ERROR: %s\n", err.Error())
        if n > 0 {
            rn, err := migrate.ExecMax(database.SQL.Dbx.DB, "mysql", database.SQL.migrations, migrate.Down, n)
            if err != nil {
                log.Printf("ROLLBACK FAILED: %s\n", err.Error())
                return err
            }
            log.Printf("Rolled back %d migrations.\n", rn)
            return err
        } else {
            log.Println("No rollback required.")
            return err
        }
    }
    if n > 0 {
        log.Printf("Applied %d migrations to %s. Database up to date. \n", n, tableName)
    }
    return nil
}