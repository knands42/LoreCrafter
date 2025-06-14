package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Profile     string        `mapstructure:"PROFILE"`
	Hostname    string        `mapstructure:"APP_HOSTNAME"`
	ServerPort  string        `mapstructure:"SERVER_PORT"`
	TokenExpiry time.Duration `mapstructure:"TOKEN_EXPIRY"`
	PrivateKey  string        `mapstructure:"PASETO_PRIVATE_KEY"`
	PublicKey   string        `mapstructure:"PASETO_PUBLIC_KEY"`

	// API Keys
	GoogleAPIKey string `mapstructure:"GOOGLE_API_KEY"`
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`

	// Database configuration
	PostgresURL string `mapstructure:"POSTGRES_URL"`
}

// LoadConfig loads the configuration from .env file and environment variables
func LoadConfig(path string) (config Config, err error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName(".env")
	v.SetConfigType("env")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("error reading config file: %v", err)
		}
	}

	keys := []string{
		"PROFILE",
		"APP_HOSTNAME",
		"SERVER_PORT",
		"TOKEN_EXPIRY",
		"PASETO_PRIVATE_KEY",
		"PASETO_PUBLIC_KEY",
		"GOOGLE_API_KEY",
		"OPENAI_API_KEY",
		"POSTGRES_URL",
	}
	for _, key := range keys {
		err := v.BindEnv(key)
		if err != nil {
			return config, fmt.Errorf("error binding env var %s: %w", key, err)
		}
	}

	// Unmarshal the configuration
	err = v.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode into config struct: %w", err)
	}

	return config, nil
}
