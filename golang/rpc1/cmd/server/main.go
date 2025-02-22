package main

import (
	"log"
	"net"
	"net/rpc"
	"runtime"
	"time"
)

func main() {
	go monitorGoroutines()

	rpc.RegisterName("/ns/HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		// log.Println("new client connected", conn.RemoteAddr())
		// time.Sleep(10 * time.Millisecond)

		rpc.ServeConn(conn)
	}
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "server: " + request
	return nil
}

func monitorGoroutines() {
	for {
		time.Sleep(1 * time.Second)
		log.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
	}
}
