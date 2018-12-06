package model

import "app/common/litedb"

type OrderEntry struct {
	ID       int    `db:"OrderID"`
	Products string `db:"ProductName"`
	Date     int    `db:"Order_Date"`
	Status   int    `db:"Order_Status"`
	Quantity int    `db:"Order_Quantity"`
}

func GetOrders() ([]OrderEntry, error) {
	var orders []OrderEntry
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return orders, err
	}

	query := `SELECT * FROM Orders`
	err = db.Select(&orders, query)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
