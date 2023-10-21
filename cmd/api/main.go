package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	"github.com/marcussss1/fio_service/internal/config"
	fio_delivery "github.com/marcussss1/fio_service/internal/fio_service/delivery/http"
	"github.com/marcussss1/fio_service/internal/fio_service/repository"
	fio_usecase "github.com/marcussss1/fio_service/internal/fio_service/usecase"
	"github.com/marcussss1/fio_service/internal/middleware"
	third_usecase "github.com/marcussss1/fio_service/internal/third_service/usecase"
	log "github.com/sirupsen/logrus"
)

func init() {
	envPath := ".env"
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(".env не найден")
	}
}

func main() {
	config := config.NewConfig()

	db, err := sqlx.Open(config.Postgres.DB, config.Postgres.ConnectionToDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	e := echo.New()
	e.Use(middleware.LoggerMiddleware)

	fioRepository := repository.NewFioRepository(db)
	thirdUsecase := third_usecase.NewThirdUsecase()
	fioUsecase := fio_usecase.NewFioUsecase(fioRepository, thirdUsecase)
	fio_delivery.NewFioHandler(e, fioUsecase)

	e.Logger.Fatal(e.Start(config.Server.Port))
}
