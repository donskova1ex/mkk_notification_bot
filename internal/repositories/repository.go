package repositories

import (
	"database/sql"
	"log/slog"
)

type SQLRepository struct {
	db     *sql.DB
	logger *slog.Logger
}

func NewSQLRepository(db *sql.DB, logger *slog.Logger) *SQLRepository {
	return &SQLRepository{
		db:     db,
		logger: logger,
	}
}
