package model

import "app/common/litedb"

type OrderEntry struct {
	ID       int    `db:"OrderID"`
	Date     string `db:"Order_Date"`
	Product  string `db:"ProductName"`
	ItemID   int    `db:"ItemID"`
	Status   int    `db:"Order_Status"`
	Quantity int    `db:"Order_Quantity"`
}

type OrderQuantity struct {
	Id       int `db:"OrderID"`
	Quantity int `db:"quantity"`
}

func GetOrders() ([]OrderEntry, error) {
	var orders []OrderEntry
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return orders, err
	}

	// query := `SELECT Orders.OrderID, Orders.ProductName, Orders.ItemID, Orders.Order_Date, Orders.Order_Status, Orders.Order_Quantity
	// FROM Orders
	// INNER JOIN Item ON Orders.ItemID=Item.ItemID`
	query := `SELECT Orders.OrderID, Product.ProductName, Item.ItemID, Orders.Order_Date, Orders.Order_Status, Orders.Order_Quantity 
	FROM Item
	INNER JOIN Orders ON Orders.OrderID=Item.OrderID
	INNER JOIN Product ON Item.ProductKey=Product.ProductKey
	ORDER BY Orders.OrderID`

	err = db.Select(&orders, query)
	if err != nil {
		return orders, err
	}
	return orders, nil
}

func CreateOrder(status int, quantity int, date string) (int64, error) {
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO Orders (Order_Status, Order_Quantity, Order_Date) VALUES ($1, $2, $3)`
	res, err := db.Exec(query, status, quantity, date)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func DeleteOrder(id int) error {
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return err
	}

	query := `DELETE FROM orders WHERE OrderID=$1;`
	_ = db.MustExec(query, id)
	return err
}

func OrderQuantities() ([]OrderQuantity, error) {
	var orders []OrderQuantity
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return orders, err
	}

	query := `SELECT OrderID, (SELECT IFNULL(COUNT(*),0) FROM Item i WHERE i.OrderID=o.OrderID GROUP BY OrderID) as quantity FROM Orders o`
	err = db.Select(&orders, query)
	if err != nil {
		return orders, err
	}
	return orders, nil
}
