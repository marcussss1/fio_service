package usecase

import (
	"context"
	"fmt"
	"github.com/marcussss1/fio_service/internal/models"
)

func (u usecase) GetFullPeopleInfo(ctx context.Context, req models.AbbreviatedPeopleRequest) (models.People, error) {
	age, err := u.GetPeopleAge(ctx, req.Name)
	if err != nil {
		return models.People{}, fmt.Errorf("getPeopleAge: %w", err)
	}

	gender, err := u.GetPeopleGender(ctx, req.Name)
	if err != nil {
		return models.People{}, fmt.Errorf("getPeopleGender: %w", err)
	}

	nationality, err := u.GetPeopleNationality(ctx, req.Name)
	if err != nil {
		return models.People{}, fmt.Errorf("getPeopleNationality: %w", err)
	}

	return models.People{
		Name:        req.Name,
		Surname:     req.Surname,
		Patronymic:  req.Patronymic,
		Age:         age,
		Gender:      gender,
		Nationality: nationality,
	}, nil
}
