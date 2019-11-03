package CO2Handlers

import "net/http"

func IssuesHandler (w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
}
