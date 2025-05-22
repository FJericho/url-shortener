package handler

import (
	"github.com/FJericho/url-shortener/dto"
	"github.com/FJericho/url-shortener/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type URLHandler struct {
	URLService service.URLService
	Log        *logrus.Logger
}

func NewURLHandler(s service.URLService, log *logrus.Logger) *URLHandler {
	return &URLHandler{
		URLService: s,
		Log:        log,
	}
}

// Shorten godoc
// @Summary      Shorten a URL
// @Description  Accepts a long URL and returns the shortened version
// @Tags         URL
// @Accept       json
// @Produce      json
// @Param        request body dto.ShortenRequest true "Request body"
// @Success      200 {object} dto.WebResponseSwagger
// @Failure      400 {object} dto.ErrorResponseSwagger
// @Router       /shorten [post]
func (h *URLHandler) Shorten(ctx *fiber.Ctx) error {
	var request dto.ShortenRequest

	if err := ctx.BodyParser(&request); err != nil {
		h.Log.Warnf("Failed to parse request body: %+v", err)
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body.")
	}

	url, err := h.URLService.ShortenURL(request)
	if err != nil {
		return err
	}

	return ctx.JSON(dto.WebResponse[*dto.ShortenResponse]{
		Message: "Short url successfully",
		Data:    url,
	})

}

// @Summary Redirect short URL
// @Description Redirects to original URL using short code
// @Tags URL
// @Param short_code path string true "Short Code"
// @Success 301 "Redirect to original URL"
// @Failure 404 {object} dto.ErrorResponseSwagger
// @Router /{short_code} [get]
func (h *URLHandler) Redirect(c *fiber.Ctx) error {
	shortCode := c.Params("short_code")

	url, err := h.URLService.GetOriginalURL(shortCode)

	if err != nil {
		return err
	}

	return c.Redirect(url.Original, fiber.StatusMovedPermanently)
}

// GetOriginal godoc
// @Summary      Get original URL
// @Description  Retrieves the original URL from the short code
// @Tags         URL
// @Produce      json
// @Param        short_code path string true "Short Code"
// @Success      200 {object} dto.WebResponseSwagger
// @Failure      404 {object} dto.ErrorResponseSwagger
// @Router       /api/url/{short_code} [get]
func (h *URLHandler) GetOriginal(c *fiber.Ctx) error {
	shortCode := c.Params("short_code")

	url, err := h.URLService.GetOriginalURL(shortCode)
	if err != nil {
		return err
	}

	return c.JSON(dto.WebResponse[*dto.ShortenResponse]{
		Message: "Get original url successfully",
		Data:    url,
	})
}
