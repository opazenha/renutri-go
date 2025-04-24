package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config holds application configuration loaded from environment.
type Config struct {
    MONGO_URL string
    Port  string
}

// Load reads environment variables and returns a Config.
func Load() (*Config, error) {
    // Load environment variables from .env file, if present
    _ = godotenv.Load()
    MONGO_URL := os.Getenv("MONGO_URL")
    if MONGO_URL == "" {
        return nil, fmt.Errorf("DB_URL environment variable required")
    }
    port := os.Getenv("PORT")
    if port == "" {
        port = "8881"
    }
    return &Config{MONGO_URL: MONGO_URL, Port: port}, nil
}