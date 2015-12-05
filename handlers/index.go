package handlers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Rooooaaar!!\n"))

	appendLionHeaders(w)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
