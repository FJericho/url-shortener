package config

import (
	"fmt"

	"github.com/FJericho/url-shortener/entity"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConfig(v *viper.Viper, log *logrus.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		v.GetString("DB_HOST"),
		v.GetString("DB_USER"),
		v.GetString("DB_PASSWORD"),
		v.GetString("DB_NAME"),
		v.GetInt("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.WithError(err).Fatal("Failed to connect database")
	}

	log.Info("Database Connected")

	err = db.AutoMigrate(
		&entity.URL{},
	)

	if err != nil {
		log.Fatalf("error on running migration : %+v", err)
	}

	return db
}
