package controller

import (
	"app/common/session"
	"app/model"
	"encoding/json"
	"log"
	"net/http"
)

type InventoryItem struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	SubType    string `json:"SubType"`
	Type       string `json:"Type"`
	ItemStatus int    `json:"ItemStatus"`
}

//CreateItem Creates a new item in the DB
func CreateItem(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)

	if session.Values["id"] != nil {
		decoder := json.NewDecoder(r.Body)

		var item InventoryItem
		err := decoder.Decode(&item)
		if err != nil {
			Error500(w, r)
		}

		var subCatId int
		subCat, err := model.GetSubCategoryByType(item.SubType)
		if err != nil {
			id, err := model.CreateSubCategory(item.SubType)
			if err != nil {
				Error500(w, r)
			}
			subCatId = int(id)
		} else {
			subCatId = subCat.ID
		}

		var catId int
		cat, err := model.GetCategoryByType(item.Type)
		if err != nil {
			id, err := model.CreateCategory(item.Type)
			if err != nil {
				Error500(w, r)
			}
			catId = int(id)
		} else {
			catId = cat.ID
		}

		var productKey int
		product, err := model.GetProductByType(item.Name)
		if err != nil {
			id, err := model.CreateProduct(item.Name, catId, subCatId)
			if err != nil {
				Error500(w, r)
			}
			productKey = int(id)
		} else {
			productKey = product.ID
		}

		model.CreateItem(productKey, 0)

		log.Println(item)

	} else {
		Error500(w, r)
	}
}
