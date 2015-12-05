package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/op/go-logging"
	"github.com/sebdah/lion/routers"
	"github.com/spf13/viper"
)

var (
	log    logging.Logger
	config *viper.Viper // Viper configuration object
)

func init() {
	setupConfiguration()

	// Configure logging
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	format := logging.MustStringFormatter("%{color}%{time:15:04:05.000} %{shortfunc} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}")
	backendFormatter := logging.NewBackendFormatter(backend, format)
	logging.SetBackend(backendFormatter)
}

func main() {
	log.Info("Starting the Lion Service")

	port := config.GetInt("port")
	router := routers.NewRouter()
	log.Info("Starting the web server at port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

// Set up configuration management
func setupConfiguration() {
	// Find the config file path
	_, filename, _, _ := runtime.Caller(1)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		panic(err)
	}
	dir = path.Join(dir, "..")

	// Set up the viper object
	config = viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath(dir)

	// Read the configuration
	err = config.ReadInConfig()
	if err != nil {
		if err.Error() == "open : no such file or directory" {
			panic("Could not find the configuration file")
		}
		panic(err)
	}
}
