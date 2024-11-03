package libs

import (
	"encoding/json"
	"os"
)

type Config struct {
	DB       string `json:"db"`
	DSN      string `json:"dsn"`
	LogLevel string `json:"log_level"`
	LogPath  string `json:"log_path"`
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

var config_ Config
var InitConfig_ = 0

func InitConfig() {
	InitConfig_ = 1
	var configFile string
	var dev = "config.dev.json"
	var prod = "config.prod.json"
	if fileExists(prod) {
		configFile = prod
	} else {
		configFile = dev
	}

	file, _ := os.Open(configFile)
	defer file.Close()
	decoder := json.NewDecoder(file)
	_ = decoder.Decode(&config_)
}

func LoadConfig() (Config, error) {
	if InitConfig_ != 1 {
		InitConfig()
	}
	return config_, nil
}
