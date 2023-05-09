package zlog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	prefix       string
	stdoutLogger *zap.Logger
	fileLogger   *zap.Logger
}

func NewLogger(prefix string) *Logger {
	stdoutEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	stdoutConfig := zap.Config{
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:      true,
		Encoding:         "console",
		EncoderConfig:    stdoutEncoderConfig,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	stdoutLogger, err := stdoutConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil
	}

	fileEncoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	fileConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.InfoLevel),
		Development:       true,
		Encoding:          "json",
		EncoderConfig:     fileEncoderConfig,
		OutputPaths:       []string{"log.txt"},
		ErrorOutputPaths:  []string{"log.txt"},
		InitialFields:     map[string]interface{}{"app": "has"},
		DisableCaller:     true,
		DisableStacktrace: true,
	}
	fileLogger, err := fileConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		return nil
	}

	return &Logger{prefix: prefix, stdoutLogger: stdoutLogger, fileLogger: fileLogger}
}

func (m *Logger) Info(msg string, fields ...zap.Field) {
	m.stdoutLogger.Info(m.prefix+msg, fields...)
	m.fileLogger.Info(m.prefix+msg, fields...)
}

func (m *Logger) Debug(msg string, fields ...zap.Field) {
	m.stdoutLogger.Debug(m.prefix+msg, fields...)
	m.fileLogger.Debug(m.prefix+msg, fields...)
}

func (m *Logger) Error(msg string, fields ...zap.Field) {
	m.stdoutLogger.Error(m.prefix+msg, fields...)
	m.fileLogger.Error(m.prefix+msg, fields...)
}
