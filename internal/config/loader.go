package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Printf("Warning: Could not read .env file: %v", err)
	}

	err = viper.Unmarshal(&config)
	return
}
