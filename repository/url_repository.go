package repository

import (
	"github.com/FJericho/url-shortener/entity"
	"gorm.io/gorm"
)

type URLRepository interface {
	Save(url *entity.URL) error
	FindByAlias(alias string) (*entity.URL, error)
	IncrementClicks(alias string) error
}

type urlRepository struct {
	DB *gorm.DB
}

func NewURLRepository(db *gorm.DB) URLRepository {
	return &urlRepository{
		DB: db,
	}
}

func (r *urlRepository) Save(url *entity.URL) error {
	return r.DB.Create(url).Error
}

func (r *urlRepository) FindByAlias(alias string) (*entity.URL, error) {
	var url entity.URL

	err := r.DB.Where("short_code = ?", alias).First(&url).Error
	
	return &url, err
}

func (r *urlRepository) IncrementClicks(alias string) error {
	return r.DB.Model(&entity.URL{}).
		Where("short_code = ?", alias).
		UpdateColumn("click_count", gorm.Expr("click_count + ?", 1)).Error
}
