package utils

import (
	"context"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/e"
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

func SendGetRequestToThirdService(ctx context.Context, name, thirdServiceUrl string) (*http.Response, error) {
	getParams := url.Values{}
	getParams.Add("name", name)
	url := thirdServiceUrl + getParams.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{Timeout: time.Second * 5}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func MostLikelyNationality(countries []models.Country) string {
	var (
		nationality    string
		maxProbability float64
	)

	for _, country := range countries {
		if country.Probability > maxProbability {
			nationality = country.CountryID
			maxProbability = country.Probability
		}
	}

	return nationality
}

func CreateGetPeopleRequest(ctx echo.Context) (models.GetPeopleRequest, error) {
	// конкретно в этом кейсе ошибку можно не проверять
	startAge, _ := strconv.ParseInt(ctx.QueryParam("start_age"), 10, 32)
	endAge, _ := strconv.ParseInt(ctx.QueryParam("end_age"), 10, 32)
	limit, _ := strconv.ParseUint(ctx.QueryParam("limit"), 10, 64)
	offset, _ := strconv.ParseUint(ctx.QueryParam("offset"), 10, 64)

	return models.GetPeopleRequest{
		Name:        ctx.QueryParam("name"),
		Surname:     ctx.QueryParam("surname"),
		Patronymic:  ctx.QueryParam("patronymic"),
		StartAge:    int(startAge),
		EndAge:      int(endAge),
		Gender:      ctx.QueryParam("gender"),
		Nationality: ctx.QueryParam("nationality"),
		Limit:       limit,
		Offset:      offset,
	}, nil
}
