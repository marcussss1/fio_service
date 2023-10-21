package http

import (
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/fio_service/internal/models"
	"net/http"
)

func (h handler) UpdatePeopleHandler(ctx echo.Context) error {
	peopleID := ctx.Param("peopleID")
	var req models.AbbreviatedPeopleRequest

	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	people, err := h.fioUsecase.UpdatePeopleByID(ctx.Request().Context(), req, peopleID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, people)
}
