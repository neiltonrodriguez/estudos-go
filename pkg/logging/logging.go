// internal/infrastructure/logging/logger.go
package logging

import (
	"log"
)

type EchoLogger struct{}

func NewLogger() *EchoLogger {
	return &EchoLogger{}
}

func (l *EchoLogger) Info(msg string, args ...any) {
	log.Printf("[INFO] "+msg, args...)
}

func (l *EchoLogger) Error(msg string, args ...any) {
	log.Printf("[ERROR] "+msg, args...)
}

func (l *EchoLogger) Debug(msg string, args ...any) {
	log.Printf("[DEBUG] "+msg, args...)
}
