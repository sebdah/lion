package handlers

import "net/http"

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Rooooaaar!!\n"))

	appendLionHeaders(w)
	w.WriteHeader(http.StatusOK)
}
