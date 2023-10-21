package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcussss1/fio_service/internal/pkg/utils"
)

// @Summary		Get people
// @Tags			People
// @Description	Get people
// @Produce		json
// @Param			name		query		string	false	"name"
// @Param			surname		query		string	false	"surname"
// @Param			patronymic	query		string	false	"patronymic"
// @Param			start_age	query		int		false	"start_age"
// @Param			end_age		query		int		false	"end_age"
// @Param			gender		query		string	false	"gender"
// @Param			nationality	query		string	false	"nationality"
// @Param			limit		query		uint64	false	"limit"
// @Param			offset		query		uint64	false	"offset"
// @Success		200			{array}		models.People
// @Failure		400			{object}	error
// @Failure		500			{object}	error
// @Router			/api/v1/people [get]
func (h handler) GetPeopleHandler(ctx echo.Context) error {
	req, err := utils.CreateGetPeopleRequest(ctx)
	if err != nil {
		return fmt.Errorf("utils.CreateGetPeopleRequest: %w", err)
	}

	people, err := h.fioUsecase.GetPeople(ctx.Request().Context(), req)
	if err != nil {
		return fmt.Errorf("fioUsecase.GetPeople: %w", err)
	}

	return ctx.JSON(http.StatusOK, people)
}
