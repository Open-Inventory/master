package model

import (
	"app/common/litedb"
)

type Item struct {
	Id         int    `db:"ItemID"`
	ProductKey int    `db:"ProductKey"`
	ItemStatus int    `db:"ItemStatus"`
	Created    string `db:Created_on`
	Updated    string `db:Updated_on`
}

func CreateItem(productKey int, status int) (int64, error) {
	var err error
	db, err := litedb.Connect()
	defer db.Close()
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

	query := `SELECT Item.ItemID, Product.ProductName, Category.CategoryType, SubCategory.SubCategoryType, Item.ItemStatus, Item.Created_On, Item.Updated_On FROM Item INNER JOIN Product ON Product.ProductKey=Item.ProductKey INNER JOIN Category ON Category.CategoryID=Product.CategoryID INNER JOIN SubCategory ON SubCategory.SubCategoryID=Product.SubCategoryID WHERE ItemID=$1`
	err = db.Select(&item, query, id)
	if err != nil {
		return item, err
	}
	return item, nil
}

func UpdateItemOrderId(id int, orderId int) error {
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return err
	}

	query := `UPDATE item SET OrderID=$1 WHERE ItemID=$2;`
	_ = db.MustExec(query, orderId, id)
	return err
}

func DeleteItem(id int) error {
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return err
	}

	query := `DELETE FROM item WHERE ItemID=$1;`
	_ = db.MustExec(query, id)
	return err
}

func UpdateOrderIdTo0(orderId int) error {
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return err
	}

	query := `UPDATE item SET OrderID=0 WHERE OrderID=$1;`
	_ = db.MustExec(query, orderId)
	return err
}
