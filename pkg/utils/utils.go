package utils

import (
	"log"
	"os"
)

// GetEnv gets the environment variable by key, if not found returns the default value
func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		log.Printf("Environment variable %s not found, setting default value: %s", key, defaultValue)
		return defaultValue
	}
	return value
}