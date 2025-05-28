package repositories

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func NewSQLDB(sqlDSN string) (*sql.DB, error) {
	db, err := sql.Open("sqlserver", sqlDSN)
	if err != nil {
		return nil, fmt.Errorf("error connecting to SQLite database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging SQL database: %w", err)
	}
	return db, nil
}
