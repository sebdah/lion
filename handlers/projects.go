package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sebdah/lion/cfg"
)

func ProjectsList(w http.ResponseWriter, r *http.Request) {
	js, err := json.Marshal(cfg.Config.Projects)
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
	vars := mux.Vars(r)

	project := cfg.Config.Projects.GetBySlug(vars["slug"])
	if project == nil {
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
