package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/op/go-logging"
	"github.com/sebdah/lion/config"
	"github.com/sebdah/lion/routers"
)

var (
	log logging.Logger
)

func init() {
	// Configure logging
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

func main() {
	log.Info("Starting the Lion Service")

	port := config.Config.GetInt("port")
	router := routers.NewRouter()
	log.Info("Starting the web server at port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
