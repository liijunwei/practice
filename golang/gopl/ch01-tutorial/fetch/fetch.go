package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const httpPrefix = "http://"

func main() {
	for _, u := range os.Args[1:] {
		url := checkAndPrependHttpPrefix(u)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		b := new(strings.Builder)
		written, err := io.Copy(b, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "copy: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("%d bytes written\n\n", written)

		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		fmt.Printf("status: %s\n", resp.Status)
		fmt.Printf("%s", b)
	}
}

func checkAndPrependHttpPrefix(rawUrl string) string {
	if strings.HasPrefix(rawUrl, httpPrefix) {
		return rawUrl
	}

	return fmt.Sprintf("%s%s", httpPrefix, rawUrl)
}
