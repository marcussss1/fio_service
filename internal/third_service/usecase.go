package third_service

import (
	"context"
	"github.com/marcussss1/fio_service/internal/models"
)

type Usecase interface {
	GetFullPeopleInfo(ctx context.Context, req models.AbbreviatedPeopleRequest) (models.People, error)
	GetPeopleAge(ctx context.Context, name string) (int, error)
	GetPeopleGender(ctx context.Context, name string) (string, error)
	GetPeopleNationality(ctx context.Context, name string) (string, error)
}
