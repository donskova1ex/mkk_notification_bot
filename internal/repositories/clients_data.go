package repositories

import (
	"context"
	"database/sql"
	"fmt"
)

func (r *SQLRepository) ClientNumberFound(ctx context.Context, number string) (bool, error) {
	var finder int32
	query := `SELECT CAST(
						CASE WHEN EXISTS (SELECT 1 FROM Clients c  WHERE c.MobileNumber  = @number) 
							 THEN 1 ELSE 0 END AS INTEGER
					) AS ClientExists`

	row := r.db.QueryRowContext(ctx, query, sql.Named("number", number))
	err := row.Scan(&finder)
	if err != nil {
		return false, fmt.Errorf("error scan client number: %w", err)
	}

	return finder > 0, nil
}
