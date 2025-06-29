package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Profile      string        `mapstructure:"PROFILE"`
	ServerPort   string        `mapstructure:"SERVER_PORT"`
	TokenExpiry  time.Duration `mapstructure:"TOKEN_EXPIRY"`
	PrivateKey   string        `mapstructure:"PASETO_PRIVATE_KEY"`
	PublicKey    string        `mapstructure:"PASETO_PUBLIC_KEY"`
	PasswordSalt string        `mapstructure:"PASSWORD_SALT"`

	// API Keys
	GoogleAPIKey string `mapstructure:"GOOGLE_API_KEY"`
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`

	// Database configuration
	PostgresURL string `mapstructure:"POSTGRES_URL"`

	// SMTP configuration
	SMTPServer   string `mapstructure:"SMTP_SERVER"`
	SMTPPort     int    `mapstructure:"SMTP_PORT"`
	SMTPUsername string `mapstructure:"SMTP_USERNAME"`
	SMTPPassword string `mapstructure:"SMTP_PASSWORD"`
	SMTPFrom     string `mapstructure:"SMTP_FROM"`
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
		"SERVER_PORT",
		"TOKEN_EXPIRY",
		"PASETO_PRIVATE_KEY",
		"PASETO_PUBLIC_KEY",
		"PASSWORD_SALT",
		"GOOGLE_API_KEY",
		"OPENAI_API_KEY",
		"POSTGRES_URL",
		"SMTP_SERVER",
		"SMTP_PORT",
		"SMTP_USERNAME",
		"SMTP_PASSWORD",
		"SMTP_FROM",
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
