package repository

import (
	"database/sql"
)

type MainRepository struct {
	DB *sql.DB
}