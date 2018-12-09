package controller

import (
	"app/common/session"
	"app/model"
	"log"
	"net/http"
	"strconv"
)

type InventoryItem struct {
	Id         int    `json:"Id"`
	Name       string `json:"Name"`
	SubType    string `json:"SubType"`
	Type       string `json:"Type"`
	ItemStatus int    `json:"ItemStatus"`
	Order      int    `json:"Order"`
}

//CreateItem Creates a new item in the DB
func CreateItem(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)

	queryValues := r.URL.Query()
	status, _ := strconv.Atoi(queryValues.Get("status"))
	orderId, _ := strconv.Atoi(queryValues.Get("order"))
	id := queryValues.Get("id")
	idint, _ := strconv.Atoi(id)

	item := InventoryItem{
		Id:         idint,
		Name:       queryValues.Get("name"),
		SubType:    queryValues.Get("subtype"),
		Type:       queryValues.Get("type"),
		ItemStatus: status,
		Order:      orderId,
	}

	if session.Values["id"] != nil {
		if id == "N/A" {
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
			model.UpdateItemOrderId(item.Id, item.Order)

			log.Println(item)
		} else {
			log.Println(item)
			err := model.UpdateItemOrderId(item.Id, item.Order)
			if err != nil {
				log.Println("Unuccessfully updated order ID")
			} else {
				log.Println("Successfully updated order ID")
			}
		}
	} else {
		Error500(w, r)
	}
}
