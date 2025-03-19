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

	// from proxy to pg
	go func() {
		if _, err := io.Copy(pgConn, conn); err != nil {
			log.Println("err", err)
		}
	}()

	// from pg to proxy
	if _, err := io.Copy(conn, pgConn); err != nil {
		log.Println("err", err)
	}

	return nil
}
