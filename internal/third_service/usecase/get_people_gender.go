package usecase

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/marcussss1/fio_service/internal/config"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/utils"
)

func (u usecase) GetPeopleGender(ctx context.Context, name string) (string, error) {
	resp, err := utils.SendGetRequestToThirdService(ctx, name, config.GetGenderUrl)
	if err != nil {
		return "", fmt.Errorf("sendGetRequestToThirdService: %w", err)
	}

	var thirdServiceGenderResponse models.ThirdServiceGenderResponse
	err = json.NewDecoder(resp.Body).Decode(&thirdServiceGenderResponse)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return thirdServiceGenderResponse.Gender, nil
}
