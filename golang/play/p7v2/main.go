package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// this is so inconvenient in go
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		sum := 0

		r := strings.NewReader(scanner.Text())

		for {
			var num int
			n, _ := fmt.Fscan(r, &num)
			if n == 0 {
				break
			}
			sum += num
		}

		fmt.Println(sum)
	}
}
