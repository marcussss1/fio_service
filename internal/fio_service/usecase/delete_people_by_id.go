package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/marcussss1/fio_service/internal/pkg/e"
	"github.com/marcussss1/fio_service/internal/pkg/validation"
)

func (u usecase) DeletePeopleByID(ctx context.Context, peopleID string) error {
	err := validation.IsUUID(peopleID)
	if err != nil {
		return err
	}

	err = u.fioRepository.DeletePeopleByID(ctx, peopleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("fioRepository.DeletePeopleByID: %w", e.ErrPeopleNotFound)
		}

		return fmt.Errorf("fioRepository.DeletePeopleByID: %w", err)
	}

	return nil
}
