package config

import (
	"fmt"
	"path/filepath"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	Storage StorageConfig `mapstructure:"storage"`
	App     AppConfig     `mapstructure:"app"`
}

// StorageConfig defines storage-related configuration
type StorageConfig struct {
	Type     string `mapstructure:"type"`     // memory, json, gorm
	FilePath string `mapstructure:"filepath"` // for json and gorm storage
}

// AppConfig defines application-level configuration
type AppConfig struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

// defaultConfig returns the default configuration
func defaultConfig() Config {
	return Config{
		Storage: StorageConfig{
			Type:     "memory",
			FilePath: "contacts.json",
		},
		App: AppConfig{
			Name:    "Mini CRM",
			Version: "2.0.0",
		},
	}
}

// Load loads configuration from file or uses defaults
func Load(configPath string) (*Config, error) {
	// Set up viper
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if configPath != "" {
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.mini-crm")
		viper.AddConfigPath("/etc/mini-crm")
	}

	// Set default values
	defaults := defaultConfig()
	viper.SetDefault("storage.type", defaults.Storage.Type)
	viper.SetDefault("storage.filepath", defaults.Storage.FilePath)
	viper.SetDefault("app.name", defaults.App.Name)
	viper.SetDefault("app.version", defaults.App.Version)

	// Read configuration file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, use defaults
			fmt.Printf("Config file not found, using defaults\n")
		} else {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
	}

	// Unmarshal into struct
	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// GetStorageFilePath returns the full path for storage files
func (c *Config) GetStorageFilePath() string {
	if filepath.IsAbs(c.Storage.FilePath) {
		return c.Storage.FilePath
	}

	// Make it relative to current directory for simplicity
	switch c.Storage.Type {
	case "json":
		if c.Storage.FilePath == "" {
			return "contacts.json"
		}
		return c.Storage.FilePath
	case "gorm":
		if c.Storage.FilePath == "" {
			return "contacts.db"
		}
		return c.Storage.FilePath
	}

	return c.Storage.FilePath
}

// Validate validates the configuration
func (c *Config) Validate() error {
	validStorageTypes := map[string]bool{
		"memory": true,
		"json":   true,
		"gorm":   true,
	}

	if !validStorageTypes[c.Storage.Type] {
		return fmt.Errorf("invalid storage type: %s (valid options: memory, json, gorm)", c.Storage.Type)
	}

	return nil
}
