package fio_service

import (
	"context"

	"github.com/marcussss1/fio_service/internal/models"
)

type Usecase interface {
	GetPeople(ctx context.Context, req models.GetPeopleRequest) ([]models.People, error)
	CreatePeople(ctx context.Context, req models.AbbreviatedPeopleRequest) (models.People, error)
	UpdatePeopleByID(ctx context.Context, req models.AbbreviatedPeopleRequest, peopleID string) (models.People, error)
	DeletePeopleByID(ctx context.Context, peopleID string) error
}
