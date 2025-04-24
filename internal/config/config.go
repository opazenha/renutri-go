package config

import (
    "fmt"
    "github.com/joho/godotenv"
    "os"
)

// Config holds application configuration loaded from environment.
type Config struct {
    DBUrl string
    Port  string
}

// Load reads environment variables and returns a Config.
func Load() (*Config, error) {
    // Load environment variables from .env file, if present
    _ = godotenv.Load()
    dbUrl := os.Getenv("DB_URL")
    if dbUrl == "" {
        return nil, fmt.Errorf("DB_URL environment variable required")
    }
    port := os.Getenv("PORT")
    if port == "" {
        port = "8881"
    }
    return &Config{DBUrl: dbUrl, Port: port}, nil
}