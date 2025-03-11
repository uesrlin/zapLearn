package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"os"
)

// 这里自定义的前提是text文本  而不是json格式

// 定义前缀
const logPrefix = "[MyApp] "

type prefixedEncoder struct {
	zapcore.Encoder
}

func (pe *prefixedEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buf, err := pe.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	// 在日志行的最前面添加前缀
	logLine := buf.String()
	buf.Reset()
	buf.AppendString(logPrefix + logLine)
	return buf, nil

}
func main() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建自定义的 Encoder
	encoder := &prefixedEncoder{
		Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}
	// 只有在console下才可以使用
	// 创建 Core
	core := zapcore.NewCore(
		encoder,                    // 使用自定义的 Encoder
		zapcore.AddSync(os.Stdout), // 输出到控制台
		zapcore.DebugLevel,         // 设置日志级别
	)

	// 创建 Logger
	logger := zap.New(core, zap.AddCaller())

	logger.Info("dev this is info")
	logger.Warn("dev this is warn")
	logger.Error("dev this is error")

	// [MyApp] 2025-03-11 17:25:10	info	prefix/main.go:46	dev this is info
	//[MyApp] 2025-03-11 17:25:10	warn	prefix/main.go:47	dev this is warn
	//[MyApp] 2025-03-11 17:25:10	error	prefix/main.go:48	dev this is error
}
