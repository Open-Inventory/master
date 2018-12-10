package model

import (
	"app/common/litedb"
)

type Category struct {
	ID   int    `db:"CategoryID"`
	Type string `db:"CategoryType"`
}

type SubCategory struct {
	ID   int    `db:"SubCategoryID"`
	Type string `db:"SubCategoryType"`
}

func CreateCategory(categoryType string) (int64, error) {
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO Category (CategoryType) VALUES ($1)`
	res, err := db.Exec(query, categoryType)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetCategoryById(id int) (Category, error) {
	var category Category
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return category, err
	}

	query := `SELECT * FROM Category WHERE CategoryId=$1 LIMIT 1`
	err = db.Select(&category, query, id)
	if err != nil {
		return category, err
	}
	return category, nil
}

func GetCategoryByType(catType string) (Category, error) {
	var category Category
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return category, err
	}

	query := `SELECT * FROM Category WHERE CategoryType LIKE $1 LIMIT 1`
	err = db.Select(&category, query, catType)
	if err != nil {
		return category, err
	}
	return category, nil
}

func CreateSubCategory(categoryType string) (int64, error) {
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO SubCategory (SubCategoryType) VALUES ($1)`
	res, err := db.Exec(query, categoryType)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func GetSubCategoryById(id int) (SubCategory, error) {
	var subcategory SubCategory
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return subcategory, err
	}

	query := `SELECT * FROM SubCategory WHERE SubCategoryId=$1 LIMIT 1`
	err = db.Select(&subcategory, query, id)
	if err != nil {
		return subcategory, err
	}
	return subcategory, nil
}

func GetSubCategoryByType(subType string) (SubCategory, error) {
	var subcategory SubCategory
	var err error
	db, err := litedb.Connect()
	defer db.Close()
	if err != nil {
		return subcategory, err
	}

	query := `SELECT * FROM SubCategory WHERE SubCategoryType LIKE $1 LIMIT 1`
	err = db.Select(&subcategory, query, subType)
	if err != nil {
		return subcategory, err
	}
	return subcategory, nil
}
