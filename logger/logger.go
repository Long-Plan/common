package logger

import (
	"fmt"
	"os"

	"github.com/Long-Plan/common/environment"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Debugf(template string, args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
	Panicf(template string, args ...interface{})
	Fatalf(template string, args ...interface{})
	BuildFields(args ...interface{}) []zap.Field
	With(fields ...zap.Field) ILogger
	Close()
}

type logger struct {
	logger *zap.Logger
}

func NewLogger(option *LoggerOption) ILogger {
	option = validateLoggerOption(option)
	var encoderConfig zapcore.EncoderConfig
	if option.Mode == environment.Dev {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.TimeKey = "datetime"
	encoderConfig.LevelKey = "level"
	encoderConfig.CallerKey = "at"
	encoderConfig.MessageKey = "msg"
	var encoder zapcore.Encoder
	if option.JsonEncoding {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	priority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.Level(option.LogLevel)
	})

	writer := zapcore.AddSync(option.Writer)

	core := zapcore.NewCore(encoder, zapcore.Lock(zapcore.AddSync(os.Stdout)), priority)
	if option.Writer != nil && option.Mode == environment.Prod {
		core = zapcore.NewTee(
			core,
			zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), writer, priority),
		)
	}

	instance := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(option.SkipCaller))
	if option.Mode == environment.Dev {
		instance = instance.WithOptions(zap.AddStacktrace(zapcore.DebugLevel))
	}
	return &logger{logger: instance}

}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.logger.Debug(msg, fields...)
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, fields...)
}

func (l *logger) Panic(msg string, fields ...zap.Field) {
	l.logger.Panic(msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, fields...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.logger.Debug(getMessage(template, args...))
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.logger.Info(getMessage(template, args...))
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.logger.Warn(getMessage(template, args...))
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.logger.Error(getMessage(template, args...))
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.logger.Panic(getMessage(template, args...))
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatal(getMessage(template, args...))
}

func (l *logger) Close() {
	l.logger.Sync()
}

func getMessage(template string, args ...interface{}) string {
	if len(args) == 0 {
		return template
	}

	if template != "" {
		return fmt.Sprintf(template, args...)
	}

	if len(args) == 1 {
		return fmt.Sprint(args[0])
	}

	return fmt.Sprint(args...)
}

func (l *logger) BuildFields(args ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0)
	isEven := len(args)%2 == 0
	if !isEven {
		return fields
	}

	for i := 0; i < len(args); i += 2 {
		fields = append(fields, zap.Any(args[i].(string), args[i+1]))
	}
	return fields
}

func (l *logger) With(fields ...zap.Field) ILogger {
	return &logger{logger: l.logger.With(fields...)}
}
