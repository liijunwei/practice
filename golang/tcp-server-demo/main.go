package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	server := newServer("localhost:3000")
	go func() {
		for msg := range server.msgCh {
			fmt.Printf("%s> %s", msg.from, string(msg.payload))
		}
		fmt.Println("msgch closed...")
	}()

	log.Fatal(server.start())
}

type message struct {
	from    string
	payload []byte
}

type server struct {
	listenAddr string
	listener   net.Listener
	quitCh     chan struct{}
	msgCh      chan message
}

func newServer(addr string) *server {
	return &server{
		listenAddr: addr,
		quitCh:     make(chan struct{}),
		msgCh:      make(chan message, 10),
	}
}

func (s *server) start() error {
	ln, err := net.Listen("tcp", s.listenAddr)
	if err != nil {
		return fmt.Errorf("failed to start the tcp listener: %w", err)
	}

	defer ln.Close()

	s.listener = ln

	go s.acceptLoop()

	fmt.Println("server started", s.listenAddr)
	info := strings.Split(s.listenAddr, ":")
	fmt.Println("echo demo | nc", info[0], info[1])

	<-s.quitCh
	close(s.msgCh)

	return nil
}

func (s *server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("failed to accept new connection: ", err)
			continue
		}

		fmt.Println("new connection to server:", conn.RemoteAddr())
		conn.Write([]byte("welcome!\n"))
		go s.readLoop(conn)
	}
}

func (s *server) readLoop(conn net.Conn) {
	defer conn.Close()

	const bufSize = 2048
	buf := make([]byte, bufSize)

	for {
		n, err := conn.Read(buf)
		if err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Printf("closed connection: %s, error: %+v\n", conn.RemoteAddr().String(), err)
				return
			}

			fmt.Printf("failed to read from connection: %s, error: %+v\n", conn.RemoteAddr().String(), err)
			continue
		}

		msg := buf[:n]
		s.msgCh <- message{
			from:    conn.RemoteAddr().String(),
			payload: msg,
		}

		if _, err := conn.Write([]byte("ack\n")); err != nil {
			fmt.Printf("failed to reply to client: %s, error: %+v\n", conn.RemoteAddr().String(), err)
			continue
		}
	}
}

func must(ok bool, msgs ...string) {
	if !ok {
		panic(msgs)
	}
}
