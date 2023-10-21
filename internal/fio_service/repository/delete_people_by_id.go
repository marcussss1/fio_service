package repository

import (
	"context"
	"database/sql"
)

func (r repository) DeletePeopleByID(ctx context.Context, peopleID string) error {
	result, err := r.db.ExecContext(ctx, `DELETE FROM people WHERE id=$1`, peopleID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}
