package service

import (
	"github.com/FJericho/url-shortener/dto"
	"github.com/FJericho/url-shortener/entity"
	"github.com/FJericho/url-shortener/helper"
	"github.com/FJericho/url-shortener/repository"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type URLService interface {
	ShortenURL(request dto.ShortenRequest) (*dto.ShortenResponse, error)
	GetOriginalURL(shortCode string) (*dto.ShortenResponse, error)
}

type urlService struct {
	URLRepository repository.URLRepository
	Log           *logrus.Logger
	Validate      *validator.Validate
}

func NewURLService(repo repository.URLRepository, log *logrus.Logger, validate *validator.Validate) URLService {
	return &urlService{
		URLRepository: repo,
		Log:           log,
		Validate:      validate,
	}
}

func (u urlService) ShortenURL(request dto.ShortenRequest) (*dto.ShortenResponse, error) {
	if err := u.Validate.Struct(request); err != nil {
		u.Log.Warnf("Invalid request body  : %+v", err)
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid request. Please check your input.")
	}

	var shortCode string
	for {
		shortCode = helper.GenerateRandomCode(6)
		existing, _ := u.URLRepository.FindByAlias(shortCode)
		if existing.ID == 0 {
			break
		}
	}

	url := &entity.URL{
		Original:  request.Original,
		ShortCode: shortCode,
	}

	if err := u.URLRepository.Save(url); err != nil {
		u.Log.Errorf("Failed to save URL: %+v", err)
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Failed to save shortened URL.")
	}

	u.Log.Infof("Shortened URL created: %s -> %s", url.ShortCode, url.Original)
	return &dto.ShortenResponse{
		Original:  url.Original,
		ShortCode: url.ShortCode,
	}, nil
}

func (u urlService) GetOriginalURL(shortCode string) (*dto.ShortenResponse, error) {
	url, err := u.URLRepository.FindByAlias(shortCode)

	if err != nil {
		u.Log.Warnf("Short code not found: %s", shortCode)
		return nil, fiber.NewError(fiber.StatusNotFound, "Short URL not found")
	}

	_ = u.URLRepository.IncrementClicks(shortCode)

	u.Log.Infof("Redirected short code: %s -> %s", url.ShortCode, url.Original)

	return &dto.ShortenResponse{
		Original:  url.Original,
		ShortCode: url.ShortCode,
	}, nil
}
