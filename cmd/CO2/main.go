package main

import (
	"github.com/NTNU-sondrbaa-2019/CLOUD-O2/pkg/CO2Handlers"
	"log"
	"net/http"
	"os"
)

func main () {

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("Found no port!")
		return
	}

	http.HandleFunc(COMMITS_PATH, CO2Handlers.CommitsHandler)
	http.HandleFunc(LANGUAGES_PATH, CO2Handlers.LanguagesHandler)
	http.HandleFunc(ISSUES_PATH, CO2Handlers.IssuesHandler)
	http.HandleFunc(STATUS_PATH, CO2Handlers.StatusHandler)
	
	http.ListenAndServe(":" + port, nil)
}