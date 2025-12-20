package helpers

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/mobile-app-react-native/config"
)

func LoadConfig() (*config.Config, error) {
	configPath := getConfigPath()
	if configPath == "" {
		return nil, errors.New("config path is not set")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var cfg config.Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getConfigPath() string {
	path := os.Getenv("MOBILE_APP_CONFIG_PATH")
	if path != "" {
		return path
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return ""
	}

	configPath := filepath.Join(currentDir, "config.json")
	if _, err := os.Stat(configPath); err == nil {
		return configPath
	}

	return ""
}

func GetEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value != "" {
		return value
	}

	return defaultValue
}

func IsDevelopmentEnvironment() bool {
	env := GetEnvOrDefault("NODE_ENV", "")
	return strings.ToLower(env) == "development"
}