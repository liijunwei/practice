// Thanks Frank
// https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go

// it(struct tags) has no effect on the operation of this Go code.
// To use struct tags to accomplish something, other Go code must be written to examine structs at runtime.

// gofmt -w golang/*.go && go run golang/golang-tag-demo2.go

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type User struct {
	name          string
	password      string
	preferredFish []string
	CreatedAt     time.Time
}

func main() {
	u := &User{
		name:      "Sammy the Shark",
		password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
