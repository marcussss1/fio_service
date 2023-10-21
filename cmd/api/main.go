package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	_ "github.com/marcussss1/fio_service/docs"
	"github.com/marcussss1/fio_service/internal/config"
	fio_delivery "github.com/marcussss1/fio_service/internal/fio_service/delivery/http"
	"github.com/marcussss1/fio_service/internal/fio_service/repository"
	fio_usecase "github.com/marcussss1/fio_service/internal/fio_service/usecase"
	"github.com/marcussss1/fio_service/internal/middleware"
	third_usecase "github.com/marcussss1/fio_service/internal/third_service/usecase"
	log "github.com/sirupsen/logrus"
	echo_swagger "github.com/swaggo/echo-swagger"
)

func init() {
	envPath := ".env"
	if err := godotenv.Load(envPath); err != nil {
		log.Fatal(".env не найден")
	}
}

// @title			FIO API
// @version		1.0.1
// @description	Server API for FIO Service Application
// @contact.name	FIO API Support
// @contact.email	danilakalash60@gmail.com
// @host			localhost:8080
// @BasePath		/
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
	e.GET("/docs/*", echo_swagger.WrapHandler)

	fioRepository := repository.NewFioRepository(db)
	thirdUsecase := third_usecase.NewThirdUsecase()
	fioUsecase := fio_usecase.NewFioUsecase(fioRepository, thirdUsecase)
	fio_delivery.NewFioHandler(e, fioUsecase)

	e.Logger.Fatal(e.Start(config.Server.Port))
}
