package http

import (
	"bytes"
	"encoding/json"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	mock_fio_service "github.com/marcussss1/fio_service/internal/mocks/fio_service"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/http_utils"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"

	bussines_errors "github.com/marcussss1/fio_service/internal/pkg/e"
)

func TestHandler_UpdatePeopleHandler_InvalidJSON(t *testing.T) {
	var (
		id = uuid.NewString()
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	r := httptest.NewRequest(http.MethodPut, "/api/v1/people", bytes.NewReader([]byte(`
		{"name":"Ivan","surname":"Ivanov", ovich"""""""}
	`)))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	ctx.SetParamNames("peopleID")
	ctx.SetParamValues(id)

	mockFioUsecase := mock_fio_service.NewMockUsecase(ctrl)
	h := NewFioHandler(e, mockFioUsecase)

	err := h.UpdatePeopleHandler(ctx)
	require.Error(t, err)
	require.EqualValues(t, http.StatusInternalServerError, http_utils.StatusCode(err))
}

func TestHandler_UpdatePeopleHandler_UpdatePeopleByID_Error(t *testing.T) {
	var (
		id = uuid.NewString()
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	r := httptest.NewRequest(http.MethodPut, "/api/v1/people", bytes.NewReader([]byte(`
		{"name":"Ivan","surname":"Ivanov","patronymic":"IvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovich"}
	`)))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	ctx.SetParamNames("peopleID")
	ctx.SetParamValues(id)

	mockFioUsecase := mock_fio_service.NewMockUsecase(ctrl)
	h := NewFioHandler(e, mockFioUsecase)

	mockFioUsecase.EXPECT().
		UpdatePeopleByID(gomock.Any(), models.AbbreviatedPeopleRequest{
			Name:       "Ivan",
			Surname:    "Ivanov",
			Patronymic: "IvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovichIvanovich",
		}, id).
		Return(models.People{}, bussines_errors.ErrLongPatronymic).
		Times(1)

	err := h.UpdatePeopleHandler(ctx)
	require.ErrorIs(t, err, bussines_errors.ErrLongPatronymic)
	require.EqualValues(t, http.StatusBadRequest, http_utils.StatusCode(err))
}

func TestHandler_UpdatePeopleHandler_OK(t *testing.T) {
	var (
		id     = uuid.NewString()
		people models.People
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	e := echo.New()
	r := httptest.NewRequest(http.MethodPut, "/api/v1/people", bytes.NewReader([]byte(`
		{"name":"Ivan","surname":"Ivanov","patronymic":"Ivanovich"}
	`)))
	r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	w := httptest.NewRecorder()
	ctx := e.NewContext(r, w)
	ctx.SetParamNames("peopleID")
	ctx.SetParamValues(id)

	mockFioUsecase := mock_fio_service.NewMockUsecase(ctrl)
	h := NewFioHandler(e, mockFioUsecase)

	mockFioUsecase.EXPECT().
		UpdatePeopleByID(gomock.Any(), models.AbbreviatedPeopleRequest{
			Name:       "Ivan",
			Surname:    "Ivanov",
			Patronymic: "Ivanovich",
		}, id).
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

	err := h.UpdatePeopleHandler(ctx)
	require.NoError(t, err)
	require.EqualValues(t, http.StatusOK, w.Code)

	err = json.NewDecoder(w.Body).Decode(&people)
	require.NoError(t, err)

	require.EqualValues(t, id, people.ID)
	require.EqualValues(t, "Ivan", people.Name)
	require.EqualValues(t, "Ivanov", people.Surname)
	require.EqualValues(t, "Ivanovich", people.Patronymic)
	require.EqualValues(t, 55, people.Age)
	require.EqualValues(t, "male", people.Gender)
	require.EqualValues(t, "RU", people.Nationality)
}
