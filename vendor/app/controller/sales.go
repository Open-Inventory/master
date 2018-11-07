package controller

import (
	"app/common/session"
	"app/common/view"
	"net/http"

	"github.com/josephspurrier/csrfbanana"
)

// SalesGET displays the home page
func SalesGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	session := session.Instance(r)

	if session.Values["id"] != nil {
		// Display the view
		v := view.New(r)
		v.Name = "index/sales"
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
