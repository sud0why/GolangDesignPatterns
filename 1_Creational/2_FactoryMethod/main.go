package main

import "fmt"

// 工厂方法
// 在简单工厂基础上，把工厂也抽象了
// 简单工厂+开闭原则=工厂方法

// 场景：
// 多种日志记录器，都实现了DoLog方法，支持输出到不同的位置

// 抽象层

type AbstractLogger interface {
	DoLog()
}

type AbstractFactory interface {
	NewLogger() AbstractLogger
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

type StdioLoggerFactory struct {
}

func (s StdioLoggerFactory) NewLogger() AbstractLogger {
	return new(StdioLogger)
}

type FileLoggerFactory struct {
}

func (f FileLoggerFactory) NewLogger() AbstractLogger {
	return new(FileLogger)
}

type MemoryLoggerFactory struct {
}

func (m MemoryLoggerFactory) NewLogger() AbstractLogger {
	return new(MemoryLogger)
}

// 业务代码

func main() {
	stdioLoggerFactory := new(StdioLoggerFactory)
	stdioLogger := stdioLoggerFactory.NewLogger()
	stdioLogger.DoLog()

	fileLoggerFactory := new(FileLoggerFactory)
	fileLogger := fileLoggerFactory.NewLogger()
	fileLogger.DoLog()

}
