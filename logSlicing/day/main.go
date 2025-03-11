package main

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// 按天进行分割

// 思路就是根据之前写的文件可以判断出newCore中的zapcore.AddSync(writeSyncer),方法实现写的方法
// 可以判断就是实现write接口就可

type dynamicLogWriter struct {
	mu         sync.Mutex
	currentDay string
	file       *os.File
	logDir     string
}

/*

func AddSync(w io.Writer) WriteSyncer {
	switch w := w.(type) {
	case WriteSyncer: // 如果对象本身已实现 Sync()
		return w    // 直接使用原对象
	default:         // 如果对象只有 io.Writer
		return writerWrapper{w} // 包装成带空实现的 Sync()
	}
}
文件同步需求 当使用 zap.Logger.Sync() 时，需要确保缓冲区内容完全写入磁盘。
如果使用默认的 writerWrapper，其 Sync() 是空实现

通过自定义 Sync() 方法，可以调用 file.Sync() 强制将文件系统缓冲区的内容写入磁盘，避免程序崩溃时丢失日志

*/

func (w *dynamicLogWriter) Sync() error {
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.file != nil {
		return w.file.Sync()
	}
	return nil
}

func (w *dynamicLogWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// 获取当前日期
	currentDay := time.Now().Format("2006-01-02")
	// 检查日期是否发生变化
	if currentDay != w.currentDay {
		if w.file != nil {
			if err := w.file.Close(); err != nil {
				return 0, fmt.Errorf("关闭日志文件失败: %w", err)
			}

		}

		filePath := filepath.Join(w.logDir, "app-"+currentDay+".log")
		file, er := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if er != nil {
			return 0, er
		}
		w.file = file
		w.currentDay = currentDay
	}
	// 如果日期没变的话就写入日志
	return w.file.Write(p)
}

// 初始化全局日志
func initLogger() {
	// 使用 zap 的 NewDevelopmentConfig 快速配置
	cfg := zap.NewDevelopmentConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建 Logger
	writer := &dynamicLogWriter{
		logDir: "logs",
	}
	// 提前创建日志目录
	if err := os.MkdirAll(writer.logDir, 0755); err != nil {
		panic(err)
	}
	// 确保初始文件存在
	_, _ = writer.Write([]byte{})
	// 创建 Core
	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout), // 输出到控制台
		zapcore.DebugLevel,         // 设置日志级别
	)
	fileCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writer), // 输出到文件
		zapcore.DebugLevel,      // 设置日志级别
	)
	core := zapcore.NewTee(consoleCore, fileCore)
	// 创建 Logger
	logger := zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(logger)
}

func main() {
	initLogger()
	zap.L().Info("dev this is info")
	zap.L().Warn("dev this is warn")
	zap.L().Error("dev this is error")
	// zap.S()返回的是SugaredLogger实例
	zap.S().Infof("dev this is info %s", "xxx")
	zap.S().Warnf("dev this is warn %s", "xxx")
	zap.S().Errorf("dev this is error %s", "xxx")

}
