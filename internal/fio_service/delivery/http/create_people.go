package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcussss1/fio_service/internal/models"
)

// @Summary		Create people
// @Tags			People
// @Description	Create people
// @Produce		json
// @Param			req	body		models.AbbreviatedPeopleRequest	true	"Create people"
// @Success		201	{object}	models.People
// @Failure		400	{object}	error
// @Failure		500	{object}	error
// @Router			/api/v1/people [post]
func (h handler) CreatePeopleHandler(ctx echo.Context) error {
	var req models.AbbreviatedPeopleRequest

	err := ctx.Bind(&req)
	if err != nil {
		return fmt.Errorf("ctx.Bind: %w", err)
	}

	people, err := h.fioUsecase.CreatePeople(ctx.Request().Context(), req)
	if err != nil {
		return fmt.Errorf("fioUsecase.CreatePeople: %w", err)
	}

	return ctx.JSON(http.StatusCreated, people)
}
