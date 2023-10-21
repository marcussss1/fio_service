package http_utils

import (
	"errors"
	"github.com/marcussss1/fio_service/internal/pkg/e"
	"net/http"
)

func StatusCode(err error) int {
	switch {
	case errors.Is(err, e.ErrLongName):
		return http.StatusBadRequest
	case errors.Is(err, e.ErrLongSurname):
		return http.StatusBadRequest
	case errors.Is(err, e.ErrLongPatronymic):
		return http.StatusBadRequest
	case errors.Is(err, e.ErrInvalidAge):
		return http.StatusBadRequest
	case errors.Is(err, e.ErrInvalidGender):
		return http.StatusBadRequest
	case errors.Is(err, e.ErrLongNationality):
		return http.StatusBadRequest
	case errors.Is(err, e.ErrInvalidPeopleID):
		return http.StatusBadRequest
	case errors.Is(err, e.ErrPeopleNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
