package main

import "fmt"

var (
	CommitHash = "unset"
	CommitTime = "unset"
)

func main() {
	fmt.Println("CommitHash", CommitHash)
	fmt.Println("CommitTime", CommitTime)
}
