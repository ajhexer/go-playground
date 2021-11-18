package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)
var Log ZapLogger

type ZapLogger struct {
	logger *zap.Logger
}

func (z ZapLogger) Info(message string)  {
	z.logger.Info(message)
}

func (z ZapLogger) Debug(message string)  {
	z.logger.Debug(message)
}

func (z ZapLogger) Error(message string)  {
	z.logger.Error(message)
}

func init(){

	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger , err := logConfig.Build()
	if err!=nil {
		panic(err)
	}
	Log = ZapLogger{logger: logger}
}

