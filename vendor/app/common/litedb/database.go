package litedb

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

const dbpath = "./OpenInventory.db"

func Connect() (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", "./OpenInventory.db")
	if err != nil {
		return nil, err
	}
	return db, nil

}
