package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiberConfig(v *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      "URL Shortener API",
		Prefork:      false,
		ErrorHandler: NewErrorHandler(),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(ctx *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError

		if fiberErr, ok := err.(*fiber.Error); ok {
			code = fiberErr.Code
		}

		return ctx.Status(code).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
