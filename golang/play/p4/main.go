package main

import (
	"fmt"
)

func main() {
	for {
		var count int
		n, _ := fmt.Scan(&count)
		if count == 0 {
			break
		}

		if n != 1 {
			panic("boom")
		}

		sum := 0

		for i := 0; i < count; i++ {
			var a int
			n, _ := fmt.Scan(&a)

			if n != 1 {
				panic("boom")
			}

			sum += a
		}

		fmt.Println(sum)
	}
}
