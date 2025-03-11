package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"strings"
	"time"
)

func main() {
	dl := &dynamicLogger{base: "app"}
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 在 initLogger 函数中添加：
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(&lumberjack.Logger{
				// 自定义文件名生成器
				// 但是这样的额话不压缩就正常  如果压缩的话 压缩文件的名字会出现双时间的问题 如果修改的话 需要修改整个lumberjack 不推荐
				// 而且由于是按照大小进行分割app-20250311-193915-2025-03-11T11-39-15.835.log  这样会发现前后的时间格式不一样 不推荐
				Filename: dl.Filename(), // 日志文件名
				//Filename:   "app.log",
				MaxSize:    1,     // 单个文件最大尺寸 (MB)
				MaxBackups: 3,     // 保留旧文件的最大个数
				MaxAge:     7,     // 保留旧文件的最大天数
				Compress:   false, // 是否压缩/归档旧文件
			}),
		),
		zapcore.DebugLevel,
	)
	logger := zap.New(core, zap.AddCaller())
	// 全局设置
	zap.ReplaceGlobals(logger)

	// 测试写入日志
	for i := 0; i < 2000; i++ {
		zap.L().Info(strings.Repeat("A", 1024)) // 每次写入 1KB
	}

}

// 方法 2：动态生成文件名
type dynamicLogger struct {
	base string
}

func (d *dynamicLogger) Filename() string {
	return fmt.Sprintf("%s-%s.log",
		d.base,
		time.Now().Format("20060102-150405"))
}
