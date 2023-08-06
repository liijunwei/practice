package main

// go build -o a.out -ldflags "-X main.CommitHash=$(git rev-parse --short HEAD) -X main.CommitTime=$(git log main -n1 --format='%cd' --date='iso-strict')" go-build-demo.go

import "fmt"

var (
	CommitHash = "unset"
	CommitTime = "unset"
)

func main() {
	fmt.Println("CommitHash", CommitHash)
	fmt.Println("CommitTime", CommitTime)
}

