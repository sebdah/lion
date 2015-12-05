package routers

import (
	"net/http"

	"github.com/sebdah/lion/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", "GET", "/", handlers.Index},
}
