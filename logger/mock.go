package logger

import (
	"go.uber.org/zap"
)

type mockLogger struct {
	logger *zap.Logger
}

func NewMockLogger() ILogger {
	return &mockLogger{
		logger: zap.NewExample(),
	}
}

func (l *mockLogger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *mockLogger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *mockLogger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *mockLogger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *mockLogger) Panic(msg string, fields ...zap.Field) {
	l.logger.Panic(msg, fields...)
}

func (l *mockLogger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *mockLogger) Debugf(template string, args ...interface{}) {
	l.logger.Sugar().Debugf(template, args...)
}

func (l *mockLogger) Infof(template string, args ...interface{}) {
	l.logger.Sugar().Infof(template, args...)
}

func (l *mockLogger) Warnf(template string, args ...interface{}) {
	l.logger.Sugar().Warnf(template, args...)
}

func (l *mockLogger) Errorf(template string, args ...interface{}) {
	l.logger.Sugar().Errorf(template, args...)
}

func (l *mockLogger) Panicf(template string, args ...interface{}) {
	l.logger.Sugar().Panicf(template, args...)
}

func (l *mockLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Sugar().Fatalf(template, args...)
}

func (l *mockLogger) BuildFields(args ...interface{}) []zap.Field {
	return []zap.Field{}
}

func (l *mockLogger) With(fields ...zap.Field) ILogger {
	return &mockLogger{
		logger: l.logger.With(fields...),
	}
}

func (l *mockLogger) Close() {
	l.logger.Sync()
}
