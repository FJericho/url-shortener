package main

import (
	"fmt"

	"github.com/FJericho/url-shortener/config"
)

func main() {
	viper := config.NewViperConfig()
	app := config.NewFiberConfig(viper)
	log := config.NewLoggerConfig()
	validate := config.NewValidatorConfig(viper)
	db := config.NewDatabaseConfig(viper, log)

	config.StartServer(&config.AppConfig{
		App:      app,
		DB:       db,
		Validate: validate,
		Viper:    viper,
		Log:      log,
	})

	appPort := viper.GetInt("APP_PORT")

	err := app.Listen(fmt.Sprintf(":%d", appPort))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
