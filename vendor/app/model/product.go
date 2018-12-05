package model

import "app/common/litedb"

type Product struct {
	ID          int    `db:"ProductKey"`
	Name        string `db:"ProductName"`
	Category    string `db:"CategoryID"`
	SubCategory string `db:'SubCategoryID"`
}

func CreateProduct(name string, catID int, subCatID int) (int64, error) {
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO Product (ProductName, CategoryID, SubCategoryID) VALUES ($1, $2, $3)`
	res, err := db.Exec(query, name, catID, subCatID)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetProductById(id int) (Product, error) {
	var product Product
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return product, err
	}

	query := `SELECT * FROM Product WHERE ProductKey=$1 LIMIT 1`
	err = db.Select(&product, query, id)
	if err != nil {
		return product, err
	}
	return product, nil
}

func GetProductByType(name string) (Product, error) {
	var product Product
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return product, err
	}

	query := `SELECT * FROM Product WHERE ProductName=$1 LIMIT 1`
	err = db.Select(&product, query, name)
	if err != nil {
		return product, err
	}
	return product, nil
}
