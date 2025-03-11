package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")

	// 为输出控制台创建Core
	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		zapcore.DebugLevel,
	)
	// 为输出文件创建Core
	//指定输出的文件
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开日志文件失败")
		panic(err)
	}
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(file),
		zapcore.DebugLevel,
	)
	// 合并Core
	core := zapcore.NewTee(consoleCore, fileCore)
	// 这里把zap.AddCaller()加上去，这样就可以显示调用的文件名和行号了
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)

}

func main() {
	InitLogger()
	// zap.L()返回的是标准的zap实例
	zap.L().Info("dev this is info")
	zap.L().Warn("dev this is warn")
	zap.L().Error("dev this is error")
	// zap.S()返回的是SugaredLogger实例
	zap.S().Infof("dev this is info %s", "xxx")
	zap.S().Warnf("dev this is warn %s", "xxx")
	zap.S().Errorf("dev this is error %s", "xxx")

}
