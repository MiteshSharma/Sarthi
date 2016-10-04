package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	ServerConfig   ServerConfig
	DatabaseConfig DatabaseConfig
	Config         map[string]interface{}
}

type ServerConfig struct {
	Port string
}

type DatabaseConfig struct {
	Host   string
	DbName string
}

func (o *Config) SaveDefaultConfigParams() {
	if o.ServerConfig.Port == "" {
		o.ServerConfig.Port = ":8080"
	}
	if o.DatabaseConfig.Host == "" {
		o.DatabaseConfig.Host = "localhost"
	}
	if o.DatabaseConfig.DbName == "" {
		o.DatabaseConfig.DbName = "sarthi"
	}
}

var ConfigParam *Config = &Config{}

func findConfigFile(fileName string) string {
	if _, error := os.Stat("./" + fileName); error == nil {
		fileName, _ = filepath.Abs("./" + fileName)
	} else if _, error := os.Stat("./config/" + fileName); error == nil {
		fileName, _ = filepath.Abs("./config/" + fileName)
	}
	return fileName
}

func LoadConfig(fileName string) {
	filePath := findConfigFile(fileName)

	file, error := os.Open(filePath)

	if error != nil {
		panic("Error occured during config file reading " + error.Error())
	}

	jsonParser := json.NewDecoder(file)

	config := Config{}

	if jsonErr := jsonParser.Decode(&config); jsonErr != nil {
		panic("Json parsing error" + jsonErr.Error())
	}

	config.SaveDefaultConfigParams()

	ConfigParam = &config
}
