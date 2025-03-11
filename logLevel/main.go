package main

import "go.uber.org/zap"

func main() {
	// 使用 zap 的 NewProductionConfig快速配置
	config := zap.NewProductionConfig()
	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logger, _ := config.Build()
	logger.Info("hello world")
	logger.Warn("hello world")
	logger.Error("hello world")

}
