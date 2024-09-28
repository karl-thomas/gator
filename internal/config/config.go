package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	DBUrl    string `json:"db_url"`
	Username string `json:"username"`
}

func Read() Config {
	filePath, error := getConfigFilePath()
	if error != nil {
		panic(error)
	}
	file, error := os.Open(filePath)
	if error != nil {
		panic(error)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	cfg := Config{}
	error = decoder.Decode(&cfg)
	if error != nil {
		panic(error)
	}
	return cfg
}

func SetUser(username string) error {
	cfg := Read()
	cfg.Username = username
	return write(cfg)
}

func getConfigFilePath() (string, error) {
	dirName, error := os.UserHomeDir()
	if error != nil {
		return "", error
	}

	return dirName + "/" + configFileName, nil
}

func write(cfg Config) error {
	filePath, error := getConfigFilePath()
	if error != nil {
		return error
	}

	file, error := os.Create(filePath)
	if error != nil {
		return error
	}

	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(cfg)
}

const configFileName = ".gatorconfig.json"
