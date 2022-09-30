package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// jj "Hello, Reader!".chars.index_by {|c| c.bytes}
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 10)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
