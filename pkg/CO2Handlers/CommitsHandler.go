package CO2Handlers

import (
	"encoding/json"
	"github.com/NTNU-sondrbaa-2019/CLOUD-O2/pkg/CO2Struct"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func CommitsHandler (w http.ResponseWriter, r *http.Request){
	//limitParam := r.URL.Query().Get("limit")
	authParam := r.URL.Query().Get("auth")
	authString := ""
	if authParam != "" {
		authString = "private_token=" + authParam
	}
	log.Print(authString)
	response, err := http.Get(API_PATH + "projects?" + authString + "&per_page=100")
	var project CO2Struct.Project
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
	} else {
		body, error := ioutil.ReadAll(response.Body)
		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		err1 := json.Unmarshal(body, &project.Repos)
		if err1 != nil {
			log.Print(err)
		}
		log.Print(project)
		for i := 0; i < len(project.Repos) ; i++  {
			commit, _ := http.Get(API_PATH + "projects/" + strconv.FormatInt(project.Repos[i].ID, 10) + "/repository/commits?" + authString)
			project.Repos[i].Repos, _ = strconv.ParseInt(commit.Header.Get("X-Total"), 10, 64)
		}
		if authString != "" {
			project.Auth = true
		} else {
			project.Auth = false
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(project)
	}

}
