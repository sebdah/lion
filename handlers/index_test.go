package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	indexHandle := http.HandlerFunc(Index)
	req, _ := http.NewRequest("GET", "", nil)
	w := httptest.NewRecorder()
	indexHandle.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Bad response code: %v", http.StatusOK)
	}

	if w.Header().Get("Lion-Api-Version") != "v1" {
		t.Errorf("Bad Lion-Api-Version header: '%v'",
			w.Header().Get("Lion-Api-Version"))
	}
}
