package repository

import (
	"context"

	"github.com/marcussss1/fio_service/internal/models"
)

func (r repository) GetPeopleByID(ctx context.Context, peopleID string) (models.People, error) {
	var people models.People

	err := r.db.GetContext(ctx, &people, `SELECT * FROM people WHERE id=$1`, peopleID)
	if err != nil {
		return models.People{}, err
	}

	return people, nil
}
