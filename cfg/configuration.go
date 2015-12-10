package cfg

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/sebdah/lion/models"
)

var Config *Configuration

func init() {
	loadConfiguration()
}

type Configuration struct {
	Port int `json:"port"`
	*models.Projects
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func loadConfiguration() {
	_, filename, _, _ := runtime.Caller(1)
	dir, err := filepath.Abs(filepath.Dir(filename))
	if err != nil {
		panic(err)
	}
	fullPath := path.Join(dir, "../config.json")

	f, err := os.Open(fullPath)
	if err != nil {
		panic(fmt.Sprintf("Configuration file %s not found", fullPath))
	}
	defer f.Close()

	decoder := json.NewDecoder(f)
	Config = NewConfiguration()
	err = decoder.Decode(&Config)
	if err != nil {
		fmt.Printf("Error reading configuration: %s", err.Error())
	}
}
