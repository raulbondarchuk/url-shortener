package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var (
	defaultPath = "./config/local.yaml" // Example default value
)

func SetupConfigPath() {
	// Search .env file
	envPath, err := findEnvFile(".")
	if err != nil {
		log.Fatalf("Error finding .env file: %v", err)
	} else {
		// Load found .env file
		if err := godotenv.Load(envPath); err != nil {
			log.Printf("Error loading .env file: %v", err)
		}
	}

	// Get CONFIG_PATH from .env
	configPath := os.Getenv("CONFIG_PATH")

	// If CONFIG_PATH is not set, use the default value
	if configPath == "" {
		log.Printf("CONFIG_PATH is not set in .env, using default value: %s", defaultPath)
		configPath = defaultPath
	}

	// Set CONFIG_PATH in environment variable
	if err := os.Setenv("CONFIG_PATH", configPath); err != nil {
		log.Fatalf("Error setting CONFIG_PATH: %v", err)
	}
	log.Printf("CONFIG_PATH set: %s", configPath)
}

func findEnvFile(root string) (string, error) {
	var envPath string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Name() == ".env" {
			envPath = path
			return filepath.SkipAll
		}
		return nil
	})
	if envPath == "" {
		return "", fmt.Errorf(".env file not found")
	}
	return envPath, err
}
