package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	indexHandle := http.HandlerFunc(Index)
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	indexHandle.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "v1", w.Header().Get("Lion-Api-Version"))
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
