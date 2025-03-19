package main

import (
	"io"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", "localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			if err == io.EOF {
				return
			}

			log.Println("err", err)
		}

		go echo(conn)
	}
}

func echo(conn net.Conn) {
	io.Copy(conn, conn)
}
