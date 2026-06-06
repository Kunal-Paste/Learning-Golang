package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURI    string
	MongoDBname string
	JWTsecret   string
}

func Load() (Config, error) {
	_ = godotenv.Load()

	cfg := Config{
		MongoURI:    strings.TrimSpace(os.Getenv("MONGO_URL")),
		MongoDBname: strings.TrimSpace(os.Getenv("MONGO_DB_URL")),
		JWTsecret:   strings.TrimSpace(os.Getenv("JWT_SECRET")),
	}

	if cfg.MongoURI == "" {
		return Config{}, fmt.Errorf("missing mongodb link")
	}

	if cfg.MongoDBname == "" {
		return Config{}, fmt.Errorf("missing mongodb name")
	}

	if cfg.JWTsecret == "" {
		return Config{}, fmt.Errorf("missing jwt secret")
	}

	return cfg, nil

}
