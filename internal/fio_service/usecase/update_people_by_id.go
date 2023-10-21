package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/e"
	"github.com/marcussss1/fio_service/internal/pkg/validation"
)

func (u usecase) UpdatePeopleByID(ctx context.Context, req models.AbbreviatedPeopleRequest, peopleID string) (models.People, error) {
	err := validation.IsUUID(peopleID)
	if err != nil {
		return models.People{}, err
	}

	err = validation.ValidateAbbreviatedPeopleRequest(req)
	if err != nil {
		return models.People{}, fmt.Errorf("validation.ValidateAbbreviatedPeopleRequest: %w", err)
	}

	people, err := u.fioRepository.GetPeopleByID(ctx, peopleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.People{}, fmt.Errorf("fioRepository.GetPeopleByID: %w", e.ErrPeopleNotFound)
		}

		return models.People{}, fmt.Errorf("fioRepository.GetPeopleByID: %w", err)
	}

	if people.Name != req.Name {
		people, err = u.thirdUsecase.GetFullPeopleInfo(ctx, req)
		if err != nil {
			return models.People{}, fmt.Errorf("thirdUsecase.GetFullPeopleInfo: %w", err)
		}

		people, err = u.fioRepository.UpdatePeopleByID(ctx, people, peopleID)
		if err != nil {
			return models.People{}, fmt.Errorf("fioRepository.UpdatePeopleByID: %w", err)
		}
	}

	return people, nil
}
