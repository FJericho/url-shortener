package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLoggerConfig() *logrus.Logger {
	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	log.SetOutput(os.Stdout)

	log.SetLevel(logrus.Level(logrus.InfoLevel))

	return log
}
