package main

import (
	"encoding/hex"
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

	go spy(pgConn, conn, "pg client->server")
	spy(conn, pgConn, "pg server->client")

	return nil
}

func spy(dst io.Writer, src io.Reader, connDesc string) {
	hexer := logHexer(connDesc)
	teeReader := io.TeeReader(src, hexer)
	io.Copy(dst, teeReader)
}

// https://gist.github.com/dustin/5478818
type logHexer string

func (lh logHexer) Write(b []byte) (int, error) {
	log.Printf("%v packet(%d bytes):\n%v\n", string(lh), len(b), hex.Dump(b))
	return len(b), nil
}
