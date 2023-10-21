package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h handler) DeletePeopleHandler(ctx echo.Context) error {
	peopleID := ctx.Param("peopleID")

	err := h.fioUsecase.DeletePeopleByID(ctx.Request().Context(), peopleID)
	if err != nil {
		return fmt.Errorf("fioUsecase.DeletePeopleByID: %w", err)
	}

	return ctx.NoContent(http.StatusNoContent)
}
