package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// make it private so other files won't call logger.log.
// instead, other function will call logger.Info, or logger.Debug
var log *zap.Logger

func init() {
	var err error

	config := zap.NewProductionConfig()

	encoderConfig := zap.NewProductionEncoderConfig()
	// change from "ts" to "timestamp"
	encoderConfig.TimeKey = "timestamp"
	// change time format from unix epoch to ISO8601
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// disable stacktrace
	encoderConfig.StacktraceKey = ""
	config.EncoderConfig = encoderConfig

	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
}

// Info function can be called by logger.Info
func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

// Error function can be called by logger.Error
func Error(message string, fields ...zap.Field) {
	log.Error(message, fields...)
}

// Fatal function can be called by logger.Fatal
func Fatal(message string, fields ...zap.Field) {
	log.Fatal(message, fields...)
}

// Debug function can be called by logger.Debug
func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}
