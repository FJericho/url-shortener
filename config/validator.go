package config

import (
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func NewValidatorConfig(v *viper.Viper) *validator.Validate {
	return validator.New()
}
