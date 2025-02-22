package logger

import (
	"io"

	"github.com/Long-Plan/common/environment"
)

type LoggerOption struct {
	Mode         environment.Mode
	LogLevel     LogLevel
	SkipCaller   int
	JsonEncoding bool
	Writer       io.Writer
}

var LoggerDefaultOption LoggerOption = LoggerOption{
	Mode:         environment.Dev,
	LogLevel:     DebugLevel,
	JsonEncoding: false,
	SkipCaller:   1,
	Writer:       nil,
}

func validateLoggerOption(opt *LoggerOption) *LoggerOption {
	if opt == nil {
		return &LoggerDefaultOption
	}

	return opt
}
