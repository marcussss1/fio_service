package usecase

import (
	"context"
	"fmt"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/validation"
)

func (u usecase) CreatePeople(ctx context.Context, req models.AbbreviatedPeopleRequest) (models.People, error) {
	err := validation.ValidateAbbreviatedPeopleRequest(req)
	if err != nil {
		return models.People{}, fmt.Errorf("validation.ValidateAbbreviatedPeopleRequest: %w", err)
	}

	people, err := u.thirdUsecase.GetFullPeopleInfo(ctx, req)
	if err != nil {
		return models.People{}, fmt.Errorf("thirdUsecase.GetFullPeopleInfo: %w", err)
	}

	people, err = u.fioRepository.CreatePeople(ctx, people)
	if err != nil {
		return models.People{}, fmt.Errorf("fioRepository.CreatePeople: %w", err)
	}

	return people, nil
}
