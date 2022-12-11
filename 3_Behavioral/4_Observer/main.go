package main

//观察者模式
//订阅、推送模式
import (
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"
)

// Bus 总线接口
type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// AsyncEventBus 异步事件总线
type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

// NewAsyncEventBus new
func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

// Subscribe 订阅
func (bus *AsyncEventBus) Subscribe(topic string, f interface{}) error {
	//上锁
	bus.lock.Lock()
	defer bus.lock.Unlock()

	//检查f类型
	v := reflect.ValueOf(f)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	//将函数放置到handlers
	handler, ok := bus.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler, v)
	bus.handlers[topic] = handler

	return nil
}

// Publish 发布
func (bus *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := bus.handlers[topic]
	if !ok {
		fmt.Println("not found handlers in topic:", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}

	//异步执行
	for i := range handlers {
		go handlers[i].Call(params)
	}
}

//观察者

func sub1(msg1, msg2 string) {
	time.Sleep(1 * time.Microsecond)
	fmt.Printf("sub1, %s %s\n", msg1, msg2)
}

func sub2(msg1, msg2 string) {
	fmt.Printf("sub2, %s %s\n", msg1, msg2)
}

func main() {
	bus := NewAsyncEventBus()
	//向主题"topic:1"注册了两个观察者
	err := bus.Subscribe("topic:1", sub1)
	if err != nil {
		log.Printf("bus.Subscribe fail: %s", err)
	}
	err = bus.Subscribe("topic:1", sub2)
	if err != nil {
		log.Printf("bus.Subscribe fail: %s", err)
	}
	//向"topic:1"发送消息
	bus.Publish("topic:1", "test1", "test2")
	bus.Publish("topic:1", "testA", "testB")
	time.Sleep(1 * time.Second)
}
