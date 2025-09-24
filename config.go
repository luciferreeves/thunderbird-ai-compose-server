package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config ServerConfig

func init() {
	godotenv.Load()

	Config = ServerConfig{
		Port:     getEnvAsInt("PORT"),
		Provider: Provider(getEnv("PROVIDER")),
		Model:    getEnv("MODEL"),
		APIKey:   getEnv("API_KEY"),
	}

	if Config.Port == 0 {
		Config.Port = 3000
	}

	if Config.Provider == "" {
		Config.Provider = Gemini
	}

	if Config.Model == "" {
		Config.Model = "gemini-2.5-flash"
	}

	if Config.APIKey == "" {
		log.Fatal("API_KEY environment variable is required")
	}

	log.Printf("Configuration loaded: %+v\n", Config)
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
