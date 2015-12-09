package handlers_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

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
	assert.Equal(t, "{\"projects\":[{\"name\":\"android\"},{\"name\":\"ios\"},{\"name\":\"web\"}]}", w.Body.String())
}

func TestProjectGet(t *testing.T) {
	req, _ := http.NewRequest("GET", "/project/android", nil)
	w := httptest.NewRecorder()
	routers.NewRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "v1", w.Header().Get("Lion-Api-Version"))
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"name\":\"android\"}", w.Body.String())
}

func TestProjectGetNotFound(t *testing.T) {
	req, _ := http.NewRequest("GET", "/project/does_not_exist", nil)
	w := httptest.NewRecorder()
	routers.NewRouter().ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "v1", w.Header().Get("Lion-Api-Version"))
}
