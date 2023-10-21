package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/marcussss1/fio_service/internal/models"
)

func (r repository) UpdatePeopleByID(ctx context.Context, req models.People, peopleID string) (models.People, error) {
	var people models.People

	builder := squirrel.Update("people").
		Where(squirrel.Eq{"id": peopleID}).
		Set("name", req.Name).
		Set("surname", req.Surname).
		Set("patronymic", req.Patronymic).
		Set("age", req.Age).
		Set("gender", req.Gender).
		Set("nationality", req.Nationality).
		Suffix(`RETURNING *`)

	query, args, err := builder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return models.People{}, err
	}

	err = r.db.GetContext(ctx, &people, query, args...)
	if err != nil {
		return models.People{}, err
	}

	return people, nil
}
