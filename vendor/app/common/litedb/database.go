package litedb

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const dbpath = "./Open_Inventory.db"

func Connect() (*sqlx.DB, error) {
	var err error
	if db, err := sqlx.Connect("sqlite3", dbpath); err == nil {
		return db, err
	}

	log.Panic(err.Error())
	return nil, err
}
