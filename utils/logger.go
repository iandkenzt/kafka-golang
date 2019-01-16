package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// Logger ...
var Logger = log.New()

// InitLogger ...
func InitLogger() {
	Logger.Formatter = &log.JSONFormatter{}
	Logger.Out = os.Stdout
}
