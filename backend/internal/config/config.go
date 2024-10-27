package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"time"
)

type Config struct {
	StoragePath string `yaml:"storage_path" env-required:"true"`
	LogsPath    string `yaml:"logs_path" env-required:"true"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address string        `yaml:"address" env-default:"localhost:8080"`
	Timeout time.Duration `yaml:"timeout" env-default:"10s"`
}

func LoadConfig() (*Config, error) {
	errorStatement := "config.go: "
	configPath := "config/config.yaml"

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return &Config{}, fmt.Errorf("%s config file not found - %s", errorStatement, err.Error())
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		return &Config{}, fmt.Errorf("%s failed to load config from file - %s", errorStatement, err.Error())
	}

	return &config, nil
}
