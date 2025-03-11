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
	// 这里将输出的时间key 由原来的ts 改为 time
	cfg.EncoderConfig.TimeKey = "time"

	file, err := os.OpenFile(getPath()+"/log/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开日志文件失败")
		panic(err)
	}
	//
	writeSyncer := zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(file),
	)

	core := zapcore.NewCore(
		// 注意NewConsoleEncode 是控制台输出，控制台输出的话是文本格式，NewJSONEncoder 是json格式输出
		//zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(writeSyncer),
		zapcore.DebugLevel,
	)
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

func getPath() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("获取当前路径失败:", err)
		return ""
	}
	return dir
}
