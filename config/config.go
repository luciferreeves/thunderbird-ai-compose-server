package config

import (
	"log"
	"os"
	"strconv"
	"thunderbird-ai-compose-server/types"

	"github.com/joho/godotenv"
)

var Config types.ServerConfig

func init() {
	godotenv.Load()

	Config = types.ServerConfig{
		Port:     getEnvAsInt("PORT"),
		Provider: types.Provider(getEnv("PROVIDER")),
		Model:    getEnv("MODEL"),
		APIKey:   getEnv("API_KEY"),
	}

	if Config.Port == 0 {
		Config.Port = 3000
	}

	if Config.Provider == "" {
		Config.Provider = types.Gemini
	}

	if Config.Model == "" {
		Config.Model = "gemini-2.5-flash"
	}

	if Config.APIKey == "" {
		log.Fatal("API_KEY environment variable is required")
	}
}

func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return ""
}

func getEnvAsInt(key string) int {
	valueStr := getEnv(key)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return 0
}
