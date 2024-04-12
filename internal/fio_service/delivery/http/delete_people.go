package http

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary		Delete people
// @Tags			People
// @Description	Delete people
// @Param			peopleID	query		string	true	"peopleID"
// @Success		204			{object}	models.NoContentResponse
// @Failure		400			{object}	error
// @Failure		500			{object}	error
// @Router			/api/v1/people [delete]
func (h handler) DeletePeopleHandler(ctx echo.Context) error {
	peopleID := ctx.QueryParam("peopleID")

	err := h.fioUsecase.DeletePeopleByID(ctx.Request().Context(), peopleID)
	if err != nil {
		return fmt.Errorf("fioUsecase.DeletePeopleByID: %w", err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
