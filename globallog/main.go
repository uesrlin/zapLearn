package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func main() {
	InitLogger()
	dev()

}

func InitLogger() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	logger, _ := cfg.Build()
	zap.ReplaceGlobals(logger)
}
func dev() {
	// zap.L()返回的是标准的zap实例
	zap.L().Info("dev this is info")
	zap.L().Warn("dev this is warn")
	zap.L().Error("dev this is error")
	// zap.S()返回的是SugaredLogger实例
	zap.S().Infof("dev this is info %s", "xxx")
	zap.S().Warnf("dev this is warn %s", "xxx")
	zap.S().Errorf("dev this is error %s", "xxx")
}
