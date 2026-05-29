package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	MongoUrl   string
	MongoDb    string
	ServerPort string
}

func Load() (Config, error) {

	// godotenv.Load() reads the .env and sets them into process env
	// os.getenv -> reads those values

	if err := godotenv.Load(); err != nil {
		return Config{}, fmt.Errorf("failed to load env")
	}

	mongoUrl, err := extractEnv("MONGO_URL")
	if err != nil {
		return Config{}, err
	}

	mongoDb, err := extractEnv("MONGO_DB_NAME")
	if err != nil {
		return Config{}, err
	}

	port, err := extractEnv("PORT")
	if err != nil {
		return Config{}, err
	}

	return Config{
		MongoUrl:   mongoUrl,
		MongoDb:    mongoDb,
		ServerPort: port,
	}, nil

}

func extractEnv(key string) (string, error) {
	val := os.Getenv(key)
	if val == "" {
		return "", fmt.Errorf("missing required env")
	}

	return val, nil
}
