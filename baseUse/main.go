package main

import (
	"go.uber.org/zap"
)

func main() {
	prod()

}

func dev() {
	// 默认模式 text 格式
	logger, _ := zap.NewDevelopment()
	logger.Info("hello world")
	logger.Warn("hello world")
	logger.Error("hello world")
	/*
		2025-03-11T13:09:56.431+0800	INFO	baseUse/main.go:14	hello world
		2025-03-11T13:09:56.455+0800	WARN	baseUse/main.go:15	hello world
	*/
}

func test() {
	// json格式 但是缺少时间戳 以及行号
	logger := zap.NewExample()
	logger.Info("hello world")
	logger.Warn("hello world")
	logger.Error("hello world")
	/*
		{"level":"info","msg":"hello world"}
		{"level":"warn","msg":"hello world"}
		{"level":"error","msg":"hello world"}
	*/
}

func prod() {
	// 生产模式 一般采用这个 json格式 并且有时间戳 以及行号
	logger, _ := zap.NewProduction()
	logger.Info("hello world")
	logger.Warn("hello world")
	logger.Error("hello world")
	/*
		{"level":"info","ts":1741669865.5499392,"caller":"baseUse/main.go:37","msg":"hello world"}
		{"level":"warn","ts":1741669865.5499392,"caller":"baseUse/main.go:38","msg":"hello world"}
		{"level":"error","ts":1741669865.5499392,"caller":"baseUse/main.go:39","msg":"hello world","stacktrace":"main.prod\n\tE:/go_example/zapLearn/baseUse/main.go:39\nmain.main\n\tE:/go_example/zapLearn/baseUse/main.go:8\nruntime.main\n\tD:/go/go1.21.11/src/runtime/proc.go:267"}

	*/
}
