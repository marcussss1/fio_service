package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/marcussss1/fio_service/internal/fio_service"
)

type repository struct {
	db *sqlx.DB
}

func NewFioRepository(db *sqlx.DB) fio_service.Repository {
	return &repository{db: db}
}
