package rest

import (
	"fmt"
	"net/http"
)

func HandleGetNear(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Always set content type and status code
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "Hello World!")
}
