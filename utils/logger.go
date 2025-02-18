package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// Logger instance
var Logger = logrus.New()

func InitLogger() {
	Logger.SetFormatter(&logrus.JSONFormatter{}) // Log in JSON format
	Logger.SetOutput(os.Stdout)                  // Print logs to console
	Logger.SetLevel(logrus.InfoLevel)            // Set default level to INFO
}
