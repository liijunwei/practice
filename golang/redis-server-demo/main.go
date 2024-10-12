package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
			fmt.Println("falied to accept request", err.Error())
			continue
		}

		fmt.Println("new connection to server", conn.RemoteAddr().String())

		go func(conn net.Conn) {
			defer conn.Close()

			if err := s.handleConnection(conn); err != nil {
				fmt.Println("falied to handle connection", err.Error())
			}
		}(conn)
	}
}

// goal#1: handle ping
// goal#2: handle set a 1
// goal#3: handle get a
// other goals: TBC
func (s *server) handleConnection(conn net.Conn) error {
	const bufSize = 2048
	// buf := make([]byte, bufSize)

	for {
		resp := newResp(conn)

		value, err := resp.read()
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("failed to read message", conn.RemoteAddr(), err)
			return err
		}

		if len(value.Array) == 0 {
			fmt.Println("invalid request, expected array length > 0")
			continue
		}

		data, err := json.Marshal(value)
		if err != nil {
			log.Panic(err)
		}
		fmt.Println("value", string(data))

		command := strings.ToUpper(value.Array[0].Bulk)
		args := value.Array[1:]

		fmt.Println("command", command)
		fmt.Println("args", args)

		// msg := string(buf[:n])
		// fmt.Println("client request:", strconv.Quote(msg))

		conn.Write([]byte("+PONG\r\n"))
	}

	return nil
}
