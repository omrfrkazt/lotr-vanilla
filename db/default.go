package db

import (
	"database/sql"
)

type DBModel struct {
	DB *sql.DB
}

//NewModels returns model with dbpool
func NewModels(db *sql.DB) DBModel {
	return DBModel{
		DB: db,
	}
}

