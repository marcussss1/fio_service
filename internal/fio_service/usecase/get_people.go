package usecase

import (
	"context"
	"fmt"

	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/validation"
)

func (u usecase) GetPeople(ctx context.Context, req models.GetPeopleRequest) ([]models.People, error) {
	err := validation.ValidateGetPeopleRequest(req)
	if err != nil {
		return nil, fmt.Errorf("validation.ValidateGetPeopleRequest: %w", err)
	}

	if req.Limit == 0 {
		req.Limit = 100
	}

	people, err := u.fioRepository.GetPeople(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("GetPeople: %w", err)
	}

	return people, nil
}
