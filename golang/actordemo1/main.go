package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建第一个Actor
	actor1 := NewActor(func(msg Message) {
		switch data := msg.Data.(type) {
		case string:
			fmt.Printf("Actor1 收到消息: %s\n", data)
			if data == "你好" {
				// 回复消息
				msg.Sender.Tell("你好，我是Actor1")
			}
		}
	})

	// 创建第二个Actor
	actor2 := NewActor(func(msg Message) {
		switch data := msg.Data.(type) {
		case string:
			fmt.Printf("Actor2 收到消息: %s\n", data)
		default:
			panic("unsupported data type")
		}
	})

	// 发送消息
	actor1.ref.TellWithSender("你好", actor2.ref)

	// 给一点时间让消息处理完成
	time.Sleep(100 * time.Millisecond)

	// 停止Actor
	actor1.Stop()
	actor2.Stop()
}

// Message 表示发送给Actor的消息
type Message struct {
	Data   any
	Sender *ActorRef
}

// ActorRef 是对Actor的引用，用于发送消息
type ActorRef struct {
	mailbox chan Message
}

// Actor 表示一个Actor实体
type Actor struct {
	behavior func(Message)
	ref      *ActorRef
}

// NewActor 创建一个新的Actor
func NewActor(behavior func(Message)) *Actor {
	mailbox := make(chan Message, 100) // 带缓冲的邮箱
	ref := &ActorRef{mailbox: mailbox}
	actor := &Actor{
		behavior: behavior,
		ref:      ref,
	}

	go actor.run()
	return actor
}

// run 启动Actor的消息处理循环
func (a *Actor) run() {
	for msg := range a.ref.mailbox {
		a.behavior(msg)
	}
}

// Tell 发送消息给Actor（异步）
func (ref *ActorRef) Tell(data any) {
	ref.mailbox <- Message{Data: data}
}

// TellWithSender 发送消息并包含发送者信息
func (ref *ActorRef) TellWithSender(data any, sender *ActorRef) {
	ref.mailbox <- Message{Data: data, Sender: sender}
}

// Stop 停止Actor
func (a *Actor) Stop() {
	close(a.ref.mailbox)
}
