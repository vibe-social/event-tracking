package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// Load configuration from the file (if present)
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Config file not found, using environment variables.")
	}

	// Load configuration from environment variables
	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("SERVER_PORT", 8080)
	viper.SetDefault("SERVER_MODE", "debug")
}
