// Thanks Frank
// https://www.digitalocean.com/community/tutorials/how-to-use-struct-tags-in-go

// it(struct tags) has no effect on the operation of this Go code.
// To use struct tags to accomplish something, other Go code must be written to examine structs at runtime.

// Camel casing fields properly requires that the first character be lower-cased.
// While JSON doesn’t care how you name your fields, Go does, as it indicates the visibility of the field outside of the package.

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
	Name          string    `json:"name1"`
	Password      string    `json:"password"`
	PreferredFish []string  `json:"preferredFish"`
	CreatedAt     time.Time `json:"createdAt"`
}

func main() {
	u := &User{
		Name:      "Sammy the Shark",
		Password:  "fisharegreat",
		CreatedAt: time.Now(),
	}

	out, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))
}
