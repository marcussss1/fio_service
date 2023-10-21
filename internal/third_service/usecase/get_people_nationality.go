package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/marcussss1/fio_service/internal/config"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/utils"
	"golang.org/x/net/context"
)

func (u usecase) GetPeopleNationality(ctx context.Context, name string) (string, error) {
	resp, err := utils.SendGetRequestToThirdService(ctx, name, config.GetNationalityUrl)
	if err != nil {
		return "", fmt.Errorf("sendGetRequestToThirdService: %w", err)
	}

	var thirdServiceGenderResponse models.ThirdServiceNationalityResponse
	err = json.NewDecoder(resp.Body).Decode(&thirdServiceGenderResponse)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return utils.MostLikelyNationality(thirdServiceGenderResponse.Countries), nil
}
