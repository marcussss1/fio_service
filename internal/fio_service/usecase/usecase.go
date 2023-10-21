package usecase

import (
	desc "github.com/marcussss1/fio_service/internal/fio_service"
	"github.com/marcussss1/fio_service/internal/third_service"
)

type usecase struct {
	fioRepository desc.Repository
	thirdUsecase  third_service.Usecase
}

func NewFioUsecase(fioRepository desc.Repository, thirdUsecase third_service.Usecase) desc.Usecase {
	return usecase{fioRepository: fioRepository, thirdUsecase: thirdUsecase}
}
