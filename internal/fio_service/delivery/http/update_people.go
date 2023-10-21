package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/marcussss1/fio_service/internal/models"
)

// @Summary		Update people
// @Tags			People
// @Description	Update people
// @Produce		json
// @Param			peopleID	query		string							true	"peopleID"
// @Param			req			body		models.AbbreviatedPeopleRequest	true	"New people info"
// @Success		200			{object}	models.People
// @Failure		400			{object}	error
// @Failure		500			{object}	error
// @Router			/api/v1/people [put]
func (h handler) UpdatePeopleHandler(ctx echo.Context) error {
	var req models.AbbreviatedPeopleRequest
	peopleID := ctx.QueryParam("peopleID")

	err := ctx.Bind(&req)
	if err != nil {
		return fmt.Errorf("ctx.Bind: %w", err)
	}

	people, err := h.fioUsecase.UpdatePeopleByID(ctx.Request().Context(), req, peopleID)
	if err != nil {
		return fmt.Errorf("fioUsecase.UpdatePeopleByID: %w", err)
	}

	return ctx.JSON(http.StatusOK, people)
}
