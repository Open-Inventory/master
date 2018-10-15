package controller

import (
	"app/common/session"
	"app/common/view"
	"net/http"
)

//AppGET displays the app
func AppGET(w http.ResponseWriter, r *http.Request) {
	// Get session
	sess := session.Instance(r)

	//userID := fmt.Sprintf("%s", sess.Values["id"])

	/*
		notes, err := model.NotesByUserID(userID)
		if err != nil {
			log.Println(err)
			notes = []model.Note{}
		}
	*/

	// Display the view
	v := view.New(r)
	v.Name = "app"
	v.Vars["first_name"] = sess.Values["first_name"]
	//v.Vars["notes"] = notes
	v.Render(w)
}
