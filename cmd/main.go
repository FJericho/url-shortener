package main

import (
	"fmt"

	"github.com/FJericho/url-shortener/config"
	_ "github.com/FJericho/url-shortener/docs"
)

// @title       URL Shortener API
// @version     1.0
// @description URL shortener backend service
// @host        localhost:3000
// @BasePath    /
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
