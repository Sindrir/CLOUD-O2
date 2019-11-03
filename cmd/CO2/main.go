package main

import (
	"github.com/Sindrir/CLOUD-O2/pkg/CO2Handlers"
	"log"
	"net/http"
	"os"
)

const COMMITS_PATH		= "/repocheck/v1/commits"
const LANGUAGES_PATH	= "/repocheck/v1/languages"
const ISSUES_PATH		= "/repocheck/v1/issues"
const STATUS_PATH		= "/repocheck/v1/status"

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
