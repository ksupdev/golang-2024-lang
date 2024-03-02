package logs

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var log *zap.Logger

func init() {
	//Run before main func()

	config := zap.NewProductionConfig()
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.StacktraceKey = "" // Disable StacktraceKey

	var err error
	log, err = config.Build(zap.AddCallerSkip(1))

	if err != nil {
		panic(err)
	}
	// Log, _ = zap.NewProduction()
}

func Info(message string, fields ...zap.Field) {
	log.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	log.Debug(message, fields...)
}

func Error(message interface{}, fields ...zap.Field) {
	// msg, ok := message.(error)
	// if ok {
	// 	log.Error(msg.Error(), fields...)
	// }

	switch v := message.(type) {
	case error:
		log.Error(v.Error(), fields...)
	case string:
		log.Error(v, fields...)
	}

}
