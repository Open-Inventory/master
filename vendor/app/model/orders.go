package model

import "app/common/litedb"

type OrderEntry struct {
	ID       int    `db:"OrderID"`
	Products string `db:"ProductName"`
	ItemID   int    `db:"ItemID"`
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

	query := `SELECT Orders.OrderID, Orders.ProductName, Orders.ItemID, Orders.Order_Date, Orders.Order_Status, Orders.Order_Quantity 
	FROM Orders 
	INNER JOIN Item ON Orders.ItemID=Item.ItemID`
	err = db.Select(&orders, query)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
