package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	mock_fio_service "github.com/marcussss1/fio_service/internal/mocks/fio_service"
	mock_third_service "github.com/marcussss1/fio_service/internal/mocks/third_service"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/e"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUsecase_CreatePeople_Validation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	thirdUsecase := mock_third_service.NewMockUsecase(ctrl)
	fioRepository := mock_fio_service.NewMockRepository(ctrl)
	fioUsecase := NewFioUsecase(fioRepository, thirdUsecase)

	tests := []struct {
		expectedError error
		req           models.AbbreviatedPeopleRequest
		name          string
	}{
		{
			expectedError: e.ErrLongName,
			req: models.AbbreviatedPeopleRequest{
				Name:       "IvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvanIvan",
				Surname:    "Ivanov",
				Patronymic: "Ivanovich",
			},
			name: "длинное имя",
		},
		{
			expectedError: e.ErrLongSurname,
			req: models.AbbreviatedPeopleRequest{
				Name:       "Ivan",
				Surname:    "IvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanovIvanov",
				Patronymic: "Ivanovich",
			},
			name: "длинная фамилия",
		},
		{
			expectedError: e.ErrLongPatronymic,
			req: models.AbbreviatedPeopleRequest{
				Name:       "Ivan",
				Surname:    "Ivanov",
				Patronymic: "IvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovich",
			},
			name: "длинное отчество",
		},
	}

	for _, test := range tests {
		_, err := fioUsecase.CreatePeople(context.TODO(), test.req)
		require.ErrorIs(t, err, test.expectedError, test.name)
	}
}

func TestUsecase_CreatePeople_GetFullPeopleInfo_Error(t *testing.T) {
	var (
		expectedError = errors.New("ошибка похода в сторонний сервис")
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	thirdUsecase := mock_third_service.NewMockUsecase(ctrl)
	fioRepository := mock_fio_service.NewMockRepository(ctrl)
	fioUsecase := NewFioUsecase(fioRepository, thirdUsecase)

	thirdUsecase.EXPECT().
		GetFullPeopleInfo(gomock.Any(), models.AbbreviatedPeopleRequest{
			Name:       "Ivan",
			Surname:    "Ivanov",
			Patronymic: "Ivanovich",
		}).
		Return(models.People{}, expectedError).
		Times(1)

	_, err := fioUsecase.CreatePeople(context.TODO(), models.AbbreviatedPeopleRequest{
		Name:       "Ivan",
		Surname:    "Ivanov",
		Patronymic: "Ivanovich",
	})
	require.ErrorIs(t, err, expectedError)
}

func TestUsecase_CreatePeople_CreatePeopleRepository_Error(t *testing.T) {
	var (
		expectedError = errors.New("ошибка БД")
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	thirdUsecase := mock_third_service.NewMockUsecase(ctrl)
	fioRepository := mock_fio_service.NewMockRepository(ctrl)
	fioUsecase := NewFioUsecase(fioRepository, thirdUsecase)

	thirdUsecase.EXPECT().
		GetFullPeopleInfo(gomock.Any(), models.AbbreviatedPeopleRequest{
			Name:       "Ivan",
			Surname:    "Ivanov",
			Patronymic: "Ivanovich",
		}).
		Return(models.People{
			Name:        "Ivan",
			Surname:     "Ivanov",
			Patronymic:  "Ivanovich",
			Age:         55,
			Gender:      "male",
			Nationality: "RU",
		}, nil).
		Times(1)

	fioRepository.EXPECT().
		CreatePeople(gomock.Any(), models.People{
			Name:        "Ivan",
			Surname:     "Ivanov",
			Patronymic:  "Ivanovich",
			Age:         55,
			Gender:      "male",
			Nationality: "RU",
		}).
		Return(models.People{}, expectedError).
		Times(1)

	_, err := fioUsecase.CreatePeople(context.TODO(), models.AbbreviatedPeopleRequest{
		Name:       "Ivan",
		Surname:    "Ivanov",
		Patronymic: "Ivanovich",
	})
	require.ErrorIs(t, err, expectedError)
}

func TestUsecase_CreatePeople_OK(t *testing.T) {
	var (
		id = uuid.NewString()
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	thirdUsecase := mock_third_service.NewMockUsecase(ctrl)
	fioRepository := mock_fio_service.NewMockRepository(ctrl)
	fioUsecase := NewFioUsecase(fioRepository, thirdUsecase)

	thirdUsecase.EXPECT().
		GetFullPeopleInfo(gomock.Any(), models.AbbreviatedPeopleRequest{
			Name:       "Ivan",
			Surname:    "Ivanov",
			Patronymic: "Ivanovich",
		}).
		Return(models.People{
			Name:        "Ivan",
			Surname:     "Ivanov",
			Patronymic:  "Ivanovich",
			Age:         55,
			Gender:      "male",
			Nationality: "RU",
		}, nil).
		Times(1)

	fioRepository.EXPECT().
		CreatePeople(gomock.Any(), models.People{
			Name:        "Ivan",
			Surname:     "Ivanov",
			Patronymic:  "Ivanovich",
			Age:         55,
			Gender:      "male",
			Nationality: "RU",
		}).
		Return(models.People{
			ID:          id,
			Name:        "Ivan",
			Surname:     "Ivanov",
			Patronymic:  "Ivanovich",
			Age:         55,
			Gender:      "male",
			Nationality: "RU",
		}, nil).
		Times(1)

	people, err := fioUsecase.CreatePeople(context.TODO(), models.AbbreviatedPeopleRequest{
		Name:       "Ivan",
		Surname:    "Ivanov",
		Patronymic: "Ivanovich",
	})
	require.NoError(t, err)
	require.EqualValues(t, id, people.ID)
	require.EqualValues(t, "Ivan", people.Name)
	require.EqualValues(t, "Ivanov", people.Surname)
	require.EqualValues(t, "Ivanovich", people.Patronymic)
	require.EqualValues(t, 55, people.Age)
	require.EqualValues(t, "male", people.Gender)
	require.EqualValues(t, "RU", people.Nationality)
}
