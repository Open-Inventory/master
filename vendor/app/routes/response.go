package route

import "net/http"

func WriteError(w http.ResponseWriter, description string) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500: " + description))
}
