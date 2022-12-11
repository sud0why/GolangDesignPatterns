package main

import (
	"fmt"
	"strings"
)

//职责链模式
//通过职责链实现认证，授权，限流，过滤等操作，例如Gin中的中间件

type Handler interface {
	Process(*MyGinContext)
}

type MyGinContext struct {
	builder  strings.Builder
	handlers []Handler
	index    int
}

func (g *MyGinContext) Use(handlers ...Handler) {
	for _, handler := range handlers {
		g.handlers = append(g.handlers, handler)
	}
}

func (g *MyGinContext) Next() {
	g.index++
	for g.index < len(g.handlers) {
		handler := g.handlers[g.index]
		handler.Process(g)
		g.index++
	}
}

type AuthHandler struct {
}

func (a AuthHandler) Process(context *MyGinContext) {
	context.builder.WriteString("Auth Start\n")
	context.builder.WriteString("Auth Done\n")
}

type SaveHandler struct {
}

func (s SaveHandler) Process(context *MyGinContext) {
	context.builder.WriteString("Save Start\n")
	context.builder.WriteString("Save Done\n")
}

type LogALLHandler struct {
}

func (l LogALLHandler) Process(context *MyGinContext) {
	context.builder.WriteString("Log Start\n")
	context.Next()
	context.builder.WriteString("Log Done\n")
}

type LogHandler struct {
}

func (l LogHandler) Process(context *MyGinContext) {
	context.builder.WriteString("Log Start\n")
	context.builder.WriteString("Log Done\n")
}

func main() {
	ginContext := &MyGinContext{
		builder:  strings.Builder{},
		handlers: nil,
		index:    -1,
	}
	ginContext.Use(LogALLHandler{}, AuthHandler{}, SaveHandler{})
	ginContext.Next()
	fmt.Println(ginContext.builder.String())

	ginContext1 := &MyGinContext{
		builder:  strings.Builder{},
		handlers: nil,
		index:    -1,
	}
	ginContext1.Use(LogHandler{}, AuthHandler{}, SaveHandler{})
	ginContext1.Next()
	fmt.Println(ginContext1.builder.String())
}
