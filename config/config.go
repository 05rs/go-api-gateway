package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Route struct {
	Path string `mapstructure:"path"`
}

type Service struct {
	BaseURL string  `mapstructure:"base_url"`
	Routes  []Route `mapstructure:"routes"`
}

type Config struct {
	RateLimitWindow time.Duration      `mapstructure:"rateLimitWindow"`
	RateLimitCount  int                `mapstructure:"rateLimitCount"`
	Services        map[string]Service `mapstructure:"services"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigType("yaml")   // Set configuration file format (replace as needed)
	viper.SetConfigName("config") // Set configuration file name
	viper.AddConfigPath(".")      // Search for config file in current directory

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	return &config, nil
}
