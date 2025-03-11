package main

import "go.uber.org/zap"

func main() {
	New()

}

func Old() {
	//适合对性能要求较高的场景，尤其是在生产环境中
	logger, _ := zap.NewProduction()
	logger.Info("hello ",
		zap.String("name", "zhangsan"),
		zap.Int("age", 18),
		zap.Bool("active", true))
	//{"level":"info","ts":1741681047.912754,"caller":"structLog/main.go:13","msg":"hello ","name":"zhangsan","age":18,"active":true}
}

func New() {
	//对性能要求不是特别高的场景下使用  但是可以支持自定义输出
	logger, _ := zap.NewProduction()
	sl := logger.Sugar()
	sl.Info("hello ", "name", "zhangsan", "age", 18, "active", true)
	//{"level":"info","ts":1741681076.6357212,"caller":"structLog/main.go:24","msg":"hello namezhangsanage18activetrue"}

}
