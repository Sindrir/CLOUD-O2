package CO2Handlers

import (
	"fmt"
	"net/http"
)

func CommitsHandler (w http.ResponseWriter, r *http.Request){
	limitParam := r.URL.Query().Get("limit")
	authParam := r.URL.Query().Get("auth")
	fmt.Print(limitParam)
	fmt.Print(authParam)
}
