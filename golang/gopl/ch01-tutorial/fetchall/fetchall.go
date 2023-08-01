package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const httpPrefix = "http://"
const httpsPrefix = "https://"

// go run fetchall/fetchall.go https://formulae.brew.sh/api/cask.json https://formulae.brew.sh/api/formula.json gopl.io
func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] { // ignore the loop variable if it's not used
		fmt.Println(<-ch) // n urls -> n channel message -> receive n times
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	url = checkAndPrependHttpPrefix(url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	written, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %10d  %s", secs, written, url)
}

func checkAndPrependHttpPrefix(rawUrl string) string {
	if strings.HasPrefix(rawUrl, httpPrefix) || strings.HasPrefix(rawUrl, httpsPrefix) {
		return rawUrl
	}

	return fmt.Sprintf("%s%s", httpPrefix, rawUrl)
}
