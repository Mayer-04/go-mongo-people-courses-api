package config

import (
	"fmt"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	AppPort uint64
	Mongo   MongoConfig
}

type MongoConfig struct {
	URI      string
	Database string
}

func LoadConfig() (*Config, error) {
	mongoURI := os.Getenv("MONGODB_URI")
	if mongoURI == "" {
		return nil, fmt.Errorf("set the environment variable 'MONGODB_URI'")
	}

	mongoDB := getEnv("MONGODB_DATABASE", "people_courses")

	portStr := getEnv("PORT", "8080")
	port, err := strconv.ParseUint(portStr, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("the port must be a valid number")
	}

	cfg := &Config{
		AppPort: port,
		Mongo: MongoConfig{
			URI:      mongoURI,
			Database: mongoDB,
		},
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
