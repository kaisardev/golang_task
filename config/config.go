package config

import (
	"sync"
)

var (
	onceConfig sync.Once
	config     *Config
)

type Config struct {
	ServerPort     string `env:"HTTP_SERVER_PORT"`
	AllowedOrigins string `env:"HTTP_ALLOWED_ORIGINS"`
}

func GetConfig() *Config {
	onceConfig.Do(func() {
		// Implementation of reading .env file
		conf := &Config{
			ServerPort:     "8000",
			AllowedOrigins: "*",
		}

		config = conf
	})

	return config
}
