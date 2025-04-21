package repository

import (
	"database/sql"
)

type MainRepository struct {
	DB *sql.DB
}

func NewMainRepository(db *sql.DB) *MainRepository {
	return &MainRepository{
		DB: db,
	}
}