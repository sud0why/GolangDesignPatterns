package main

import "fmt"

//状态模式
//将事件触发的状态转移和动作指定拆分到不同的状态类中，避免分支判断逻辑
//常用于实现状态机

//容器
//TCP连接状态：关闭、等待连接、连接
//关闭，发送SYN=>等待连接
//等待连接，接收ACK=>连接
//等待连接，接收FIN=>关闭
//连接，接收FIN=>关闭

// TCPState 状态接口
type TCPState interface {
	SendSYN(*TCPConn)
	RecvFIN(*TCPConn)
	RecvACK(*TCPConn)
	String() string
}

type CloseState struct {
}

func (c *CloseState) String() string {
	return "Close"
}

func (c *CloseState) SendSYN(conn *TCPConn) {
	conn.StateNow = conn.States["Wait"]
	fmt.Printf("%s SendSYN => %s\n", c.String(), conn.StateNow.String())
}

func (c *CloseState) RecvFIN(conn *TCPConn) {
	fmt.Printf("%s RecvFIN => %s\n", c.String(), c.String())
}

func (c *CloseState) RecvACK(conn *TCPConn) {
	fmt.Printf("%s RecvACK => %s\n", c.String(), c.String())
}

type WaitState struct {
}

func (w *WaitState) SendSYN(conn *TCPConn) {
	fmt.Printf("%s SendSYN => %s\n", w.String(), conn.StateNow.String())
}

func (w *WaitState) RecvFIN(conn *TCPConn) {
	conn.StateNow = conn.States["Close"]
	fmt.Printf("%s RecvFIN => %s\n", w.String(), conn.StateNow.String())
}

func (w *WaitState) RecvACK(conn *TCPConn) {
	conn.StateNow = conn.States["Connect"]
	fmt.Printf("%s RecvACK => %s\n", w.String(), conn.StateNow.String())
}

func (w *WaitState) String() string {
	return "Wait"
}

type ConnectState struct {
}

func (c *ConnectState) SendSYN(conn *TCPConn) {
	fmt.Printf("%s SendSYN => %s\n", c.String(), conn.StateNow.String())
}

func (c *ConnectState) RecvFIN(conn *TCPConn) {
	conn.StateNow = conn.States["Close"]
	fmt.Printf("%s RecvFIN => %s\n", c.String(), conn.StateNow.String())
}

func (c *ConnectState) RecvACK(conn *TCPConn) {
	fmt.Printf("%s RecvACK => %s\n", c.String(), conn.StateNow.String())
}

func (c *ConnectState) String() string {
	return "Connect"
}

type TCPConn struct {
	StateNow TCPState
	States   map[string]TCPState
}

func (t *TCPConn) SendSYN() {
	t.StateNow.SendSYN(t)
}

func (t *TCPConn) RecvFIN() {
	t.StateNow.RecvFIN(t)
}

func (t *TCPConn) RecvACK() {
	t.StateNow.RecvACK(t)
}

func NewTCPConn(initState string) *TCPConn {
	tcpConn := &TCPConn{
		States: map[string]TCPState{
			"Wait":    &WaitState{},
			"Connect": &ConnectState{},
			"Close":   &CloseState{},
		},
	}
	tcpConn.StateNow = tcpConn.States[initState]
	return tcpConn
}

func main() {
	tcpConn := NewTCPConn("Close")
	tcpConn.RecvACK()
	tcpConn.RecvACK()
	// => wait
	tcpConn.SendSYN()
	// => close
	tcpConn.RecvFIN()
	// => wait
	tcpConn.SendSYN()
	// => conn
	tcpConn.RecvACK()
	// => close
	tcpConn.RecvFIN()
}
