package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

// go run db-proxy1/main.go
// psql 'postgres://postgres:123@localhost:9999/dummydb?sslmode=disable&application_name=dummytester'
func main() {
	l, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("db proxy listening on", l.Addr())
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Println("err", err)
		}

		log.Println("conn", conn.RemoteAddr())

		if err := proxy(conn); err != nil {
			log.Println(err)
		}
	}
}

func proxy(conn net.Conn) error {
	defer conn.Close()

	pg := "localhost:5432"

	pgConn, err := net.Dial("tcp", pg)
	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", pg, err)
	}

	defer pgConn.Close()

	// Data flow from client to postgres
	go func() {
		// Use custom data capture function
		if err := copyAndCapture(pgConn, conn, "[client -> postgres]"); err != nil {
			log.Println("client -> postgres err:", err)
		}
	}()

	if err := copyAndCapture(conn, pgConn, "[postgres -> client]"); err != nil {
		log.Println("postgres -> client err:", err)
	}

	return nil
}

// copyAndCapture copies data while capturing and analyzing traffic
func copyAndCapture(dst io.Writer, src io.Reader, direction string) error {
	buf := make([]byte, 4096)
	for {
		// Read data
		nr, err := src.Read(buf)
		if nr > 0 {
			capturePacket(buf[:nr], direction)

			// Write to destination, maintaining normal flow
			nw, err := dst.Write(buf[:nr])
			if err != nil {
				return err
			}
			if nr != nw {
				return io.ErrShortWrite
			}
		}
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
}

// capturePacket analyzes captured packets
func capturePacket(data []byte, direction string) {
	log.Println(direction, "size =", len(data), "bytes", string(data))
}
