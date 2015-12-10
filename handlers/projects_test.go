package handlers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sebdah/lion/models"
	"github.com/sebdah/lion/routers"
	"github.com/stretchr/testify/assert"
)

func TestProjectsList(t *testing.T) {
	req, _ := http.NewRequest("GET", "/projects", nil)
	w := httptest.NewRecorder()
	routers.NewRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "v1", w.Header().Get("Lion-Api-Version"))
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	projects := models.NewProjects()
	err := json.Unmarshal(w.Body.Bytes(), projects)
	assert.Nil(t, err)
	assert.Equal(t, "android", projects.GetBySlug("android").Slug)
	assert.Equal(t, "Android", projects.GetBySlug("android").Name)
	assert.Equal(t, "ios", projects.GetBySlug("ios").Slug)
	assert.Equal(t, "iOS", projects.GetBySlug("ios").Name)
	assert.Equal(t, "web", projects.GetBySlug("web").Slug)
	assert.Equal(t, "Web project", projects.GetBySlug("web").Name)
}

func TestProjectGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/project/android", nil)
	w := httptest.NewRecorder()
	routers.NewRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "v1", w.Header().Get("Lion-Api-Version"))
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

	project := models.NewProject()
	err := json.Unmarshal(w.Body.Bytes(), project)
	assert.Nil(t, err)
	assert.Equal(t, "android", project.Slug)
	assert.Equal(t, "Android", project.Name)
}

func TestProjectGetNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/project/does_not_exist", nil)
	w := httptest.NewRecorder()
	routers.NewRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "v1", w.Header().Get("Lion-Api-Version"))
}
