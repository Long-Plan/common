package logger

var LoggerInstance ILogger = NewMockLogger()

func InitLogger(option *LoggerOption) {
	LoggerInstance = NewLogger(option)
}

func CloseLogger() {
	LoggerInstance.Close()
}
