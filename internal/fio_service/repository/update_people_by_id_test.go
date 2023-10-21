package repository

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestRepository_UpdatePeopleByID_OK(t *testing.T) {
	var (
		id = uuid.NewString()
	)

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	row := sqlmock.NewRows([]string{"id", "name", "surname", "patronymic", "age", "gender", "nationality"}).
		AddRow(id, "Ivan", "Ivanov", "Ivanovich", 55, "male", "RU")

	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE people SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nationality = $6 WHERE id = $7 RETURNING *`)).
		WithArgs("Ivan", "Ivanov", "Ivanovich", 55, "male", "RU", id).
		WillReturnRows(row)

	dbx := sqlx.NewDb(db, "sqlmock")
	fioRepository := NewFioRepository(dbx)

	people, err := fioRepository.UpdatePeopleByID(context.TODO(), models.People{
		Name:        "Ivan",
		Surname:     "Ivanov",
		Patronymic:  "Ivanovich",
		Age:         55,
		Gender:      "male",
		Nationality: "RU",
	}, id)
	require.NoError(t, err)
	require.EqualValues(t, id, people.ID)
	require.EqualValues(t, "Ivan", people.Name)
	require.EqualValues(t, "Ivanov", people.Surname)
	require.EqualValues(t, "Ivanovich", people.Patronymic)
	require.EqualValues(t, 55, people.Age)
	require.EqualValues(t, "male", people.Gender)
	require.EqualValues(t, "RU", people.Nationality)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}

func TestRepository_UpdatePeopleByID_Error(t *testing.T) {
	var (
		id            = uuid.NewString()
		expectedError = errors.New("ошибка БД")
	)

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`UPDATE people SET name = $1, surname = $2, patronymic = $3, age = $4, gender = $5, nationality = $6 WHERE id = $7 RETURNING *`)).
		WithArgs("Ivan", "Ivanov", "Ivanovich", 55, "male", "RU", id).
		WillReturnError(expectedError)

	dbx := sqlx.NewDb(db, "sqlmock")
	fioRepository := NewFioRepository(dbx)

	_, err = fioRepository.UpdatePeopleByID(context.TODO(), models.People{
		Name:        "Ivan",
		Surname:     "Ivanov",
		Patronymic:  "Ivanovich",
		Age:         55,
		Gender:      "male",
		Nationality: "RU",
	}, id)
	require.Error(t, err)
	require.ErrorIs(t, err, expectedError)

	err = mock.ExpectationsWereMet()
	require.NoError(t, err)
}
