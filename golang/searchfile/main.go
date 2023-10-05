package main

import (
	"log"
	"os"
	"time"
)

var query = "test"
var matches int

func main() {
	start := time.Now()

	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	startPoint := dirname + "/"
	search(startPoint)

	log.Println(matches, "matches")

	log.Println(time.Since(start))
}

func search(path string) {
	log.Print("searching in", path)

	files, err := os.ReadDir(path)
	if err != nil {
		log.Println("cannot read dir: ", path)

		return
	}

	for _, file := range files {
		name := file.Name()

		if name == query {
			matches++
		}

		if file.IsDir() {
			search(path + name + "/")
		}
	}
}
