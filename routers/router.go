package routers

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/op/go-logging"
)

var log logging.Logger

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler http.Handler
		handler = requestLogger(route.HandlerFunc)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func requestLogger(inner http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now().UTC()

		inner.ServeHTTP(w, r)

		log.Info(
			"%s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(start))
	})
}
