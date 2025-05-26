package repositories

import (
	"context"
	"fmt"
)

func (r *SQLRepository) ClientNumberFound(ctx context.Context, number string) (bool, error) {
	var finder int64
	query := `SELECT CAST(
						CASE WHEN EXISTS (SELECT 1 FROM Clients c  WHERE c.MobileNumber  = $1) 
							 THEN 1 ELSE 0 END AS BIT
					) AS ClientExists`

	err := r.db.QueryRowContext(ctx, query, number).Scan(&finder)
	if err != nil {
		return false, fmt.Errorf("error finding client number: %w", err)
	}

	return finder > 0, nil
}
