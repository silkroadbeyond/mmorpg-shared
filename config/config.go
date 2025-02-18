package config

import (
	"github.com/spf13/viper"
)

func Load() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic("Failed to load config: " + err.Error())
	}
}

func GetString(key string) string {
	return viper.GetString(key)
}