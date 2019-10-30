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
		port = "8080"
	}

	log.Printf("Initializing handlers...")

	http.HandleFunc(COMMITS_PATH, CO2Handlers.CommitsHandler)
	http.HandleFunc(LANGUAGES_PATH, CO2Handlers.LanguagesHandler)
	http.HandleFunc(ISSUES_PATH, CO2Handlers.IssuesHandler)
	http.HandleFunc(STATUS_PATH, CO2Handlers.StatusHandler)

	log.Printf("Running server...")

	http.ListenAndServe(":" + port, nil)

}
