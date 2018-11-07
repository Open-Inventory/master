package model

import (
	"app/common/litedb"
)

type Item struct {
	Id int `db:"id"`
}

func CreateItem(name, category, subcategory, location, status string) error {
	var err error
	db, err := litedb.Connect()
	if err != nil {
		return err
	}

	query := `INSERT INTO Product(ProductName, Category, 
}
