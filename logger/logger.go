package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	var err error
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config := zap.NewProductionConfig()
	config.EncoderConfig = encoderConfig
	if log, err = config.Build(zap.AddCallerSkip(1)); err != nil {
		panic("Something is really wrong" + err.Error())
	}
}

func Info(msg string, fields ...zapcore.Field) {
	log.Info(msg, fields...)
}
func Debug(msg string, fields ...zapcore.Field) {
	log.Debug(msg, fields...)
}
func Error(msg string, fields ...zapcore.Field) {
	log.Error(msg, fields...)
}
func Fatal(msg string, fields ...zapcore.Field) {
	log.Fatal(msg, fields...)
}
