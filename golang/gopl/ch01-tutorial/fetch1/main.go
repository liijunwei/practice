package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// go run ch01-tutorial/fetch1/main.go https://go.dev/ https://www.gopl.io/ https://go.dev/doc/
// go run ch01-tutorial/fetch1/main.go https://www.baidu.com/ https://cn.bing.com/ https://www.bilibili.com/

func main() {
	start := time.Now()

	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf(err.Error())
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf(err.Error())
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%2fs %7d  %s", secs, nbytes, url)
}
