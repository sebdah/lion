package handlers

import (
	"net/http"

	"github.com/op/go-logging"
)

var log logging.Logger

func appendLionHeaders(w http.ResponseWriter) {
	w.Header().Add("Lion-Api-Version", "v1")
}
