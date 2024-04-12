package usecase

import (
	desc "github.com/marcussss1/fio_service/internal/third_service"
)

type usecase struct{}

func NewThirdUsecase() desc.Usecase {
	return usecase{}
}
