package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
)

// 按级别进行分割
// 前缀也是同样在这里 所以和实现添加前缀的实现方式一样
// 自定义 Encoder
type levelEncoder struct {
	zapcore.Encoder
	errFile *os.File
}

func (e *levelEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// 先调用原始的 EncodeEntry 方法生成日志行
	buf, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	switch entry.Level {
	case zapcore.ErrorLevel:
		if e.errFile == nil {
			file, _ := os.OpenFile("err.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
			e.errFile = file
		}
		e.errFile.WriteString(buf.String())

	}
	// 这里可以根据日志级别进行不同的处理
	// 比如将不同级别的日志写入不同的文件或进行其他操作
	// 这里只是一个示例，你可以根据自己的需求进行修改

	return buf, nil
}

func initLogger() {
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05") // 替换时间格式化方式
	// 创建自定义的 Encoder
	encoder := &levelEncoder{
		Encoder: zapcore.NewJSONEncoder(cfg.EncoderConfig), // 使用 Console 编码器
	}
	// 创建 Core
	core := zapcore.NewCore(
		encoder, // 使用自定义的 Encoder
		// 关闭控制台输出
		zapcore.AddSync(zapcore.AddSync(io.Discard)),
		//zapcore.AddSync(os.Stdout), // 输出到控制台
		zapcore.DebugLevel, // 设置日志级别
	)

	// 创建 Logger
	logger := zap.New(core, zap.AddCaller())
	// 全局设置
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
