package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	sep := ","

	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), sep)

		sort.Strings(fields)
		fmt.Println(strings.Join(fields, sep))
	}
}
