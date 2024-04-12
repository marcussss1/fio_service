package fio_service

import (
	"context"

	"github.com/marcussss1/fio_service/internal/models"
)

type Repository interface {
	GetPeople(ctx context.Context, req models.GetPeopleRequest) ([]models.People, error)
	CreatePeople(ctx context.Context, people models.People) (models.People, error)
	GetPeopleByID(ctx context.Context, peopleID string) (models.People, error)
	UpdatePeopleByID(ctx context.Context, req models.People, peopleID string) (models.People, error)
	DeletePeopleByID(ctx context.Context, peopleID string) error
}
