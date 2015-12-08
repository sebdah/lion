package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectList(t *testing.T) {
	handle := http.HandlerFunc(ProjectsList)
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	handle.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "v1", w.Header().Get("Lion-Api-Version"))
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "{\"projects\":[{\"name\":\"android\"},{\"name\":\"ios\"},{\"name\":\"web\"}]}", w.Body.String())
}
