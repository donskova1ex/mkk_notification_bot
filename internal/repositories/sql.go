package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

func NewSQLDB(ctx context.Context, sqlDSN string) (*sql.DB, error) {
	dbDNS := ""
	db, err := sql.Open("sqlserver", dbDNS)
	if err != nil {
		return nil, fmt.Errorf("error connecting to SQLite database: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging SQL database: %w", err)
	}
	return db, nil
}
