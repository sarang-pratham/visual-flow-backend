package config

import (
    "os"
)

type Config struct {
    ServerPort string
}

func LoadConfig() Config {
    port := os.Getenv("SERVER_PORT")
    if port == "" {
        port = ":8000"
    }
    
    return Config{ServerPort: port}
}
