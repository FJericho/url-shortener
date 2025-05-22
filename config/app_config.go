package config

import (
	"github.com/FJericho/url-shortener/handler"
	"github.com/FJericho/url-shortener/repository"
	"github.com/FJericho/url-shortener/router"
	"github.com/FJericho/url-shortener/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AppConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Viper    *viper.Viper
}

func StartServer(config *AppConfig) {

	urlRepository := repository.NewURLRepository(config.DB)

	urlService := service.NewURLService(urlRepository, config.Log, config.Validate)

	urlHandler := handler.NewURLHandler(urlService, config.Log)

	routeConfig := router.RouteConfig{
		App:        config.App,
		URLHandler: urlHandler,
	}

	routeConfig.Setup()
}
