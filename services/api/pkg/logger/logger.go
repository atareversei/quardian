package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Level      string `koanf:"level"`
	Encoding   string `koanf:"encoding"`
	OutputPath string `koanf:"output_path"`
	MaxSize    int    `koanf:"max_size"`
	MaxBackups int    `koanf:"max_backups"`
	MaxAge     int    `koanf:"max_age"`
	Compress   bool   `koanf:"compress"`
}

var globalLogger *zap.Logger
var globalDevLogger *zap.Logger

func Init(cfg Config, env string) {
	var core zapcore.Core
	switch env {
	case "prod":
		core = getProdCore(cfg, env)
	case "dev":
		core = getDevCore(cfg)
	default:
		core = zapcore.NewTee()
	}

	globalLogger = zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.FatalLevel),
	)
}

func L() *zap.Logger {
	if globalLogger == nil {
		panic("main logger not initialized")
	}
	return globalLogger
}

func getProdCore(cfg Config, env string) zapcore.Core {
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		fmt.Printf("invalid log level %s, defaulting to info: %v\n", cfg.Level, err)
		level = zapcore.InfoLevel
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(
			zapcore.EncoderConfig{
				TimeKey:        "timestamp",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "message",
				StacktraceKey:  "stacktrace",
				FunctionKey:    zapcore.OmitKey,
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.LowercaseLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			}),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.OutputPath,
			MaxSize:    cfg.MaxSize,
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,
			Compress:   cfg.Compress,
		}), level)
}

func getDevCore(cfg Config) zapcore.Core {
	return zapcore.NewCore(
		zapcore.NewConsoleEncoder(
			zapcore.EncoderConfig{
				TimeKey:       "timestamp",
				LevelKey:      "level",
				NameKey:       "logger",
				CallerKey:     zapcore.OmitKey,
				MessageKey:    zapcore.OmitKey,
				StacktraceKey: zapcore.OmitKey,
				FunctionKey:   zapcore.OmitKey,
				LineEnding:    zapcore.DefaultLineEnding,
			}),
		zapcore.AddSync(os.Stdout),
		zapcore.ErrorLevel,
	)
}
