package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

func main() {
	log.Fatal(newServer("localhost", 6378).run())
}

type server struct {
	ip   string
	port int
	quit chan struct{}
}

func newServer(ip string, port int) *server {
	return &server{
		ip:   ip,
		port: port,
		quit: make(chan struct{}),
	}
}

func (s *server) run() error {
	addr := fmt.Sprintf("%s:%d", s.ip, s.port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	go s.acceptLoop(listener)
	fmt.Println("redis server started", addr)

	<-s.quit

	return nil
}

func (s *server) acceptLoop(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("falied to accept request: %v", err)
			continue
		}

		fmt.Println("new connection to server", conn.RemoteAddr().String())

		go func(conn net.Conn) {
			defer conn.Close()

			if err := s.handleConnection(conn); err != nil {
				fmt.Println("falied to handle connection: %v", err)
			}
		}(conn)
	}
}

func (s *server) handleConnection(conn net.Conn) error {
	const bufSize = 2048
	buf := make([]byte, bufSize)

	n, err := conn.Read(buf)
	if err != nil {
		if err == io.EOF {
			return nil
		}

		fmt.Println("failed to read message", conn.RemoteAddr(), err)
		return err
	}

	msg := string(buf[:n])
	fmt.Println("client request:", strconv.Quote(msg))

	conn.Write([]byte("+OK\r\n"))

	return nil
}
