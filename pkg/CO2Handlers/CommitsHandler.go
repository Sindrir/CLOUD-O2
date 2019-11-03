package CO2Handlers

import (
	"encoding/json"
	"github.com/Sindrir/CLOUD-O2/pkg/CO2Struct"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
)

func CommitsHandler (w http.ResponseWriter, r *http.Request){
	limitParam := r.URL.Query().Get("limit")
	authParam := r.URL.Query().Get("auth")
	authString := ""
	if authParam != "" {
		authString = "private_token=" + authParam
	}
	log.Print(authString)
	response, err := http.Get(API_PATH + "projects?" + authString + "&per_page=100&page=1")
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
		totalPages, _ := strconv.Atoi(response.Header.Get("X-Total-Pages"))

		for j := 0; j < totalPages ; j++  {
			var repo CO2Struct.Commits
			for i := 0 + (j * 100); i < len(project.Repos) ; i++  {
				commit, _ := http.Get(API_PATH + "projects/" + strconv.FormatInt(project.Repos[i].ID, 10) + "/repository/commits?" + authString)
				project.Repos[i].Repos, _ = strconv.ParseInt(commit.Header.Get("X-Total"), 10, 64)
			}
			if j < (totalPages-1){
				response, err = http.Get(API_PATH + "projects?" + authString + "&per_page=100&page=" + strconv.Itoa(j+2))
				body, error := ioutil.ReadAll(response.Body)
				if error != nil {
					w.WriteHeader(http.StatusInternalServerError)
				}

				json.Unmarshal(body, &repo)
			}
			project.Repos = append(project.Repos, repo)
		}

		if authString != "" {
			project.Auth = true
		} else {
			project.Auth = false
		}

		log.Print(len(project.Repos))

		sort.Slice(project.Repos, func(i, j int) bool {
			return project.Repos[i].Repos > project.Repos[j].Repos
		})
		limit, _ := strconv.ParseInt(limitParam, 10, 64)
		log.Print(limitParam)
		log.Print(limit)
		project.Repos = project.Repos[:limit]

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(project)
	}

}
