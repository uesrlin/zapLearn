package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	logger, _ := config.Build()
	logger.Info("hello world")
	logger.Warn("hello world")
	logger.Error("hello world")

}
