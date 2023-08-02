package configs

import (
	"log"

	"github.com/spf13/viper"
)

var Configs *configs

func InitConfigs() {
	Configs = loadEnvVariables()
}

type configs struct {
	Token string `mapstructure:"TOKEN"`
}

func loadEnvVariables() (config *configs) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
