package logging

import (
	"log"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

// Init initializes logger with the help of pre-defined configuration.
func Init() {
	logWriter := zapcore.AddSync(os.Stderr)

	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.NameKey = "name"
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.CallerKey = "caller"
	cfg.EncoderConfig.StacktraceKey = "stacktrace"

	err := cfg.Level.UnmarshalText([]byte("info"))
	if err != nil {
		log.Fatalln(err)
	}

	cfg.Encoding = "console"
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	Logger, err = cfg.Build(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewCore(zapcore.NewConsoleEncoder(cfg.EncoderConfig), logWriter, cfg.Level)
	}))
	if err != nil {
		log.Fatalln(err)
	}
}
