package http

import (
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/fio_service/internal/models"
	"net/http"
)

func (h handler) GetPeopleHandler(ctx echo.Context) error {
	var req models.GetPeopleRequest

	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	people, err := h.fioUsecase.GetPeople(ctx.Request().Context(), req)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, people)
}
