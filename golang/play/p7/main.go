package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// this is so inconvenient in go
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		sum := 0
		fields := strings.Fields(scanner.Text())

		for _, field := range fields {
			n, err := strconv.Atoi(field)
			boom(err)

			sum += n
		}

		fmt.Println(sum)
	}
}

func boom(err error) {
	if err != nil {
		panic(err)
	}
}
