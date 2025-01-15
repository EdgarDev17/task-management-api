package repositories

import "go.uber.org/zap"

// Logger es una interfaz para el logger
type Logger interface {
	Info(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}
