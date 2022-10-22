package main

import (
	"fmt"
	"os"
)

// gofmt -w golang/*.go && go build -o a.out golang/cli-demo.go && ./a.out a b c d 3 f
func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
