package repository

import (
	"context"

	"github.com/marcussss1/fio_service/internal/models"
)

func (r repository) CreatePeople(ctx context.Context, people models.People) (models.People, error) {
	row, err := r.db.NamedQueryContext(ctx, `INSERT INTO people (name, surname, patronymic, age, gender, nationality) 
		VALUES (:name, :surname, :patronymic, :age, :gender, :nationality) RETURNING id`, &people)
	if err != nil {
		return models.People{}, err
	}

	row.Next()
	err = row.Scan(&people.ID)
	if err != nil {
		return models.People{}, err
	}

	return people, nil
}
