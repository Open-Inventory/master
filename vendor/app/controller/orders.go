package controller

import (
	"app/common/session"
	"app/common/view"
	"app/model"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/josephspurrier/csrfbanana"
)

// OrdersGET displays the home page
func OrdersGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)

	if session.Values["id"] != nil {
		// Display the view
		v := view.New(r)
		v.Name = "index/orders"
		v.Vars["first_name"] = session.Values["first_name"]
		v.Render(w)
	} else {
		// Display the view
		v := view.New(r)
		v.Name = "index/anon"
		v.Vars["token"] = csrfbanana.Token(w, r, session)
		v.Render(w)
		return
	}
}

func GetOrders(w http.ResponseWriter, r *http.Request) {
	session := session.Instance(r)

	if session.Values["id"] != nil {
		response, err := model.GetOrders()
		if err != nil {
			log.Println("Error: " + err.Error())
			Error500(w, r)
		}
		data, _ := json.Marshal(response)
		w.Write(data)
	} else {
		Error500(w, r)
	}

}

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)
	if session.Values["id"] != nil {
		queryValues := r.URL.Query()

		id := queryValues.Get("id")
		idint, _ := strconv.Atoi(id)

		err := model.DeleteOrder(idint)
		if err != nil {
			log.Println("Unuccessfully deleted item")
		} else {
			log.Println("Successfully deleted item")
		}
	} else {
		Error500(w, r)
	}
}
