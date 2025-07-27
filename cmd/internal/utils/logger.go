package utils

import (
	"log"
	"os"
	"strings"
)

var level string

func InitLogging() {
	level = strings.ToUpper(os.Getenv("LOG_LEVEL"))
	log.SetFlags(2)
}

func LogInfo(message string) {
	if level == "INFO" || level == "DEBUG" {
		log.Println("[INFO]", message)
	}
}

func LogError(message string) {
	log.Println("[ERROR]", message)
}
