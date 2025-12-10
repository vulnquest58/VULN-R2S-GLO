package utils

import (
	"log"
	"os"
)

var logger *log.Logger

func InitLogger() {
	logger = log.New(os.Stdout, "", 0)
}

func LogInfo(msg string) {
	logger.Printf("\033[96m[INFO]\033[0m %s", msg)
}

func LogWarn(msg string) {
	logger.Printf("\033[93m[WARN]\033[0m %s", msg)
}

func LogVuln(msg string) {
	logger.Printf("\033[91m%s\033[0m", msg)
}
