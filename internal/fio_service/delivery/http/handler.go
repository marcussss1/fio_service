package http

import (
	"github.com/labstack/echo/v4"
	"github.com/marcussss1/fio_service/internal/fio_service"
)

type handler struct {
	fioUsecase fio_service.Usecase
}

func NewFioHandler(e *echo.Echo, fioUsecase fio_service.Usecase) handler {
	h := handler{fioUsecase: fioUsecase}

	e.GET("api/v1/people", h.GetPeopleHandler)
	e.POST("api/v1/people", h.CreatePeopleHandler)
	e.PUT("api/v1/people", h.UpdatePeopleHandler)
	e.DELETE("api/v1/people", h.DeletePeopleHandler)

	return h
}
