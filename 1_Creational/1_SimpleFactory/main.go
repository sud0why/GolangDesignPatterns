package main

import "fmt"

// 简单工厂
// 违背开闭原则，比如支持更多的种类，需要修改工厂类逻辑
// 工厂类职责过重

// 场景：
// 多种日志记录器，都实现了DoLog方法，支持输出到不同的位置

// 抽象层

type AbstractLogger interface {
	DoLog()
}

// 实现层

type StdioLogger struct {
	AbstractLogger
}

func (sl *StdioLogger) DoLog() {
	fmt.Println("Log to stdio")
}

type FileLogger struct {
}

func (fl *FileLogger) DoLog() {
	fmt.Println("Log to file")
}

type MemoryLogger struct {
}

func (ml *MemoryLogger) DoLog() {
	fmt.Println("Log to memory")
}

// 工厂类

type LoggerFactory struct {
}

func (lf *LoggerFactory) NewLogger(target string) AbstractLogger {
	switch target {
	case "stdio":
		return new(StdioLogger)
	case "file":
		return new(FileLogger)
	case "memory":
		return new(MemoryLogger)
	default:
		return new(StdioLogger)
	}
}

// 业务代码

func main() {
	loggerFactory := new(LoggerFactory)
	stdioLogger := loggerFactory.NewLogger("stdio")
	stdioLogger.DoLog()

	memoryLogger := loggerFactory.NewLogger("memory")
	memoryLogger.DoLog()
}
