package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sebdah/lion/models"
)

func ProjectsList(w http.ResponseWriter, r *http.Request) {
	var projects models.Projects

	for _, project := range models.PopulateAllProjectsFromConfig() {
		projects.Projects = append(projects.Projects, project)
	}

	js, err := json.Marshal(projects)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	appendLionHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}

func ProjectGet(w http.ResponseWriter, r *http.Request) {
	var project *models.Project
	vars := mux.Vars(r)

	project = models.NewProject()
	err := project.PopulateFromConfig(vars["project_name"])
	if err != nil {
		appendLionHeaders(w)
		http.NotFound(w, r)
		return
	}

	js, err := json.Marshal(project)
	if err != nil {
		appendLionHeaders(w)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	appendLionHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
}
