package main

// from chatGPT
import (
	"fmt"
	"strings"
	"time"
)

func main() {
	for i := 0; i <= 100; i += 10 {
		fmt.Printf("\rProgress: [%-10s] %d%%", strings.Repeat("=", i/10), i)
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("\nDone!")
}
