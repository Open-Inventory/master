package model

import (
	"app/common/litedb"
)

type Item struct {
	Id         int `db:"id"`
	ProductKey int `db:"ProductKey"`
	ItemStatus int `db:"ItemStatus"`
}

func CreateItem(name string, productKey int, status string) (int64, error) {
	var err error
	db, err := litedb.Connect()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO Item (ProductKey, ItemStatus) VALUES ($1, $2)`
	res, err := db.Exec(query, productKey, status)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetItemById(id int) (Item, error) {
	var item Item
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return item, err
	}

	query := `SELECT * FROM Item WHERE ItemId=$1 LIMIT 1`
	err = db.Select(&item, query, id)
	if err != nil {
		return item, err
	}
	return item, nil
}
