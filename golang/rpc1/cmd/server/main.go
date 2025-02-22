package main

import (
	"log"
	"net"
	"net/rpc"
)

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		log.Println("new client connected", conn.RemoteAddr())

		rpc.ServeConn(conn)
	}
}

type HelloService struct{}

func (p *HelloService) Hello(request string, reply *string) error {
	*reply = "server: " + request
	return nil
}
