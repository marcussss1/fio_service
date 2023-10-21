package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/marcussss1/fio_service/internal/models"
)

func (r repository) GetPeople(ctx context.Context, req models.GetPeopleRequest) ([]models.People, error) {
	var people []models.People

	builder := squirrel.Select("*").
		Limit(req.Limit).
		Offset(req.Offset).
		From("people")

	if req.Name != "" {
		builder = builder.Where(squirrel.Eq{"name": req.Name})
	}
	if req.Surname != "" {
		builder = builder.Where(squirrel.Eq{"surname": req.Surname})
	}
	if req.Patronymic != "" {
		builder = builder.Where(squirrel.Eq{"patronymic": req.Patronymic})
	}
	if req.StartAge != 0 && req.EndAge != 0 {
		builder = builder.Where(squirrel.GtOrEq{"age": req.StartAge})
		builder = builder.Where(squirrel.LtOrEq{"age": req.EndAge})
	}
	if req.Gender != "" {
		builder = builder.Where(squirrel.Eq{"gender": req.Gender})
	}
	if req.Nationality != "" {
		builder = builder.Where(squirrel.Eq{"nationality": req.Nationality})
	}

	query, args, err := builder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	err = r.db.SelectContext(ctx, &people, query, args...)
	if err != nil {
		return nil, err
	}

	return people, nil
}
