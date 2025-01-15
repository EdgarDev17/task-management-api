package logger

import "go.uber.org/zap"

// ZapLogger implementa la interfaz Logger usando Zap
type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() *ZapLogger {
	// logger, _ := zap.NewProduction()
	logger, _ := zap.NewDevelopment()
	return &ZapLogger{logger: logger}
}

func (l *ZapLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *ZapLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg)
}
