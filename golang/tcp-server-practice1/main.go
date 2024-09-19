package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	log.Fatal(newServer("localhost", 3000).start())
}

type server struct {
	ip   string
	port int
	quit chan struct{}
}

func newServer(addr string, port int) *server {
	return &server{
		ip:   addr,
		port: port,
		quit: make(chan struct{}),
	}
}

func (s *server) start() error {
	addr := fmt.Sprintf("%s:%d", s.ip, s.port)

	listner, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("fail to list %s, %w", addr, err)
	}

	go s.acceptLoop(listner)
	fmt.Println("tcp server started", addr)
	fmt.Println("nc", s.ip, s.port)

	<-s.quit

	return nil
}

func (s *server) acceptLoop(listener net.Listener) net.Conn {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept", err)
			continue
		}

		fmt.Println("new connection to server", conn.RemoteAddr().String())
		conn.Write([]byte("welcome!\n"))

		go s.readLoop(conn)
	}
}

func (s *server) readLoop(conn net.Conn) {
	defer func() {
		fmt.Println("connection closed", conn.RemoteAddr())
		conn.Close()
	}()

	const bufSize = 2048
	buf := make([]byte, bufSize)

	for {
		n, err := conn.Read(buf)
		if err != nil && errors.Is(err, io.EOF) {
			return
		}

		if err != nil {
			fmt.Println("failed to read message", conn.RemoteAddr(), err)
			continue
		}

		msg := string(buf[:n])
		fmt.Printf("received message: %s", msg)

		var out bytes.Buffer
		out.WriteString(msg)

		conn.Write(out.Bytes())
	}
}
