package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `json:"name"`
		Host string `json:"host"`
		Port string `json:"port"`
	}
}

var AppConfig *Config

func InitConfig() {
	// Set the file name of the configurations file
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	// Read the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}
	AppConfig = &Config{}
	err = viper.Unmarshal(AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
}
