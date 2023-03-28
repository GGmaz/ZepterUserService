package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Port string
	Host string
}

func NewConfig() *Config {
	return &Config{
		Port: goDotEnvVariable("USER_SERVICE_PORT"),
		Host: goDotEnvVariable("USER_SERVICE_HOST"),
	}
}

func goDotEnvVariable(key string) string {
	godotenv.Load(".env")
	return os.Getenv(key)
}
