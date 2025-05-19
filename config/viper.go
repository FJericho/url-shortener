package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func NewViperConfig() *viper.Viper {
	v := viper.New()

	v.SetConfigFile(".env")
	v.SetConfigType("env")

	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf(".env file found: %w \n", err))
	}

	v.AutomaticEnv()

	v.SetDefault("APP_PORT", "3000")

	return v
}
