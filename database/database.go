package database

import (
    _ "github.com/go-sql-driver/mysql"
    // "github.com/myanrichal/VeWeb/database/sql"
)

// Database - interface for sql type
type Database struct {
    SQL         *SQL
}

// Default - connects database 
func Default() *Database {

    database := Database {
        SQL: DefaultSQL(),
    }

    return &database
}