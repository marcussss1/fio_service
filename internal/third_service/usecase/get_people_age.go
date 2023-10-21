package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/marcussss1/fio_service/internal/config"
	"github.com/marcussss1/fio_service/internal/models"
	"github.com/marcussss1/fio_service/internal/pkg/utils"
)

func (u usecase) GetPeopleAge(ctx context.Context, name string) (int, error) {
	resp, err := utils.SendGetRequestToThirdService(ctx, name, config.GetAgeUrl)
	if err != nil {
		return 0, fmt.Errorf("sendGetRequestToThirdService: %w", err)
	}

	var thirdServiceAgeResponse models.ThirdServiceAgeResponse
	err = json.NewDecoder(resp.Body).Decode(&thirdServiceAgeResponse)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return thirdServiceAgeResponse.Age, nil
}
