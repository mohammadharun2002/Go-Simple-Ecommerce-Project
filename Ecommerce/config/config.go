package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations *Config

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int64
	JwtSecretKey string
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load the env variables", err)
	}

	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("ServiceName is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("HTTP PORT is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		fmt.Errorf("Cant Convert Http-Port to Integer")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecretKey,
	}
}

func GetConfig() *Config {
	if configurations == nil {
		loadConfig()

	}
	return configurations
}
