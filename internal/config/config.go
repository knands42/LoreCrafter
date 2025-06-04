package config

import (
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	ServerPort  string        `mapstructure:"SERVER_PORT"`
	TokenExpiry time.Duration `mapstructure:"TOKEN_EXPIRY"`
	PrivateKey  string        `mapstructure:"PASETO_PRIVATE_KEY"`
	PublicKey   string        `mapstructure:"PASETO_PUBLIC_KEY"`

	// API Keys
	GoogleAPIKey string `mapstructure:"GOOGLE_API_KEY"`
	OpenAIAPIKey string `mapstructure:"OPENAI_API_KEY"`

	// Database configuration
	DBConfig DBConfig
}

// DBConfig holds the database configuration
type DBConfig struct {
	Host     string `mapstructure:"POSTGRES_HOST"`
	Port     string `mapstructure:"POSTGRES_PORT"`
	User     string `mapstructure:"POSTGRES_USER"`
	Password string `mapstructure:"POSTGRES_PASSWORD"`
	DBName   string `mapstructure:"POSTGRES_DBNAME"`
	SSLMode  string `mapstructure:"POSTGRES_SSLMODE"`
	URL      string `mapstructure:"POSTGRES_URL"`
}

// LoadConfig loads the configuration from .env file and environment variables
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// Read from the .env file
	viper.AutomaticEnv()

	// Try to read the config file, but don't return an error if it doesn't exist
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Printf("error reading config file: %v", err)
			return config, nil
		}
	}

	// Unmarshal the configuration
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, fmt.Errorf("unable to decode into config struct: %w", err)
	}

	// Unmarshal the database configuration
	dbConfig := DBConfig{}
	err = viper.Unmarshal(&dbConfig)
	if err != nil {
		return config, fmt.Errorf("unable to decode into db config struct: %w", err)
	}
	config.DBConfig = dbConfig

	return config, nil
}

// GetDBConfig returns the database configuration in the format expected by the database package
func (c *Config) GetDBConfig() DBConfig {
	return c.DBConfig
}
