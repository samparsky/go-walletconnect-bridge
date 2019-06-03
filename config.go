package main

import (
)

// Config app configuration
type Config struct {
	Environment string
	Port int
	Host string
	Debug bool
}

// InitConfig initialise config
func InitConfig() *Config {
	// environment := os.Getenv("ENV") || "development"
	return nil
}
