package model

import "app/common/litedb"

type InventoryItem struct {
	Id         int    `db:"ItemID"`
	Name       string `db:"ProductName"`
	SubType    string `db:"SubCategoryType"`
	Type       string `db:"CategoryType"`
	ItemStatus int    `db:"ItemStatus"`
	Order      int    `db:"OrderID"`
	Created    string `db:"Created_On"`
	Updated    string `db:"Updated_On"`
}

func GetInventory() ([]InventoryItem, error) {
	var items []InventoryItem
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return items, err
	}

	query := `SELECT Item.ItemID, Product.ProductName, Category.CategoryType, SubCategory.SubCategoryType, Item.ItemStatus, Item.OrderID, Item.Created_On, Item.Updated_On FROM Item INNER JOIN Product ON Product.ProductKey=Item.ProductKey INNER JOIN Category ON Category.CategoryID=Product.CategoryID INNER JOIN SubCategory ON SubCategory.SubCategoryID=Product.SubCategoryID`
	err = db.Select(&items, query)
	if err != nil {
		return items, err
	}
	return items, nil
}
