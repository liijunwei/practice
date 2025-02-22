package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	defer client.Close()

	assert(len(os.Args) == 2)

	msg := os.Args[1]

	var reply string

	err = client.Call("/ns/HelloService.Hello", msg, &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}

func assert(cond bool, msg ...string) {
	if !cond {
		log.Fatal(strings.Join(msg, " "))
	}
}
