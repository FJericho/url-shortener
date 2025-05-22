package router

import (
	"github.com/FJericho/url-shortener/handler"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type RouteConfig struct {
	App        *fiber.App
	URLHandler *handler.URLHandler
}

func (c *RouteConfig) Setup() {
	c.SetupPublicRoute()
}

func (c *RouteConfig) SetupPublicRoute() {

	c.App.Get("/swagger/*", fiberSwagger.WrapHandler)

	c.App.Post("/shorten", c.URLHandler.Shorten)
	c.App.Get("/:short_code", c.URLHandler.Redirect)
	c.App.Get("/api/url/:short_code", c.URLHandler.GetOriginal)
}
