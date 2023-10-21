package utils

import (
	"context"
	"github.com/marcussss1/fio_service/internal/models"
	"net/http"
	"net/url"
	"time"
)

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
