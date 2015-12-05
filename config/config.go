package config

import (
	"path"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	// Find the config file path
	_, filename, _, _ := runtime.Caller(1)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		panic(err)
	}
	dir = path.Join(dir, "..")

	// Set up the viper object
	Config = viper.New()
	Config.SetConfigName("config")
	Config.SetConfigType("json")
	Config.AddConfigPath(dir)

	// Read the configuration
	err = Config.ReadInConfig()
	if err != nil {
		if err.Error() == "open : no such file or directory" {
			panic("Could not find the configuration file")
		}
		panic(err)
	}
}
