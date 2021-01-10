package config

import (
	"fmt"
	"os"
	"strings"
)

type Config struct {
	DBEngine string

	ExtHostname string

	LoggerConfig string
}

func CollectConfig() *Config {
	var missingEnv []string
	var config Config

	// DB_ENGINE
	config.DBEngine = os.Getenv("DB_ENGINE")
	if config.DBEngine == "" {
		missingEnv = append(missingEnv, "DB_ENGINE")
	}

	// EXT_HOSTNAME
	config.ExtHostname = os.Getenv("EXT_HOSTNAME")
	if config.ExtHostname == "" {
		missingEnv = append(missingEnv, "EXT_HOSTNAME")
	}

	// LOG_LEVEL
	var envLoggerLevel = os.Getenv("LOG_LEVEL")
	if envLoggerLevel == "" {
		config.LoggerConfig = "<root>=INFO"
	} else {
		config.LoggerConfig = fmt.Sprintf("<root>=%s", strings.ToUpper(envLoggerLevel))
	}

	// Validation
	if len(missingEnv) > 0 {
		msg := fmt.Sprintf("Environment variables missing: %v", missingEnv)
		panic(fmt.Sprint(msg))
	}

	return &config
}