package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// 定义颜色
const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"

	// 颜色重置 也可以理解为颜色结尾的字段
	colorReset = "\033[0m"
)

func main() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 1. 设置日志记录级别（决定哪些级别的日志会被记录）
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	// 2. 设置日志级别的编码方式（决定如何显示日志级别）
	cfg.EncoderConfig.EncodeLevel = colorLevelEncoder
	// 关键修改：改用控制台编码器   必须使用控制台编码器（Console Encoder）而非 JSON 编码器 否则颜色无法生效
	// 但是不是json的话 颜色就可以生效  是json的话 颜色就无法生效
	cfg.Encoding = "console"
	logger, _ := cfg.Build()
	logger.Info("hello world")
	logger.Warn("hello world")
	logger.Error("hello world")
	zapcore.Encoder()

}

// 根据不同的颜色编码日志级别
func colorLevelEncoder(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.DebugLevel:
		enc.AppendString(colorBlue + "DEBUG" + colorReset)
	case zapcore.InfoLevel:
		enc.AppendString(colorGreen + "INFO" + colorReset)
	case zapcore.WarnLevel:
		enc.AppendString(colorYellow + "WARN" + colorReset)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(colorRed + "ERROR" + colorReset)
	default:
		enc.AppendString(level.String()) // 默认行为
	}

}
