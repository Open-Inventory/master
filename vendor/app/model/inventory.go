package model

import "app/common/litedb"

func GetInventory() ([]Item, error) {
	var items []Item
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return items, err
	}

	query := `SELECT * FROM Item`
	err = db.Select(&items, query)
	if err != nil {
		return items, err
	}
	return items, nil
}
