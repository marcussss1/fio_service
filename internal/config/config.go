package config

import (
	"os"

	"github.com/labstack/gommon/log"
)

var (
	GetAgeUrl         = "https://api.agify.io/?"
	GetGenderUrl      = "https://api.genderize.io/?"
	GetNationalityUrl = "https://api.nationalize.io/?"
)

type Config struct {
	Server   Server   `yaml:"Server"`
	Postgres Postgres `yaml:"Postgres"`
}

type Server struct {
	Port string `yaml:"port"`
}

type Postgres struct {
	DB             string `yaml:"db"`
	ConnectionToDB string `yaml:"connectionToDB"`
}

func NewConfig() Config {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		log.Fatal("порт сервера в .env не найден")
	}

	db, exists := os.LookupEnv("DB")
	if !exists {
		log.Fatal("БД в .env не найден")
	}

	dsn, exists := os.LookupEnv("CONNECTION_TO_DB")
	if !exists {
		log.Fatal("dsn в .env не найден")
	}

	return Config{
		Server: Server{
			Port: port,
		},
		Postgres: Postgres{
			DB:             db,
			ConnectionToDB: dsn,
		},
	}
}
